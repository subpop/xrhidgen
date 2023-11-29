package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
	"go.openly.dev/pointy"
)

func TestNewIdentity(t *testing.T) {
	type Tests struct {
		description string
		seed        int64
		input       Identity
		want        *identity.XRHID
		wantError   error
	}
	tests := []Tests{
		{
			description: "empty template",
			seed:        100,
			input:       Identity{},
			want: &identity.XRHID{
				Identity: identity.Identity{
					AccountNumber:         "",
					Associate:             identity.Associate{},
					AuthType:              "basic-auth",
					EmployeeAccountNumber: "02299",
					Internal:              identity.Internal{},
					OrgID:                 "41123",
					System:                identity.System{},
					Type:                  "50cQB",
					User:                  identity.User{},
					X509:                  identity.X509{},
				},
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: Identity{
				AccountNumber: pointy.String("1234"),
			},
			want: &identity.XRHID{
				Identity: identity.Identity{
					AccountNumber:         "1234",
					Associate:             identity.Associate{},
					AuthType:              "cert-auth",
					EmployeeAccountNumber: "00229",
					Internal:              identity.Internal{},
					OrgID:                 "94112",
					System:                identity.System{},
					Type:                  "M5",
					User:                  identity.User{},
					X509:                  identity.X509{},
				},
			},
		},
		{
			description: "full template",
			seed:        100,
			input: Identity{
				AccountNumber:         pointy.String("10001"),
				AuthType:              pointy.String("basic-auth"),
				EmployeeAccountNumber: pointy.String("112233"),
				OrgID:                 pointy.String("111111"),
				Type:                  pointy.String("universal"),
			},
			want: &identity.XRHID{
				Identity: identity.Identity{
					AccountNumber:         "10001",
					Associate:             identity.Associate{},
					AuthType:              "basic-auth",
					EmployeeAccountNumber: "112233",
					Internal:              identity.Internal{},
					OrgID:                 "111111",
					System:                identity.System{},
					Type:                  "universal",
					User:                  identity.User{},
					X509:                  identity.X509{},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)
			got, err := NewIdentity(test.input)

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
