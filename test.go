package main

import (
	"log"
	"net/http"

	export "ac9/glad/services/rds_to_sf"
	handler "ac9/glad/services/sf_handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/account", handler.AccountHandler)
	router.HandleFunc("/course", handler.CourseHandler)
	router.HandleFunc("/product", handler.ProductHandler)
	router.HandleFunc("/timing", handler.TimingHandler)
	router.HandleFunc("/center", handler.CenterHandler)
	// router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	parsed, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		log.Println("there was an error parsing the body", err)
	// 	}
	// 	var result []test.Data
	// 	err = json.Unmarshal(parsed, &result)
	// 	if err != nil {
	// 		log.Println("there was an error parsing the result", err)
	// 	}
	// 	for _, record := range result {
	// 		value := record.Object
	// 		switch value {
	// 		case "Event__c":
	// 			// Test_Parser(record.Items)
	// 		case "Timing__c":
	// 			handler.TimingHandler(w, r)
	// 		}

	// 	}
	// })

	// router.HandleFunc("/export", func(w http.ResponseWriter, r *http.Request) {
	// 	parsed, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		log.Println("there was an error parsing the body")
	// 	}
	// 	type q struct {
	// 		Id int `json:"id"`
	// 	}
	// 	var tester q
	// 	err = json.Unmarshal(parsed, &tester)
	// 	if err != nil {
	// 		log.Println("there was an error in the unmarshal process")
	// 	}
	// 	export.Export(entity.ID(tester.Id))
	// })
	router.HandleFunc("/rds/export/{id}", export.ExportHandler)
	log.Println("now listening at port 4001")
	log.Println(http.ListenAndServe(":4001", router))
}
