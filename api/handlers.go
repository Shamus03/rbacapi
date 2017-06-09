package api

import (
    "encoding/json"
    "net/http"
	
	"rbacapi/rbac"
	. "rbacapi/types"

    "github.com/gorilla/mux"
)

func HasAccessHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    user := User(vars["userId"])
    permission := Permission(vars["permissionId"])
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	
	hasAccess := rbac.HasAccess(user, permission)
    if err := json.NewEncoder(w).Encode(hasAccess); err != nil {
        panic(err)
    }
}