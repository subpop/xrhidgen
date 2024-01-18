package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/v2/identity"
)

// ServiceAccount holds values to be used as input when generating a service
// account identity record.
type ServiceAccount struct {
	ClientID *string
	Username *string
}

// NewServiceAccount will build and return a fully populated ServiceAccount data
// structure, using any values that are present in template.
func NewServiceAccount(template ServiceAccount) (*identity.ServiceAccount, error) {
	var id identity.ServiceAccount

	if template.ClientID != nil {
		id.ClientId = *template.ClientID
	} else {
		id.ClientId = faker.String()
	}

	if template.Username != nil {
		id.Username = *template.Username
	} else {
		id.Username = faker.Username()
	}

	return &id, nil
}
