package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/v2/identity"
	"go.openly.dev/pointy"
)

// ServiceDetails holds values to be used as input when generating a service
// details record.
type ServiceDetails struct {
	IsEntitled *bool
	IsTrial    *bool
}

// NewServiceDetail will build and return a fully populated ServiceDetails data
// structure, using any values that are present in template.
func NewServiceDetail(template ServiceDetails) (*identity.ServiceDetails, error) {
	var id identity.ServiceDetails

	if template.IsEntitled != nil {
		id.IsEntitled = *template.IsEntitled
	} else {
		id.IsEntitled = faker.Bool()
	}

	if template.IsTrial != nil {
		id.IsTrial = *template.IsTrial
	} else {
		id.IsTrial = faker.Bool()
	}

	return &id, nil
}

type Entitlements map[string]ServiceDetails

// NewEntitlements will build and return a fully populated map of ServiceDetails
// data structures, using any values that are present in template.
func NewEntitlements(template Entitlements) (map[string]identity.ServiceDetails, error) {
	entitlements := make(map[string]identity.ServiceDetails)

	if template == nil {
		template = Entitlements{}
		for i := 0; i <= faker.IntInRange(0, 2); i++ {
			sd := ServiceDetails{
				IsEntitled: pointy.Bool(faker.Bool()),
				IsTrial:    pointy.Bool(faker.Bool()),
			}

			template[faker.StringWithSize(4)] = sd
		}
	}

	for k, v := range template {
		sd, err := NewServiceDetail(v)
		if err != nil {
			return nil, err
		}

		entitlements[k] = *sd
	}

	return entitlements, nil
}
