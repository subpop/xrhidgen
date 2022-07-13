package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/subpop/xrhidgen"
)

var userFlags struct {
	email      StringFlag
	firstName  StringFlag
	isActive   BoolFlag
	isInternal BoolFlag
	isOrgAdmin BoolFlag
	lastName   StringFlag
	locale     StringFlag
	userID     StringFlag
	username   StringFlag
}

func NewUserFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	fs.Var(&userFlags.email, "email", "set the identity.user.email field (string)")
	fs.Var(&userFlags.firstName, "first-name", "set the identity.user.first_name field (string)")
	fs.Var(&userFlags.isActive, "is-active", "set the identity.user.is_active field (bool)")
	fs.Var(&userFlags.isInternal, "is-internal", "set the identity.user.is_internal field (bool)")
	fs.Var(&userFlags.isOrgAdmin, "is-org-admin", "set the identity.user.is_org_admin field (bool)")
	fs.Var(&userFlags.lastName, "lastname", "set the identity.user.last_name field (string)")
	fs.Var(&userFlags.locale, "locale", "set the identity.user.locale field (string)")
	fs.Var(&userFlags.userID, "user-id", "set the identity.user.user_id field (string)")
	fs.Var(&userFlags.username, "username", "set the identity.user.username field (string)")

	return fs
}

var userCommand = &ffcli.Command{
	Name:       "user",
	ShortUsage: "user [flags]",
	ShortHelp:  "generate a user identity JSON record",
	LongHelp:   WordWrap("Generate a user identity record, populating fields with values provided by the matching flag. Any omitted flags will have their corresponding fields populated with a suitable random value.", 72),
	FlagSet:    NewUserFlagSet("user", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		user := xrhidgen.User{
			Email:      userFlags.email.Value,
			FirstName:  userFlags.firstName.Value,
			IsActive:   userFlags.isActive.Value,
			IsInternal: userFlags.isInternal.Value,
			IsOrgAdmin: userFlags.isOrgAdmin.Value,
			LastName:   userFlags.lastName.Value,
			Locale:     userFlags.locale.Value,
			UserID:     userFlags.userID.Value,
			Username:   userFlags.username.Value,
		}

		id, err := xrhidgen.NewUserIdentity(mainIdentity(), user)
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
