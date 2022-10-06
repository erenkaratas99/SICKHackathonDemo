package configs

import (
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
	"net/http"
	"time"
)

type ServiceConfigs struct {
	MongoDuration      time.Duration
	MongoClientURI     string
	DBName             string
	ColName            string
	Port               string
	BaseUrl            string
	QConfig            QueueConfig
	CustomerServiceURL string
}

type QueueConfig struct {
	QMongoClientURI string
	KafkaConfig     KafkaExternalConfig
	RabbitConfig    RabbitExternalConfig
	MongoDuration   time.Duration
}

type RabbitExternalConfig struct {
	RabbitDBName  string
	RabbitColName string
}

type KafkaExternalConfig struct {
	KDBName  string
	KColName string
}

//
//mongodb://root:root1234@mongodb_docker:27017
var cfgs = map[string]ServiceConfigs{
	"prod": {
		MongoDuration:  time.Second * 10,
		MongoClientURI: "mongodb+srv://root:tesodevpair@cluster0.7vdm7jd.mongodb.net/test",
		DBName:         "SAPServDB",
		ColName:        "keywords",
		Port:           ":8000",
		BaseUrl:        "0.0.0.0",
		QConfig: QueueConfig{
			KafkaConfig: KafkaExternalConfig{
				KDBName:  "-",
				KColName: "-",
			},
			RabbitConfig: RabbitExternalConfig{
				RabbitDBName:  "-",
				RabbitColName: "-",
			},
			QMongoClientURI: "-",
			MongoDuration:   time.Second * 10,
		},
		CustomerServiceURL: "customer_app:8001",
	},
	"dev": {
		MongoDuration:  time.Second * 10,
		MongoClientURI: "mongodb+srv://root:tesodevpair@cluster0.7vdm7jd.mongodb.net/test",
		DBName:         "SAPServDB",
		ColName:        "keywords",
		Port:           ":8001",
		BaseUrl:        "127.0.0.1",
		QConfig: QueueConfig{
			KafkaConfig: KafkaExternalConfig{
				KDBName:  "-",
				KColName: "-",
			},
			RabbitConfig: RabbitExternalConfig{
				RabbitDBName:  "-",
				RabbitColName: "-",
			},
			QMongoClientURI: "-",
			MongoDuration:   time.Second * 10,
		},
		CustomerServiceURL: "localhost:8000",
	},
}

func GetConfig(env string) (*ServiceConfigs, error) {
	config, isExist := cfgs[env]
	if !isExist {
		return nil, customErrors.NewHTTPError(http.StatusInternalServerError,
			"ConfigErr",
			"Service configs could not have fetched correctly.")
	}
	return &config, nil
}

//"mongodb+srv://root:tesodevpair@cluster0.7vdm7jd.mongodb.net/test"
