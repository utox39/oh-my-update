package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"path"
)

const omhGenericError = "omu: error:"

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

func UpdateFolder(folder string) {
	fmt.Printf("Updating %s...\n\n", folder)

	homeFolder, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(omhGenericError, err)
		os.Exit(1)
	}

	omzFolder := path.Join(homeFolder, ".oh-my-zsh/custom", folder)

	if _, err := os.Stat(omzFolder); errors.Is(err, os.ErrNotExist) {
		fmt.Println(omhGenericError, err)
		os.Exit(1)
	}

	entries, err := os.ReadDir(omzFolder)
	if err != nil {
		fmt.Println(omhGenericError, err)
		os.Exit(1)
	}

	for _, e := range entries {
		if e.IsDir() && e.Name() != "example" {
			fmt.Println("Updating:", e.Name())
			gitFolder := path.Join(omzFolder, e.Name())
			if updateErr := UpdateRepo(gitFolder); updateErr != nil {
				HandleError(updateErr, e.Name())
			} else {
				fmt.Println(e.Name(), "updated successfully")
				fmt.Println()
			}
		}
	}
}

func HandleError(err error, pluginName string) {
	if errors.Is(err, git.NoErrAlreadyUpToDate) {
		fmt.Printf("%s is %s\n\n", pluginName, err)
	} else {
		fmt.Printf("omu: %s: %s\n\n", pluginName, err)
	}
}

func main() {
	UpdateFolder("plugins")
	UpdateFolder("themes")
}
