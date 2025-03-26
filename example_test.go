package xrhidgen_test

import (
	"encoding/json"
	"fmt"

	"github.com/subpop/xrhidgen"
)

func ExampleAssociate() {
	xrhidgen.SetSeed(100)
	id, err := xrhidgen.NewAssociateIdentity(xrhidgen.Identity{}, xrhidgen.Associate{}, nil)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"employee_account_number":"02299","org_id":"41123","internal":{"org_id":""},"associate":{"Role":null,"email":"winnifredwinning@shred.org","givenName":"Cameron","rhatUUID":"00e3c758-1d7d-4ecd-98a2-997157e2d05c","surname":"Swift"},"type":"Associate","auth_type":"basic-auth"},"entitlements":null}
}

func ExampleInternal() {
	xrhidgen.SetSeed(101)
	id, err := xrhidgen.NewInternalIdentity(xrhidgen.Identity{}, xrhidgen.Internal{}, nil)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"org_id":"03797","internal":{"org_id":"08321","auth_time":-2978345600000000000},"type":"Internal","auth_type":"basic-auth"},"entitlements":null}
}

func ExampleSystem() {
	xrhidgen.SetSeed(102)
	id, err := xrhidgen.NewSystemIdentity(xrhidgen.Identity{}, xrhidgen.System{}, nil)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"account_number":"16398","org_id":"57572","internal":{"org_id":"57572"},"system":{"cn":"It6P","cert_type":"consumer","cluster_id":"x8LdjPo"},"type":"System","auth_type":"basic-auth"},"entitlements":null}
}

func ExampleUser() {
	xrhidgen.SetSeed(103)
	id, err := xrhidgen.NewUserIdentity(xrhidgen.Identity{}, xrhidgen.User{}, nil)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"org_id":"23807","internal":{"org_id":"23807"},"user":{"username":"tycoon","email":"fransen@crump.biz","first_name":"Frankie","last_name":"Collins","is_active":false,"is_org_admin":false,"is_internal":true,"locale":"pi","user_id":"backset"},"type":"User","auth_type":"cert-auth"},"entitlements":null}
}

func ExampleX509() {
	xrhidgen.SetSeed(103)
	id, err := xrhidgen.NewX509Identity(xrhidgen.Identity{}, xrhidgen.X509{}, nil)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"org_id":"23807","internal":{"org_id":""},"x509":{"subject_dn":"2","issuer_dn":"3sfSj"},"type":"X509","auth_type":"cert-auth"},"entitlements":null}
}

func ExampleServiceAccount() {
	xrhidgen.SetSeed(103)
	id, err := xrhidgen.NewServiceAccountIdentity(xrhidgen.Identity{}, xrhidgen.ServiceAccount{}, nil)
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"org_id":"23807","internal":{"org_id":""},"service_account":{"client_id":"2","username":"crump"},"type":"ServiceAccount","auth_type":"cert-auth"},"entitlements":null}
}
