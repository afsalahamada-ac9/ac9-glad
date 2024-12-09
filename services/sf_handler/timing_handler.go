package sf_handler

import (
	test_entity "ac9/glad/entity/sf_entity"
	tapi "ac9/glad/services/tapi"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func TimingHandler(w http.ResponseWriter, r *http.Request) {
	var response []test_entity.Timing
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
		_, err := tapi.WriteToDB(record.NewTiming(value.Course_id, value.Ext_id, value.Course_date, value.Start_time, value.End_time, value.Updated_at, value.Created_at))
		if err == nil {
			json.NewEncoder(w).Encode(record)
		} else {
			json.NewEncoder(w).Encode(err)
		}
	}
	json.NewEncoder(w).Encode(response)
}
