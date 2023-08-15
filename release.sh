#!/bin/sh
#
# Create release tarballs/zip-files for Rust using cargo
#

rustup_targets="
aarch64-unknown-linux-gnu
i686-unknown-linux-gnu
x86_64-apple-darwin
x86_64-unknown-linux-gnu
aarch64-apple-darwin
aarch64-unknown-linux-musl
arm-unknown-linux-gnueabi
arm-unknown-linux-gnueabihf
armv7-unknown-linux-gnueabihf
riscv64gc-unknown-linux-gnu
x86_64-unknown-freebsd
x86_64-unknown-linux-musl
x86_64-unknown-netbsd
"

for target in $rustup_targets; do
  rustup target add "$target"
done

platforms="
aarch64-unknown-linux-gnu,linux_aarch64,tar.xz
i686-unknown-linux-gnu,linux_i686,tar.xz
x86_64-apple-darwin,macos_x86_64,tar.gz
x86_64-unknown-linux-gnu,linux_x86_64,tar.xz
aarch64-apple-darwin,macos_aarch64,tar.gz
aarch64-unknown-linux-musl,linux_aarch64_musl,tar.xz
arm-unknown-linux-gnueabi,linux_armv6,tar.xz
arm-unknown-linux-gnueabihf,linux_armv6hf,tar.xz
armv7-unknown-linux-gnueabihf,linux_armv7,tar.xz
riscv64gc-unknown-linux-gnu,linux_riscv64,tar.xz
x86_64-unknown-freebsd,freebsd_x86_64,tar.xz
x86_64-unknown-linux-musl,linux_x86_64_musl,tar.xz
x86_64-unknown-netbsd,netbsd_x86_64,tar.xz
"

name='in'
version='1.7.1'  # This is now set as a variable to be reused

compile_and_compress() {
  target="$1"
  platform="$2"
  compression="$3"

  echo "Compiling $name.$platform..."

  cargo build --release --target="$target" || {
    echo "Error: failed to compile for $platform"
    echo "Platform string: $p"
    echo "Target: $target"
    exit 1
  }

  echo "Compressing $name-$version.$platform.$compression"
  mkdir "$name-$version-$platform"
  cp ../$name.1 "$name-$version-$platform/"
  gzip "$name-$version-$platform/$name.1"
  cp "target/$target/release/$name" "$name-$version-$platform/"
  cp ../LICENSE "$name-$version-$platform/"

  case "$compression" in
    tar.xz)
      tar Jcf "$name-$version-$platform.$compression" "$name-$version-$platform"
      ;;
    tar.gz)
      tar zcf "$name-$version-$platform.$compression" "$name-$version-$platform"
      ;;
  esac

  rm -r "$name-$version-$platform"
}

echo 'Compiling...'
while read -r p; do
  [ -z "$p" ] && continue
  IFS=',' read -r target platform compression << EOF
$p
EOF
  compile_and_compress "$target" "$platform" "$compression
  #&
done << EOF
$platforms
EOF

wait

mkdir -p release
mv -v $name-$version* release
