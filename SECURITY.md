# Security Policy

## Supported versions

Only the latest GitHub Release receives security fixes.

## How to report

Do not open a public issue, discussion, or pull request for a security vulnerability.

Use **Report a vulnerability** on this repository's Security tab.

Please include:

* The version or git tag you used
* What went wrong
* How to reproduce it
* What an attacker could do with it

There is no paid bug bounty and no guaranteed response time.

Please keep vulnerability details private until a fix or appropriate mitigation is available.

## What is not a vulnerability

Telegram Commander runs the shell commands you configure. Executing those commands is the core purpose of the project.

The following are not security vulnerabilities in Telegram Commander by themselves:

* A command in your configuration is dangerous
* A user listed in `allowed_users` can run the commands they are authorized to run
* Your bot token has been leaked or exposed
* `allowed_users` is empty, missing, or incorrectly configured
* A command you configured allows access to files, services, or systems that you intentionally made available

Configuration and deployment mistakes may still create security risks, but they are not necessarily vulnerabilities in Telegram Commander itself.
