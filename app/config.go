package app

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	App struct {
		Name string `json:"name"`
		Port int    `json:"port"`
	}
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
	}
}

func whichConfig() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	return env
}

func (a *App) InitConfig() error {
	filename := fmt.Sprintf("cmd/configs/config-%s.json", whichConfig())
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var cfg config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return err
	}
	a.config = &cfg
	return nil
}

func (a *App) dBConnectionString(cfg *config) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable&client_encoding=UTF8",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
}
