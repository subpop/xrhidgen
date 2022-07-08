package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/pioz/faker"
	"github.com/redhatinsights/module-update-router/identity"
)

var id identity.Identity

func main() {
	fs := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)

	id.Identity.AccountNumber = fs.String("account-number", faker.Pick("", faker.DigitsWithSize(6)), "set the identity.account_number field to `NUMBER`")
	fs.StringVar(&id.Identity.AuthType, "auth-type", "", "set the identity.authtype field to `STRING`")
	id.Identity.EmployeeAccountNumber = fs.String("employe-account-number", faker.Pick("", faker.DigitsWithSize(6)), "set the identity.employee_account_number field to `STRING`")
	fs.StringVar(&id.Identity.OrgID, "org-id", faker.DigitsWithSize(5), "set the identity.org_id field to `STRING`")
	id.Identity.Type = fs.String("type", "", "set the identity.type field to `STRING`")

	root := &ffcli.Command{
		ShortUsage: fmt.Sprintf("%v [flags] <subcommand>", fs.Name()),
		FlagSet:    fs,
		Subcommands: []*ffcli.Command{
			userCommand,
			internalCommand,
			systemCommand,
			associateCommand,
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
