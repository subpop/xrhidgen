package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
	"go.openly.dev/pointy"
)

func TestNewUser(t *testing.T) {
	type Tests struct {
		description string
		seed        int64
		input       User
		want        *identity.User
		wantError   error
	}
	tests := []Tests{
		{
			description: "empty template",
			seed:        100,
			input:       User{},
			want: &identity.User{
				Email:     "ammon@benison.biz",
				FirstName: "Augustine",
				Active:    true,
				Internal:  false,
				OrgAdmin:  false,
				LastName:  "Stracke",
				Locale:    "fo",
				UserID:    "salify",
				Username:  "platas",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: User{
				Email:     pointy.String("jsmith@redhat.com"),
				FirstName: pointy.String("John"),
				LastName:  pointy.String("Smith"),
			},
			want: &identity.User{
				Email:     "jsmith@redhat.com",
				FirstName: "John",
				Active:    false,
				Internal:  true,
				OrgAdmin:  true,
				LastName:  "Smith",
				Locale:    "ik",
				UserID:    "goniometer",
				Username:  "waldner",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: User{
				Email:      pointy.String("jsmith@redhat.com"),
				FirstName:  pointy.String("John"),
				IsActive:   pointy.Bool(true),
				IsInternal: pointy.Bool(true),
				IsOrgAdmin: pointy.Bool(false),
				LastName:   pointy.String("Smith"),
				Locale:     pointy.String("en"),
				UserID:     pointy.String("jsmith1"),
				Username:   pointy.String("jsm"),
			},
			want: &identity.User{
				Email:     "jsmith@redhat.com",
				FirstName: "John",
				Active:    true,
				Internal:  true,
				OrgAdmin:  false,
				LastName:  "Smith",
				Locale:    "en",
				UserID:    "jsmith1",
				Username:  "jsm",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)

			got, err := NewUser(test.input)

			if test.wantError != nil {
				if !cmp.Equal(err, test.wantError, cmpopts.EquateErrors()) {
					t.Errorf("%#v != %#v", err, test.wantError)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if !cmp.Equal(got, test.want) {
					t.Errorf("%v", cmp.Diff(got, test.want))
				}
			}
		})
	}
}
