package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ldez/go-git-cmd-wrapper/add"
	"github.com/ldez/go-git-cmd-wrapper/checkout"
	"github.com/ldez/go-git-cmd-wrapper/clone"
	"github.com/ldez/go-git-cmd-wrapper/commit"
	"github.com/ldez/go-git-cmd-wrapper/config"
	"github.com/ldez/go-git-cmd-wrapper/git"
	"github.com/ldez/go-git-cmd-wrapper/push"
	"github.com/ldez/go-git-cmd-wrapper/types"
	"github.com/traefik/mixtus/file"
)

func run(cfg Config) error {
	dir, err := ioutil.TempDir("", "mixtus-*")
	if err != nil {
		return err
	}

	defer func() { _ = os.RemoveAll(dir) }()

	docTarget := filepath.Join(dir, cfg.Target.DocPath)

	docGitURL := makeRemoteURL(cfg.Target, cfg.Token, true)

	output, err := git.Clone(clone.Repository(docGitURL), clone.Depth("1"), clone.Directory(dir), git.Debugger(cfg.Debug))
	if err != nil {
		log.Println(output)
		return fmt.Errorf("failed to clone documentation repository: %w", err)
	}

	// clean doc
	err = os.RemoveAll(docTarget)
	if err != nil {
		return fmt.Errorf("failed to reset documentation: %w", err)
	}

	// copy source docs into targeted directory
	err = file.Copy(cfg.Source.DocPath, docTarget)
	if err != nil {
		return fmt.Errorf("failed to copy documentation: %w", err)
	}

	// move to temp dir
	err = os.Chdir(dir)
	if err != nil {
		return err
	}

	// setup git user info
	output, err = setupGitUserInfo(cfg.Git)
	if err != nil {
		fmt.Println(output)
		return fmt.Errorf("failed to setup Git user configuration: %w", err)
	}

	// check the git status of the dir
	output, err = git.Raw("status", func(g *types.Cmd) { g.AddOptions("-s") })
	if err != nil {
		fmt.Println(output)
		return fmt.Errorf("failed to get Git status: %w", err)
	}

	if len(output) == 0 {
		log.Println("Nothing to commit.")
		return nil
	}

	branchName := filepath.Base(dir)

	// checkout a new branch
	output, err = git.Checkout(checkout.NewBranch(branchName), git.Debugger(cfg.Debug))
	if err != nil {
		log.Println(output)
		return fmt.Errorf("failed to create a new branch: %w", err)
	}

	// add target doc path to the index
	output, err = git.Add(add.PathSpec(cfg.Target.DocPath), git.Debugger(cfg.Debug))
	if err != nil {
		log.Println(output)
		return fmt.Errorf("failed to add files: %w", err)
	}

	// create a commit
	output, err = git.Commit(commit.Message(fmt.Sprintf("Update documentation for %s", cfg.Source.RepoName)), git.Debugger(cfg.Debug))
	if err != nil {
		log.Println(output)
		return fmt.Errorf("failed to commit: %w", err)
	}

	// push the branch to the target git repo
	output, err = git.Push(push.Remote("origin"), push.Repo(cfg.Target.RepoName), git.Debugger(cfg.Debug))
	if err != nil {
		log.Println(output)
		return fmt.Errorf("failed to push: %s", err)
	}

	ctx := context.Background()

	return createPullRequest(ctx, cfg, branchName)
}

func makeRemoteURL(target RepoInfo, token string, ssh bool) string {
	if ssh {
		return fmt.Sprintf("git@github.com:%s/%s.git", target.Owner, target.RepoName)
	}

	prefix := "https://"
	if len(token) > 0 {
		prefix += token + "@"
	}

	return fmt.Sprintf("%sgithub.com/%s/%s.git", prefix, target.Owner, target.RepoName)
}

func setupGitUserInfo(gitInfo GitInfo) (string, error) {
	if len(gitInfo.UserEmail) != 0 {
		output, err := git.Config(config.Entry("user.email", gitInfo.UserEmail))
		if err != nil {
			return output, err
		}
	}

	if len(gitInfo.UserName) != 0 {
		output, err := git.Config(config.Entry("user.name", gitInfo.UserName))
		if err != nil {
			return output, err
		}
	}

	return "", nil
}
