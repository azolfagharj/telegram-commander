# Install

## Download release pack

```bash
wget -O telegram-commander.tar.gz https://github.com/azolfagharj/telegram-commander/releases/latest/download/telegram-commander.tar.gz
tar -xzf telegram-commander.tar.gz
```

Archive root contains:

- `telegram-commander-linux-amd64`
- `telegram-commander-linux-arm64`
- `config.minimal.yaml`
- `config.full.yaml`
- `telegram-commander.service`

## Manual run

```bash
# amd64
cp telegram-commander-linux-amd64 telegram-commander && chmod +x telegram-commander
# arm64: cp telegram-commander-linux-arm64 telegram-commander && chmod +x telegram-commander

# edit config.minimal.yaml (token + user id)
./telegram-commander validate --config config.minimal.yaml
./telegram-commander run --config config.minimal.yaml
```

## systemd (root)

No Linux user is created. The service runs as root.

```bash
sudo mkdir -p /etc/telegram-commander
# amd64:
sudo cp telegram-commander-linux-amd64 /etc/telegram-commander/telegram-commander
# arm64: sudo cp telegram-commander-linux-arm64 /etc/telegram-commander/telegram-commander
sudo chmod +x /etc/telegram-commander/telegram-commander
sudo cp config.minimal.yaml /etc/telegram-commander/config.yaml
# edit /etc/telegram-commander/config.yaml (token + user id)
sudo cp telegram-commander.service /etc/systemd/system/
sudo systemctl daemon-reload && sudo systemctl enable --now telegram-commander
```

Useful commands:

```bash
sudo systemctl status telegram-commander
sudo journalctl -u telegram-commander -f
```

## Build from source

```bash
git clone https://github.com/azolfagharj/telegram-commander.git
cd telegram-commander
go build -o telegram-commander ./cmd/telegram-commander
sudo mkdir -p /etc/telegram-commander
sudo cp telegram-commander /etc/telegram-commander/telegram-commander
sudo chmod +x /etc/telegram-commander/telegram-commander
```
