package client

import "sync"

type ConfigClient interface {
}

var (
	configClient ConfigClient
	configOnce   sync.Once
)

func GetConfigClient() ConfigClient {
	return db
}

func InitConfigClient(client ConfigClient) {
	configOnce.Do(
		func() {
			configClient = client
		},
	)
}
