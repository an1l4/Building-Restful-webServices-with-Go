package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func pingTime(req *restful.Request, resp *restful.Response) {

	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))

}

func main() {
	//create a web service
	webservice := new(restful.WebService)

	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))

	// Add the service to application
	restful.Add(webservice)

	fmt.Println("server running at port 8000")
	http.ListenAndServe(":8000", nil)
}
