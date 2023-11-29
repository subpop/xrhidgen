package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

// System holds values to be used as input when generating a system identity
// record.
type System struct {
	CertType  *string
	ClusterID *string
	CN        *string
}

// NewSystem will build and return a fully populated System data structure,
// using any values that are present in template.
func NewSystem(template System) (*identity.System, error) {
	var id identity.System

	if template.CertType != nil {
		id.CertType = template.CertType
	} else {
		id.CertType = ptrstring(faker.Pick("", "consumer", "system"))
	}

	if template.ClusterID != nil {
		id.ClusterID = template.ClusterID
	} else {
		id.ClusterID = ptrstring(faker.String())
	}

	if template.CN != nil {
		id.CN = *template.CN
	} else {
		id.CN = faker.String()
	}

	return &id, nil
}
