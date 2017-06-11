package api

import (
    "encoding/json"
    "net/http"
	
	"rbacapi/rbac"
	"rbacapi/repository"
	. "rbacapi/types"

    "github.com/gorilla/mux"
)

func HasAccessHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user := User(vars["user"])
    permission := Permission(vars["permission"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	hasAccess := rbac.HasAccess(user, permission)
    if err := json.NewEncoder(w).Encode(hasAccess); err != nil {
        panic(err)
    }
}

func GetRolesForUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user := User(vars["user"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	roles := repository.RolesForUser(user)
    if err := json.NewEncoder(w).Encode(roles); err != nil {
        panic(err)
    }
}

func PutRoleForUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user := User(vars["user"])
	role := Role(vars["role"])
	value := vars["value"] == "true"
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	repository.SetRoleForUser(user, role, value)
}

func GetPermissionsForRoleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    role := Role(vars["role"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	permissions := repository.PermissionsForRole(role)
    if err := json.NewEncoder(w).Encode(permissions); err != nil {
        panic(err)
    }
}

func PutPermissionForRoleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	role := Role(vars["role"])
    permission := Permission(vars["permission"])
	value := vars["value"] == "true"
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	repository.SetPermissionForRole(permission, role, value)
}

func GetPermissionsForKeysetHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	keyset := Keyset(vars["keyset"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	permissions := repository.PermissionsForKeyset(keyset)
    if err := json.NewEncoder(w).Encode(permissions); err != nil {
        panic(err)
    }
}

func PutPermissionForKeysetHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    permission := Permission(vars["permission"])
	keyset := Keyset(vars["keyset"])
	value := vars["value"] == "true"
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	repository.SetPermissionForKeyset(permission, keyset, value)
}

func GetRolesRequiredForKeysetHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	keyset := Keyset(vars["keyset"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	roles := repository.RolesRequiredForKeyset(keyset)
    if err := json.NewEncoder(w).Encode(roles); err != nil {
        panic(err)
    }
}

func PutRoleRequiredForKeysetHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	role := Role(vars["role"])
	keyset := Keyset(vars["keyset"])
	value := vars["value"] == "true"
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	repository.SetRoleRequiredForKeyset(keyset, role, value)
}

func GetParentRolesForRoleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	role := Role(vars["role"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	parentRoles := repository.ParentRolesForRole(role)
    if err := json.NewEncoder(w).Encode(parentRoles); err != nil {
        panic(err)
    }
}

func GetChildRolesForRoleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	role := Role(vars["role"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	childRoles := repository.ChildRolesForRole(role)
    if err := json.NewEncoder(w).Encode(childRoles); err != nil {
        panic(err)
    }
}

func PutRoleInheritedFromRoleHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	childRole := Role(vars["childrole"])
    parentRole := Role(vars["parentrole"])
	value := vars["value"] == "true"
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	repository.SetRoleInheritedFromRole(childRole, parentRole, value)
}