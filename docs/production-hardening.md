# Production hardening checklist

This repo is intentionally a starter, not a full platform. If you want to push it toward production use, this is a sensible next-step checklist.

| Area | Why it matters | Typical next step |
| --- | --- | --- |
| Request identity | makes logs and traces easier to correlate | add request ID middleware |
| Recovery | prevents panics from taking down a request path silently | add panic recovery middleware |
| Metrics | helps you see traffic, latency, and failure modes | export Prometheus metrics |
| Tracing | makes cross-service debugging practical | add OpenTelemetry tracing |
| Config validation | catches bad deploys early | validate required env vars at startup |
| Dependency health | improves readiness accuracy | check DB / cache / upstream readiness |
| Auth | protects non-public endpoints | add authn / authz middleware |
| Error contracts | improves client integration and debugging | standardize JSON error responses |
| Background work | avoids dropped tasks on shutdown | drain workers with context-aware shutdown |
| Release metadata | improves debugging in production | inject version, commit, and build time |
| Container delivery | simplifies deployment | maintain Dockerfile and release notes |
| CI / policy | protects the starter from regressions | add lint, formatting, and docs checks |

## Suggested order

If you are evolving this into a real internal service, a practical order is:

1. request IDs + recovery
2. config validation
3. metrics and tracing
4. persistence wiring
5. auth
6. CI + release process

## Related reading

- [`../README.md`](../README.md)
- [`./architecture.md`](./architecture.md)
- [`backend-engineer-checklist`](https://github.com/happysnaker/backend-engineer-checklist)
