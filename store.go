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

	// Persisted-query hash for vtex.wish-list's ViewLists operation, which
	// backs the heart icons and /account/#/wishlist. Tied to the installed
	// app version (1.19.1 as of 2026-08-07); if `frescatto fav` starts
	// reporting a rejected request, re-capture it from a browser session.
	WishlistHash: "16643ee1547d54d81f6c73462be5d18b16297b00c2c100a94f62bc83fe3e1aea",
}
