# Bootstrap golang project

See https://github.com/golang-standards/project-layout for project structure
## Badges

[![Build Status](https://github.com/thibaultserti/bootstrap-golang/actions/workflows/release.yaml/badge.svg)](https://github.com/thibaultserti/bootstrap-golang/actions/workflows/release.yaml)
[![License](https://img.shields.io/github/license/thibaultserti/bootstrap-golang)](/LICENSE)
[![Release](https://img.shields.io/github/release/thibaultserti/bootstrap-golang.svg)](https://github.com/thibaultserti/bootstrap-golang/releases/latest)
[![GitHub Releases Stats of bootstrap-golang](https://img.shields.io/github/downloads/thibaultserti/bootstrap-golang/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=thibaultserti&repository=bootstrap-golang)

[![Maintainability](https://api.codeclimate.com/v1/badges/4133d7da3d73fa0c0884/maintainability)](https://codeclimate.com/github/thibaultserti/bootstrap-golang/maintainability)
[![codecov](https://codecov.io/gh/thibaultserti/bootstrap-golang/branch/main/graph/badge.svg?token=5BO47LR632)](https://codecov.io/gh/thibaultserti/bootstrap-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/thibaultserti/test-saas-ci)](https://goreportcard.com/report/github.com/thibaultserti/bootstrap-golang)

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
