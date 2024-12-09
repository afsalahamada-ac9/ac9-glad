package tapi

import "log"

type EntityType string

func EntityCreationHandler(entity EntityType) (string, string) {
	_, err := WriteToDB(entity)
	if err != nil {
		log.Println("there was an error writing the entity to the DB")
		return "", ""
	}
	url := "test"
	id := "test"
	// todo: implement the logic for url and id generation
	return url, id
}
