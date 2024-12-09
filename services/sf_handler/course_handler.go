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

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	var courses []entity.Course
	var repo repository.Mongo
	collection := repo.Connect()
	parsed_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("there was an error parsing the body")
	}
	err = json.Unmarshal(parsed_body, &courses)
	if err != nil {
		log.Println("there was an error unmarshalling the body")
	}
	for _, course := range courses {
		value := course.Value
		err := tapi.WriteToDB(course.NewCourse(value.Url, value.Max_attendees, value.Address, value.Tenant_id, value.Ext_id, value.Name, value.Timezone, value.Mode, value.Center_id, value.Status, value.Created_at, value.Num_attendees, value.Product_id, value.Updated_at, value.Notes, value.Short_url))
		if err != nil {
			json.NewEncoder(w).Encode(err)
			_, insertErr := collection.InsertOne(context.Background(), err)
			if insertErr != nil {
				log.Println("there was an error creating the log", insertErr)
			}
		} else {
			json.NewEncoder(w).Encode(course)
		}
	}
	log.Println("this is what is being parsed:", courses)

}
