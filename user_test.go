package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/module-update-router/identity"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		description string
		seed        int64
		input       User
		want        *identity.User
		wantError   error
	}{
		{
			description: "empty template",
			seed:        100,
			input:       User{},
			want: &identity.User{
				Email:      "ammon@benison.biz",
				FirstName:  "Augustine",
				IsActive:   true,
				IsInternal: false,
				IsOrgAdmin: false,
				LastName:   "Stracke",
				Locale:     "fo",
				UserID:     "salify",
				Username:   "platas",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: User{
				Email:     ptrstring("jsmith@redhat.com"),
				FirstName: ptrstring("John"),
				LastName:  ptrstring("Smith"),
			},
			want: &identity.User{
				Email:      "jsmith@redhat.com",
				FirstName:  "John",
				IsActive:   false,
				IsInternal: true,
				IsOrgAdmin: true,
				LastName:   "Smith",
				Locale:     "ik",
				UserID:     "goniometer",
				Username:   "waldner",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: User{
				Email:      ptrstring("jsmith@redhat.com"),
				FirstName:  ptrstring("John"),
				IsActive:   ptrbool(true),
				IsInternal: ptrbool(true),
				IsOrgAdmin: ptrbool(false),
				LastName:   ptrstring("Smith"),
				Locale:     ptrstring("en"),
				UserID:     ptrstring("jsmith1"),
				Username:   ptrstring("jsm"),
			},
			want: &identity.User{
				Email:      "jsmith@redhat.com",
				FirstName:  "John",
				IsActive:   true,
				IsInternal: true,
				IsOrgAdmin: false,
				LastName:   "Smith",
				Locale:     "en",
				UserID:     "jsmith1",
				Username:   "jsm",
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
