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

var user identity.User

func NewUserFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fs := flag.NewFlagSet(name, errorHandling)

	fs.StringVar(&user.Email, "email", faker.Email(), "set the identity.user.email field to `STRING`")
	fs.StringVar(&user.FirstName, "first-name", faker.FirstName(), "set the identity.user.first_name field to `STRING`")
	fs.BoolVar(&user.IsActive, "is-active", faker.Bool(), "set the identity.user.is_active field to `BOOL`")
	fs.BoolVar(&user.IsInternal, "is-internal", faker.Bool(), "set the identity.user.is_internal field to `BOOL`")
	fs.BoolVar(&user.IsOrgAdmin, "is-org-admin", faker.Bool(), "set the identity.user.is_org_admin field to `BOOL`")
	fs.StringVar(&user.LastName, "lastname", faker.LastName(), "set the identity.user.last_name field to `STRING`")
	fs.StringVar(&user.Locale, "locale", faker.LangCode(), "set the identity.user.locale field to `STRING`")
	fs.StringVar(&user.UserID, "user-id", faker.Username(), "set the identity.user.user_id field to `STRING`")
	fs.StringVar(&user.Username, "username", faker.Username(), "set the identity.user.username field to `STRING`")

	return fs
}

var userCommand = &ffcli.Command{
	Name:       "user",
	ShortUsage: "user [flags]",
	ShortHelp:  "generate a user identity JSON object",
	FlagSet:    NewUserFlagSet("user", flag.ExitOnError),
	Exec: func(ctx context.Context, args []string) error {
		identityType := "User"

		id.Identity.Type = &identityType
		id.Identity.User = &user
		id.Identity.Internal = &identity.Internal{
			OrgID: id.Identity.OrgID,
		}

		data, err := json.Marshal(id)
		if err != nil {
			return fmt.Errorf("cannot marshal data: %w", err)
		}

		fmt.Println(string(data))

		return nil
	},
}
