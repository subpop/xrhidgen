package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/pioz/faker"
	"github.com/redhatinsights/module-update-router/identity"
)

var internal identity.Internal

func NewInternalFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	internal.AuthTime = fs.Float64("auth-time", float64(faker.Duration()), "set the identity.internal.auth_time field to `FLOAT`")
	internal.CrossAccess = fs.Bool("cross-access", faker.Bool(), "set the identity.internal.cross_access field to `BOOL`")
	fs.StringVar(&internal.OrgID, "org-id", faker.DigitsWithSize(5), "set the identity.internal.org_id field to `STRING`")

	return fs
}

var internalCommand = &ffcli.Command{
	Name:       "internal",
	ShortUsage: "internal [flags]",
	ShortHelp:  "generate an internal identity JSON object",
	FlagSet:    NewInternalFlagSet("internal", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		id.Identity.Internal = &internal

		data, err := json.Marshal(id)
		if err != nil {
			return fmt.Errorf("cannot marshal data: %w", err)
		}

		fmt.Println(string(data))

		return nil
	},
}
