package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/pioz/faker"
	"github.com/redhatinsights/module-update-router/identity"
	"github.com/sgreben/flagvar"
)

var associate identity.Associate

var role flagvar.Strings

func NewAssociateFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	fs.StringVar(&associate.Email, "email", faker.Email(), "set the identity.associate.email field to `STRING`")
	fs.StringVar(&associate.GivenName, "givenname", faker.FirstName(), "set the identity.associate.givenName field to `STRING`")
	fs.StringVar(&associate.RHatUUID, "rhatuuid", faker.UUID(), "set the identity.associate.rhatUUID field to `STRING`")
	fs.Var(&role, "role", "set the identity.associate.Role field to `STRING` (can be set multiple times)")
	fs.StringVar(&associate.Surname, "surname", faker.LastName(), "set the identity.associate.surname field to `STRING`")

	return fs
}

var associateCommand = &ffcli.Command{
	Name:       "associate",
	ShortUsage: "associate [flags]",
	ShortHelp:  "generate an associate identity JSON object",
	FlagSet:    NewAssociateFlagSet("associate", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		associate.Role = role.Values
		id.Identity.Associate = &associate

		data, err := json.Marshal(id)
		if err != nil {
			return fmt.Errorf("cannot marshal data: %w", err)
		}

		fmt.Println(string(data))

		return nil
	},
}
