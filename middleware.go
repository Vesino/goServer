package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() MiddleWare {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("CHeckingk Auth")
			if !flag {
				return
			}
			hf(w, r)
		}
	}
}

func Loggin() MiddleWare {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			hf(w, r)
		}
	}
}
