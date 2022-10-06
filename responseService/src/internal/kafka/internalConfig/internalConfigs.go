package internalConfig

type KafkaConfig struct {
	BootstrapServers string
	SecurityProtocol string
	SaslMech         string
	SaslUsername     string
	SaslPassword     string
	GroupID          string
	AutoOffsetReset  string
}

var kafkaCfgs = map[string]KafkaConfig{
	"responseServiceKafka": {
		BootstrapServers: "pkc-zpjg0.eu-central-1.aws.confluent.cloud:9092",
		SecurityProtocol: "SASL_SSL",
		SaslMech:         "PLAIN",
		SaslUsername:     "D7XDTBINNHQF2RLW",
		SaslPassword:     "cgAh7Rd1reUNHj6IdCB62MVvRGvoZtzSF0648r6jvchX/lmtBwX4RjYQ5ynMttFz",
		GroupID:          "myGroup",
		AutoOffsetReset:  "earliest",
	},
}

func GetKafkaConfig() *KafkaConfig {
	config, isExist := kafkaCfgs["responseServiceKafka"]
	if !isExist {
		return nil
	}
	return &config
}
