`xrhidgen` generates X-Rh-Identity JSON records suitable for passing into HTTP
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

xrhidgen can be used to generate JSON records suitable for passing in to
the X-Rh-Identity header. Each subcommand will generate a record of the
specified type. Any flag set will be inserted instead of a random value.
All remaining fields will be filled with a suitably random value.

SUBCOMMANDS
  user       generate a user identity JSON record
  internal   generate an internal identity JSON record
  system     generate a system identity JSON record
  associate  generate an associate identity JSON record

FLAGS
  -account-number ...          set the identity.account_number field (string)
  -auth-type ...               set the identity.authtype field (string)
  -employe-account-number ...  set the identity.employee_account_number field (string)
  -org-id ...                  set the identity.org_id field (string)
  -type ...                    set the identity.type field (string)
```

# Examples

```
$ xrhidgen user -email someuser@redhat.com
{"identity":{"account_number":"16349","auth_type":"cert","employee_account_number":"06900","internal":{"org_id":"51818"},"org_id":"51818","type":"User","user":{"email":"someuser@redhat.com","first_name":"Quinn","is_active":true,"is_internal":true,"is_org_admin":true,"last_name":"Runolfsdottir","locale":"se","user_id":"taps","username":"dunstable"}}}
```

```
$ xrhidgen system | base64 -w0
eyJpZGVudGl0eSI6eyJhY2NvdW50X251bWJlciI6IjQ0NDY4OCIsImVtcGxveWVlX2FjY291bnRfbnVtYmVyIjoiIiwib3JnX2lkIjoiODEzNTIiLCJzeXN0ZW0iOnsiY2VydF90eXBlIjoiIiwiY2x1c3Rlcl9pZCI6ImNjWWJhTllCIiwiY24iOiJhYzRlM2RmYy1kOGU3LTQwODUtYjg3YS0zMTcyZjU1M2I3M2UifSwidHlwZSI6IiJ9fQo=
```

```
ht GET http://localhost:8080/api/module-update-router/v1/channel?module=insights-core "X-Rh-Identity: $(xrhidgen system | base64 -w0)"
```
