package config

import (
	"flag"
	"os"
	"fmt"
	"github.com/naoina/toml"
)

type DataModel struct {
	itemid 		string
	type_ 		string
	hostid		string
	name 		string
	key_		string
	lastvalue 	string
	lastclock 	string
	value_type 	string
	datatype    string
	description	string
}

type ZabbixConfig struct {
	Zabbix ZabbixInfo
}

type ZabbixInfo struct {
	Ip 			string
	Port 		string
	User 		string
	Password 	string
}

type OpenfalconConfig struct {
	Openfalcon  OpenfalconInfo
}

type OpenfalconInfo struct {
	Ip          string
	Port        string
	Name        string
	Password    string
}

type Database struct {
	Mysql MysqlInfo
}

type MysqlInfo struct {
	Ip          string
	Port        string
	Name        string
	Password    string
	Tabel		string
}

func Get_ZabbixConf() (config ZabbixConfig, err error) {
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


func Get_OpenfalconConf() (config OpenfalconConfig, err error) {
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

func Get_MysqlConf() (config MysqlInfo, err error) {
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
