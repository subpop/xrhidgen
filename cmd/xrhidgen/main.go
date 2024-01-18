// xrhidgen generates X-Rh-Identity records.
//
//	USAGE
//	  xrhidgen [flags] <subcommand>
//
//	xrhidgen can be used to generate JSON records suitable for passing in to
//	the X-Rh-Identity header. Each subcommand will generate a record of the
//	specified type. Any flag set will be inserted instead of a random value.
//	All remaining fields will be filled with a suitably random value.
//
//	SUBCOMMANDS
//	  user             generate a user identity JSON record
//	  internal         generate an internal identity JSON record
//	  system           generate a system identity JSON record
//	  associate        generate an associate identity JSON record
//	  service-account  generate a service account identity JSON record
//
//	FLAGS
//	  -account-number value          set the identity.account_number field (string)
//	  -auth-type value               set the identity.authtype field (string)
//	  -employe-account-number value  set the identity.employee_account_number field (string)
//	  -org-id value                  set the identity.org_id field (string)
//	  -type value                    set the identity.type field (string)
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/subpop/xrhidgen"
)

var mainFlags struct {
	accountNumber         StringFlag
	authType              StringFlag
	employeeAccountNumber StringFlag
	orgID                 StringFlag
	Type                  StringFlag
}

func main() {
	fs := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	fs.Var(&mainFlags.accountNumber, "account-number", "set the identity.account_number field (string)")
	fs.Var(&mainFlags.authType, "auth-type", "set the identity.authtype field (string)")
	fs.Var(&mainFlags.employeeAccountNumber, "employe-account-number", "set the identity.employee_account_number field (string)")
	fs.Var(&mainFlags.orgID, "org-id", "set the identity.org_id field (string)")
	fs.Var(&mainFlags.Type, "type", "set the identity.type field (string)")

	root := &ffcli.Command{
		ShortUsage: fmt.Sprintf("%v [flags] <subcommand>", fs.Name()),
		LongHelp:   WordWrap("xrhidgen can be used to generate JSON records suitable for passing in to the X-Rh-Identity header. Each subcommand will generate a record of the specified type. Any flag set will be inserted instead of a random value. All remaining fields will be filled with a suitably random value.", 72),
		FlagSet:    fs,
		Subcommands: []*ffcli.Command{
			userCommand,
			internalCommand,
			systemCommand,
			associateCommand,
			serviceAccountCommand,
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	if val, has := os.LookupEnv("SEED"); has {
		seed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		xrhidgen.SetSeed(seed)
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func mainIdentity() xrhidgen.Identity {
	return xrhidgen.Identity{
		AccountNumber:         mainFlags.accountNumber.Value,
		AuthType:              mainFlags.authType.Value,
		EmployeeAccountNumber: mainFlags.employeeAccountNumber.Value,
		OrgID:                 mainFlags.orgID.Value,
		Type:                  mainFlags.Type.Value,
	}
}
