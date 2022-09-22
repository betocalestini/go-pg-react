package db

import (
	"context"
	"go-pg-react/util"
	"testing"

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

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
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
