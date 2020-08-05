package main

// Config is the bot configuration.
type Config struct {
	Token  string
	Git    GitInfo
	Source RepoInfo
	Target RepoInfo
	Debug  bool
}

// GitInfo represents the Git user configuration used for commit.
type GitInfo struct {
	UserName  string
	UserEmail string
}

// RepoInfo represents the information about a GitHub repository.
type RepoInfo struct {
	Owner    string
	RepoName string
	DocPath  string
}
