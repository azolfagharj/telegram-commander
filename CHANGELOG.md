# Changelog

## Core

### Added
- Command results can be delivered as a `.txt` file. Root keys `output` (`auto` / `text` / `file`, default `auto`) and `max_output_messages` (default `2`) control when a long result becomes a file instead of being cut after ten chat messages. Each button can override `output`.

## Documentation

### Changed
- The configuration page now documents `output` and `max_output_messages`, and explains when results arrive as a file
- The button page now describes the optional `output` override for long results
- The configuration and button pages now say you can omit `output` and `max_output_messages`, and that a button without `output` uses the root setting
