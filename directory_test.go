package main

import (
	"os"
	"path/filepath"
	"testing"
)

const verbose = true

func TestCreateAndEnter(t *testing.T) {
	t.Run("enter existing directory", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		created, err := createAndEnter(tmpDir, verbose)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if created != 0 {
			t.Error("Expected no directories to be created since it exists")
		}
	})

	t.Run("create and enter new directory", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		newDir := filepath.Join(tmpDir, "new-dir")

		created, err := createAndEnter(newDir, verbose)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if created == 0 {
			t.Error("Expected at least one directory to be created since it doesn't exist")
		}
	})

	t.Run("error for invalid directory", func(t *testing.T) {
		invalidDir := "" // assuming empty string is an invalid directory

		created, err := createAndEnter(invalidDir, verbose)
		if err == nil {
			t.Fatalf("Expected an error for invalid directory")
		}
		if created != 0 {
			t.Error("Expected no directories to be created due to an error")
		}
	})

	t.Run("create multiple directories and remove if empty", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "test")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		multiDirPath := filepath.Join(tmpDir, "a/b/c")

		created, err := createAndEnter(multiDirPath, verbose)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if created == 0 {
			t.Error("Expected directories to be created since they don't exist")
		}

		cwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		if cwd != multiDirPath {
			t.Errorf("Expected current directory to be %s, got %s", multiDirPath, cwd)
		}

		err = removeIfEmpty(multiDirPath, created, verbose) // start removal from 'b', as 'c' is our current directory
		if err != nil {
			t.Fatalf("Expected no error while removing, got %v", err)
		}

		_, err = os.Stat(filepath.Join(tmpDir, "a"))
		if !os.IsNotExist(err) {
			t.Error("Expected directory 'a' to be removed as it is empty")
		}
	})
}
