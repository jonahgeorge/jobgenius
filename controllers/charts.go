package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/jonahgeorge/jobgenius.net/models/blocks"
)

type Chart struct{}

func (c Chart) GroupWork() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		data := blocks.GroupworkChart{}.RetrieveById(params.Get("i"))

		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, string(b))
	}
}

func (c Chart) Fulfillment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		data := blocks.FulfillmentChart{}.RetrieveById(params.Get("i"))

		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, string(b))
	}
}

func (c Chart) Breakdown() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		data := blocks.BreakdownChart{}.RetrieveById(params.Get("i"))

		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, string(b))
	}
}
