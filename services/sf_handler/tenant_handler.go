package sf_handler

import (
	"ac9/glad/entity"
	"ac9/glad/services/tapi"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func TenantHandler(w http.ResponseWriter, r *http.Request) {
	var tenants []entity.Tenant
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
		_, err := tapi.WriteToDB(&record)
		if err == nil {
			json.NewEncoder(w).Encode(record)
			log.Println("insertion successful")
		} else {
			json.NewEncoder(w).Encode(err)
		}
	}
	log.Println(tenants)
}
