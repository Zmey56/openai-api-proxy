package proxy

import (
	"fmt"
)

type Configuration struct {
	OpenaiToken   string
	OpenaiAddress string
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
