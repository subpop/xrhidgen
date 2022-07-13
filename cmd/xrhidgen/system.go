package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/subpop/xrhidgen"
)

var systemFlags struct {
	certType  StringFlag
	clusterID StringFlag
	cn        StringFlag
}

func NewSystemFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	fs.Var(&systemFlags.certType, "cert-type", "set the identity.system.cert_type field (string)")
	fs.Var(&systemFlags.clusterID, "cluster-id", "set the identity.system.cluster_id field (string)")
	fs.Var(&systemFlags.cn, "cn", "set the identity.system.cn field (string)")

	return fs
}

var systemCommand = &ffcli.Command{
	Name:       "system",
	ShortUsage: "system [flags]",
	ShortHelp:  "generate a system identity JSON record",
	LongHelp:   WordWrap("Generate a system identity record, populating fields with values provided by the matching flag. Any omitted flags will have their corresponding fields populated with a suitable random value.", 72),
	FlagSet:    NewSystemFlagSet("system", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		system := xrhidgen.System{
			CertType:  systemFlags.certType.Value,
			ClusterID: systemFlags.clusterID.Value,
			CN:        systemFlags.cn.Value,
		}

		id, err := xrhidgen.NewSystemIdentity(mainIdentity(), system)
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
