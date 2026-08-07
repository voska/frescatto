// Package frescatto describes the Frescatto storefront (www.frescatto.com),
// a fish and seafood retailer in Rio de Janeiro.
package frescatto

import "github.com/voska/vtexkit/store"

// Store is the Frescatto descriptor.
//
// Everything not listed here is discovered from the store's own API at
// runtime: the login method (classic password and emailed access code are
// both enabled), the payment systems, the seller for each SKU, and the
// delivery SLAs.
//
// No MinOrder: storePreferencesData.minimumOrderValue is null and the site
// advertises free shipping above R$400, which is not a minimum order.
//
// No Quirks: there is no evidence Frescatto's gateway requires ClearSale
// fingerprinting, unlike Zona Sul's.
//
// Search is left at the default SearchAuto, which uses Intelligent Search
// REST and falls back to the catalog API. Both were verified working
// against this store on 2026-08-07.
var Store = store.Store{
	Name:        "frescatto",
	DisplayName: "Frescatto",
	BaseURL:     "https://www.frescatto.com",
}
