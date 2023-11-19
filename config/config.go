package config

import (
	"github.com/alexflint/go-arg"
)

const RootPath = "/v1/"

// Config contains environment variables
type Config struct {
	LogLevel             string `arg:"env:LOG_LEVEL"`
	Port                 int    `arg:"env:PORT"`
	DBUser               string `arg:"env:DB_USER"`
	DBName               string `arg:"env:DB_NAME"`
	DBPass               string `arg:"env:DB_PASS"`
	DBUrl                string `arg:"env:DB_URL"`
	LoanO8gClient        string `arg:"env:LOAN_O8G_CLIENT"`
	LoanDictClient       string `arg:"env:LOAN_DICT_CLIENT"`
	LoanNotification     string `arg:"env:LOAN_NOTIFICATION_CLIENT"`
	LoanOfferClient      string `arg:"env:LOAN_OFFER_CLIENT"`
	LoanCollateralClient string `arg:"env:LOAN_COLLATERAL_CLIENT"`
	DictClient           string `arg:"env:DICT_CLIENT"`
	CustomerClient       string `arg:"env:CUSTOMER_CLIENT"`
	UfcUrl               string `arg:"env:UFC_URL"`
	NotificationQ        string `arg:"env:NOTIFICATION_Q"`
	RabbitMqHost         string `arg:"env:RABBITMQ_HOST"`
	RabbitMqPort         string `arg:"env:RABBITMQ_PORT"`
	RabbitMqUser         string `arg:"env:RABBITMQ_USER"`
	RabbitMqPass         string `arg:"env:RABBITMQ_PASS"`
}

// Conf Props is global variable for environment variables usage
var Conf Config

// LoadConfig loads service configuration into environment
func LoadConfig() {
	_ = arg.Parse(&Conf)
}
