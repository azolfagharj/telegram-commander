# Changelog

## Added

- MIT license, README badges (CI, docs, license, donate), and a donate footer on the docs site
- Material documentation with a step-by-step user guide (install, concepts, buttons, functions, CLI)
- Extra custom function examples in the release pack (`curl-url`, `journal-unit`, `disk-path`, `ping-host`)
- GitHub Release now also attaches the linux-amd64 and linux-arm64 binaries next to the `.tar.gz` pack

## Changed

- Telegram menu uses a reply keyboard (bottom of the chat) instead of an inline keyboard, with Home / Back / paging and confirm buttons
- MkDocs config and Python requirements live in `docs/`; the docs workflow publishes the built site to the `gh-pages` branch
- systemd unit runs as root; Docker packaging was removed
- GitHub Release body includes `CHANGELOG.md` as-is before the Assets section, then CI empties `CHANGELOG.md` on the repo

## Fixed

- GitHub Actions updated to Node.js 24-compatible action versions
