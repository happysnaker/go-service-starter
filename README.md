# go-service-starter

A minimal but production-minded **Go HTTP service starter** for backend engineers.

This repo is designed to be a clean starting point for small services and internal APIs that need better structure than a one-file demo, without immediately pulling in a large framework.

## What it includes

- standard-library-first HTTP server
- environment-based configuration loading
- structured JSON logging with `log/slog`
- health and readiness endpoints
- version endpoint with build metadata placeholders
- graceful shutdown handling
- simple project layout you can extend

## Why this exists

A lot of Go starter repos are either:

- too tiny to be reused in a real service, or
- so heavy that they feel like adopting a framework.

This project tries to stay in the middle:

- small enough to understand in one sitting
- structured enough to use as a real internal-service base

## Endpoints

- `GET /` — basic service response
- `GET /healthz` — liveness signal
- `GET /readyz` — readiness signal
- `GET /version` — build metadata placeholder

## Project layout

```text
cmd/api/                 service entrypoint
internal/config/         config loading
internal/httpserver/     HTTP server and middleware
internal/buildinfo/      version / commit placeholders
configs/                 example env config
docs/                    architecture notes
```

## Quick start

1. Copy `configs/service.env.example` into your own environment setup.
2. Adjust values such as service name, address, and timeouts.
3. Extend handlers, middleware, and internal modules for your own domain.

## Good next steps

You can build on top of this starter by adding:

- request IDs
- panic recovery middleware
- metrics and tracing
- persistence layer wiring
- authentication / authorization
- background workers
- deployment manifests

## Who this is for

- backend engineers starting a new internal service
- Go learners graduating from toy demos to service structure
- teams that want a lightweight starter instead of a full framework

## Support

If this starter saves you time, consider:

- starring the repo
- sharing it with other backend engineers
- supporting my open-source work via my GitHub profile: [happysnaker](https://github.com/happysnaker#support-my-open-source-work)

## License

MIT
