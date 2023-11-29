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
		id.CertType = *template.CertType
	} else {
		id.CertType = faker.Pick("", "consumer", "system")
	}

	if template.ClusterID != nil {
		id.ClusterId = *template.ClusterID
	} else {
		id.ClusterId = faker.String()
	}

	if template.CN != nil {
		id.CommonName = *template.CN
	} else {
		id.CommonName = faker.String()
	}

	return &id, nil
}
