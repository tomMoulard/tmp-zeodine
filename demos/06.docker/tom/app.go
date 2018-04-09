package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"time"

	"errors"
	"net"
)

//go get github.com/julienschmidt/httprouter

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func gettime(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, time.Now())
}

func reverse(str string) (res string) {
	for _, s := range str {
		res = string(s) + res
	}
	return
}

func getreverse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, reverse(ps.ByName("str")))
}

func getold(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	i, err := strconv.Atoi(ps.ByName("year"))
	i = time.Now().Year() - i
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Fprint(w, i)
	}
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/gettime", gettime)
	router.GET("/getreverse/:str", getreverse)
	router.GET("/getold/:year", getold)

	fmt.Println("listening on port 9000")
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
	log.Fatal(http.ListenAndServe(":9000", router))
}
