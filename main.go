package main

import (
	"os"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

type Weter struct {
	Url 		string
	Api 		string
	Location 	string
}

func main() {
	weter := Weter{}
	// Get weter service url
	fmt.Println("Get variables from os.envs")
	weter.Url = os.Getenv("PLUGIN_URL")
	if weter.Url == "" {
		weter.Url = "http://samples.openweathermap.org"
	}
	weter.Api = os.Getenv("PLUGIN_API")
	if weter.Api == "" {
		weter.Api = "/data/2.5/weather?q="
	}
	weter.Location = os.Getenv("PLUGIN_LOCATION")
	if weter.Location == "" {
		weter.Location = "Katowice"
	}
	fmt.Printf("We got \nURL: %s \nAPI: %s \nLOCATION: %s\n", weter.Url, weter.Api, weter.Location)
	request := get_data(weter)
	fmt.Printf("%+v", request)
}

func get_data(weter Weter) string {
	body := strings.NewReader("Plugin test\n")
	url := fmt.Sprintf("%s%s?q=%s,uk&appid=b1b15e88fa797225412429c1c50c122a1", weter.Url, weter.Api, weter.Location)
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		fmt.Printf("NewRequest Error %v \n", err)
		os.Exit(1)
	}
	c := &http.Client{}
	response, err := c.Do(req)
	if err != nil {
		fmt.Printf("Do Error %v \n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	return string(data)
}