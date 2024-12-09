package sf_handler

import (
	entity "ac9/glad/entity/sf_entity"
	"ac9/glad/repository"
	"ac9/glad/services/tapi"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CenterHandler(w http.ResponseWriter, r *http.Request) {
	var centers []entity.Center
	var repo repository.Mongo
	collection := repo.Connect()
	parsed_response, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("an error occurred")
	}
	err = json.Unmarshal(parsed_response, &centers)
	if err != nil {
		log.Println("an error occurred in the unmarshalling opf the centers")
	}
	for _, record := range centers {
		value := record.Value
		err := tapi.WriteToDB(record.NewCenter(value.Ext_id, value.Tenant_id, value.Ext_name, value.Address, value.Geo_Location, value.Capacity, value.Mode, value.Webpage, value.Is_national_center, value.Is_enabled, value.Created_at, value.Updated_at))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			_, insertErr := collection.InsertOne(context.Background(), err)
			if insertErr != nil {
				log.Println("there was an error creating the record", insertErr)
			}
		} else {
			json.NewEncoder(w).Encode(record)
		}
	}
	log.Println(centers)
}
