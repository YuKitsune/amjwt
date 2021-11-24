<h1 align="center">
    🍎🎵
    <br />
    AMJWT
</h1>

<h3 align="center">
  An Apple Music JWT generator.

  [![GitHub Workflow Status](https://img.shields.io/github/workflow/status/yukitsune/amjwt/CI)](https://github.com/yukitsune/amjwt/actions?query=workflow:CI)
  [![Go Report Card](https://goreportcard.com/badge/github.com/yukitsune/amjwt)](https://goreportcard.com/report/github.com/yukitsune/amjwt)
  [![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/yukitsune/amjwt)](https://pkg.go.dev/mod/github.com/yukitsune/amjwt)
</h3>

# Usage

## Pre-requisites
Before you can generate a JWT, you'll need a **key id**, **team id**, and **private key file**.
Rather than re-explaining that process, he's Apples docs: https://developer.apple.com/documentation/applemusicapi/getting_keys_and_creating_tokens

Once you have those ready, you can use `amjwt` to generate the JWT.

## As a CLI

`amjwt` can be used as a CLI. You can provide the key and team IDs via the `-k` and `-t` flags respectively. `-f` can be used to specify the path to the private key file.
```
amjwt -k <key-id> -t <team-id> -f ./MyPrivateKey.p8
```

You can also pipe the private key in if necessary:
```
cat ./MyPrivateKey.p8 | amjwt -k <key-id> -t <team-id>
```

## As a Package

You can also import amjwt as a go package:
```go
import "github.com/yukitsune/amjwt"
keyId := "foo"
teamId := "bar"
privateKeyBytes, err = ioutil.ReadFile("./somewhere/MyPrivateKey.p8")
jwtString, err := amjwt.CreateJwt(keyId, teamId, expiryDays, privateKeyBytes)
```

# Contributing

Contributions are what make the open source community such an amazing place to be, learn, inspire, and create.
Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`feature/AmazingFeature`)
3. Commit your Changes
4. Push to the Branch
5. Open a Pull Request

# But... Why?
Because for some reason, [Apple wants you to hand craft the JWT](https://developer.apple.com/documentation/applemusicapi/getting_keys_and_creating_tokens) to authenticate with the Apple Music API.
This is much easier than hand crafting it every time.
