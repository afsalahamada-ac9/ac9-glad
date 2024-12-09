package sf_handler

import (
	test_entity "ac9/glad/entity/sf_entity"
	"ac9/glad/repository"
	tapi "ac9/glad/services/tapi"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func TimingHandler(w http.ResponseWriter, r *http.Request) {
	var response []test_entity.Timing
	var repo repository.Mongo
	collection := repo.Connect()
	parse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("there was an error parsing the request body", err)
	}
	defer r.Body.Close()
	err = json.Unmarshal(parse, &response)
	if err != nil {
		log.Println("there was an error unmarshalling the request body", err)
	}
	for _, record := range response {
		value := record.Value
		err := tapi.WriteToDB(record.NewTiming(value.Course_id, value.Ext_id, value.Course_date, value.Start_time, value.End_time, value.Updated_at, value.Created_at))
		if err == nil {
			json.NewEncoder(w).Encode(record)
		} else {
			json.NewEncoder(w).Encode(err)
			_, insertErr := collection.InsertOne(context.Background(), err)
			if insertErr != nil {
				log.Println("there was an error saving the log", insertErr)
			}
		}
	}
	json.NewEncoder(w).Encode(response)
}
