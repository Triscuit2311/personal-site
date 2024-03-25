echo "> Removing old templ generated files"
find . -type f -name '*templ.go' -delete
echo "> Running templ generate"
templ generate
echo "> Running server"
go run cmd/server/main.go

