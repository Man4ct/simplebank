package db

import (
	"context"
	"testing"

	"github.com/Man4ct/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
}
