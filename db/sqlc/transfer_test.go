package db

import (
	"context"
	"testing"
	"time"

	"github.com/danyouknowme/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAccount Account, toAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	createRandomTransfer(t, fromAccount, toAccount)
}

func TestGetTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	randomTransfer := createRandomTransfer(t, fromAccount, toAccount)

	transfer, err := testQueries.GetTransfer(context.Background(), randomTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, randomTransfer.ID, transfer.ID)
	require.Equal(t, randomTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, randomTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, randomTransfer.Amount, transfer.Amount)
	require.WithinDuration(t, randomTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}
