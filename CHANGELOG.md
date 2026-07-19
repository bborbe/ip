# Changelog

## v2.0.4

- fix(deps): bump x/text v0.39.0 (CVE-2026-56852) + Go 1.26.5 (GO-2026-5856); suppress unreachable/unfixable transitive CVEs (containerd, x/crypto/openpgp)

## v2.0.3

- Bump Go base image to 1.26.2 and Alpine to 3.23
- Update bborbe/* dependencies (errors, http, run, sentry, service, argument, collection, etc.)
- Update golangci-lint to v2.11.4 and simplify linter config
- Update osv-scanner, gosec, and other tooling dependencies
- Add new CVE ignores for bbolt and aws-sdk-go-v2 vulnerabilities

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
