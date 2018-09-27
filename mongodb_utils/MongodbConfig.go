// Author: Paweł Konopko
// License: MIT

package mongodb_utils

type MongodbConfig struct {
	Host                       string `yaml:"host"`
	Port                       uint16 `yaml:"port"`
	DatabaseName               string `yaml:"name"`
	User                       string `yaml:"user"`
	Password                   string `yaml:"password"`

	MaxRetries                 uint   `yaml:"maxRetries"`
	MillisecondsBetweenRetries uint   `yaml:"milisecondsBetweenRetries"`
}

func LocalMongodbConfig() *MongodbConfig {
	return &MongodbConfig{
		Host: "127.0.0.1",
		Port: uint16(27017),
		DatabaseName: "test",
		User: "",
		Password: "",

		MaxRetries: 3,
		MillisecondsBetweenRetries: 1000,
	}
}
