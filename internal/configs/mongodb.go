package configs

import "strings"

type MongoDBConfig struct {
	Host     string
	User     string
	Pass     string
	Protocol string
	Port     string
	Database string
}

func (m *MongoDBConfig) GetConnectionString() string {
	if strings.Contains(m.Protocol, "srv") {
		return m.Protocol + "://" + m.User + ":" + m.Pass + "@" + m.Host + "/" + m.Database +
			"?retryWrites=true&w=majority&appName=Develop"
	} else {
		return m.Protocol + "://" + m.User + ":" + m.Pass + "@" + m.Host + ":" + m.Port + "/" + m.Database +
			"?retryWrites=true&w=majority&appName=Develop"
	}
}
