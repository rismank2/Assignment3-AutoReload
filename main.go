package main

import (
	"Assignment3-AutoReload/helper"
	"fmt"
	"net/http"
)

func main() {
	go helper.CreateJson()
	http.HandleFunc("/", helper.ReloadWeb)
	fmt.Println("Server aktif di http://127.0.0.1:9999")
	fmt.Println("Keluar Server CTRL+ C")
	http.ListenAndServe(":9999", nil)
}
