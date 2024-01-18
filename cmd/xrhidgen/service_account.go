package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/subpop/xrhidgen"
)

var serviceAccountFlags struct {
	clientID StringFlag
	username StringFlag
}

func NewServiceAccountFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	fs.Var(&serviceAccountFlags.clientID, "client-id", "set the identity.service_account.client_id field (string)")
	fs.Var(&serviceAccountFlags.username, "username", "set the identity.service_account.username field (string)")

	return fs
}

var serviceAccountCommand = &ffcli.Command{
	Name:       "service-account",
	ShortUsage: "service-account [flags]",
	ShortHelp:  "generate a service account identity JSON record",
	LongHelp:   WordWrap("Generate a service account identity record, populating fields with values provided by the matching flag. Any omitted flags will have their corresponding fields populated with a suitable random value.", 72),
	FlagSet:    NewServiceAccountFlagSet("service-account", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		serviceAccount := xrhidgen.ServiceAccount{
			ClientID: serviceAccountFlags.clientID.Value,
			Username: serviceAccountFlags.username.Value,
		}

		id, err := xrhidgen.NewServiceAccountIdentity(mainIdentity(), serviceAccount)
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
