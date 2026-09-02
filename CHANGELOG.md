# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v2.1.1

- chore: update github.com/bborbe/errors to v1.6.0, github.com/bborbe/http to v1.26.25, github.com/bborbe/run to v1.10.1, github.com/bborbe/sentry to v1.10.0, github.com/bborbe/service to v1.10.10

## v2.1.0

- feat: opt into `autoMerge.trivial` for mechanically-trivial update PRs

## v2.0.11

- chore: update github.com/bborbe/errors to v1.5.21

## v2.0.10

- chore: update go module dependencies

## v2.0.9

- chore: update Go to 1.27.0 and github.com/bborbe/errors to v1.5.20, github.com/bborbe/http to v1.26.24, github.com/bborbe/run to v1.9.37, github.com/bborbe/sentry to v1.9.27, github.com/bborbe/service to v1.10.9

## v2.0.8

- exclude no-fix docker/containerd advisories in checker config (GO-2026-4883/4887/5064/5338/5622/5932 v1 no-fix)
## v2.0.7

- Migrate to tools.env + Makefile @version pattern; remove tools.go and obsolete replace block. go.mod reduced from 478 to 48 lines (errcheck folded into golangci-lint per go-tools-versioning guide).

## v2.0.6

- chore(security): bump Go toolchain 1.26.5 -> 1.26.6 in `go.mod` and Dockerfile (stdlib GO-2026-5026 / GO-2026-5972 / GO-2026-6090 / GO-2026-6218)
- chore(security): bump `golang.org/x/mod` v0.37.0 -> v0.40.0 (GO-2026-6179 / GO-2026-6180, CVE-2026-56864 / CVE-2026-56865)
- chore: update `golang.org/x/*` — `crypto` v0.53.0 -> v0.55.0, `net` v0.56.0 -> v0.58.0, `sync` v0.21.0 -> v0.22.0, `sys` v0.46.0 -> v0.47.0, `term` v0.44.0 -> v0.45.0, `text` v0.39.0 -> v0.41.0, `tools` v0.47.0 -> v0.49.0, `telemetry`

## v2.0.5

- fix(deps): bump go-git/go-git/v5 to v5.19.2 (GHSA-hc8v-wwc9-vgxm, GHSA-qgq7-7hm3-q39j)
- fix(deps): bump google.golang.org/grpc to v1.82.1 (GHSA-hrxh-6v49-42gf)

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
