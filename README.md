# frescatto

CLI for ordering fish and seafood from [Frescatto](https://www.frescatto.com) in Rio de
Janeiro. Built for humans and AI agents: data to stdout, hints to stderr, stable exit
codes, structured output.

Built on [vtexkit](https://github.com/voska/vtexkit), a shared library for Brazilian VTEX
storefronts.

## Install

```sh
make build   # -> bin/frescatto
```

## Use

Start every session with `doctor`. It answers "can I order right now?" and, if not,
prints the exact fix for each problem.

```sh
frescatto doctor
```

```
ok    store reachable      https://www.frescatto.com
ok    catalog search       Kit Peixe Fresco - Salmão, Tilápia e Camarão
ok    logged in            you@example.com
ok    delivery address     2 on file
ok    saved cards          Visa ************1234
ok    order history        3 orders

Ready to order. Next: frescatto cart add <sku>, then frescatto checkout
```

Then the whole flow is five commands:

```sh
frescatto fav                            # your wishlist from the website
frescatto fav add 67                     # save a product to it
frescatto search salmao --limit 5        # find a SKU
frescatto cart add 42 --qty 2            # seller is looked up for you
frescatto delivery windows               # pick a window number
frescatto checkout --window 0            # preview — places nothing
frescatto checkout --window 0 --confirm  # actually orders
```

`checkout` never places an order without `--confirm`. Every mutating command takes
`--dry-run`.

### First-time setup

A brand-new Frescatto account cannot order from the API until it has a profile, a
delivery address, and one completed order — VTEX will not expose a saved card to the
checkout API before that. **Place your first order on the website.** Everything after
that works from the CLI. `doctor` tells you exactly where you are.

## Output

| Flag | Behavior |
|---|---|
| default | Aligned, colored table |
| `--json` | Structured JSON to stdout |
| `--plain` | Tab-separated, stable column order |
| `--quiet` | Bare values, one per line |
| `--select sku,name` | Field projection, dot paths supported |
| `--results-only` | Strip the metadata envelope |

Env equivalents: `FRESCATTO_JSON`, `FRESCATTO_PLAIN`, `FRESCATTO_NO_INPUT`, and so on.

## Exit codes

`frescatto exit-codes --json`. 3 means empty, 4 means login required, 7 and 8 are
transient and safe to retry, 9 means a store rule refused the operation.

## Status

The unauthenticated path — search, cart, delivery simulation, checkout preview — is
verified against the live store. The authenticated legs are implemented against the
vanilla VTEX standard and unit-tested, but have not yet run against a real Frescatto
account.
