[![Go Reference](https://pkg.go.dev/badge/github.com/subpop/xrhidgen.svg)](https://pkg.go.dev/github.com/subpop/xrhidgen)

# xrhidgen

`xrhidgen` generates X-Rh-Identity JSON records suitable for passing into HTTP
requests to console.redhat.com services. Any field not explicitly set via a
command line flag will be populated by an appropriate random value.

## Command Line

### Installation

```
go install github.com/subpop/xrhidgen/cmd/xrhidgen@latest
```

### Usage

```
USAGE
  xrhidgen [flags] <subcommand>

xrhidgen can be used to generate JSON records suitable for passing in to
the X-Rh-Identity header. Each subcommand will generate a record of the
specified type. Any flag set will be inserted instead of a random value.
All remaining fields will be filled with a suitably random value.

SUBCOMMANDS
  user             generate a user identity JSON record
  internal         generate an internal identity JSON record
  system           generate a system identity JSON record
  associate        generate an associate identity JSON record
  service-account  generate a service account identity JSON record

FLAGS
  -account-number value          set the identity.account_number field (string)
  -auth-type value               set the identity.authtype field (string)
  -employe-account-number value  set the identity.employee_account_number field (string)
  -org-id value                  set the identity.org_id field (string)
  -type value                    set the identity.type field (string)
```

### Examples

```
$ xrhidgen user -email someuser@redhat.com
{"identity":{"org_id":"45841","internal":{"org_id":"45841"},"user":{"username":"backbencher","email":"someuser@redhat.com","first_name":"Cira","last_name":"Roob","is_active":false,"is_org_admin":true,"is_internal":false,"locale":"li","user_id":"nicole"},"system":{},"associate":{"Role":null,"email":"","givenName":"","rhatUUID":"","surname":""},"x509":{"subject_dn":"","issuer_dn":""},"service_account":{"client_id":"","username":""},"type":"User","auth_type":"cert-auth"},"entitlements":null}
```

```
$ xrhidgen system | base64 -w0
eyJpZGVudGl0eSI6eyJvcmdfaWQiOiIyMTcwMyIsImludGVybmFsIjp7Im9yZ19pZCI6IjIxNzAzIn0sInVzZXIiOnsidXNlcm5hbWUiOiIiLCJlbWFpbCI6IiIsImZpcnN0X25hbWUiOiIiLCJsYXN0X25hbWUiOiIiLCJpc19hY3RpdmUiOmZhbHNlLCJpc19vcmdfYWRtaW4iOmZhbHNlLCJpc19pbnRlcm5hbCI6ZmFsc2UsImxvY2FsZSI6IiIsInVzZXJfaWQiOiIifSwic3lzdGVtIjp7ImNuIjoiRyIsImNlcnRfdHlwZSI6ImNvbnN1bWVyIiwiY2x1c3Rlcl9pZCI6IkhKbyJ9LCJhc3NvY2lhdGUiOnsiUm9sZSI6bnVsbCwiZW1haWwiOiIiLCJnaXZlbk5hbWUiOiIiLCJyaGF0VVVJRCI6IiIsInN1cm5hbWUiOiIifSwieDUwOSI6eyJzdWJqZWN0X2RuIjoiIiwiaXNzdWVyX2RuIjoiIn0sInNlcnZpY2VfYWNjb3VudCI6eyJjbGllbnRfaWQiOiIiLCJ1c2VybmFtZSI6IiJ9LCJ0eXBlIjoiU3lzdGVtIiwiYXV0aF90eXBlIjoiYmFzaWMtYXV0aCJ9LCJlbnRpdGxlbWVudHMiOm51bGx9Cg==
```

```
ht GET http://localhost:8080/api/module-update-router/v1/channel?module=insights-core "X-Rh-Identity: $(xrhidgen system | base64 -w0)"
```

The `SEED` environment variable can be set to an integer. If set, it will be
used to initialize the generator to a deterministic state.

```
$ SEED=100 xrhidgen user
{"identity":{"employee_account_number":"02299","org_id":"41123","internal":{"org_id":"41123"},"user":{"username":"skeptic","email":"winnifredwinning@shred.org","first_name":"Cameron","last_name":"Stehr","is_active":false,"is_org_admin":false,"is_internal":false,"locale":"fi","user_id":"meredeth"},"system":{},"associate":{"Role":null,"email":"","givenName":"","rhatUUID":"","surname":""},"x509":{"subject_dn":"","issuer_dn":""},"service_account":{"client_id":"","username":""},"type":"User","auth_type":"basic-auth"},"entitlements":null}
```

## Go package

### Installation

```
go get github.com/subpop/xrhidgen/cmd/xrhidgen@latest
```

### Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/subpop/xrhidgen"
)

func main() {
	xrhidgen.SetSeed(103)
	id, err := xrhidgen.NewUserIdentity(xrhidgen.Identity{}, xrhidgen.User{})
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
```
