EXECUTABLE := codeexec

build:
	./scripts/setup.sh
	go build -o ${EXECUTABLE} ./cmd/server/main.go
run:
	./${EXECUTABLE}
clean:
	find . -type f -name 'logs' -delete
	rm ${EXECUTABLE}
	podman rmi $$(podman images -q --filter "label=app=codeexec")
test:
	./scripts/curls.sh
