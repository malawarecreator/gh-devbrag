package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	serverurl := "http://localhost:8080"
	search := flag.Bool("search", false, "To search or not to search")
	publish := flag.Bool("pub", false, "To publish or not to publish")
	like := flag.Bool("like", false, "To like or not to like")
	setserverURL := flag.Bool("ssurl", false, "To set url or not to set url")
	// search
	name := flag.String("name", "", "Name to search")

	// publish
	usrname := flag.String("username", "", "Name to publish")
	description := flag.String("description", "", "Description to publish")
	favlang := flag.String("favlang", "", "Favorite Programming lang")
	recentContributions := flag.String("contribs", "", "Recent Contributions")

	// like
	likename := flag.String("lname", "", "Name to like")

	// ssurl
	ssurl := flag.String("url", "", "Url to set")

	flag.Parse()

	if *setserverURL && *ssurl != "" {
		
		serverurl = *ssurl
	} 
	if (*search && *name != "") {
		queryName := url.QueryEscape(*name)
		resp, err := http.Get(serverurl + "/search?name=" + queryName)
		if err != nil {
			fmt.Println("Error")
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		if err != nil {
			fmt.Println("Error")
			return
		}

		fmt.Println(string(body))
	} else if *publish && *usrname != "" && *description != "" && *favlang != "" && *recentContributions != "" {
		data := map[string]string{
			"username" : *usrname,
			"description": *description,
			"favoritelang": *favlang,
			"recentcontributions": *recentContributions,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Invalid json:", err)
			return
		}

	

		resp, err := http.Post(serverurl + "/publish", "application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			fmt.Println("Error (server might be down):", err)
			return
		}

		defer resp.Body.Close()

		var resb bytes.Buffer

		_, err = resb.ReadFrom(resp.Body)

		if err != nil {
			fmt.Println("Could not read response:", err)
			return
		}

	} else if *like && *likename != "" {
		queryName := url.QueryEscape(*likename)
		resp, err := http.Get(serverurl + "/like?name=" + queryName)
		if err != nil {
			fmt.Println("Error")
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		if err != nil {
			fmt.Println("Error")
			return
		}

		fmt.Println(string(body))
	} else {
		fmt.Println("devbrag: Flex for no reason")
	}
}

