package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func indexHdndler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "This is the RESTful api")
}
func main() {

	route := httprouter.New()
	route.GET("/", indexHdndler)

	//
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in profuction mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8090", route)

}
