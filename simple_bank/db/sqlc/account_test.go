package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/useranonymous/simplebank/util"
)

func TestAccount(t *testing.T) {

	account := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	acc, err := testQueries.CreateAccount(context.Background(), account)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	// check of the value we ccreate is as expected or not
	require.Equal(t, acc.Owner, account.Owner)
	require.Equal(t, acc.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

}

func TestGetAccount(t *testing.T) {
	account, err := testQueries.GetAccount(context.Background(), 12)

	if err != nil {
		log.Fatal("err getting account detail: ", err)
	}

	t.Log(account)

}

func TestListAccounts(t *testing.T) {

	listAccParam := ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), listAccParam)

	if err != nil {
		t.Fail()
	}
	t.Log(accounts)
}

func TestUpdateAccount(t *testing.T) {
	updateAccParam := UpdateAcountParams{
		ID:      17,
		Balance: 1000,
	}

	account, err := testQueries.UpdateAcount(context.Background(), updateAccParam)

	if err != nil {
		t.Fail()
	}
	t.Log(account)
	// simulating error
	// t.Fail()
}

func TestDeleteAccount(t *testing.T) {
	if err := testQueries.DeleteAccount(context.Background(), 18); err != nil {
		t.Fail()
	}

}
