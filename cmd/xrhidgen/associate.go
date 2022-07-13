package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/sgreben/flagvar"
	"github.com/subpop/xrhidgen"
)

var associateFlags struct {
	email     StringFlag
	givenName StringFlag
	rhatUUID  StringFlag
	role      flagvar.Strings
	surName   StringFlag
}

func NewAssociateFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	fs.Var(&associateFlags.email, "email", "set the identity.associate.email field (string)")
	fs.Var(&associateFlags.givenName, "givenname", "set the identity.associate.givenName field (string)")
	fs.Var(&associateFlags.rhatUUID, "rhatuuid", "set the identity.associate.rhatUUID field (string)")
	fs.Var(&associateFlags.role, "role", "set the identity.associate.Role field (string) (can be set multiple times)")
	fs.Var(&associateFlags.surName, "surname", "set the identity.associate.surname field (string)")

	return fs
}

var associateCommand = &ffcli.Command{
	Name:       "associate",
	ShortUsage: "associate [flags]",
	ShortHelp:  "generate an associate identity JSON record",
	LongHelp:   WordWrap("Generate an associate identity record, populating fields with values provided by the matching flag. Any omitted flags will have their corresponding fields populated with a suitable random value.", 72),
	FlagSet:    NewAssociateFlagSet("associate", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		associate := xrhidgen.Associate{
			Email:     associateFlags.email.Value,
			GivenName: associateFlags.givenName.Value,
			RHatUUID:  associateFlags.rhatUUID.Value,
			Role:      &associateFlags.role.Values,
			Surname:   associateFlags.surName.Value,
		}

		id, err := xrhidgen.NewAssociateIdentity(mainIdentity(), associate)
		if err != nil {
			return err
		}

		data, err := json.Marshal(id)
		if err != nil {
			return fmt.Errorf("cannot marshal data: %w", err)
		}

		fmt.Println(string(data))

		return nil
	},
}
