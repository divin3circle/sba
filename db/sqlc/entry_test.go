package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testUserID  = 1
	testUserID2 = 2
)

func TestCreateEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: testUserID,
		Amount:    10,
	}
	createdEntry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, createdEntry)
}
