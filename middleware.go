package main

import (
	"fmt"
	"net/http"
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
