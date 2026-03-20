# CONGA 🔧

Config generator CLI. Build configs without the pain.

Then proceed with installation:

## Downloads

Get CONGA using a prebuilt binary, a package manager, or build from source. Releases are tagged on GitHub; use the tag version (for example `v1.1.0`).

### Prebuilt binaries

Download the matching binary for your platform from the `build/` folder or the GitHub Releases page:

- Linux (x86_64): `build/conga-linux`
- Linux (arm64): `build/conga-linux-arm64`
- macOS (arm64): `build/conga-macos-arm64`
- macOS (intel): `build/conga-macos-intel`

You can also find release bundles on GitHub: https://github.com/pc0stas/conga/releases

### Package managers

- Snap (amd64):

```bash
snap install conga
```

- Homebrew (macOS):

```bash
brew tap pc0stas/conga
brew install conga
```

- Fedora/Copr:

```bash
sudo dnf copr enable pablocostas/conga
sudo dnf install conga
```

Package availability varies across distributions — if you don't find a package for your platform, use the prebuilt binaries or build from source.

### Build from source

## Prerequisites ()

Make sure you have Go installed and have added the Go bin directory to your `$PATH`:

```bash
export PATH="$PATH:$HOME/go/bin"
```

```bash
go install github.com/pc0stas/conga@latest
# or to build locally
go build -o conga ./...
```

For multi-platform builds use the included `build.sh` script.

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
