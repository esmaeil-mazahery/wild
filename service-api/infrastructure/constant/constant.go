package constant

import (
	"fmt"
	"os"
)

//Performance dddd
var Performance int64 = 100

//KavenegarAPIKey kavenegar Api Key
var KavenegarAPIKey = os.Getenv("KAVENEGAR_API_KEY")

//JwtSecret use for encryption jwt token
var JwtSecret = os.Getenv("SERVICE_SECRET_KEY")

//Dburi Defult Connection String To Mongodb
func Dburi() string {
	Username := os.Getenv("MONGO_USERNAME")
	Password := os.Getenv("MONGO_PASSWORD")
	Address := os.Getenv("MONGO_ADDRESS")
	return fmt.Sprintf("mongodb://%s:%s@%s/?authSource=admin&readPreference=primary&appname=MongoDB%%20Compass&ssl=false", Username, Password, Address)
}

//Dbname Defult Database name
func Dbname() string {
	return os.Getenv("MONGO_DBNAME")
}

//EmailFrom ...
func EmailFrom() string {
	return os.Getenv("EMAIL_FROM")
}

//EmailPassword ...
func EmailPassword() string {
	return os.Getenv("EMAIL_PASSWORD")
}

//EmailSmtphost ...
func EmailSmtphost() string {
	return os.Getenv("EMAIL_SMTPHOST")
}

//EmailSmtpport ...
func EmailSmtpport() string {
	return os.Getenv("EMAIL_SMTPPORT")
}

//RabbitMQuri Defult Connection String To RabbitMQ
func RabbitMQuri() string {
	Username := os.Getenv("RABBITMQ_USERNAME")
	Password := os.Getenv("RABBITMQ_PASSWORD")
	Address := os.Getenv("RABBITMQ_ADDRESS")
	return fmt.Sprintf("amqp://%s:%s@%s/", Username, Password, Address)
}

//RedisURI get redis address server from environment
func RedisURI() string {
	return os.Getenv("REDIS_URI")
}

//RedisPassword get redis password server from environment
func RedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}
