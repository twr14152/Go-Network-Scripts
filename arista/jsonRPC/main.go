package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Parameters struct {
	Version int      `json:"version"`
	Cmds    []string `json:"cmds"`
	Format  string   `json:"format"`
}

type Request struct {
	Jsonrpc string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  Parameters `json:"params"`
	Id      string     `json:"id"`
}

var hostList []string
var un string
var pw string
var url string
var hostfile string
var cmdsfile string

func loginHosts(hostfile string) []string {
	hf, err := os.Open(hostfile)
	if err != nil {
		log.Fatal("Failed to Open file: ", err)
	}
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
	return hostList
}

func readCommands(cmdsfile string) []string {
	var cmds []string
	cf, err := os.Open(cmdsfile)
	if err != nil {
		log.Fatal("Failed to Open file: ", err)
	}
	scanner := bufio.NewScanner(cf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		cmds = append(cmds, scanner.Text())
	}
	return cmds
}
func connect(url, un, pw, cmdsfile, format string) *http.Response {
	// Build JSON-RPC payload
	var cmds []string
	cmds = readCommands(cmdsfile)
	fmt.Println(cmds)
	p := Parameters{1, cmds, format}
	req := Request{"2.0", "runCmds", p, "1"}
	buf, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	// Create HTTP client
	client := &http.Client{}
	if strings.HasPrefix(url, "https") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable ssl verification
		}
		client = &http.Client{Transport: tr}
	}

	// Create the request with POST method and the JSON body
	reqHTTP, err := http.NewRequest("POST", url, bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}

	// Set Basic Authentication header
	reqHTTP.SetBasicAuth(un, pw)

	// Set content type
	reqHTTP.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := client.Do(reqHTTP)
	if err != nil {
		panic(err)
	}

	return resp
}

func main() {
	hosts := loginHosts(os.Args[1])
	fmt.Println("hosts:", hosts)
	for host := range hosts {
		url = fmt.Sprintf("https://%s/command-api/", hosts[host])
		un = os.Args[2]
		pw = os.Args[3]
		//fmt.Println(url)
		//fmt.Println(un)
		//fmt.Println(pw)
		//This is the name of your config file
		cmdsfile = fmt.Sprintf("%s.cfg", hosts[host])
		fmt.Println(cmdsfile)

		resp := connect(url, un, pw, cmdsfile, "json")
		defer resp.Body.Close() // Always close the response body

		// Check if the request was successful
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			// Pretty print the JSON response
			var pretty bytes.Buffer
			err = json.Indent(&pretty, bodyBytes, "", "  ")
			if err != nil {
				// Print the raw response
				fmt.Println(string(bodyBytes))
			} else {
				// Pretty print
				fmt.Println(pretty.String())
			}
		}
	}
}
