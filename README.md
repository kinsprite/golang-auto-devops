# GitLab Hello World in Golang

## build

on Windows:
```cmd
set GOPROXY=https://goproxy.io
set GIN_MODE=release
go build -tags=jsoniter -o gitlab-hello-world.exe .
```

or Linux:

```shell
export GOPROXY=https://goproxy.io
export GIN_MODE=release
go build -tags=jsoniter -o gitlab-hello-world .
```
