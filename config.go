package zoom

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	SDKKey    string `split_words:"true" envconfig:"ZOOM_SDK_KEY"`
	SDKSecret string `split_words:"true" envconfig:"ZOOM_SDK_SECRET"`
}

func loadConfig() (*config, error) {
	var c config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
