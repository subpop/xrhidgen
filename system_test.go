package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
	"go.openly.dev/pointy"
)

func TestNewSystem(t *testing.T) {
	type Tests struct {
		description string
		seed        int64
		input       System
		want        *identity.System
		wantError   error
	}
	tests := []Tests{
		{
			description: "empty template",
			seed:        100,
			input:       System{},
			want: &identity.System{
				CertType:   "consumer",
				ClusterId:  "i",
				CommonName: "wu7Re",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: System{
				CertType: pointy.String("asdf"),
			},
			want: &identity.System{
				CertType:   "asdf",
				ClusterId:  "gimwu7Re",
				CommonName: "hCRM50",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: System{
				CertType:  pointy.String("consumer"),
				ClusterID: pointy.String("1234"),
				CN:        pointy.String("xyz"),
			},
			want: &identity.System{
				CertType:   "consumer",
				ClusterId:  "1234",
				CommonName: "xyz",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)

			got, err := NewSystem(test.input)

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
