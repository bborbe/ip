# Changelog

## v2.0.2

- Bump indirect dependencies (docker, containerd, moby, opentelemetry, golang.org/x/*)
- Add replace directives for charmbracelet/x/cellbuf, denis-tingaikin/go-header, nunnatsa/ginkgolinter/types, opencontainers/runtime-spec
- Update go-git, ssh_config, klauspost/compress, and various other indirect deps

## v2.0.1

- chore: verify project health — all tests pass, linting clean, precommit succeeds

## v2.0.0

### Breaking Changes

- Removed `ip-client` command
- Changed to modern service patterns

### Added

- Prometheus metrics (`ip_requests_total`)
- Health check endpoint (`/healthz`)
- Readiness endpoint (`/readiness`)
- Graceful shutdown support
- GitHub Actions CI
- OSV vulnerability scanning
- golangci-lint

### Changed

- Migrated from `dep` to Go modules
- Updated to Go 1.26
- Migrated from `flagenv` to `bborbe/argument` (via `bborbe/service`)
- Migrated from `gracehttp` to `bborbe/run`
- Modernized project structure (go-skeleton style)
- Simplified to single Dockerfile
- Updated all dependencies

### Removed

- `ip-client` command (unused)
- Vendored dependencies
- Legacy `Gopkg.toml`/`Gopkg.lock`

## v1.1.0

- Initial tagged release
