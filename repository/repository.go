package repository

import (
	. "rbacapi/types"
)

var kra map[Keyset]map[Role]bool
var pka map[Permission]map[Keyset]bool
var pra map[Permission]map[Role]bool
var rra map[Role]map[Role]bool
var ura map[User]map[Role]bool

func LoadData() { // this is where the data would be loaded from somewhere
	kra = make(map[Keyset]map[Role]bool)
	pka = make(map[Permission]map[Keyset]bool)
	pra = make(map[Permission]map[Role]bool)
	rra = make(map[Role]map[Role]bool)
	ura = make(map[User]map[Role]bool)
	
	ura["shamus"] = make(map[Role]bool)
	ura["shamus"]["reallyAwesomeGuy"] = true //shamus is a reallyAwesomeGuy
	ura["shamus"]["coolGuy"] = true //shamus is a coolGuy
	
	rra["reallyAwesomeGuy"] = make(map[Role]bool)
	rra["reallyAwesomeGuy"]["awesomeGuy"] = true //if you are a reallyAwesomeGuy, you are also an awesomeGuy
	
	pra["beAwesome"] = make(map[Role]bool)
	pra["beAwesome"]["awesomeGuy"] = true //awesomeGuy can beAwesome
	
	pra["beReallyAwesome"] = make(map[Role]bool)
	pra["beReallyAwesome"]["reallyAwesomeGuy"] = true //reallyAwesomeGuy can beReallyAwesome
	
	pka["beAwesomeAndCool"] = make(map[Keyset]bool)
	pka["beAwesomeAndCool"]["awesomeAndCoolKeyset"] = true //you can beAwesomeAndCool if you have the awesomeAndCoolKeyset
	
	kra["awesomeAndCoolKeyset"] = make(map[Role]bool)
	kra["awesomeAndCoolKeyset"]["awesomeGuy"] = true //you have the awesomeAndCoolKeyset if you are an awesomeGuy and a coolGuy
	kra["awesomeAndCoolKeyset"]["coolGuy"] = true
	
	// shamus has permissions: 
	//   beAwesome (from awesomeGuy, inherited from reallyAwesomeGuy),
	//   beReallyAwesome (from reallyAwesomeGuy),
	//   beAwesomeAndCool (from awesomeAndCoolKeyset)
}

func RolesForUser(user User) map[Role]bool {
	return ura[user]
}

func RolesWithPermission(permission Permission) map[Role]bool {
	return pra[permission]
}

func KeysetsWithPermission(permission Permission) map[Keyset]bool {
	return pka[permission]
}

func RolesInheritedFromRole(role Role) map[Role]bool {
	return rra[role]
}

func RolesRequiredForKeyset(keyset Keyset) map[Role]bool {
	return kra[keyset]
}