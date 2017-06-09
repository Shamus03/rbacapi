.PHONY: all run clean

all: main

run: main
	./main

clean:
	rm -f main

main: main.go api/router.go api/routes.go api/logger.go api/handlers.go rbac/rbac.go repository/repository.go
	go build -o main main.go
