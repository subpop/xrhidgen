package xrhidgen_test

import (
	"encoding/json"
	"fmt"

	"github.com/subpop/xrhidgen"
)

func ExampleAssociate() {
	xrhidgen.SetSeed(100)
	id, err := xrhidgen.NewAssociateIdentity(xrhidgen.Identity{}, xrhidgen.Associate{})
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"associate":{"email":"winnifredwinning@shred.org","givenName":"Cameron","rhatUUID":"00e3c758-1d7d-4ecd-98a2-997157e2d05c","Role":null,"surname":"Swift"},"auth_type":"basic","employee_account_number":"02299","org_id":"41123","type":"Associate"}}
}

func ExampleInternal() {
	xrhidgen.SetSeed(101)
	id, err := xrhidgen.NewInternalIdentity(xrhidgen.Identity{}, xrhidgen.Internal{})
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"auth_type":"basic","internal":{"auth_time":-2978345425851500500,"cross_access":false,"org_id":"08321"},"org_id":"03797","type":"Internal"}}
}

func ExampleSystem() {
	xrhidgen.SetSeed(102)
	id, err := xrhidgen.NewSystemIdentity(xrhidgen.Identity{}, xrhidgen.System{})
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"account_number":"16398","auth_type":"basic","org_id":"57572","system":{"cert_type":"consumer","cluster_id":"x8LdjPo","cn":"It6P"},"type":"System"}}
}

func ExampleUser() {
	xrhidgen.SetSeed(103)
	id, err := xrhidgen.NewUserIdentity(xrhidgen.Identity{}, xrhidgen.User{})
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"auth_type":"cert","internal":{"org_id":"23807"},"org_id":"23807","type":"User","user":{"email":"fransen@crump.biz","first_name":"Frankie","is_active":false,"is_internal":true,"is_org_admin":false,"last_name":"Collins","locale":"pi","user_id":"backset","username":"tycoon"}}}
}

func ExampleX509() {
	xrhidgen.SetSeed(103)
	id, err := xrhidgen.NewX509Identity(xrhidgen.Identity{}, xrhidgen.X509{})
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//Output: {"identity":{"auth_type":"cert","org_id":"23807","type":"X509","x509":{"subject_dn":"2","issuer_dn":"3sfSj"}}}
}
