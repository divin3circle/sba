package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTxn(t *testing.T) {
	testStore := NewStore(testDB)

	testAccount := createTestAccount(t)
	testAccount2 := createTestAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxnResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := testStore.TransferTxn(context.Background(), TransferTxnParams{
				FromAccountID: testAccount.ID,
				ToAccountID:   testAccount2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)

	for j := 0; j < n; j++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfers
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		_, err = testStore.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)
		require.Equal(t, -amount, fromEntry.Amount)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
		require.Equal(t, amount, toEntry.Amount)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, testAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, testAccount2.ID)

		fromAccountBalanceDifference := testAccount.Balance - fromAccount.Balance
		toAccountBalanceDifference := toAccount.Balance - testAccount2.Balance
		require.Equal(t, fromAccountBalanceDifference, toAccountBalanceDifference)
		require.True(t, fromAccountBalanceDifference > 0 && toAccountBalanceDifference > 0)
		require.True(t, fromAccountBalanceDifference%amount == 0)

		k := int(fromAccountBalanceDifference / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedFromAccount, err := testQueries.GetAccount(context.Background(), testAccount.ID)
	require.NoError(t, err)

	updatedToAccount, err := testQueries.GetAccount(context.Background(), testAccount2.ID)
	require.NoError(t, err)

	require.Equal(t, testAccount.Balance-int64(n)*amount, updatedFromAccount.Balance)
	require.Equal(t, updatedToAccount.Balance, testAccount2.Balance+int64(n)*amount)
}
