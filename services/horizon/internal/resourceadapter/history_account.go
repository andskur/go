package resourceadapter

import (
	"context"

	. "github.com/andskur/go/protocols/horizon"
	"github.com/andskur/go/services/horizon/internal/db2/history"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
