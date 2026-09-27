# Changelog

## Telegram

### Changed
- A brief connectivity drop no longer makes a command result disappear; the bot retries the send with a growing wait and only gives up (silently, logged on the server) after the retry window runs out

## Configuration

### Added
- Three new root settings (`delivery_retry_backoff`, `delivery_retry_backoff_max`, `delivery_retry_ttl`) to control how long the bot keeps retrying a failed send before giving up

## Documentation

### Added
- A "Delivery retries" page explaining how a short network drop is retried automatically without losing a result
