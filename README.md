# in

A command for running commands in other directories.

## Example 1

    in build cmake ..

Instead of:

    mkdir -p build
    cd build
    cmake ..

## Example 2

    in build ./configure --prefix=/usr

Instead of:

    cd build
    ./configure --prefix=/usr

## Why?

`cd build` changes the directory also after the command, you would then have to `cd ..` or `cd $srcdir` afterwards. Or use `pushd` and `popd`. Or use parenthesis, like this, which starts a subshell:

    (cd build; ./configure --prefix=/usr)

Using `in` is nicer:

    in build ./configure --prefix=/usr

Also, `in` can create the top level directory, if missing.

## Installation

Download the latest release and build it with `go build`, then install it with `install -Dm755 in /usr/bin/in`.

Or, for the development release:

    export GOPATH=~/go
    go get github.com/xyproto/in
    install -Dm755 "$GOPATH/bin/in" /usr/bin/in

## Version

1.1
