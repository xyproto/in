![logo](img/in_160.png)

# in ![Build](https://github.com/xyproto/in/workflows/Build/badge.svg)

Utility to execute commands in directories, and create directories if needed.

It will also create the directories, if missing. If the top level directory is empty after executing the command, it will be removed. This means that `in testdirectory pwd` leaves no traces.

## Example 1

    in build cmake ..

Instead of:


    mkdir -p build
    cd build
    cmake ..

Or:

    mkdir -p build
    cmake -B build -S .

## Example 2

    in project ./configure --prefix=/usr

Instead of:

    cd project
    ./configure --prefix=/usr
    cd ..

Or:

    (cd project; ./configure --prefix=/usr)

Or:

    pushd project
    ./configure --prefix=/usr
    popd

## Example 3

Globbing (note the double quotes to avoid shell expansion). No directories are created when using globbing, but the given command will be run in each directory where a matching file is found, for each matching file.

    in "./**/*pom.xml" mvn clean

## Installation

Manual installation, using `git`, `rust`, `sudo` and `install`:

    git clone https://github.com/xyproto/in
    cd in
    cargo build --release
    mkdir -p /usr/bin
    sudo install -m755 target/release/in /usr/bin/in

## General info

* Version: 1.7.0
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
