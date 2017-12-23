package router

import (
	"net/http"

	"controller"
)

func init() {
	http.HandleFunc("/", controller.GetDefaultData)
	http.HandleFunc("/hello", controller.GetHello)
}
