package main

import (
  "fmt"
  "net/http"
  "html/template"
  "strconv"
  "os"
  "bytes"
  "log"
  "time"

  "github.com/BobBurns/particle"
  "github.com/gorilla/mux"
)

var odata *particle.Event
var t *template.Template

type PData struct {
  Name  string
  Data  int
}

func routeOutput (w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  q := vars["q"]

  fmt.Println(q)

  // handle button press
  if q != "" {

    event := particle.Event{
     Name: "indata",
     Data: particle.Data{
       TTL: uint32(60),
       Private: false,
       },
    }

    // TODO handle access tokens by reading file instead of args
    at := os.Args[1]

    switch(q) {
    case "on":
      event.Data.Data = "on"
    case "off":
      event.Data.Data = "off"
    case "reverse":
      event.Data.Data = "rev"
    case "fast":
      event.Data.Data = "fast"
    case "med":
      event.Data.Data = "med"
    case "slow":
      event.Data.Data = "slow"
    }

    result, err := event.Publish(at)
    if err != nil {
      fmt.Println("error: event.Publish", err)
    }
    //log this
    if result.Error != "" || !result.OK {
      fmt.Println("got error: ", result.Error)
      fmt.Println("descripton: ", result.ErrorDescription)
    }

  }

  if odata.Name == "" {
	  odata.Name = "unknown"
  }

  intdata, _ := strconv.Atoi(odata.Data.Data)

  d := PData{
    Name: odata.Name,
    Data: intdata,
  }
  var b bytes.Buffer

  h := "html-template.html"
  err := t.ExecuteTemplate (&b, h, d)

  if err != nil {
    fmt.Fprintf(w, "Error with template: %s ", err)
    return
  }
  b.WriteTo(w)
}


func main() {
  if len(os.Args) < 2 {
    fmt.Println("usage ./moto <access-token>")
    os.Exit(-1)
  }

  at := os.Args[1]

  // parse html template
  t = template.Must(template.ParseFiles("html/html-template.html"))

  // subscribe to particle events
  myevent := particle.Subscribe("outdata", at)

  odata = new(particle.Event)

  go func() {
    for {
      e := <-myevent
      odata = &e
      }
    }()


	router := mux.NewRouter()
	/* change this to IP addr !! */
	sub := router.Host("localhost").Subrouter()
	sub.PathPrefix("/html/").Handler(http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
	sub.HandleFunc("/data", routeOutput)
	sub.HandleFunc("/data/{q}", routeOutput)

	// IdleTimeout requires go1.8
	server := http.Server{
		Addr:         ":8082",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}
	fmt.Println("Server started at localhost:8082")
	log.Fatal(server.ListenAndServe())

}
