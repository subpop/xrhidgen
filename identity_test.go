package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

func TestNewIdentity(t *testing.T) {
	tests := []struct {
		description string
		seed        int64
		input       Identity
		want        *identity.Identity
		wantError   error
	}{
		{
			description: "empty template",
			seed:        100,
			input:       Identity{},
			want: &identity.Identity{
				Identity: struct {
					AccountNumber         *string             "json:\"account_number,omitempty\""
					Associate             *identity.Associate "json:\"associate,omitempty\""
					AuthType              string              "json:\"auth_type,omitempty\""
					EmployeeAccountNumber *string             "json:\"employee_account_number,omitempty\""
					Internal              *identity.Internal  "json:\"internal,omitempty\""
					OrgID                 string              "json:\"org_id\""
					System                *identity.System    "json:\"system,omitempty\""
					Type                  *string             "json:\"type,omitempty\""
					User                  *identity.User      "json:\"user,omitempty\""
					X509                  *identity.X509      "json:\"x509,omitempty\""
				}{
					AccountNumber:         nil,
					Associate:             nil,
					AuthType:              "basic-auth",
					EmployeeAccountNumber: ptrstring("02299"),
					Internal:              nil,
					OrgID:                 "41123",
					System:                nil,
					Type:                  ptrstring("50cQB"),
					User:                  nil,
					X509:                  nil,
				},
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: Identity{
				AccountNumber: ptrstring("1234"),
			},
			want: &identity.Identity{
				Identity: struct {
					AccountNumber         *string             "json:\"account_number,omitempty\""
					Associate             *identity.Associate "json:\"associate,omitempty\""
					AuthType              string              "json:\"auth_type,omitempty\""
					EmployeeAccountNumber *string             "json:\"employee_account_number,omitempty\""
					Internal              *identity.Internal  "json:\"internal,omitempty\""
					OrgID                 string              "json:\"org_id\""
					System                *identity.System    "json:\"system,omitempty\""
					Type                  *string             "json:\"type,omitempty\""
					User                  *identity.User      "json:\"user,omitempty\""
					X509                  *identity.X509      "json:\"x509,omitempty\""
				}{
					AccountNumber:         ptrstring("1234"),
					Associate:             nil,
					AuthType:              "cert-auth",
					EmployeeAccountNumber: ptrstring("00229"),
					Internal:              nil,
					OrgID:                 "94112",
					System:                nil,
					Type:                  ptrstring("M5"),
					User:                  nil,
					X509:                  nil,
				},
			},
		},
		{
			description: "full template",
			seed:        100,
			input: Identity{
				AccountNumber:         ptrstring("10001"),
				AuthType:              ptrstring("basic-auth"),
				EmployeeAccountNumber: ptrstring("112233"),
				OrgID:                 ptrstring("111111"),
				Type:                  ptrstring("universal"),
			},
			want: &identity.Identity{
				Identity: struct {
					AccountNumber         *string             "json:\"account_number,omitempty\""
					Associate             *identity.Associate "json:\"associate,omitempty\""
					AuthType              string              "json:\"auth_type,omitempty\""
					EmployeeAccountNumber *string             "json:\"employee_account_number,omitempty\""
					Internal              *identity.Internal  "json:\"internal,omitempty\""
					OrgID                 string              "json:\"org_id\""
					System                *identity.System    "json:\"system,omitempty\""
					Type                  *string             "json:\"type,omitempty\""
					User                  *identity.User      "json:\"user,omitempty\""
					X509                  *identity.X509      "json:\"x509,omitempty\""
				}{
					AccountNumber:         ptrstring("10001"),
					Associate:             nil,
					AuthType:              "basic-auth",
					EmployeeAccountNumber: ptrstring("112233"),
					Internal:              nil,
					OrgID:                 "111111",
					System:                nil,
					Type:                  ptrstring("universal"),
					User:                  nil,
					X509:                  nil,
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
