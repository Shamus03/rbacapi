# RBAC REST API
REST API for role-based access control in Go

Use: 
```
# Run
go run main.go &

# HasAccess check
curl -X GET localhost:8080/user/{user}/hasaccess/{permission}

# Index all direct/inherited roles for user
curl -X GET localhost:8080/user/{user}/roles

# Index permissions directly granted by role
curl -X GET localhost:8080/role/{role}/permissions

# Index all permissions granted by a keyset
curl -X GET localhost:8080/keyset/{keyset}/permissions

# Index all roles required for a keyset
curl -X GET localhost:8080/keyset/{keyset}/roles

# Index all roles a role inherits from (ex. manager inherits from employee)
curl -X GET localhost:8080/role/{role}/parentroles

# Index all roles that inherit from a role (ex. manager, clerk, janitor all inherit from employee)
curl -X GET localhost:8080/role/{role}/childroles

# Set a role for a user
curl -X PUT localhost:8080/user/{user}/roles/{role}/{true/false}

# Set a permission granted by a role
curl -X PUT localhost:8080/role/{role}/permissions/{permission}/{true/false}

# Set a role required for a keyset
curl -X PUT localhost:8080/keyset/{keyset}/roles/{role}/{true/false}

# Set a permission granted by a keyset (careful here; a permission granted by a keyset with no required roles is granted to everyone)
curl -X PUT localhost:8080/keyset/{keyset}/permissions/{permission}/{true/false}

# Make a role inherit from another role (ex. role=manager, parentrole=employee) 
curl -X PUT localhost:8080/role/{role}/parentroles/{parentrole}/{true/false}
```

Future plans:
- Load data from a database
- Write to database on PUT, read from database via some sort of trigger
- Security
