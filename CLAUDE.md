# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go build          # Build the binary (outputs ./uuid)
go test ./...     # Run all tests
go test -v        # Run tests with verbose output
go test -run TestUUID  # Run a single test by name
go fmt ./...      # Format code
go vet ./...      # Static analysis
```

## Architecture

Single-package Go CLI tool (`package main`) — all logic lives in `main.go` with tests in `main_test.go`.

**Entry point:** `main()` parses CLI flags and dispatches to one generator function. Only one ID type may be generated per invocation (enforced via flag validation).

**Generator functions:**
- `createUUID()` — UUIDs v1, 2p, 2g, 4, 6, 7
- `createCuid()` — CUIDs with optional custom length (2–32)
- `createNanoid()` — NanoIDs with configurable length (default 21)
- `createUlid()` — ULIDs in fast or cryptographic mode
- `createXid()` — XIDs (always 20 chars)
- `createObjectID()` — MongoDB ObjectIDs (24 hex chars)

**Smart default behavior:** The binary detects its own executable name (`uuid`, `cuid`, `nanoid`, `ulid`, `xid`, `oid`) to set the default ID type — useful when installed via symlinks.

**Clipboard:** The last generated ID is automatically copied to the clipboard (disable with `-clip=false`). Uses `github.com/atotto/clipboard`.

**Demo mode:** `-demo` generates one of each ID type for comparison.

## Key Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/google/uuid` | UUID generation |
| `github.com/nrednav/cuid2` | CUID generation |
| `github.com/jaevor/go-nanoid` | NanoID generation |
| `github.com/oklog/ulid/v2` | ULID generation |
| `github.com/rs/xid` | XID generation |
| `go.mongodb.org/mongo-driver/v2` | MongoDB ObjectID |
| `github.com/atotto/clipboard` | Clipboard access |