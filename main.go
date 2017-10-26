package main

import (
	"os"
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
)

type Weather struct {
	Url 		string
	Api 		string
	Location 	string
}

func main() {
	weter := Weather{}
	// Get weather service url
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

	// Get external data via API
	request := get_data(weter)
	fmt.Printf("%+v\n\n", request)

	// Get internal data from workspace file
	data_from_file := read_file(os.Getenv("PLUGIN_FILENAME"))
	fmt.Printf("Data from file: %s\n", data_from_file)
}

func get_data(weather Weather) string {
	body := strings.NewReader("Plugin test\n")
	url := fmt.Sprintf("%s%s?q=%s,uk&appid=b1b15e88fa797225412429c1c50c122a1", weather.Url, weather.Api, weather.Location)
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

func read_file(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Open file error %s \n", err)
		os.Exit(1)
	}
	return string(data)
}