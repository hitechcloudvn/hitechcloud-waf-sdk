# Contributing to HiTechCloud WAF SDK

Thank you for your interest in contributing!

## Development

1. Clone the repository.
2. Pick a language SDK under `sdk/`.
3. Install dependencies and run the checks:
   - Python: `python -m py_compile hitechcloud_waf/**/*.py`
   - Node.js: `npm run lint` (tsc --noEmit)
   - PHP: `php -l src/**/*.php`
   - Go: `gofmt -l .` / `go build ./...`
   - Java: `mvn compile`

## Guidelines

- Every method must map to a real WAF endpoint.
- Keep resource grouping and naming consistent across languages.
- Update `openapi.json` and the docs when endpoints change.
- Document public methods.

## Pull Requests

1. Fork and create a feature branch.
2. Add or update examples/tests.
3. Ensure all per-language checks pass.
4. Submit a PR.

By contributing you agree your work is licensed under the MIT License.
