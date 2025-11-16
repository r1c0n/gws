# Changelog

## [v1.5.0](https://github.com/r1c0n/gws/tree/v1.5.0) (2025-11-16)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.4.0...v1.5.0)

**New Features:**

- Custom error pages with branded GWS styling (404, 403, 500, 429)
- CORS middleware with configurable origins, methods, headers, and preflight support
- Rate limiting middleware with token bucket algorithm, IP-based limiting, and path exemptions
- Headless mode support for server operation without UI
- Linux build support with platform-specific release packages
- Enhanced build script with `--debug`, `--debug-ssl`, and `--deploy` flags
- Improved port configuration (443 for SSL, 8443 for Linux SSL)
- Platform-specific release packages (Release-Windows.zip and Release-Linux.tar.gz)

**Bug Fixes:**

- Fixed 404 error pages not displaying with custom handler interceptor
- Fixed middleware order to prevent gzip compression of error pages
- Fixed CORS headers not applying to PathPrefix handlers

**Improvements:**

- Better Linux compatibility in build script
- Cleaned up Go build flags
- Enhanced configuration structure with middleware options
- UTF-8 support in config JSON output

## [v1.4.0](https://github.com/r1c0n/gws/tree/v1.4.0) (2024-01-13)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.4.0-beta.2...v1.4.0)

**Merged pull requests:**

- Bump golang.org/x/crypto from 0.10.0 to 0.17.0 in /src [\#25](https://github.com/r1c0n/gws/pull/25) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump github.com/fatih/color from 1.15.0 to 1.16.0 in /src [\#24](https://github.com/r1c0n/gws/pull/24) ([dependabot[bot]](https://github.com/apps/dependabot))
- v1.4.0 [\#26](https://github.com/r1c0n/gws/pull/26) ([r1c0n](https://github.com/r1c0n))
- add formatting workflows [\#21](https://github.com/r1c0n/gws/pull/21) ([r1c0n](https://github.com/r1c0n))
- Cleanup HTML code & update README.md [\#20](https://github.com/r1c0n/gws/pull/20) ([r1c0n](https://github.com/r1c0n))
- Format gammaws.gq source code to use prettier standards [\#19](https://github.com/r1c0n/gws/pull/19) ([r1c0n](https://github.com/r1c0n))
- Create CHANGELOG.md [\#18](https://github.com/r1c0n/gws/pull/18) ([r1c0n](https://github.com/r1c0n))
- Add building documentation in README.md [\#17](https://github.com/r1c0n/gws/pull/17) ([r1c0n](https://github.com/r1c0n))
- Remove emoji from code block in README.md [\#16](https://github.com/r1c0n/gws/pull/16) ([r1c0n](https://github.com/r1c0n))
- Add emojis to README.md headers & add star message at the end [\#15](https://github.com/r1c0n/gws/pull/15) ([r1c0n](https://github.com/r1c0n))
- Format README.md & SECURITY.md [\#14](https://github.com/r1c0n/gws/pull/14) ([r1c0n](https://github.com/r1c0n))
- Format code to prettier standards [\#13](https://github.com/r1c0n/gws/pull/13) ([r1c0n](https://github.com/r1c0n))

## [v1.4.0-beta.2](https://github.com/r1c0n/gws/tree/v1.4.0-beta.2) (2023-11-07)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.4.0-beta.1...v1.4.0-beta.2)

**Merged pull requests:**

- Bump github.com/gorilla/mux from 1.8.0 to 1.8.1 in /src [\#23](https://github.com/r1c0n/gws/pull/23) ([r1c0n](https://github.com/r1c0n))
- Bump github.com/gorilla/mux from 1.8.0 to 1.8.1 in /src [\#22](https://github.com/r1c0n/gws/pull/22) ([dependabot[bot]](https://github.com/apps/dependabot))

## [v1.4.0-beta.1](https://github.com/r1c0n/gws/tree/v1.4.0-beta.1) (2023-10-20)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.3.0...v1.4.0-beta.1)

**Merged pull requests:**

- Change thank you message in README.md @ \#contributors [\#12](https://github.com/r1c0n/gws/pull/12) ([r1c0n](https://github.com/r1c0n))
- add codacy badge [\#11](https://github.com/r1c0n/gws/pull/11) ([r1c0n](https://github.com/r1c0n))
- docs: add porokimun as a contributor for design [\#9](https://github.com/r1c0n/gws/pull/9) ([allcontributors[bot]](https://github.com/apps/allcontributors))
- docs: add r1c0n as a contributor for code, doc, and 4 more [\#7](https://github.com/r1c0n/gws/pull/7) ([allcontributors[bot]](https://github.com/apps/allcontributors))

## [v1.3.0](https://github.com/r1c0n/gws/tree/v1.3.0) (2023-06-30)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.3.0-rc2...v1.3.0)

**Merged pull requests:**

- redesigned branding [\#6](https://github.com/r1c0n/gws/pull/6) ([bonito-curt](https://github.com/bonito-curt))
- v1.3.0 [\#5](https://github.com/r1c0n/gws/pull/5) ([r1c0n](https://github.com/r1c0n))
- improve build process, restructure project, new favicon [\#4](https://github.com/r1c0n/gws/pull/4) ([r1c0n](https://github.com/r1c0n))
- add build argument: --run-dev [\#3](https://github.com/r1c0n/gws/pull/3) ([r1c0n](https://github.com/r1c0n))
- 1.3.0-rc2 - SSL suport [\#2](https://github.com/r1c0n/gws/pull/2) ([r1c0n](https://github.com/r1c0n))
- 1.3.0 rc1 [\#1](https://github.com/r1c0n/gws/pull/1) ([r1c0n](https://github.com/r1c0n))

## [v1.3.0-rc2](https://github.com/r1c0n/gws/tree/v1.3.0-rc2) (2023-06-28)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.3.0-rc1...v1.3.0-rc2)

## [v1.3.0-rc1](https://github.com/r1c0n/gws/tree/v1.3.0-rc1) (2023-06-28)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.1.0...v1.3.0-rc1)

## [v1.1.0](https://github.com/r1c0n/gws/tree/v1.1.0) (2022-12-24)

[Full Changelog](https://github.com/r1c0n/gws/compare/v1.0.0...v1.1.0)

## [v1.0.0](https://github.com/r1c0n/gws/tree/v1.0.0) (2022-12-24)

[Full Changelog](https://github.com/r1c0n/gws/compare/11a6904ce815ef6074eaf77cd78eaaf61b65497d...v1.0.0)



\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*
