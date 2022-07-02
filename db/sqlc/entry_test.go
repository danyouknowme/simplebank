package db

import (
	"context"
	"testing"
	"time"

	"github.com/danyouknowme/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	randomEntry := createRandomEntry(t, createRandomAccount(t))
	entry, err := testQueries.GetEntry(context.Background(), randomEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, randomEntry.ID, entry.ID)
	require.Equal(t, randomEntry.AccountID, entry.AccountID)
	require.Equal(t, randomEntry.Amount, entry.Amount)
	require.WithinDuration(t, randomEntry.CreatedAt, entry.CreatedAt, time.Second)
}

func TestGetListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t, createRandomAccount(t))
	}

	arg := GetListEntriesParams{
		Limit:  8,
		Offset: 8,
	}

	entries, err := testQueries.GetListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 8)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
