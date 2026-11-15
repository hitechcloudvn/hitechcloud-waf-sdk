# Changelog

> **English** | [Tiếng Việt](CHANGELOG.vi.md)

---

All notable changes to the HiTechCloud WAF SDK are documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [Semantic Versioning](https://semver.org/).

## [1.0.1] - 2026-11-16

### Fixed
- Java SDK: `WafClient` now refers to `vn.hitechcloud.waf.resources.System`
  explicitly - the bare `System` name was ambiguous with `java.lang.System`
  and broke the Maven build.
- Node.js SDK: add the missing `npm test` script and a smoke test so the
  test job runs instead of failing with "Missing script: test".
- PHP SDK: add `phpunit.xml.dist` and smoke tests so `vendor/bin/phpunit`
  runs the suite instead of printing help and failing the test job.

## [1.0.0] - 2026-11-01

### Added
- Initial release with 5 language SDKs (Node.js/TypeScript, Python, PHP, Go, Java).
- 227 endpoints across 25 resource groups.
- Automated CSRF -> login -> JWT authentication flow with optional TFA.
- OpenAPI 3.0 specification.
- Vietnamese documentation.
- MCP server exposing core WAF actions.
- Examples for every supported language.
