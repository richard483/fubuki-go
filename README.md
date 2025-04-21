To update all dependencies:
`go get -u ./...`
`go mod tidy`

To install dependencies:
`go mod tidy`

For running the application on windows:

- Set the env variable on `go-run-example.bat`
- Exec the bat

To get gcloud access token:
`gcloud auth application-default print-access-token`

To update swagger:
`swag i`

Build docker locally:
`docker build --tag <username>/fbk-go .`

Run docker image with env variable
`docker run -it -p 8080:8080 -e ... -e ... <username>fbk-go`
