package tapi

// todo: write the data to rds db
//  todo: <entity> operations are done in rds,
import (
	ops "ac9/glad/ops/db"
	"log"
)

func WriteToDB(record any) error {
	db, err := ops.GetDB()
	if err != nil {
		log.Println("there is an error fetching the db", err)
	}
	if db == nil {
		log.Println("db is nil")
	}
	log.Println("inserting record now:", record)
	result := db.Create(record) // todo: check if upsert or insert
	if result.Error != nil {
		log.Println("error occurred in the write process", result.Error)
		return result.Error
	}
	return nil
}
