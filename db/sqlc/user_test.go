package db

import (
	"context"
	"testing"

	"github.com/betocalestini/go-pg-react/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:     util.RandomName(20),
		Email:    util.RandomEmail(),
		Password: util.RandomSenha(15),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Name, arg.Name)
	require.Equal(t, user.Password, arg.Password)
	require.Equal(t, user.Email, arg.Email)
	require.NotEmpty(t, user.CreatedAt)

	return user
}

func Test_CreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func Test_DeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetDeletedUser(context.Background(), user1.ID)
	require.NotEmpty(t, user2.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, user1.ID, user2.ID)

	user3, err := testQueries.GetUser(context.Background(), user1.Email)
	require.Error(t, err)
	require.Empty(t, user3)
}

func Test_GetDeletedUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetDeletedUser(context.Background(), user1.ID)
	require.NotEmpty(t, user2.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, user1.ID, user2.ID)
}

func Test_GetDeletedUsers(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetDeletedUser(context.Background(), user1.ID)
	require.NotEmpty(t, user2.DeletedAt)
	require.NoError(t, err)
	require.Equal(t, user1.ID, user2.ID)

	users, err := testQueries.GetDeletedUsers(context.Background())
	require.NoError(t, err)
	var match bool = false
	for _, user := range users {
		require.NotEmpty(t, user.DeletedAt)
		if user1.ID == user.ID {
			match = true
			require.Equal(t, user2.ID, user.ID)
			require.Equal(t, user2.Name, user.Name)
			require.Equal(t, user2.Email, user.Email)
			require.Equal(t, user2.DeletedAt, user.DeletedAt)
		}
	}
	require.Equal(t, true, match)
}

func Test_GetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}

func Test_GetUserById(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}

func Test_GetUsers(t *testing.T) {
	user1 := CreateRandomUser(t)

	users, err := testQueries.GetUsers(context.Background())
	require.NoError(t, err)
	var match bool = false
	for _, user := range users {
		require.Empty(t, user.DeletedAt)
		if user1.ID == user.ID {
			match = true
			require.Equal(t, user1.ID, user.ID)
			require.Equal(t, user1.Name, user.Name)
			require.Equal(t, user1.Email, user.Email)
			require.Equal(t, user1.DeletedAt, user.DeletedAt)
		}
	}
	require.Equal(t, true, match)
}
