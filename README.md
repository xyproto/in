# in

A command for running commands in other directories.

Example:

    in build ./configure --prefix=/usr

Instead of:

    cd build
    ./configure --prefix=/usr

or

    cd build; ./configure --prefix=/usr

Why?

`cd build` changes the directory also after the command, you would then have to `cd ..` or `cd $srcdir` afterwards. Or use `pushd` and `popd`. Or use parenthesis, like this, which starts a subshell:

    (cd build; ./configure --prefix=/usr)

Using `in` is nicer:

    in build ./configure --prefix=/usr
