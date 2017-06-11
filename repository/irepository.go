package repository

import (
	. "rbacapi/types"
)

type IRepository interface {
	RolesForUser(User) map[Role]bool
	RolesWithPermission(Permission) map[Role]bool
	PermissionsForRole(Role) map[Permission]bool
	KeysetsWithPermission(Permission) map[Keyset]bool
	RolesInheritedFromRole(Role) map[Role]bool
	RolesThatInheritFromRole(Role) map[Role]bool
	RolesRequiredForKeyset(Keyset) map[Role]bool
	
	SetRoleForUser(User, Role, bool)
	SetPermissionForRole(Permission, Role, bool)
	SetPermissionForKeyset(Permission, Keyset, bool)
	SetRoleInheritedFromRole(Role, Role, bool)
	SetRoleRequiredForKeyset(Keyset, Role, bool)
}