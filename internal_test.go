package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
	"go.openly.dev/pointy"
)

func TestNewInternal(t *testing.T) {
	type Tests struct {
		description string
		seed        int64
		input       Internal
		want        *identity.Internal
		wantError   error
	}
	tests := []Tests{
		{
			description: "empty template",
			seed:        100,
			input:       Internal{},
			want: &identity.Internal{
				AuthTime:    float32(-1.6924639230312625e+18),
				CrossAccess: true,
				OrgID:       "00229",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: Internal{
				AuthTime: pointy.Float32(1.0),
			},
			want: &identity.Internal{
				AuthTime:    1.0,
				CrossAccess: false,
				OrgID:       "80022",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: Internal{
				AuthTime:    pointy.Float32(2.0),
				CrossAccess: pointy.Bool(true),
				OrgID:       pointy.String("123456"),
			},
			want: &identity.Internal{
				AuthTime:    2.0,
				CrossAccess: true,
				OrgID:       "123456",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)
			got, err := NewInternal(test.input)

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
