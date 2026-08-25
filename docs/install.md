# Install

## Build from source

```bash
git clone https://github.com/azolfagharj/telegram-commander.git
cd telegram-commander
go build -o telegram-commander ./cmd/telegram-commander
sudo install -m 755 telegram-commander /usr/local/bin/
```

## systemd

1. Create user and directories:

```bash
sudo useradd --system --home /var/lib/telegram-commander --shell /usr/sbin/nologin telegram-commander
sudo mkdir -p /etc/telegram-commander /var/log/telegram-commander
sudo cp examples/config.yaml /etc/telegram-commander/config.yaml
sudo chown -R telegram-commander:telegram-commander /etc/telegram-commander /var/log/telegram-commander
```

2. Install the unit:

```bash
sudo cp deployments/systemd/telegram-commander.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now telegram-commander
```

Edit `/etc/telegram-commander/config.yaml` with your bot token and allowed users.
If buttons run privileged commands (`systemctl`, etc.), either grant the service user the needed rights
or run under a user that already has them. Prefer least privilege.

## Docker

```bash
docker build -f deployments/Dockerfile -t telegram-commander .
docker run --rm -v "$PWD/config.yaml:/config.yaml:ro" telegram-commander run --config /config.yaml
```

Note: running host `systemctl` commands from inside a container usually does not make sense.
Prefer the binary + systemd unit on the host for server administration menus.

## Releases

Multi-platform binaries can be produced with [GoReleaser](https://goreleaser.com/) using `.goreleaser.yaml`.
