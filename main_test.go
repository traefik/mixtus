package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_validate(t *testing.T) {
	testCases := []struct {
		desc     string
		cfg      Config
		expected string
	}{
		{
			desc: "success with Git info",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
		},
		{
			desc: "success without Git info",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "", UserEmail: ""},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
		},
		{
			desc: "missing token",
			cfg: Config{
				Token:  "",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
			expected: "token is required",
		},
		{
			desc: "missing Git username",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
			expected: "both options (Git user's email and username) must be set together or not set at all:  - bot@example.com",
		},
		{
			desc: "missing Git email",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: ""},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
			expected: "both options (Git user's email and username) must be set together or not set at all: bot - ",
		},
		{
			desc: "missing source owner",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
			expected: "source: owner is required",
		},
		{
			desc: "missing source repo name",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
			expected: "source: repository name is required",
		},
		{
			desc: "missing source do path",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: ""},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: "./destination/"},
			},
			expected: "source: documentation path is required",
		},
		{
			desc: "missing target owner",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "", RepoName: "docs", DocPath: "./destination/"},
			},
			expected: "target: owner is required",
		},
		{
			desc: "missing target repo name",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "", DocPath: "./destination/"},
				Debug:  true,
			},
			expected: "target: repository name is required",
		},
		{
			desc: "missing target doc path",
			cfg: Config{
				Token:  "secret",
				Git:    GitInfo{UserName: "bot", UserEmail: "bot@example.com"},
				Source: RepoInfo{Owner: "lasius", RepoName: "project", DocPath: "./sources/docs/"},
				Target: RepoInfo{Owner: "lasius", RepoName: "docs", DocPath: ""},
			},
			expected: "target: documentation path is required",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			err := validate(test.cfg)

			if test.expected != "" {
				assert.EqualError(t, err, test.expected)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
