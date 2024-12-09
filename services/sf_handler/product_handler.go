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

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	var response []test_entity.Product
	var repo repository.Mongo
	collection := repo.Connect()
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("there was an error reading the body")
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Println("error in unmarshal process", err)
	}
	defer r.Body.Close()
	log.Println("response:", string(resp))
	for _, record := range response {
		value := record.Value
		err := tapi.WriteToDB(record.NewProduct(value.Updated_at, value.Created_at /*value.Is_deleted,*/, value.Format, value.Max_Attendees, value.Listing_Visibity, value.Event_Duration, value.Product, value.CType, value.Title, value.Name, value.TenantID, value.ExtID, value.Base_product_ext_id, value.Is_auto_approve))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			_, insertErr := collection.InsertOne(context.Background(), err)
			if insertErr != nil {
				log.Println("there was an error in creating the error record", insertErr)
			}
		} else {
			json.NewEncoder(w).Encode(record)
		}
	}

}
