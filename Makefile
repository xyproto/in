PREFIX ?= /usr/local
DESTDIR ?=
BIN_NAME = in
MAN_PAGE = in.1

MAN_DIR = $(DESTDIR)$(PREFIX)/share/man/man1
BIN_DIR = $(DESTDIR)$(PREFIX)/bin

.PHONY: all install clean

all:
	cargo build --release

fmt:
	rustfmt src/*

install: all
	mkdir -p $(BIN_DIR) $(MAN_DIR)
	install target/release/$(BIN_NAME) $(BIN_DIR)/$(BIN_NAME)
	install -m644 $(MAN_PAGE) $(MAN_DIR)/$(MAN_PAGE)

clean:
	cargo clean
