# in

A command for running commands in other directories.

## Example

Example:

    in build ./configure --prefix=/usr

Instead of:

    cd build
    ./configure --prefix=/usr

or

    cd build; ./configure --prefix=/usr

If the directory is missing, it will be created (but not the parent directories, if missing).

## Why?

`cd build` changes the directory also after the command, you would then have to `cd ..` or `cd $srcdir` afterwards. Or use `pushd` and `popd`. Or use parenthesis, like this, which starts a subshell:

    (cd build; ./configure --prefix=/usr)

Using `in` is nicer:

    in build ./configure --prefix=/usr

When using CMake, using `in` is handy. Instead of:

    mkdir -p build; cd build; cmake ..

One can:

    in build cmake ..

## Installation

One of many possible ways:

    export GOPATH=~/go
    go get github.com/xyproto/in
    install -Dm755 "$GOPATH/bin/in" /usr/bin/in

## Version

1.1
