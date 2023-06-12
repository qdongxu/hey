binary = hey

release: generate
	GOOS=windows GOARCH=amd64 go build  -o ./bin/$(binary)_windows_amd64 hey.go
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(binary)_linux_amd64 hey.go
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(binary)_darwin_amd64 hey.go

generate:
	go generate ./...
