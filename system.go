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

var system identity.System

func NewSystemFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	system.CertType = fs.String("cert-type", faker.Pick("", "consumer"), "set the identity.system.cert_type field to `STRING`")
	system.ClusterID = fs.String("cluster-id", faker.Pick("", faker.StringWithSize(8)), "set the identity.system.cluster_id field to `STRING`")
	fs.StringVar(&system.CN, "cn", faker.UUID(), "set the identity.system.cn field to `STRING`")

	return fs
}

var systemCommand = &ffcli.Command{
	Name:       "system",
	ShortUsage: "system [flags]",
	ShortHelp:  "generate a system identity JSON object",
	FlagSet:    NewSystemFlagSet("system", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		id.Identity.System = &system

		data, err := json.Marshal(id)
		if err != nil {
			return fmt.Errorf("cannot marshal data: %w", err)
		}

		fmt.Println(string(data))

		return nil
	},
}
