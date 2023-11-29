package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
	"go.openly.dev/pointy"
)

func TestNewX509(t *testing.T) {
	type Test struct {
		description string
		seed        int64
		input       X509
		want        *identity.X509
		wantError   error
	}
	tests := []Test{
		{
			description: "empty template",
			seed:        100,
			input:       X509{},
			want: &identity.X509{
				SubjectDN: "gimwu7Re",
				IssuerDN:  "hCRM50",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: X509{
				SubjectDN: pointy.String("CN = d6cde789-fbad-45ea-a542-30ba779aa870"),
			},
			want: &identity.X509{
				SubjectDN: "CN = d6cde789-fbad-45ea-a542-30ba779aa870",
				IssuerDN:  "gimwu7Re",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: X509{
				SubjectDN: pointy.String("CN = d6cde789-fbad-45ea-a542-30ba779aa870"),
				IssuerDN:  pointy.String("O = Foo, Inc."),
			},
			want: &identity.X509{
				SubjectDN: "CN = d6cde789-fbad-45ea-a542-30ba779aa870",
				IssuerDN:  "O = Foo, Inc.",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)

			got, err := NewX509(test.input)

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
