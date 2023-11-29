package xrhidgen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

func TestNewSystem(t *testing.T) {
	tests := []struct {
		description string
		seed        int64
		input       System
		want        *identity.System
		wantError   error
	}{
		{
			description: "empty template",
			seed:        100,
			input:       System{},
			want: &identity.System{
				CertType:  ptrstring("consumer"),
				ClusterID: ptrstring("i"),
				CN:        "wu7Re",
			},
		},
		{
			description: "partial template",
			seed:        100,
			input: System{
				CertType: ptrstring("asdf"),
			},
			want: &identity.System{
				CertType:  ptrstring("asdf"),
				ClusterID: ptrstring("gimwu7Re"),
				CN:        "hCRM50",
			},
		},
		{
			description: "full template",
			seed:        100,
			input: System{
				CertType:  ptrstring("consumer"),
				ClusterID: ptrstring("1234"),
				CN:        ptrstring("xyz"),
			},
			want: &identity.System{
				CertType:  ptrstring("consumer"),
				ClusterID: ptrstring("1234"),
				CN:        "xyz",
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
