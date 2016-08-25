install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/password_random/password_random.go
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/password_server/password_server.go
test:
	GO15VENDOREXPERIMENT=1 go test `glide novendor`
vet:
	go tool vet .
	go tool vet --shadow .
lint:
	golint -min_confidence 1 ./...
errcheck:
	errcheck -ignore '(Close|Write)' ./...
check: lint vet errcheck
run:
	password_server \
	-port 8080
open:
	open http://localhost:8080/
format:
	find . -name "*.go" -exec gofmt -w "{}" \;
	goimports -w=true .
prepare:
	npm install
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/Masterminds/glide
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	glide install
update:
	glide up
clean:
	rm -rf vendor
