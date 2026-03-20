# CONGA 🔧

Config generator CLI. Build Nginx, Docker Compose, and WireGuard configs without the pain.

## Install

```bash
go install github.com/pc0stas/conga@latest
```

```bash
snap install conga #amd64 only
apt install conga # Debian/Ubuntu
```

```bash
sudo dnf copr enable pablocostas/conga
sudo dnf install conga
```

Or grab a [binary](https://github.com/pc0stas/conga/releases).

**Coming soon:**

```bash
dnf install conga # Fedora
brew install conga # macOS
```

## Usage

```bash
conga nginx generate # Generate Nginx config
conga docker generate # Generate Docker Compose
conga wireguard generate # Generate WireGuard config
conga help # Show help
conga version # Show version
```

## What it does

**Nginx Generator**

- Interactive setup for Nginx configs
- Proxy + static file serving
- HTTPS/SSL support
- WebSocket support
- Generates `default.conf`

**Docker Compose Generator**

- Define services interactively
- Port mapping, volumes, env vars
- Generates `docker-compose.yml`

**WireGuard Generator**

- Interactive peer configuration
- Server and client setup
- Key generation
- Network setup
- Generates peer configs

## Example

```bash
$ conga nginx generate

? Main domain: example.com
? Use HTTPS? Yes
📍 Configuring routes...
? Number of routes: 2

📌 Route 1:
? Path: /
? Type: Static files
? Directory: /var/www/html

📌 Route 2:
? Path: /api
? Type: Proxy
? Proxy destination: localhost:3000
? Add WebSocket support? Yes

✅ Generated: default.conf
```

## Build

```bash
go build -o conga ./...
```

Multi-platform:

```bash
./build.sh
```

## Status

Check the [project board](https://github.com/users/PC0staS/projects/8) for what’s coming.

**Done:**

- ✅ Published in snap

- ✅ Nginx generator
- ✅ Docker Compose generator

**Coming:**

- 🔄 WireGuard generator
- 🔄 Apache generator
- 🔄 Caddy generator
- 🔄 Template system
- 📦 Package managers (apt, dnf, brew, AUR)

## License

MIT
