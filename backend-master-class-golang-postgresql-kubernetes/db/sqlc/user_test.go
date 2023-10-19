package db

import (
	"context"
	"simplebank/test"
	"simplebank/util"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	// Arrange
	hashedPassword, err := util.HashPassword(test.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       test.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          test.RandomEmail(),
		FullName:       test.RandomOwner(),
	}
	// Act
	user, err := testStore.CreateUser(context.Background(), arg)
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Zero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	// Arrange
	user1 := createRandomUser(t)
	// Act
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserOnlyFullName(t *testing.T) {
	// Arrange
	oldUser := createRandomUser(t)
	newFullName := test.RandomOwner()
	// Act
	updateUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		FullName: pgtype.Text{
			String: newFullName,
			Valid:  true,
		},
	})
	// Assert
	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, updateUser.FullName)
	require.Equal(t, newFullName, updateUser.FullName)
	require.Equal(t, oldUser.Email, updateUser.Email)
	require.Equal(t, oldUser.HashedPassword, updateUser.HashedPassword)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	// Arrange
	oldUser := createRandomUser(t)
	newEmail := test.RandomEmail()
	// Act
	updateUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	})
	// Assert
	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updateUser.Email)
	require.Equal(t, newEmail, updateUser.Email)
	require.Equal(t, oldUser.FullName, updateUser.FullName)
	require.Equal(t, oldUser.HashedPassword, updateUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	// Arrange
	oldUser := createRandomUser(t)
	newPassword := test.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)
	// Act
	updateUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		HashedPassword: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
	})
	// Assert
	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updateUser.HashedPassword)
	require.Equal(t, newHashedPassword, updateUser.HashedPassword)
	require.Equal(t, oldUser.FullName, updateUser.FullName)
	require.Equal(t, oldUser.Email, updateUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	// Arrange
	oldUser := createRandomUser(t)
	newFullName := test.RandomOwner()
	newEmail := test.RandomEmail()
	newPassword := test.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)
	// Act
	updateUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: oldUser.Username,
		FullName: pgtype.Text{
			String: newFullName,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
		HashedPassword: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
	})
	// Assert
	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, updateUser.FullName)
	require.Equal(t, newFullName, updateUser.FullName)
	require.NotEqual(t, oldUser.Email, updateUser.Email)
	require.Equal(t, newEmail, updateUser.Email)
	require.NotEqual(t, oldUser.HashedPassword, updateUser.HashedPassword)
	require.Equal(t, newHashedPassword, updateUser.HashedPassword)
}
