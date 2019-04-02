go test -coverprofile=cover.out ./greet/
go tool cover -html=cover.out -o cover.html
