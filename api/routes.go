package api

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HasAccess",
		"GET",
		"/user/{user}/hasaccess/{permission}",
		HasAccessHandler,
	},
	Route{
		"GetRolesForUser",
		"GET",
		"/user/{user}/roles",
		GetRolesForUserHandler,
	},
	Route{
		"PutRoleForUser",
		"PUT",
		"/user/{user}/roles/{role}/{value}",
		PutRoleForUserHandler,
	},
	Route{
		"GetPermissionsForRole",
		"GET",
		"/role/{role}/permissions",
		GetPermissionsForRoleHandler,
	},
	Route{
		"PutPermissionForRole",
		"PUT",
		"/role/{role}/permissions/{permission}/{value}",
		PutPermissionForRoleHandler,
	},
	Route{
		"GetPermissionsForKeyset",
		"GET",
		"/keyset/{keyset}/permissions",
		GetPermissionsForKeysetHandler,
	},
	Route{
		"PutPermissionForKeyset",
		"PUT",
		"/keyset/{keyset}/permissions/{permission}/{value}",
		PutPermissionForKeysetHandler,
	},
	Route{
		"GetRolesRequiredForKeyset",
		"GET",
		"/keyset/{keyset}/roles",
		GetRolesRequiredForKeysetHandler,
	},
	Route{
		"PutPermissionForKeyset",
		"PUT",
		"/keyset/{keyset}/roles/{role}/{value}",
		PutRoleRequiredForKeysetHandler,
	},
	Route{
		"GetParentRolesForRole",
		"GET",
		"/role/{role}/parentroles",
		GetParentRolesForRoleHandler,
	},
	Route{
		"GetChildRolesForRole",
		"GET",
		"/role/{role}/childroles",
		GetChildRolesForRoleHandler,
	},
	Route{
		"PutRoleInheritedFromRole",
		"PUT",
		"/role/{childrole}/parentroles/{parentrole}/{value}",
		PutRoleInheritedFromRoleHandler,
	},
}