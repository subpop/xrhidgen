package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/v2/identity"
	"go.openly.dev/pointy"
)

func TestNewEntitlements(t *testing.T) {
	type Tests struct {
		description string
		seed        int64
		input       Entitlements
		want        map[string]identity.ServiceDetails
		wantError   error
	}
	tests := []Tests{
		{
			description: "absent template",
			seed:        100,
			input:       nil,
			want: map[string]identity.ServiceDetails{
				"mwu7": {
					IsEntitled: true,
					IsTrial:    true,
				},
			},
			wantError: nil,
		},
		{
			description: "empty template",
			seed:        100,
			input:       Entitlements{},
			want:        map[string]identity.ServiceDetails{},
			wantError:   nil,
		},
		{
			description: "partial template",
			seed:        100,
			input: Entitlements{
				"service1": {
					IsEntitled: pointy.Bool(true),
					IsTrial:    pointy.Bool(false),
				},
			},
			want: map[string]identity.ServiceDetails{
				"service1": {
					IsEntitled: true,
					IsTrial:    false,
				},
			},
			wantError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			faker.SetSeed(test.seed)

			got, err := NewEntitlements(test.input)

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
