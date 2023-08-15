package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

const versionString = "in 1.6.0"

func main() {
	app := &cli.App{
		Name:    "in",
		Usage:   "for executing commands within directories and create them if needed",
		Version: versionString,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "print verbose logs",
				Aliases: []string{"V"},
			},
		},
		Action: func(c *cli.Context) error {
			verbose := c.Bool("verbose")

			// If no arguments, print the current directory
			if c.NArg() == 0 {
				dir, err := os.Getwd()
				if err != nil {
					return err
				}
				resolvedDir, err := filepath.EvalSymlinks(dir)
				if err != nil {
					return err
				}
				fmt.Println(resolvedDir)
				return nil
			}

			// Target directory to operate within or glob pattern
			dirOrPattern := c.Args().Get(0)

			if verbose {
				fmt.Println("Directory or pattern:", dirOrPattern)
			}

			// If it's a glob pattern, run the command in all matching directories
			if strings.Contains(dirOrPattern, "*") {
				return runInAllMatching(dirOrPattern, c.Args().Tail(), verbose)
			}

			absDir, err := filepath.Abs(dirOrPattern)
			if err != nil {
				return err
			}

			// Otherwise, it's treated as a directory
			dirCreated, err := createAndEnter(dirOrPattern, verbose)
			if err != nil {
				return err
			}
			defer func() {
				if verbose {
					fmt.Printf("directories created: %d\n", dirCreated)
				}
				if dirCreated > 0 {
					if err := removeIfEmpty(dirOrPattern, dirCreated, verbose); err != nil {
						log.Println("Failed to remove the directory:", err)
					}
				}
			}()

			// If there's a command to run, run it
			if c.NArg() > 1 {
				if verbose {
					fmt.Println("Command:", c.Args().Tail())
				}
				return run(absDir, c.Args().Tail(), verbose)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
