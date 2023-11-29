package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

func TestNewInternal(t *testing.T) {
	tests := []struct {
		description string
		seed        int64
		input       Internal
		want        *identity.Internal
		wantError   error
	}{
		{
			description: "empty template",
			seed:        100,
			input:       Internal{},
			want: &identity.Internal{
				AuthTime:    ptrfloat64(-1.6924639230312625e+18),
				CrossAccess: ptrbool(true),
				OrgID:       "00229",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: Internal{
				AuthTime: ptrfloat64(1.0),
			},
			want: &identity.Internal{
				AuthTime:    ptrfloat64(1.0),
				CrossAccess: ptrbool(false),
				OrgID:       "80022",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: Internal{
				AuthTime:    ptrfloat64(2.0),
				CrossAccess: ptrbool(true),
				OrgID:       ptrstring("123456"),
			},
			want: &identity.Internal{
				AuthTime:    ptrfloat64(2.0),
				CrossAccess: ptrbool(true),
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
