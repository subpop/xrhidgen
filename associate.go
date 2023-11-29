package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

// Associate holds values to be used as input when generating an associate
// identity record.
type Associate struct {
	Email     *string
	GivenName *string
	RHatUUID  *string
	Role      *[]string
	Surname   *string
}

// NewAssociate will build and return a fully populated Associate data
// structure, using any values that are present in template.
func NewAssociate(template Associate) (*identity.Associate, error) {
	var id identity.Associate

	if template.Email != nil {
		id.Email = *template.Email
	} else {
		id.Email = faker.Email()
	}

	if template.GivenName != nil {
		id.GivenName = *template.GivenName
	} else {
		id.GivenName = faker.FirstName()
	}

	if template.RHatUUID != nil {
		id.RHatUUID = *template.RHatUUID
	} else {
		id.RHatUUID = faker.UUID()
	}

	if template.Role != nil {
		id.Role = *template.Role
	} else {
		slice := faker.Slice(faker.IntInRange(0, 3), func() interface{} { return faker.String() })
		for _, v := range slice {
			id.Role = append(id.Role, v.(string))
		}
	}

	if template.Surname != nil {
		id.Surname = *template.Surname
	} else {
		id.Surname = faker.LastName()
	}

	return &id, nil
}
