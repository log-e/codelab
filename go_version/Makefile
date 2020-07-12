BUILD_DATE := `date +%FT%T%z`
LD_FLAGS := "-w -X 'go_version/version.buildDate=$(BUILD_DATE)'"
build:
	@go build -ldflags ${LD_FLAGS} main.go
