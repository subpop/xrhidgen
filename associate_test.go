package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/module-update-router/identity"
)

func TestNewAssociate(t *testing.T) {
	tests := []struct {
		description string
		seed        int64
		input       Associate
		want        *identity.Associate
		wantError   error
	}{
		{
			description: "empty template",
			seed:        100,
			input:       Associate{},
			want: &identity.Associate{
				Email:     "ammon@benison.biz",
				GivenName: "Augustine",
				RHatUUID:  "18af1019-8c37-449c-89a3-51fdb7504295",
				Role:      nil,
				Surname:   "Haag",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: Associate{
				Email:     ptrstring("jsmith@redhat.com"),
				GivenName: ptrstring("John"),
				Surname:   ptrstring("Smith"),
			},
			want: &identity.Associate{
				Email:     "jsmith@redhat.com",
				GivenName: "John",
				RHatUUID:  "d2ba8e70-0729-4338-8203-c438d4e94bf3",
				Role:      nil,
				Surname:   "Smith",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: Associate{
				Email:     ptrstring("jsmith@redhat.com"),
				GivenName: ptrstring("John"),
				RHatUUID:  ptrstring("6208f878-b405-498e-979f-e85cd68d8a18"),
				Role:      ptrslicestring([]string{"a", "b"}),
				Surname:   ptrstring("Smith"),
			},
			want: &identity.Associate{
				Email:     "jsmith@redhat.com",
				GivenName: "John",
				RHatUUID:  "6208f878-b405-498e-979f-e85cd68d8a18",
				Role:      []string{"a", "b"},
				Surname:   "Smith",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)

			got, err := NewAssociate(test.input)

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
