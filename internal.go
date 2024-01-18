package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/v2/identity"
)

// Internal holds values to be used as input when generating an internal
// identity record.
type Internal struct {
	AuthTime    *float32
	CrossAccess *bool
	OrgID       *string
}

// NewInternal will build and return a fully populated Internal data structure,
// using any values that are present in template.
func NewInternal(template Internal) (*identity.Internal, error) {
	var id identity.Internal

	if template.AuthTime != nil {
		id.AuthTime = *template.AuthTime
	} else {
		id.AuthTime = float32(faker.Duration())
	}

	if template.CrossAccess != nil {
		id.CrossAccess = *template.CrossAccess
	} else {
		id.CrossAccess = faker.Bool()
	}

	if template.OrgID != nil {
		id.OrgID = *template.OrgID
	} else {
		id.OrgID = faker.DigitsWithSize(5)
	}

	return &id, nil
}
