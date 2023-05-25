package proxy

import (
	"database/sql"
	"fmt"
)

type Configuration struct {
	OpenaiToken   string
	OpenaiAddress string
	DBConnection  *sql.DB
}

func (conf Configuration) Validate() error {
	if conf.OpenaiToken == "" {
		return fmt.Errorf("openai token is required")
	}
	if conf.OpenaiAddress == "" {
		return fmt.Errorf("openai address is required")
	}
	if conf.DBConnection == nil {
		return fmt.Errorf("openai DB is required")
	}
	return nil
}
