package resourceadapter

import (
	. "github.com/andskur/go/protocols/horizon"
	"github.com/andskur/go/services/horizon/internal/db2/core"
)

func PopulateAccountThresholds(dest *AccountThresholds, row core.Account) {
	dest.LowThreshold = row.Thresholds[1]
	dest.MedThreshold = row.Thresholds[2]
	dest.HighThreshold = row.Thresholds[3]
}
