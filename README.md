[![Coverage](https://gocover.io/_badge/github.com/rsjethani/secret)](https://gocover.io/github.com/rsjethani/secret) [![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/rsjethani/secret/v2)
# secret v2

## What secret is?
- It provides simple Go types like [secret.Text](https://pkg.go.dev/github.com/rsjethani/secret/v2#Text) to encapsulate your secret.
- The encapsulated secret remains inaccessible to operations like printing, logging, and JSON serializtion, a redact hint like `*****` is returned instead.
- The only way to access the actual secret value is by asking explicitly via the `.Value()` method.

## What secret is not?
- It is not a secret management service or your local password manager.
- It is not a Go client to facilitate communication with third party secret managers like Hashicorp's Vault, AWS secret Manager etc. Checkout [teller](https://github.com/spectralops/teller) if that is what you are looking for.

## Installation
```
go get github.com/rsjethani/secret/v2
```
NOTE: v1 is deprectated now.

## Usage
See godoc reference for usage examples.

