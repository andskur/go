package resourceadapter

import (
	. "github.com/andskur/go/protocols/horizon"
	"github.com/andskur/go/services/horizon/internal/db2/core"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
}
