# Frescatto CLI

Go CLI for ordering fish and seafood from Frescatto (`www.frescatto.com`) in Rio de Janeiro.

All logic lives in the shared library **`github.com/voska/vtexkit`**, which also powers the
Zona Sul CLI. This repo holds only the store descriptor.

## Project Structure

```
store.go             # The Frescatto descriptor — 3 fields
cmd/frescatto/       # Entry point, ~15 lines
skills/frescatto/    # Claude Code agent skill
docs/                # Verified API notes, design and plan docs
```

Everything else — VTEX client, auth strategies, search, cart, checkout, output modes,
exit codes — is in vtexkit. Change behavior there, not here.

## Platform

- VTEX account `frescatto`, store `https://www.frescatto.com`
- Auth: classic password and emailed access code, both enabled. Discovered at runtime.
- Search: Intelligent Search REST with catalog REST fallback. No persisted GraphQL hash.
- Seller, payment systems, and delivery SLAs are all discovered, never hardcoded.
- No minimum order. No ClearSale quirk.
- Secrets: macOS Keychain service `frescatto-cli`. Config: `~/.config/frescatto/`.

## Output Modes

`--json` `--plain` `--quiet` `--results-only` `--select f1,f2` — data to stdout,
hints and errors to stderr, always. `NO_COLOR` respected.

## Exit Codes

0=success 1=error 2=usage 3=empty 4=auth 5=not_found 6=forbidden 7=rate_limited
8=retryable 9=domain_error 10=config. Run `frescatto exit-codes --json`.

## Development

```sh
make build   # bin/frescatto
make test
make lint
make ci
```

vtexkit is resolved through a `replace` directive to `../vtexkit` until it is
published. A `go.work` at `~/Developer/` spans both modules for local work.

## Important

- **Never run `checkout --confirm` during development or verification.** Never call the
  transaction endpoint manually.
- Search, cart, delivery simulation, payment discovery, and checkout preview must all
  keep working without authentication.
- Never hardcode credentials or tokens.
- The authenticated legs (login, saved cards, place order, order history) are built to
  the vanilla VTEX standard and unit-tested, but have not been verified against a real
  Frescatto account. See `docs/frescatto-api.md`.
