package proxy

import (
	"fmt"
	"github.com/Zmey56/openai-api-proxy/repository"
)

type Configuration struct {
	OpenaiToken   string
	OpenaiAddress string
	DBConnection  *repository.DBImpl
}

func (conf Configuration) Validate() error {
	if conf.OpenaiToken == "" {
		return fmt.Errorf("openai token is required")
	}
	if conf.OpenaiAddress == "" {
		return fmt.Errorf("openai address is required")
	}
	return nil
}
