package repository

import (
	. "rbacapi/types"
)

var kra map[Keyset]map[Role]bool
var pka map[Permission]map[Keyset]bool
var pra map[Permission]map[Role]bool
var rra map[Role]map[Role]bool
var ura map[User]map[Role]bool

func LoadData() {
	// Load the data from somewhere.  Here it is loaded from the CSV files in data/
	kra = loadKRA("data/KRA.csv")
	pka = loadPKA("data/PKA.csv")
	pra = loadPRA("data/PRA.csv")
	rra = loadRRA("data/RRA.csv")
	ura = loadURA("data/URA.csv")
}

func RolesForUser(user User) map[Role]bool {
	result := ura[user]
	if result != nil {
	    return result
	}
	result = make(map[Role]bool)
	ura[user] = result
	return result
}

func RolesWithPermission(permission Permission) map[Role]bool {
	result := pra[permission]
	if result != nil {
	    return result
	}
	result = make(map[Role]bool)
	pra[permission] = result
	return result
}

func PermissionsForRole(role Role) map[Permission]bool {
	result := make(map[Permission]bool)
	for perm, _ := range pra {
		if (pra[perm][role] == true) {
			result[perm] = true
		}
	}
	return result
}

func KeysetsWithPermission(permission Permission) map[Keyset]bool {
	result := pka[permission]
	if result != nil {
	    return result
	}
	result = make(map[Keyset]bool)
	pka[permission] = result
	return result
}

func PermissionsForKeyset(keyset Keyset) map[Permission]bool {
	result := make(map[Permission]bool)
	for perm, _ := range pka {
		if (pka[perm][keyset] == true) {
			result[perm] = true
		}
	}
	return result
}

func ParentRolesForRole(role Role) map[Role]bool {
	result := rra[role]
	if result != nil {
	    return result
	}
	result = make(map[Role]bool)
	rra[role] = result
	return result
}

func ChildRolesForRole(role Role) map[Role]bool {
	result := make(map[Role]bool)
	for inheritedRole, _ := range rra {
		if (rra[inheritedRole][role] == true) {
			result[inheritedRole] = true
		}
	}
	return result
}

func RolesRequiredForKeyset(keyset Keyset) map[Role]bool {
	result := kra[keyset]
	if result != nil {
	    return result
	}
	result = make(map[Role]bool)
	kra[keyset] = result
	return result
}


func SetRoleForUser(user User, role Role, value bool) {
	if value {
		if ura[user] == nil {
			ura[user] = make(map[Role]bool)
		}
		ura[user][role] = value
	} else {
		delete(ura[user], role)
	}
}

func SetPermissionForRole(permission Permission, role Role, value bool) {
	if value {
		if pra[permission] == nil {
			pra[permission] = make(map[Role]bool)
		}
		pra[permission][role] = value
	} else {
		delete(pra[permission], role)
	}
}

func SetPermissionForKeyset(permission Permission, keyset Keyset, value bool) {
	if value {
		if pka[permission] == nil {
			pka[permission] = make(map[Keyset]bool)
		}
		pka[permission][keyset] = value
	} else {
		delete(pka[permission], keyset)
	}
}

func SetRoleInheritedFromRole(childRole Role, parentRole Role, value bool) {
	if value {
		if rra[childRole] == nil {
			rra[childRole] = make(map[Role]bool)
		}
		rra[childRole][parentRole] = value
	} else {
		delete(rra[childRole], parentRole)
	}
}

func SetRoleRequiredForKeyset(keyset Keyset, role Role, value bool) {
	if value {
		if kra[keyset] == nil {
			kra[keyset] = make(map[Role]bool)
		}
		kra[keyset][role] = value
	} else {
		delete(kra[keyset], role)
	}
}
