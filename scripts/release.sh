#!/usr/bin/env bash
set -euo pipefail

# Usage: ./scripts/release.sh v1.2.3
# Accepts a version with or without leading 'v'. Commits changes, pushes, and creates tag.

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 vX.Y.Z"
  exit 1
fi

RAW_VER="$1"
TAG_VER="$RAW_VER"
# normalize without leading v for files that don't use the 'v' prefix
NO_V_VER="${RAW_VER#v}"

if [[ ! "$RAW_VER" =~ ^v?[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Invalid version format: $RAW_VER"
  echo "Expected vMAJOR.MINOR.PATCH or MAJOR.MINOR.PATCH"
  exit 2
fi

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT_DIR"

# Ensure working tree is clean
if [ -n "$(git status --porcelain)" ]; then
  echo "Working tree not clean. Commit or stash changes first."
  git status --porcelain
  exit 3
fi

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "Current branch: $CURRENT_BRANCH"

# Check tag doesn't already exist
if git rev-parse "refs/tags/$TAG_VER" >/dev/null 2>&1; then
  echo "Tag $TAG_VER already exists. Aborting."
  exit 4
fi

echo "Updating version to $NO_V_VER (tag: $TAG_VER)"

# Update main.go (var Version = "x.y.z")
sed -E -i.bak "s/^(var Version = \").*(\")/\1${NO_V_VER}\2/" main.go

# Update conga.spec (Version:        x.y.z)
sed -E -i.bak "s/^(Version:)[[:space:]]+.*/Version:        ${NO_V_VER}/" conga.spec

# Update snapcraft.yaml (version: "x.y.z")
sed -E -i.bak "s/^version:[[:space:]]+.*/version: \"${NO_V_VER}\"/" snapcraft.yaml

echo "Files updated. Showing git diff for review:"
git --no-pager diff -- main.go conga.spec snapcraft.yaml || true

read -p "Continue and commit changes? [y/N] " CONFIRM
if [[ "$CONFIRM" != "y" && "$CONFIRM" != "Y" ]]; then
  echo "Aborted by user. Restoring backups."
  mv -f main.go.bak main.go || true
  mv -f conga.spec.bak conga.spec || true
  mv -f snapcraft.yaml.bak snapcraft.yaml || true
  exit 5
fi

# Remove backup files after committing
git add main.go conga.spec snapcraft.yaml
git commit -m "Release ${TAG_VER}"
git push origin "$CURRENT_BRANCH"

git tag -a "${TAG_VER}" -m "Release ${TAG_VER}"
git push origin "${TAG_VER}"

echo "Release ${TAG_VER} created and pushed. Cleaning backups."
rm -f main.go.bak conga.spec.bak snapcraft.yaml.bak

echo "Done."
