# Golang

## Install golang version 1.13.9
```
$ wget https://dl.google.com/go/go1.13.9.linux-amd64.tar.gz
$ tar xf go1.13.9.linux-amd64.tar.gz
$ sudo mv go /usr/local/go-1.13
```
## Export paths
```
export GOROOT=/usr/local/go-1.13
export PATH=$GOROOT/bin:$PATH
```

## Verify version and env
```
$ go version
$ go env

"Output of go env"
GOROOT="/usr/local/go-1.13"
```

### Create new project and create server
```
$ mkdir myapp && cd myapp
$ go get github.com/labstack/echo/v4

```

## Run the project

```
go run main.go
```
