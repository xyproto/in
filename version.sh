#!/bin/sh -e
#
# Self-modifying script that updates the version numbers
#

# The current version goes here, as the default value
VERSION=${1:-'1.6.0'}

if [ -z "$1" ]; then
  echo "The current version is $VERSION, pass the new version as the first argument if you wish to change it"
  exit 0
fi

echo "Setting the version to $VERSION"

# Update the date and version in the man page, README.md file and also this script
d=$(LC_ALL=C date +'%d %b %Y')

# macOS
sed -E -i '' "s/1\.[[:digit:]]+\.[[:digit:]]+/$VERSION/g" README.md "$0" main.go 2> /dev/null || true

# Linux
sed -r -i "s/1\.[[:digit:]]+\.[[:digit:]]+/$VERSION/g" o.1 README.md "$0" main.go 2> /dev/null || true
