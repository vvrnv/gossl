# gossl

<p align="center">
<img src="https://img.shields.io/github/downloads/vvrnv/gossl/total" alt="Total Downloads">
<img src="https://img.shields.io/github/go-mod/go-version/vvrnv/gossl" alt="Go Version">
<a href="https://pkg.go.dev/github.com/vvrnv/gossl"><img src="https://pkg.go.dev/badge/github.com/vvrnv/gossl.svg" alt="Go Version"></a><br>
</p>

Simple CLI app for checking SSL certificates written with Go

## Installation

### Homebrew

```sh
brew install vvrnv/tap/gossl
```

### Go

```sh
go install github.com/vvrnv/gossl@latest
```

### Download binary

[release page link](https://github.com/vvrnv/gossl/releases)

## Commands

### help

`help` Help about any command.

```sh
gossl help
gossl verify -h
gossl verify --help
```

### version

`version` Print the version information of gossl

```sh
gossl verion
```

### completion

`completion` Generate the autocompletion script for the specified shell

```sh
gossl completion [bash | fish | powershell | zsh]
```

### verify

`verify` verify SSL certificate

```sh
gossl verify -s [dnsName | ipAddress]
gossl verify --server [dnsName | ipAddress]
gossl verify -s [dnsName | ipAddress] -p 8443
gossl verify -s [dnsName | ipAddress] --timeout 25
```

### Usage

```sh
gossl verify -s [dnsName | ipAddress]
```

![image](https://user-images.githubusercontent.com/40491079/210393898-118958e2-0365-47bc-8323-764a43f07c0c.png)
