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

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	var response []test_entity.Account
	var repo repository.Mongo
	collection := repo.Connect()
	parse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("there was an error reading the request body", err)
	}
	err = json.Unmarshal(parse, &response)
	if err != nil {
		log.Println("there was an error unmarshalling the body", err)
	}
	defer r.Body.Close()
	for _, record := range response {
		values := record.Value
		err := tapi.WriteToDB(record.NewAccount(values.Ext_Id, values.Tenant_Id, values.Cognito_Id, values.Name, values.First_Name, values.Last_Name, values.Phone, values.Email, values.Type, values.Updated_at, values.Created_at))
		if err == nil {
			json.NewEncoder(w).Encode(record.Value)
			log.Println("insertion was successful")
		} else {
			json.NewEncoder(w).Encode(err)
			json.NewEncoder(w).Encode("failed")
			_, insertErr := collection.InsertOne(context.Background(), err)
			if insertErr != nil {
				log.Println("there was an error in the log storage", insertErr)
			} else {
				log.Println("the log was stored successfully")
			}
		}
	}

	log.Println("you sent the following:", response)

}
