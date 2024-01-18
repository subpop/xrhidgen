package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/v2/identity"
)

// Identity holds values to be used as input when generating a main identity
// record.
type Identity struct {
	AccountNumber         *string
	AuthType              *string
	EmployeeAccountNumber *string
	OrgID                 *string
	Type                  *string
}

// NewIdentity will build and return a partially populated Identity data
// structure, using any values that are present in template.
func NewIdentity(template Identity) (*identity.XRHID, error) {
	var id identity.XRHID

	if template.AccountNumber != nil {
		id.Identity.AccountNumber = *template.AccountNumber
	} else {
		if faker.Bool() {
			id.Identity.AccountNumber = faker.DigitsWithSize(5)
		}
	}

	if template.AuthType != nil {
		id.Identity.AuthType = *template.AuthType
	} else {
		id.Identity.AuthType = faker.Pick("basic-auth", "cert-auth")
	}

	if template.EmployeeAccountNumber != nil {
		id.Identity.EmployeeAccountNumber = *template.EmployeeAccountNumber
	} else {
		if faker.Bool() {
			id.Identity.EmployeeAccountNumber = faker.DigitsWithSize(5)
		}
	}

	if template.OrgID != nil {
		id.Identity.OrgID = *template.OrgID
	} else {
		id.Identity.OrgID = faker.DigitsWithSize(5)
	}

	if template.Type != nil {
		id.Identity.Type = *template.Type
	} else {
		id.Identity.Type = faker.String()
	}

	return &id, nil
}

// NewAssociateIdentity will build and return a fully populated Associate
// identity record, using any values that are present in identityTemplate and
// associateTemplate.
func NewAssociateIdentity(identityTemplate Identity, associateTemplate Associate) (*identity.XRHID, error) {
	id, err := NewIdentity(identityTemplate)
	if err != nil {
		return nil, err
	}

	associate, err := NewAssociate(associateTemplate)
	if err != nil {
		return nil, err
	}

	id.Identity.Associate = *associate

	id.Identity.Type = "Associate"

	return id, nil
}

// NewInternalIdentity will build and return a fully populated Internal identity
// record, using any values that are present in identityTemplate and
// internalTemplate.
func NewInternalIdentity(identityTemplate Identity, internalTemplate Internal) (*identity.XRHID, error) {
	id, err := NewIdentity(identityTemplate)
	if err != nil {
		return nil, err
	}

	internal, err := NewInternal(internalTemplate)
	if err != nil {
		return nil, err
	}

	id.Identity.Internal = *internal

	id.Identity.Type = "Internal"

	return id, nil
}

// NewSystemIdentity will build and return a fully populated System identity
// record, using any values that are present in identityTemplate and
// systemTemplate.
func NewSystemIdentity(identityTemplate Identity, systemTemplate System) (*identity.XRHID, error) {
	id, err := NewIdentity(identityTemplate)
	if err != nil {
		return nil, err
	}

	system, err := NewSystem(systemTemplate)
	if err != nil {
		return nil, err
	}

	id.Identity.System = *system

	id.Identity.Type = "System"
	id.Identity.Internal = identity.Internal{
		OrgID: id.Identity.OrgID,
	}

	return id, nil
}

// NewX509Identity will build and return a fully populated X509 identity record,
// using any values that are present in identityTemplate and x509Template.
func NewX509Identity(identityTemplate Identity, x509Template X509) (*identity.XRHID, error) {
	id, err := NewIdentity(identityTemplate)
	if err != nil {
		return nil, err
	}

	x509, err := NewX509(x509Template)
	if err != nil {
		return nil, err
	}

	id.Identity.X509 = *x509

	id.Identity.Type = "X509"

	return id, nil
}

// NewUserIdentity will build and return a fully populated User identity record,
// using any values that are present in identityTemplate and userTemplate.
func NewUserIdentity(identityTemplate Identity, userTemplate User) (*identity.XRHID, error) {
	id, err := NewIdentity(identityTemplate)
	if err != nil {
		return nil, err
	}

	user, err := NewUser(userTemplate)
	if err != nil {
		return nil, err
	}

	id.Identity.User = *user

	id.Identity.Type = "User"
	id.Identity.Internal = identity.Internal{
		OrgID: id.Identity.OrgID,
	}

	return id, nil
}
