# Lasius Mixtus - Publish Documentation on a GitHub Repository

[![GitHub release](https://img.shields.io/github/release/traefik/mixtus.svg)](https://github.com/traefik/mixtus/releases/latest)
[![Build Status](https://travis-ci.com/traefik/mixtus.svg?branch=master)](https://travis-ci.com/traefik/mixtus)
[![Docker Build Status](https://img.shields.io/docker/build/traefik/mixtus.svg)](https://hub.docker.com/r/traefik/mixtus/builds/)

## Description

```
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
        GitUserEmail
  -git-user-name string
        GitUserName
  -h    Show this help.
  -src-doc-path string
        Path to the documentation. (default "./docs/site")
  -src-owner string
        Owner of the source repository. (default "traefik")
  -src-repo-name string
        Name of the source repo. (default "traefik")
  -token string
        GitHub Token
```

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

## What does Lasius Mixtus mean?

![Lasius Mixtus](https://antwiki.org/wiki/images/0/00/Lasius_mixtus_casent0172710_head_1.jpg)
