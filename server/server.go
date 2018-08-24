package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

// Errors
var (
	errNoObjectKey = errors.New("no object key in request")
	errNoEntryKey  = errors.New("no entry key in request")
)

// Start inits web server on provided port.
func Start(port string) {
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        &handler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// header := r.Header
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "body reading error: %v", err)
		return
	}

	val, err := handleRequest(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "handling error: %v", err)
		return

	}

	fmt.Fprintf(w, val)
}

func handleRequest(body []byte) (string, error) {
	values := gjson.GetManyBytes(body, "object", "entry")
	object, entry := values[0], values[1]
	obj := object.String()
	if len(obj) == 0 {
		return "", errNoObjectKey
	}

	if obj == "page" {
		// do something
	}

	if !entry.IsArray() {
		return "", errNoEntryKey
	}

	entry.ForEach(func(_, value gjson.Result) bool {
		id := value.Get("id")
		if v := id.String(); v == "something" {
			// do something
		}

		messaging := value.Get("messaging")
		if v := messaging.String(); v == "something" {
			// do something
		}

		return true // When returned value is false, will stop iterate.
	})

	return obj, nil
}
