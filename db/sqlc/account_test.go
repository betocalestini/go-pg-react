package db

import (
	"context"
	"go-pg-react/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	category := CreateRandomCategory(t)
	arg := CreateAccountParams{
		UserID:      category.UserID,
		CategoryID:  category.ID,
		Title:       util.RandomName(10),
		Type:        category.Type,
		Description: util.RandomName(50),
		Value:       util.RandomFloatString(3, 2),
		Date:        util.RandomDate(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.UserID, arg.UserID)
	require.Equal(t, account.Title, arg.Title)
	require.Equal(t, account.CategoryID, arg.CategoryID)
	require.Equal(t, account.Type, arg.Type)
	require.Equal(t, account.Description, arg.Description)
	require.Equal(t, account.Value, arg.Value)

	accountDate := account.Date.Format("02/01/2006 03:04:05")
	argDate := arg.Date.Format("02/01/2006 03:04:05")
	require.Equal(t, accountDate, argDate)
	require.NotEmpty(t, account.CreatedAt)

	return account
}

func Test_CreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func Test_DeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)

	account3, err := testQueries.GetDeletedAccount(context.Background(), account1.ID)
	require.NotEmpty(t, account3.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, account1.ID, account3.ID)
}

func Test_GetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.Title, account2.Title)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Description, account2.Description)
	require.Equal(t, account1.Value, account2.Value)
	accountDate := account1.Date.Format("02/01/2006 03:04:05")
	argDate := account2.Date.Format("02/01/2006 03:04:05")

	require.Equal(t, accountDate, argDate)
	require.NotEmpty(t, account2.CreatedAt)
}

func Test_GetAccounts(t *testing.T) {
	account1 := CreateRandomAccount(t)
	arg := GetAccountsParams{
		UserID:      account1.UserID,
		Title:       account1.Title,
		Type:        account1.Type,
		Description: account1.Description,
		CategoryID:  account1.CategoryID,
		Date:        account1.Date,
	}
	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)

	category, err := testQueries.GetCategoryById(context.Background(), account1.CategoryID)
	require.NoError(t, err)

	for _, account := range accounts {
		require.Equal(t, account1.UserID, account.UserID)
		require.Equal(t, account1.Title, account.Title)
		require.Equal(t, account1.Type, account.Type)
		require.Equal(t, account1.Description, account.Description)
		require.Equal(t, account1.Value, account.Value)
		require.Equal(t, account.CategoryTitle.String, category.Title)

		account1Date := account1.Date.Format("02/01/2006 03:04:05")
		accountDate := account.Date.Format("02/01/2006 03:04:05")
		require.Equal(t, account1Date, accountDate)

		require.NotEmpty(t, account.CreatedAt)
	}
}

func Test_GetDeletedAccounts(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetDeletedAccount(context.Background(), account1.ID)
	require.NotEmpty(t, account2.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, account1.ID, account2.ID)

	accounts, err := testQueries.GetDeletedAccounts(context.Background())
	require.NoError(t, err)
	var match bool = false
	for _, account := range accounts {
		require.NotEmpty(t, account.DeletedAt)
		if account1.ID == account.ID {
			match = true
			require.Equal(t, account2.ID, account.ID)
			require.Equal(t, account2.UserID, account.UserID)
			require.Equal(t, account2.CategoryID, account.CategoryID)
			require.Equal(t, account2.Title, account.Title)
			require.Equal(t, account2.Type, account.Type)
			require.Equal(t, account2.Description, account.Description)
			require.Equal(t, account2.Value, account.Value)
			require.Equal(t, account2.Date, account.Date)
			require.Equal(t, account2.DeletedAt, account.DeletedAt)
		}
	}
	require.Equal(t, true, match)
}

func Test_GetDeletedAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetDeletedAccount(context.Background(), account1.ID)
	require.NotEmpty(t, account2.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, account1.ID, account2.ID)
}

func Test_UpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	arg := UpdateAccountParams{
		ID:          account1.ID,
		Title:       util.RandomName(10),
		Description: util.RandomName(50),
		Value:       util.RandomFloatString(2, 2),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, account1.ID, account2.ID, account1, account2)
	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, arg.Title, account2.Title)
	require.Equal(t, arg.Description, account2.Description)
	require.Equal(t, arg.Value, account2.Value)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)

	require.NotEqual(t, account1.UpdatedAt, account2.UpdatedAt)
	require.NotEqual(t, account1.Title, account2.Title)
	require.NotEqual(t, account1.Description, account2.Description)
	require.NotEqual(t, account1.Value, account2.Value)

	require.NotEmpty(t, account2)
	require.NotEmpty(t, account2.UpdatedAt)
}

func Test_GetAccountsReports(t *testing.T) {
	account1 := CreateRandomAccount(t)
	arg := GetAccountsReportsParams{
		UserID: account1.UserID,
		Type:   account1.Type,
	}
	//para realizar este teste foram alterados os tipos de retorno em GetAccountsReports para []uint8
	sumValues, err := testQueries.GetAccountsReports(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sumValues)

}

func Test_GetAccountsGraph(t *testing.T) {
	account1 := CreateRandomAccount(t)
	arg := GetAccountsGraphParams{
		UserID: account1.UserID,
		Type:   account1.Type,
	}
	sumAccounts, err := testQueries.GetAccountsGraph(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sumAccounts)

}
