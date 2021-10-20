module github.com/monitprod/send_email

go 1.16

require (
	github.com/aws/aws-lambda-go v1.27.0
	github.com/monitprod/core v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/monitprod/core => ../../core
