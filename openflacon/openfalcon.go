package main

import (
	"net/http"
	//"bytes"
	"encoding/json"
	"fmt"
	"bytes"
)

var login_addr  = "/api/v1/user/login"
var logout_addr = "/api/v1/user/logout"
var user_list   = "/api/v1/user/users"
var user_item   = "/api/v1/user/u/4/teams"

type OpenfalconResbonse struct {
	msg string					`json:"msg"`
	data Data `json:"data"`
}

type Data struct {
	sig string `json:"sig"`
	name string `json:"name"`
}

type API struct {
	url    string
	name   string
	passwd string
	Client *http.Client
}

func NewAPI(server, user, passwd string) (*API, error) {
	return &API{server, user, passwd, &http.Client{}}, nil
}


func (api *API) OpenfalconRequest(method string, data interface{}) (bool,error) {
	params := make(map[string]string)
	params["name"] = "root"
	params["password"] = "123456"
	//encoded, err := json.Marshal(params)

	//if err != nil {
	//	return  false ,err
	//}


	//response, err := http.Get(api.url, "application/json", bytes.NewReader(encoded))
	response, err := http.Get(api.url)
	//
	//// Setup our HTTP request
	//request, err := http.NewRequest("POST", api.url, bytes.NewBuffer(encoded))
	//if err != nil {
	//	return false, err
	//}
	//
	////request.Header.Add("Content-Type", "application/json-rpc")
	//
	//// Execute the request
	//response, err := api.Client.Do(request)
	//if err != nil {
	//	return false, err
	//}
	//
	///**
	//We can't rely on response.ContentLength because it will
	//be set at -1 for large responses that are chunked. So
	//we treat each API response as streamed data.
	//*/

	//var back OpenfalconResbonse
	//var buf bytes.Buffer

	fmt.Println( response )
	fmt.Println( err )
	back := make(map[string]interface{})

	json.NewDecoder(response.Body).Decode(&back)
	//_, err = io.Copy(&buf, response.Body)
	//if err != nil {
	//	return false, err
	//}
	//
	//json.Unmarshal(buf.Bytes(), &back)

	fmt.Println( back )
	response.Body.Close()

	return true, nil
}

func (api *API)Login()(bool, error){
	params := make(map[string]string, 0)
	params["name"] = api.name
	params["password"] = api.passwd

	_, err := api.OpenfalconRequest("user.login", params)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return false, err
	}


	return true, nil
}

func (api *API)Logout()(bool, error){
	params := make(map[string]string, 0)
	params["name"] = api.name
	params["password"] = api.passwd

	_, err := api.OpenfalconRequest("user.get", params)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return false, err
	}

	return true, nil
}

func Get() {
	//api, err := NewAPI("http://"+"120.92.110.204"+":"+"8081"+login_addr,"admin", "123456")
	api, err := NewAPI("http://"+"120.92.110.204"+":"+"8081"+user_list,"admin", "123456")
	if err != nil{
		return
	}
	b,err := api.Login()
	if err != nil{
		fmt.Println("Login faile")
		return
	}
	if  b != true {
		fmt.Println("Login faile")
	}
	fmt.Println("Login seccess")

}


func main()	{

	//api, err := NewAPI("http://"+"120.92.110.204"+":"+"8081"+user_list,"admin", "123456")

	params := make(map[string]string)
	params["name"] = "root"
	params["password"] = "123456"
	encoded, err := json.Marshal(params)



	//response, err := api.Client.Post(api.url, "application/json", bytes.NewReader(encoded) )

	api, err := NewAPI("http://120.92.110.204:8080/api/v1/user/login","root", "13456")

	//if err != nil {
	//	return
	//}

	response, err := api.Client.Post(api.url,"application/json", bytes.NewReader(encoded))
	if err != nil {
		return
	}

	if err != nil {
		fmt.Println( " request error !!! ")
		return
	}

	fmt.Println( response.Status )

	fmt.Println( response.Body )

	//var back = make(map[string]interface{}, 0)
	var back =  OpenfalconResbonse{}
	json.NewDecoder(response.Body).Decode(&back)

	fmt.Println( back )

}