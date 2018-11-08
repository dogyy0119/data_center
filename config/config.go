package config

import (
	"flag"
	"os"
	"fmt"
	"github.com/naoina/toml"
)

type Config struct {
	Zabbix ZabbixInfo
}

type ZabbixInfo struct {
	Ip 			string
	Port 		string
	User 		string
	Password 	string
}


func Get_conf() (config Config, err error) {
	configFile := flag.String("c","./app.conf","test comfiguration")
	flag.Parse()

	f,err := os.Open(*configFile)
	if err != nil {
		panic(err)
		fmt.Println("file error!")
	}

	defer f.Close()


	if err = toml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
		fmt.Println("open file false!")
		return
	}

	return
}

