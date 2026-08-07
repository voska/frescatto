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

	// Persisted-query hashes for vtex.wish-list, which backs the heart
	// icons and /account/#/wishlist. There is no hash-free way to reach
	// that API. These are tied to the installed app version (1.19.1 as of
	// 2026-08-07); if `frescatto fav` starts reporting a rejected request,
	// re-capture them from a live browser session.
	Wishlist: store.WishlistHashes{
		View:   "16643ee1547d54d81f6c73462be5d18b16297b00c2c100a94f62bc83fe3e1aea",
		Add:    "0fd032b3e26dc0223a8dbfcf8629e27aa9deaa2fe064dcc8eef36a8ac70af3ee",
		Remove: "7690d9e181e5eeafc21c5d5e5cde3f2322bac1d23ad9ab9321f33bfe071a8705",
	},
}
