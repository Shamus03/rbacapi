package rbac

import (
	"rbacapi/repository"
	. "rbacapi/types"
)

func HasAccess(user User, permission Permission) bool {

	allUserRoles := allUserRoles(user)
	
	rolesWithPermission := repository.RolesWithPermission(permission)
	
	for role, _ := range allUserRoles {
		_, exists := rolesWithPermission[role]
		if exists {
			return true
		}
	}
	
	keysetsWithPermission := repository.KeysetsWithPermission(permission)
	
	for keyset, _ := range keysetsWithPermission {
		if hasKeyset(allUserRoles, keyset) {
			return true
		}
	}
	
	return false
}

func allUserRoles(user User) map[Role]bool {
	directRoles := repository.RolesForUser(user)
	
	allRoles := make(map[Role]bool)
	getAllInheritedRoles(directRoles, allRoles)
	
	return allRoles
}

func hasKeyset(roles map[Role]bool, keyset Keyset) bool {
	requiredRoles := repository.RolesRequiredForKeyset(keyset)
	
	for role, _ := range requiredRoles {
		_, exists := roles[role]
		
		if !exists {
			return false
		}
	}
	
	return true
}

func getAllInheritedRoles(roles map[Role]bool, output map[Role]bool) {
	for role, _ := range roles {
		_, exists := output[role]
		if !exists {
			output[role] = true
			
			inheritedRoles := repository.ParentRolesForRole(role)
			getAllInheritedRoles(inheritedRoles, output)
		}
	}
	return
}