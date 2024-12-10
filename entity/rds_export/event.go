package entity

type Event_item struct {
	Operation string `json:"operation"`
	Items     []Data `json:"items"`
}

type Data struct {
	Operation string `json:"operation"`
	Value     Value  `json:"value"`
}

type Value struct {
	Cloud_Db_ID__c        string  `json:"Cloud_Db_ID__c"`
	Number_Of_Students__c int     `json:"Number_Of_Students__c"`
	CType_Id__c           string  `json:"CType_Id__c"`
	Location__c           string  `json:"Location__c"`
	Timezone__c           *string `json:"Timezone__c"`
	Max_Attendees__c      int     `json:"Max_Attendees__c"`
	Country__c            string  `json:"Country__c"`
	Zip_Postal_Code__c    string  `json:"Zip_Postal_Code__c"`
	State__c              string  `json:"State__c"`
	City__c               string  `json:"City__c"`
	Street_Address_2__c   string  `json:"Street_Address_2__c"`
	Street_Address_1__c   string  `json:"Street_Address_1__c"`
	Status__c             string  `json:"Status__c"`
	Notes__c              *string `json:"Notes__c"`
	Workshop_Type__c      string  `json:"Workshop_Type__c"`
	Event_Start_Date__c   string  `json:"Event_Start_Date__c"`
	Event_End_Date__c     string  `json:"Event_End_Date__c"`
}
