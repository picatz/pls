# linux builds
GOOS=linux GOARCH=386 go build -o builds/linux/386/pls pls.go
GOOS=linux GOARCH=amd64 go build -o builds/linux/amd64/pls pls.go

# windows builds
GOOS=windows GOARCH=386 go build -o builds/windows/386/pls.exe pls.go
GOOS=windows GOARCH=amd64 go build -o builds/windows/amd64/pls.exe pls.go

# freebsd builds
GOOS=freebsd GOARCH=386 go build -o builds/freebsd/386/pls pls.go
GOOS=freebsd GOARCH=amd64 go build -o builds/freebsd/amd64/pls pls.go

# dawrin builds
GOOS=dawrin GOARCH=386 go build -o builds/dawrin/386/pls pls.go
GOOS=dawrin GOARCH=amd64 go build -o builds/dawrin/amd64/pls pls.go