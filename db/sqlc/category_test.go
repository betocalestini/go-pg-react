package db

import (
	"context"
	"testing"

	"github.com/betocalestini/go-pg-react/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomCategory(t *testing.T) Category {
	user := CreateRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       util.RandomName(10),
		Type:        util.RandomName(8),
		Description: util.RandomName(50),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, category.Title, arg.Title)
	require.Equal(t, category.Type, arg.Type)
	require.Equal(t, category.Description, arg.Description)
	require.Equal(t, category.UserID, arg.UserID)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func Test_CreateCategory(t *testing.T) {
	CreateRandomCategory(t)
}

func Test_DeleteCategory(t *testing.T) {
	category1 := CreateRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)

	category3, err := testQueries.GetDeletedCategory(context.Background(), category1.ID)
	require.NotEmpty(t, category3.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, category1.ID, category3.ID)

}

func Test_GetCategories(t *testing.T) {
	category1 := CreateRandomCategory(t)
	arg := GetCategoriesParams{
		UserID:      category1.UserID,
		Title:       category1.Title,
		Type:        category1.Type,
		Description: category1.Description,
	}
	categories, err := testQueries.GetCategories(context.Background(), arg)
	require.NoError(t, err)

	for _, category := range categories {
		require.Equal(t, category1.UserID, category.UserID)
		require.Equal(t, category1.Title, category.Title)
		require.Equal(t, category1.Type, category.Type)
		require.Equal(t, category1.Description, category.Description)
		require.NotEmpty(t, category.CreatedAt)
	}
}

func Test_GetCategoriesByTitle(t *testing.T) {
	category1 := CreateRandomCategory(t)
	categories, err := testQueries.GetCategoriesByTitle(context.Background(), category1.Title)
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.Equal(t, category1.Title, category.Title)
		require.NotEmpty(t, category.CreatedAt)
	}
}

func Test_GetCategoryById(t *testing.T) {
	category1 := CreateRandomCategory(t)
	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.UserID, category2.UserID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Type, category2.Type)
	require.Equal(t, category1.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func Test_GetDeletedCategories(t *testing.T) {
	category1 := CreateRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetDeletedCategory(context.Background(), category1.ID)
	require.NotEmpty(t, category2.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, category1.ID, category2.ID)

	categories, err := testQueries.GetDeletedCategories(context.Background())
	require.NoError(t, err)
	var match bool = false
	for _, category := range categories {
		require.NotEmpty(t, category.DeletedAt)
		if category1.ID == category.ID {
			match = true
			require.Equal(t, category2.UserID, category.UserID)
			require.Equal(t, category2.Title, category.Title)
			require.Equal(t, category2.Type, category.Type)
			require.Equal(t, category2.Description, category.Description)
			require.Equal(t, category2.DeletedAt, category.DeletedAt)
		}
	}
	require.Equal(t, true, match)
}

func Test_GetDeletedCategory(t *testing.T) {
	category1 := CreateRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.ID)
	require.NoError(t, err)

	category2, err := testQueries.GetDeletedCategory(context.Background(), category1.ID)
	require.NotEmpty(t, category2.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, category1.ID, category2.ID)
}

func Test_UpdateCategory(t *testing.T) {
	category1 := CreateRandomCategory(t)
	arg := UpdateCategoryParams{
		ID:          category1.ID,
		Title:       util.RandomName(10),
		Type:        util.RandomName(8),
		Description: util.RandomName(50),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.UserID, category2.UserID)

	require.NotEqual(t, category1.Title, category2.Title, category1, category2)
	require.NotEqual(t, category1.Type, category2.Type)
	require.NotEqual(t, category1.Description, category2.Description)
	require.NotEqual(t, category1.UpdatedAt, category2.UpdatedAt)

	require.NotEmpty(t, category2)
	require.NotEmpty(t, category2.CreatedAt)
	require.NotEmpty(t, category2.UpdatedAt)
}
