package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/identity"
)

// User holds values to be used as input when generating a user identity record.
type User struct {
	Email      *string
	FirstName  *string
	IsActive   *bool
	IsInternal *bool
	IsOrgAdmin *bool
	LastName   *string
	Locale     *string
	UserID     *string
	Username   *string
}

// NewUser will build and return a fully populated User data structure, using
// any values that are present in template.
func NewUser(template User) (*identity.User, error) {
	var id identity.User

	if template.Email != nil {
		id.Email = *template.Email
	} else {
		id.Email = faker.Email()
	}

	if template.FirstName != nil {
		id.FirstName = *template.FirstName
	} else {
		id.FirstName = faker.FirstName()
	}

	if template.IsActive != nil {
		id.IsActive = *template.IsActive
	} else {
		id.IsActive = faker.Bool()
	}

	if template.IsInternal != nil {
		id.IsInternal = *template.IsInternal
	} else {
		id.IsInternal = faker.Bool()
	}

	if template.IsOrgAdmin != nil {
		id.IsOrgAdmin = *template.IsOrgAdmin
	} else {
		id.IsOrgAdmin = faker.Bool()
	}

	if template.LastName != nil {
		id.LastName = *template.LastName
	} else {
		id.LastName = faker.LastName()
	}

	if template.Locale != nil {
		id.Locale = *template.Locale
	} else {
		id.Locale = faker.LangCode()
	}

	if template.UserID != nil {
		id.UserID = *template.UserID
	} else {
		id.UserID = faker.Username()
	}

	if template.Username != nil {
		id.Username = *template.Username
	} else {
		id.Username = faker.Username()
	}

	return &id, nil
}
