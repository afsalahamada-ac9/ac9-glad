package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Location struct {
	Street1 string `json:"Street_Address_1__c"`
	Street2 string `json:"Street_Address_2__c"`
	City    string `json:"City__c"`
	State   string `josn:"State__c"`
	Zip     string `json:"Postal_Or_Zip_Code__c"`
	Country string `json:"Country__c"`
}

func (l Location) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *Location) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source")
	}
	return json.Unmarshal(bytes, &l)
}
