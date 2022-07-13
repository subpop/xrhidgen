package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/subpop/xrhidgen"
)

var internalFlags struct {
	authType    Float64Flag
	crossAccess BoolFlag
	orgID       StringFlag
}

func NewInternalFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	fs.Var(&internalFlags.authType, "auth-time", "set the identity.internal.auth_time field (float)")
	fs.Var(&internalFlags.crossAccess, "cross-access", "set the identity.internal.cross_access field (bool)")
	fs.Var(&internalFlags.orgID, "org-id", "set the identity.internal.org_id field (string)")

	return fs
}

var internalCommand = &ffcli.Command{
	Name:       "internal",
	ShortUsage: "internal [flags]",
	ShortHelp:  "generate an internal identity JSON record",
	LongHelp:   WordWrap("Generate an internal identity record, populating fields with values provided by the matching flag. Any omitted flags will have their corresponding fields populated with a suitable random value.", 72),
	FlagSet:    NewInternalFlagSet("internal", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		internal := xrhidgen.Internal{
			AuthTime:    internalFlags.authType.Value,
			CrossAccess: internalFlags.crossAccess.Value,
			OrgID:       internalFlags.orgID.Value,
		}

		id, err := xrhidgen.NewInternalIdentity(mainIdentity(), internal)
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
