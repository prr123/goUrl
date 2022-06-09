// general internet client like curl
// author: prr, azul software
// date: 9 June 2022
// copyright 2022 prr, azul software
//

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
	OPTIONS = iota
	GET
	HEAD
	POST
	PUT
	DELETE
	TRACE
	CONNECT
)

type GoUrl struct {
	HttpCmd [8]string
}

func (gourl *GoUrl) initHttp() {
	gourl.HttpCmd[OPTIONS] = "OPTIONS"
	gourl.HttpCmd[GET] = "GET"
	gourl.HttpCmd[HEAD] = "HEAD"
	gourl.HttpCmd[POST] = "POST"
	gourl.HttpCmd[PUT] = "PUT"
	gourl.HttpCmd[DELETE] = "DELETE"
	gourl.HttpCmd[TRACE] = "TRACE"
	gourl.HttpCmd[CONNECT] = "CONNECT"
}

func main() {
	var gourl GoUrl
	gourl.initHttp()

    c := http.Client{Timeout: time.Duration(1) * time.Second}
    req, err := http.NewRequest("GET", "http://www.google.fr", nil)
    if err != nil {
        fmt.Printf("error %s", err)
        return
    }
    req.Header.Add("Accept", `application/json`)
    resp, err := c.Do(req)
    if err != nil {
        fmt.Printf("error %s", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("Body : %s", body)
}
