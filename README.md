[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/rsjethani/secret/v2)
[![Build Status](https://github.com/rsjethani/secret/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/rsjethani/secret/actions)

# secret v2

### What secret is?
- It provides simple Go types like [secret.Text](https://pkg.go.dev/github.com/rsjethani/secret/v2#Text) to encapsulate your secret. Example:
```
type Login struct {
    User string
    Password secret.Text
}
```
See godev reference for more examples and usage information.
- The encapsulated secret remains inaccessible to operations like printing, logging, and JSON serializtion, a redact hint like `*****` is returned instead.
- The only way to access the actual secret value is by asking explicitly via the `.Value()` method.

### What secret is not?
- It is not a secret management service or your local password manager.
- It is not a Go client to facilitate communication with secret managers like Hashicorp Vault, AWS secret Manager etc. Checkout [teller](https://github.com/spectralops/teller) if that is what you are looking for.

### Installation
```
go get github.com/rsjethani/secret/v2
```
NOTE: v1 is deprectated now.

