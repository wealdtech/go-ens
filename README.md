# go-ens

[![Tag](https://img.shields.io/github/tag/wealdtech/go-ens.svg)](https://github.com/wealdtech/go-ens/releases/)
[![License](https://img.shields.io/github/license/wealdtech/go-ens.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wealdtech/go-ens?status.svg)](https://godoc.org/github.com/wealdtech/go-ens)
[![Travis CI](https://img.shields.io/travis/wealdtech/go-ens.svg)](https://travis-ci.org/wealdtech/go-ens)
[![codecov.io](https://img.shields.io/codecov/c/github/wealdtech/go-ens.svg)](https://codecov.io/github/wealdtech/go-ens)

Go module to simplify interacting with the [Ethereum Name Service](https://ens.domains/) contracts.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-ens` is a standard Go module which can be installed with:

```sh
go get github.com/wealdtech/go-ens
```

## Usage

`go-ens` provides simple access to the [Ethereum Name Service](https://ens.domains/) contracts.

### Example

```go
package main

import (
    "github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens"
)

func main() {
    client, err := ethclient.Dial("https://infura.io/v3/SECRET")
    if err != nil {
        panic(err)
    }

    // Resolve a name to an address
    domain := "wealdtech.eth"
    address, err := ens.Resolve(client, domain)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Address of %s is %s\n", domain, address.Hex())

    // Reverse resolve an address to a name
    reverse, err := ens.ReverseResolve(client, address)
    if reverse == "" {
      fmt.Printf("%s has no reverse lookup\n", address.Hex())
    } else {
      fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
    }
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/wealdtech/go-ens/issues).

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd
