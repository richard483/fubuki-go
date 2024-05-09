To install dependencies:
`go mod tidy`

To get gcloud access token:
`gcloud auth application-default print-access-token`

To update swagger:
`swag i`

Build docker locally:
`docker build --tag <username>/fbk-go .`

Run docker image with env variable
`docker run -it -p 8080:8080 -e ... -e ... <username>fbk-go`