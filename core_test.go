package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasDiff(t *testing.T) {
	testCases := []struct {
		desc   string
		output string
	}{
		{
			desc:   "empty",
			output: "",
		},
		{
			desc: "only sitemap.xml.gz",
			output: ` M traefik/master/sitemap.xml.gz
 M traefik/sitemap.xml.gz
 M traefik/v2.0/sitemap.xml.gz
 M traefik/v2.1/sitemap.xml.gz
 M traefik/v2.2/sitemap.xml.gz
 M traefik/v2.3/sitemap.xml.gz`,
		},
		{
			desc: "only sitemap.xml",
			output: ` M traefik/master/sitemap.xml
 M traefik/sitemap.xml
 M traefik/v2.0/sitemap.xml
 M traefik/v2.1/sitemap.xml
 M traefik/v2.2/sitemap.xml
 M traefik/v2.3/sitemap.xml`,
		},
		{
			desc: "mixed sitemap.xml and sitemap.xml.gz",
			output: ` M traefik/master/sitemap.xml.gz
 M traefik/master/sitemap.xml
 M traefik/sitemap.xml
 M traefik/sitemap.xml.gz
 M traefik/v2.0/sitemap.xml
 M traefik/v2.0/sitemap.xml.gz
 M traefik/v2.1/sitemap.xml
 M traefik/v2.1/sitemap.xml.gz
 M traefik/v2.2/sitemap.xml
 M traefik/v2.3/sitemap.xml`,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			diff := hasDiff(test.output)

			assert.False(t, diff)
		})
	}

	testCases = []struct {
		desc   string
		output string
	}{
		{
			desc: "simple",
			output: `M traefik/master/aaa.html
 M traefik/bbb.html
 M traefik/v2.0/ccc.html
 M traefik/v2.1/ddd.html
 M traefik/v2.2/eee.html
 M traefik/v2.3/fff.html`,
		},
		{
			desc: "sitemap.xml.gz but with other files",
			output: ` M traefik/master/sitemap.xml.gz
 M traefik/sitemap.xml.gz
 M traefik/v2.0/sitemap.xml.gz
 M traefik/v2.1/sitemap.xml.gz
 M traefik/v2.2/sitemap.xml.gz
 M traefik/v2.3/sitemap.xml.gz
?? foobar.txt`,
		},
		{
			desc: "sitemap.xml but with other files",
			output: ` M traefik/master/sitemap.xml
 M traefik/sitemap.xml
 M traefik/v2.0/sitemap.xml
 M traefik/v2.1/sitemap.xml
 M traefik/v2.2/sitemap.xml
 M traefik/v2.3/sitemap.xml
?? foobar.txt`,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			diff := hasDiff(test.output)

			assert.True(t, diff)
		})
	}
}

func Test_makeRemoteURL(t *testing.T) {
	testCases := []struct {
		desc     string
		target   RepoInfo
		token    string
		ssh      bool
		expected string
	}{
		{
			desc: "https with secret",
			target: RepoInfo{
				Owner:    "owner",
				RepoName: "repo",
			},
			token:    "secret",
			ssh:      false,
			expected: "https://secret@github.com/owner/repo.git",
		},
		{
			desc: "https without secret",
			target: RepoInfo{
				Owner:    "owner",
				RepoName: "repo",
			},
			token:    "",
			ssh:      false,
			expected: "https://github.com/owner/repo.git",
		},
		{
			desc: "ssh with secret",
			target: RepoInfo{
				Owner:    "owner",
				RepoName: "repo",
			},
			token:    "secret",
			ssh:      true,
			expected: "git@github.com:owner/repo.git",
		},
		{
			desc: "ssh without secret",
			target: RepoInfo{
				Owner:    "owner",
				RepoName: "repo",
			},
			token:    "",
			ssh:      true,
			expected: "git@github.com:owner/repo.git",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			remoteURL := makeRemoteURL(test.target, test.token, test.ssh)

			assert.Equal(t, test.expected, remoteURL)
		})
	}
}
