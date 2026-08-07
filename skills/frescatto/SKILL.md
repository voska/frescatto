---
name: frescatto
description: >-
  Order fish and seafood from Frescatto (www.frescatto.com) using the `frescatto` CLI.
  Use for searching products, managing a cart, checking delivery, and placing orders.
allowed-tools: Bash, Read
---

# frescatto

Order from Frescatto using the `frescatto` CLI.

## Always start here

```bash
frescatto doctor
```

Exit 0 means ordering will work. Any other exit code means it will not. Each
failed line prints the exact fix. Do that fix, or report it to the user. Do not
retry the command that failed.

## The flow

```bash
frescatto search salmao --limit 5     # 1. find a SKU
frescatto cart add 42 --qty 2         # 2. add it
frescatto delivery windows            # 3. pick a window number
frescatto checkout --window 0         # 4. preview — places nothing
frescatto checkout --window 0 --confirm   # 5. order
```

Between steps 4 and 5: **show the preview to the user and get explicit
approval.** `--confirm` spends real money.

## Details

**Search** takes Portuguese terms. The first column of the output is the SKU.

**Cart** never needs a `--seller`; it is looked up automatically.

```bash
frescatto cart show
frescatto cart update 0 --qty 3    # index from 'cart show'
frescatto cart remove 0
frescatto cart clear
```

**Payment** defaults to pix. `frescatto checkout payments` lists what the store
accepts and any card saved on the account. For a card, add `--cvv 123`.

**Favorites** are the store's own wishlist — the same list the heart icons on
the website use. Changes here show up on the website.

```bash
frescatto fav                   # show the wishlist
frescatto fav add 67            # save a product
frescatto fav remove 67         # unsave it
frescatto cart add 42           # buy one of them
```

Never add every favorite to the cart. A wishlist is things the user likes,
not things they want to buy today. There is no bulk-order command for it.

**Lists** are curated SKU groups stored locally, and these *are* meant to be
ordered wholesale:

```bash
frescatto list add weekly 42
frescatto list order weekly     # add every SKU in that list to the cart
```

## Output for scripts and agents

```bash
frescatto search salmao --json
frescatto search salmao --json --select sku,name,price
frescatto search salmao --plain          # tab-separated
frescatto search salmao --quiet --select sku   # bare values
```

Prices in `--json` are integer centavos: `8290` is R$82,90.

Data goes to stdout; progress and errors go to stderr. Every mutating command
accepts `--dry-run`.

## Exit codes

| Code | Meaning | What to do |
|---|---|---|
| 0 | success | continue |
| 2 | bad arguments | fix the command; do not retry it unchanged |
| 3 | empty result | tell the user nothing matched |
| 4 | login required | `frescatto auth login --email <email>` |
| 5 | not found | the SKU is wrong; search again |
| 7, 8 | temporary | wait, then retry once |
| 9 | store rule refused | read the message; it names the rule |
| 10 | not set up | `frescatto doctor` and follow the fixes |

Full table: `frescatto exit-codes --json`

## Rules

- Never run `--confirm` without the user approving that exact cart and total.
- If a command fails twice the same way, stop and report it. Do not loop.
- `frescatto doctor` diagnoses anything unexpected; its output names the fix.

## First-time accounts

A new Frescatto account cannot order through the API until it has a delivery
address and one completed order — VTEX does not expose a saved card to the
checkout API before that. The first order must be placed on the website.
`doctor` reports exactly which of these is missing.
