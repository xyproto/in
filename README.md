![In Logo](img/logo.png)

[![Build](https://github.com/xyproto/in/actions/workflows/build.yml/badge.svg)](https://github.com/xyproto/in/actions/workflows/build.yml) [![License](https://img.shields.io/badge/license-BSD-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/in/main/LICENSE)

Utility to execute commands in directories, and create directories if needed.

It will also create the directories, if missing. If the top level directory is empty after executing the command, it will be removed. This means that `in testdirectory pwd` leaves no traces.

## Example 1

    in build cmake ..

Instead of:


    mkdir -p build
    cd build
    cmake ..
    cd ..

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

### Linux

Manual installation, using `cargo`, `git`, `install` and `sudo`:

    git clone https://github.com/xyproto/in
    cd in
    cargo build --release
    sudo install -Dm755 target/release/in /usr/bin/in

### FreeBSD

Manual installation, using `cargo`, `doas`, `git` and `install`:

    git clone https://github.com/xyproto/in
    cd in
    cargo build --release
    mkdir -p /usr/bin
    doas install -m755 target/release/in /usr/bin/in

## General info

* Version: 1.7.3
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
