# RBAC REST API
REST API for role-based access control in Go

Use: 
```
make run &
curl -X GET localhost:8080/user/{user}/hasaccess/{permission}
```

Future plans:
- Load data from a database
- GET and POST for user/role, role/role, permission/role, keyset/role assignment
- Write to database on POST, read from database via some sort of trigger
- Security
