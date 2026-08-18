package frescatto_test

import (
	"testing"

	"github.com/voska/frescatto"
	"github.com/voska/vtexkit/store"
)

// Frescatto order 1654880526893 was placed on 2026-08-18 with a saved card
// and printed as placed. The gateway never settled it — no tid, no
// authorizedDate — and cancelled it five minutes later. The descriptor
// carried no quirks, so checkout skipped the gatewayCallback poll that
// finalizes a card payment; the replacement order, placed with the poll
// enabled, settled.
func TestStoreNeedsTheGatewayCallback(t *testing.T) {
	if !frescatto.Store.Quirks.Has(store.GatewayCallback) {
		t.Error("a card payment here does not settle without the gatewayCallback poll")
	}
}

// ClearSale fingerprinting is Zona Sul's problem, not this store's: the
// 08-07 verification run found no sign the gateway asks for one, and
// sending one is not free.
func TestStoreDoesNotFingerprintWithClearSale(t *testing.T) {
	if frescatto.Store.Quirks.Has(store.ClearSaleFingerprint) {
		t.Error("no evidence this store's gateway requires a ClearSale fingerprint")
	}
}

func TestStoreIdentityIsUnchanged(t *testing.T) {
	// Name drives the binary, config dir, keyring service, and env prefix.
	if frescatto.Store.Name != "frescatto" || frescatto.Store.AccountName() != "frescatto" {
		t.Errorf("store = %+v", frescatto.Store)
	}
}
