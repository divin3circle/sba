package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/divin3circle/sba/util"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	_ = createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	testAccount := createTestAccount(t)
	testAccount2, err := testQueries.GetAccount(context.Background(), testAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, testAccount2)

	require.Equal(t, testAccount.ID, testAccount2.ID)
	require.Equal(t, testAccount.Owner, testAccount2.Owner)
	require.Equal(t, testAccount.Balance, testAccount2.Balance)
	require.Equal(t, testAccount.Currency, testAccount2.Currency)
}

func TestUpdateAccount(t *testing.T) {
	testAccount := createTestAccount(t)

	arg := AddToAccountParams{
		ID:     testAccount.ID,
		Amount: util.RandomBalance(),
	}

	updatedAccount, err := testQueries.AddToAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, updatedAccount.Balance, arg.Amount+testAccount.Balance)
}

func TestDeleteAccount(t *testing.T) {
	testAccount := createTestAccount(t)

	err := testQueries.DeleteAccount(context.Background(), testAccount.ID)

	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), testAccount.ID)
	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}
