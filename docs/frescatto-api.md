# Frescatto VTEX API

Verified against `https://www.frescatto.com` on 2026-08-07.

## Store identity

- Account: `frescatto`
- Seller: `1` (`FRIGORIFICO JAHU EIRELI`)
- Auth cookie: `VtexIdclientAutCookie_frescatto`

## Catalog search

```http
GET /api/catalog_system/pub/products/search/?ft=salmao&_from=0&_to=2
```

HTTP 200 and 206 are successful. The response is an array of products containing `productId`, `productName`, and `items`. Each item contains `itemId`, `name`, `measurementUnit`, `unitMultiplier`, and `sellers`. The first seller's `commertialOffer` contains decimal-real `Price` and `ListPrice`, integer `AvailableQuantity`, and boolean `IsAvailable`.

## Authentication

Start both validation flows with:

```http
GET /api/vtexid/pub/authentication/start?scope=frescatto&callbackUrl=https%3A%2F%2Fwww.frescatto.com%2Fapi%2Fvtexid%2Foauth%2Ffinish&user=&locale=pt-BR&accountName=frescatto
```

The response contains `authenticationToken`. Submit form-encoded credentials to one of:

```http
POST /api/vtexid/pub/authentication/classic/validate
POST /api/vtexid/pub/authentication/accesskey/send
POST /api/vtexid/pub/authentication/accesskey/validate
```

Successful validation returns `authStatus: "Success"`, `authCookie.Name`, `authCookie.Value`, and `expiresIn`. Check a stored token with `GET /api/vtexid/pub/authenticated/user` and session state with `GET /api/sessions?items=cookie.VtexIdclientAutCookie_frescatto,checkout.orderFormId,authentication.storeUserEmail`.

## Cart and checkout

```http
GET  /api/checkout/pub/orderForm
GET  /api/checkout/pub/orderForm/{orderFormId}
POST /api/checkout/pub/orderForm/{orderFormId}/items
POST /api/checkout/pub/orderForm/{orderFormId}/items/update
POST /api/checkout/pub/orderForm/{orderFormId}/items/removeAll
```

Order-form item prices, selling prices, and totalizer values are integer centavos.

## Delivery simulation

```http
POST /api/checkout/pub/orderForms/simulation
```

Request:

```json
{"items":[{"id":"62","quantity":1,"seller":"1"}],"postalCode":"01310-100","country":"BRA"}
```

The response contains centavo-priced `items`, `logisticsInfo[].slas[]` with delivery price, estimate, and `availableDeliveryWindows`, plus `paymentData.paymentSystems[]` with `id`, `name`, and `groupName`.

## Orders and payment gateway

Authenticated order history uses `GET /api/oms/user/orders` and `GET /api/oms/user/orders/{orderId}`. Saved-card payment submission uses `https://frescatto.vtexpayments.com.br`; the checkout callback is on `https://www.frescatto.com/checkout/gatewayCallback/{orderGroup}/{messageCode}`. These authenticated paths are not yet live-verified.

## Verification record — 2026-08-07

First run of the authenticated path against a real Frescatto account
(you@example.com). Everything below was executed live.

**Verified working:**

| Surface | Result |
|---|---|
| `auth login` (classic password) | Authenticated; `refreshable: true` |
| `auth status` | Reports account and refreshability |
| `orders` | Empty history, exit 3 — no orders on this account yet |
| `cart add` / `cart show` / `cart clear` | Added SKU 42 at R$82,90, then cleared |
| `checkout payments` | 8 accepted methods; `savedCards: []` |
| `delivery simulate` (unauthenticated) | 66 free windows to CEP 01310100 |
| `search`, `product` | Live catalog, correct centavo prices |

**Account state found:** no saved cards, no order history. Both are genuine
empties read from an authenticated order form, not fallbacks.

**NOT verified:** `PlaceOrder`, `PayWithSavedCard`, `GatewayCallback`. These
require placing a real paid order. Per AGENTS.md, `checkout --confirm` is
never run during development or verification. This boundary is deliberate —
the transaction leg stays unproven until a real order is placed manually.
Card payment additionally cannot be exercised until a card is saved on the
account.

**Bug found by this run:** the OMS order list serializes `totalValue` as a
decimal (`26511.0`) where the checkout orderForm sends an integer (`15372`).
Both mean centavos. Fixed in vtexkit's `money.Centavos.UnmarshalJSON`.
