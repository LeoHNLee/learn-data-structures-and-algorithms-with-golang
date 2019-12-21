## download go

- [official web site](https://golang.org/doc/install)

## (windows) env setting

- system variable
    - goroot: %goroot% (eg c:\go)
    - path: %gopath%\bin, %goroot%\bin (eg c:\go\bin)

- user variable
    - gopath: %go work dir% (eg c:\user\username\go)
    - path: %go%, %gopath%\bin

## install tools

- command tools: 

```shell
go get golang.org/x/tools/cmd/...
```

- unix style git in windows commands:

```shell
go get github.com/revel/revel // download revel
go get github.com/revel/cmd/revel // download revel command
```

## vs code

- extensions
    - go
    - gocode
    - gopkgs
    - go-outline
    - gocode-gomod
    - godef
    - goreturns
    - golint
    - dlv