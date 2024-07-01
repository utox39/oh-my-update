package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	"os"
	"path"
)

// LogErr is a custom logger
func LogErr(err error) {
	logger := log.New(os.Stderr, "omu: ", 0)
	logger.Fatalln(err)
}

func UpdateRepo(gitFolder string) error {
	repo, err := git.PlainOpen(gitFolder)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	var output bytes.Buffer
	err = worktree.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: "refs/heads/master",
		SingleBranch:  true,
		Progress:      &output,
	})

	if output.String() != "" {
		fmt.Println(output.String())
	}

	if err != nil {
		return err
	}

	return nil
}

func UpdateFolder(folder string) error {
	fmt.Printf("Updating %s...\n\n", folder)

	homeFolder, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	omzFolder := path.Join(homeFolder, ".oh-my-zsh/custom", folder)

	if _, err := os.Stat(omzFolder); errors.Is(err, os.ErrNotExist) {
		return err
	}

	entries, err := os.ReadDir(omzFolder)
	if err != nil {
		return err
	}

	for _, e := range entries {
		if e.IsDir() && e.Name() != "example" {
			fmt.Println("Updating:", e.Name())
			gitFolder := path.Join(omzFolder, e.Name())
			if updateErr := UpdateRepo(gitFolder); updateErr != nil {
				if errors.Is(updateErr, git.NoErrAlreadyUpToDate) {
					fmt.Printf("%s is %s\n\n", e.Name(), updateErr)
					continue
				} else {
					return fmt.Errorf(fmt.Sprintf("%s: %v", e.Name(), updateErr))
				}
			}

			fmt.Printf(e.Name(), "updated successfully\n")
		}
	}

	return nil
}

func main() {
	if err := UpdateFolder("plugins"); err != nil {
		LogErr(err)
	}
	if err := UpdateFolder("themes"); err != nil {
		LogErr(err)
	}
}
