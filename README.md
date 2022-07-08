`xrhidgen` generates X-Rh-Identity JSON objects suitable for passing into HTTP
requests to console.redhat.com services. Any field not explicitly set via a
command line flag will be populated by an appropriate random value.

# Installation

```
go install github.com/subpop/xrhidgen@latest
```

# Usage

```
USAGE
  xrhidgen [flags] <subcommand>

SUBCOMMANDS
  user       generate a user identity JSON object
  internal   generate an internal identity JSON object
  system     generate a system identity JSON object
  associate  generate an associate identity JSON object

FLAGS
  -account-number ...             set the identity.account_number field to `NUMBER`
  -auth-type ...                  set the identity.authtype field to `STRING`
  -employe-account-number ...     set the identity.employee_account_number field to `STRING`
  -org-id ...                     set the identity.org_id field to `STRING`
  -type ...                       set the identity.type field to `STRING`
```

# Examples

```
$ xrhidgen user -email someuser@redhat.com
{"identity":{"account_number":"666612","employee_account_number":"817871","org_id":"82969","type":"","user":{"email":"someuser@redhat.com","first_name":"Ari","is_active":true,"is_internal":false,"is_org_admin":false,"last_name":"Sipes","locale":"kl","user_id":"ammerman","username":"riyadh"}}}
```

```
$ xrhidgen system | base64 -w0
eyJpZGVudGl0eSI6eyJhY2NvdW50X251bWJlciI6IjQ0NDY4OCIsImVtcGxveWVlX2FjY291bnRfbnVtYmVyIjoiIiwib3JnX2lkIjoiODEzNTIiLCJzeXN0ZW0iOnsiY2VydF90eXBlIjoiIiwiY2x1c3Rlcl9pZCI6ImNjWWJhTllCIiwiY24iOiJhYzRlM2RmYy1kOGU3LTQwODUtYjg3YS0zMTcyZjU1M2I3M2UifSwidHlwZSI6IiJ9fQo=
```

```
ht GET http://localhost:8080/api/module-update-router/v1/channel?module=insights-core "X-Rh-Identity: $(xrhidgen system | base64 -w0)"
```
