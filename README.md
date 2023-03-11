# Bootstrap golang project

See https://github.com/golang-standards/project-layout for project structure

## Use this template

Change repo name in `.releaserc`

Run following commands to init repo:

```bash
go mod init $REPO_URL_WITHOUT
```

To add a dependencies used in code (and remove unused):

```bash
go mod tidy
```

To download dependencies:

```bash
go mod vendor
```