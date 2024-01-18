package xrhidgen

import (
	"github.com/pioz/faker"
	"github.com/redhatinsights/platform-go-middlewares/v2/identity"
)

// X509 holds values to be used as input when generating an x509 identity
// record.
type X509 struct {
	SubjectDN *string
	IssuerDN  *string
}

// NewX509 will build and return a fully populated X509 data structure, using
// any values that are present in template.
func NewX509(template X509) (*identity.X509, error) {
	var id identity.X509

	if template.SubjectDN != nil {
		id.SubjectDN = *template.SubjectDN
	} else {
		id.SubjectDN = faker.String()
	}

	if template.IssuerDN != nil {
		id.IssuerDN = *template.IssuerDN
	} else {
		id.IssuerDN = faker.String()
	}

	return &id, nil
}
