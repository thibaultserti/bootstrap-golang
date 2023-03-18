# Bootstrap golang project

See https://github.com/golang-standards/project-layout for project structure
## Badges

[![Build Status](https://github.com/thibaultserti/bootstrap-golang/actions/workflows/release.yaml/badge.svg)](https://github.com/thibaultserti/bootstrap-golang/actions/workflows/release.yaml)
[![License](https://img.shields.io/github/license/thibaultserti/bootstrap-golang)](/LICENSE)
[![Release](https://img.shields.io/github/release/thibaultserti/bootstrap-golang.svg)](https://github.com/thibaultserti/bootstrap-golang/releases/latest)
[![GitHub Releases Stats of bootstrap-golang](https://img.shields.io/github/downloads/thibaultserti/bootstrap-golang/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=thibaultserti&repository=bootstrap-golang)

## Use this template

Change repo name in `.releaserc`

Run following commands to init repo:

```bash
go mod init $REPO_URL
```

To add a dependencies used in code (and remove unused):

```bash
go mod tidy
```

To download dependencies:

```bash
go mod vendor
```
