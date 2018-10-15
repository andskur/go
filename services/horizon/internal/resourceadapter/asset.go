package resourceadapter

import (
	"context"

	. "github.com/andskur/go/protocols/horizon"
	"github.com/andskur/go/xdr"
)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
