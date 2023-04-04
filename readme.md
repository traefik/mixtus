# Lasius Mixtus - Publish Documentation to a GitHub Repository From Another

[![GitHub release](https://img.shields.io/github/release/traefik/mixtus.svg)](https://github.com/traefik/mixtus/releases/latest)
[![Build Status](https://github.com/traefik/mixtus/workflows/Main/badge.svg?branch=master)](https://github.com/traefik/mixtus/actions)
[![Docker Build Status](https://img.shields.io/docker/cloud/build/traefik/mixtus.svg)](https://hub.docker.com/r/traefik/mixtus/builds/)

## Description

Lasius Mixtus is a cross-ci tool (GitHub Actions, SemaphoreCI 1 and 2, TravisCI, ...) used to aggregate documentation from different projects into one repository.

It is useful for building an aggregated documentation from different sources.

It creates PRs instead of commits to avoid conflicts and be able to validate the whole documentation before the merge.

```yml
Lasius Mixtus

Flags:
  -debug
        Debug mode
  -dst-doc-path string
        Path to put the documentation. (default "./traefik")
  -dst-owner string
        Owner of the targeted doc repo. (default "traefik")
  -dst-repo-name string
        Name of the targeted doc repo. (default "doc")
  -git-user-email string
        Email used to commit the documentation. [GIT_USER_EMAIL]
  -git-user-name string
        UserName used to commit the documentation. [GIT_USER_NAME]
  -h    Show this help.
  -src-doc-path string
        Path to the documentation. (default "./docs/site")
  -src-owner string
        Owner of the source repository. (default "traefik")
  -src-repo-name string
        Name of the source repo. (default "traefik")
  -token string
        GitHub Token [GITHUB_TOKEN]
```

## Workflow Example

![mixtus-workflow](https://user-images.githubusercontent.com/5674651/110240947-993cb000-7f4e-11eb-9b23-ce429cfdebf1.png)

This workflow also uses:

- [structor](https://github.com/traefik/structor): creates multiple versions of a Mkdocs documentation 
- [seo-doc](https://github.com/traefik/seo-doc): enrich HTML files with pre-requisites for improving SEO
- [mixtus](https://github.com/traefik/mixtus): creates PRs with documentation changes
- [lobicornis](https://github.com/traefik/lobicornis): rebases and merges PRs automatically

The result is here: https://doc.traefik.io/

## Examples

```bash
GITHUB_TOKEN=xxx ./mixtus \
--src-owner=containous \
--src-repo-name=traefik \
--src-doc-path="./docs/site/" \
--dst-repo-name=doc \
--dst-doc-path="./traefik" \
--git-user-name=botname \
--git-user-email=bot@example.com
```

## The Mymirca colony

- [Myrmica Lobicornis](https://github.com/traefik/lobicornis) 🐜: Update and merge pull requests.
- [Myrmica Aloba](https://github.com/traefik/aloba) 🐜: Add labels and milestone on pull requests and issues.
- [Messor Structor](https://github.com/traefik/structor) 🐜: Manage multiple documentation versions with Mkdocs.
- [Lasius Mixtus](https://github.com/traefik/mixtus) 🐜: Publish documentation to a GitHub repository from another.
- [Myrmica Bibikoffi](https://github.com/traefik/bibikoffi) 🐜: Closes stale issues
- [Chalepoxenus Kutteri](https://github.com/traefik/kutteri) 🐜: Track a GitHub repository and publish on Slack.
- [Myrmica Gallienii](https://github.com/traefik/gallienii) 🐜: Keep Forks Synchronized

## What does Lasius Mixtus mean?

![Lasius Mixtus](https://antwiki.org/wiki/images/0/00/Lasius_mixtus_casent0172710_head_1.jpg)
