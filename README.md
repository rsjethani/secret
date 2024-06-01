[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/rsjethani/secret/v2)
[![Build Status](https://github.com/rsjethani/secret/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/rsjethani/secret/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/rsjethani/secret)](https://goreportcard.com/report/github.com/rsjethani/secret)

# Installation
```
go get github.com/rsjethani/secret/v3
```

# What secret is?
It provides simple Go types like `Text` to securely store secrets. For example:
```go
type Login struct {
    User string
    Password secret.Text
}
```
The encapsulated secret remains inaccessible to operations like printing, logging, JSON serialization etc. A (customizable) redact hint like `*****` is returned instead. The only way to access the actual secret value is by asking explicitly via the `.Secret()` method. See package documentation more reference and examples.

# What secret is not?
- It is not a secret management service or your local password manager.
- It is not a Go client to facilitate communication with secret managers like Hashicorp Vault, AWS secret Manager etc. Checkout [teller](https://github.com/spectralops/teller) if that is what you are looking for.

# Versions

### v3
Current version, same functionality as v2 but much cleaner API.

### v2
Perfectly usable but no longer maintained.

### v1
Deprecated.