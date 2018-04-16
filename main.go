package main

import (
	"app/controllers"
	"app/utils/db"
	"app/utils/server"
	"app/utils/settings"
	"app/utils/token"
	"encoding/json"
	"log"
	_ "net/http/pprof"
	"runtime"
)

var config = &configuration{}

// config struct
type configuration struct {
	Version  string        `json:"Version"`
	Database db.Database   `json:"Database"`
	Server   server.Server `json:"Server"`
	Token    token.Token   `json:"Token"`
}

// ParseJSON ...
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	settings.LoadConfig("config.json", config)
	db.LoadDb(&config.Database)
	// session.InitSession(&config.Session, config.Server.Domain)
	r := controllers.RouteInit()
	server.Run(r, config.Server)
}
