package sf_handler

import (
	"ac9/glad/entity"
	"ac9/glad/repository"
	"ac9/glad/services/tapi"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func TenantHandler(w http.ResponseWriter, r *http.Request) {
	var tenants []entity.Tenant
	var repo repository.Mongo
	collection := repo.Connect()
	parsed_response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	err = json.Unmarshal(parsed_response, &tenants)
	if err != nil {
		log.Println("there was an error unmarshalling the response")
	}
	defer r.Body.Close()
	for _, record := range tenants {
		err := tapi.WriteToDB(&record)
		if err == nil {
			json.NewEncoder(w).Encode(record)
			log.Println("insertion successful")
		} else {
			json.NewEncoder(w).Encode(err)
			_, insertErr := collection.InsertOne(context.Background(), err)
			if insertErr != nil {
				log.Println("there was an error creating the error record", insertErr)
			}
		}
	}
	log.Println(tenants)
}
