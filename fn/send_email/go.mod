module github.com/monitprod/send_email

go 1.16

require (
	github.com/aws/aws-lambda-go v1.27.0
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/monitprod/core v0.0.0-00010101000000-000000000000 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
)

replace github.com/monitprod/core => ../../core
