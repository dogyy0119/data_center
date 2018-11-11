package main

import (
	"data_center/zabbix"
	"fmt"
	"data_center/config"
	"github.com/robfig/cron"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main()  {

	var db  *sql.DB
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/test")
	if err != nil{
		fmt.Println("connet faile")
	}

	result, err := db.Exec(
		"INSERT INTO user (name, age) VALUES (?, ?)",
		"liuh",
		27,
	)

	fmt.Println(result)
}

func test() {
	c := cron.New()
	spec := "0 */1 * * * *"
	c.AddFunc(spec, get_zabbix_data)
	c.Start()
	select {}
}


/**
	get zabbix  item  by using hostid
 */
func  get_zabbix_data() (){

	conf, err := config.Get_ZabbixConf()
	if err != nil {
		return
	}

	api, err := zabbix.NewAPI("http://"+conf.Zabbix.Ip+":"+conf.Zabbix.Port+"/api_jsonrpc.php", conf.Zabbix.User, conf.Zabbix.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = api.Login()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to API")

	var host []zabbix.ZabbixHost
	host, err = api.GetAllHost()

	if err != nil {
		fmt.Println( err )
		return
	}

	for _, ihost := range host{
		fmt.Println( ihost["name"] )

		var items []zabbix.ZabbixItems
		items,err = api.GetHostItems( ihost["hostid"].(string) )

		if err != nil {
			fmt.Println( err )
			return
		}

		for _,item := range items {
			fmt.Println(item["itemid"])
			fmt.Println(item["hostid"])
			fmt.Println(item["key_"])
			fmt.Println(item["name"])
			fmt.Println(item["description"])
			fmt.Println(item["status"])
			fmt.Println(item["lastclock"])
			fmt.Println(item["datatype"])
		}

	}
}
