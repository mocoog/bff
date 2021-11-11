# bff
Backend for Frontend

## Test

- test command
- you will get `cover.html`

`go test ./... -v -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html`

## Swagger

- generate interface code command

`swagger generate server --exclude-main --strict-additional-properties -t ./interface/gen -s restapiv1 -f ./interface/swagger/v1.yaml`
