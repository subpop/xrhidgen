package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/v2/identity"
	"go.openly.dev/pointy"
)

func TestNewServiceAccount(t *testing.T) {
	type Tests struct {
		description string
		seed        int64
		input       ServiceAccount
		want        *identity.ServiceAccount
		wantError   error
	}
	tests := []Tests{
		{
			description: "empty template",
			seed:        100,
			input:       ServiceAccount{},
			want: &identity.ServiceAccount{
				ClientId: "gimwu7Re",
				Username: "boltrope",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: ServiceAccount{
				ClientID: pointy.String("12345"),
			},
			want: &identity.ServiceAccount{
				ClientId: "12345",
				Username: "ammon",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: ServiceAccount{
				ClientID: pointy.String("12345"),
				Username: pointy.String("jsmith"),
			},
			want: &identity.ServiceAccount{
				ClientId: "12345",
				Username: "jsmith",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)

			got, err := NewServiceAccount(test.input)

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
