package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg := Config{}

	flag.StringVar(&cfg.Token, "token", os.Getenv("GITHUB_TOKEN"), "GitHub Token [GITHUB_TOKEN]")

	flag.StringVar(&cfg.Source.Owner, "src-owner", "traefik", "Owner of the source repository.")
	flag.StringVar(&cfg.Source.RepoName, "src-repo-name", "traefik", "Name of the source repo.")
	flag.StringVar(&cfg.Source.DocPath, "src-doc-path", "./docs/site", "Path to the documentation.")

	flag.StringVar(&cfg.Target.Owner, "dst-owner", "traefik", "Owner of the targeted doc repo.")
	flag.StringVar(&cfg.Target.RepoName, "dst-repo-name", "doc", "Name of the targeted doc repo.")
	flag.StringVar(&cfg.Target.DocPath, "dst-doc-path", "./traefik", "Path to put the documentation.")

	flag.StringVar(&cfg.Git.UserName, "git-user-name", os.Getenv("GIT_USER_NAME"), "UserName used to commit the documentation. [GIT_USER_NAME]")
	flag.StringVar(&cfg.Git.UserEmail, "git-user-email", os.Getenv("GIT_USER_EMAIL"), "Email used to commit the documentation. [GIT_USER_EMAIL]")

	flag.BoolVar(&cfg.Debug, "debug", false, "Debug mode")

	version := flag.Bool("v", false, "Show version.")
	help := flag.Bool("h", false, "Show this help.")

	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
	}

	if *version {
		displayVersion()
		return
	}

	nArgs := flag.NArg()
	if nArgs > 0 {
		usage()
	}

	err := validate(cfg)
	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}

	err = run(cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func usage() {
	_, _ = os.Stderr.WriteString(fmt.Sprintf("Lasius Mixtus (%s)\n\nFlags:\n", version))
	flag.PrintDefaults()
	os.Exit(2)
}

func validate(cfg Config) error {
	if cfg.Token == "" {
		return errors.New("token is required")
	}

	err := validateRepoInfo(cfg.Source)
	if err != nil {
		return fmt.Errorf("source: %w", err)
	}

	err = validateRepoInfo(cfg.Target)
	if err != nil {
		return fmt.Errorf("target: %w", err)
	}

	if (cfg.Git.UserName != "" && cfg.Git.UserEmail == "") || (cfg.Git.UserName == "" && cfg.Git.UserEmail != "") {
		return fmt.Errorf("both options (Git user's email and username) must be set together or not set at all: %s - %s",
			cfg.Git.UserName, cfg.Git.UserEmail)
	}

	return nil
}

func validateRepoInfo(info RepoInfo) error {
	if info.Owner == "" {
		return errors.New("owner is required")
	}

	if info.RepoName == "" {
		return errors.New("repository name is required")
	}

	if info.DocPath == "" {
		return errors.New("documentation path is required")
	}

	return nil
}
