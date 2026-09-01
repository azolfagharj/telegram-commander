# Changelog

## Added
- Long command output is now delivered as several messages instead of one. Each part is sent as a reply to the part before it, so the whole result stays grouped and in order in the chat.
- The menu buttons appear on the last part, so you can keep using the bot right where the output ends.
- Output is split on line boundaries whenever possible, so a line is not cut in half between two messages.
- When a result is still too long after 10 messages, the last message ends with a note telling you how much of the output was shown.

## Changed
- The full example config now includes the System and Network sample menu together with the nginx, Docker, and SSH demos
- Example configs turn on the Run Command button

## Fixed
- A command whose output was a little over Telegram's message size used to show only `Running …` and then nothing. The result message was cut in the middle and Telegram refused it. Results now always arrive complete.
