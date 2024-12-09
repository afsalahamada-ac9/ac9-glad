package entity

type SFPayload struct {
	Object string     `json:"object"`
	Items  []SFRecord `json:"items"`
}

type SFRecord struct {
	Operation string      `json:"operation"`
	Value     SFEventData `json:"value"`
}

type SFEventData struct {
	ExtId          string  `json:"Ext_Id,omitempty"`
	NumStudents    int     `json:"Number_Of_Students__c"`
	CTypeId        string  `json:"CType_Id__c"`
	Location       string  `json:"Location__c"`
	Timezone       *string `json:"Timezone__c"`
	MaxAttendees   int     `json:"Max_Attendees__c"`
	Country        string  `json:"Country__c"`
	ZipCode        string  `json:"Zip_Postal_Code__c"`
	State          string  `json:"State__c"`
	City           string  `json:"City__c"`
	StreetAddress1 *string `json:"Street_Address_1__c"`
	StreetAddress2 string  `json:"Street_Address_2__c"`
	Status         string  `json:"Status__c"`
	Notes          *string `json:"Notes__c"`
	WorkshopType   string  `json:"Workshop_Type__c"`
	EventStartDate string  `json:"Event_Start_Date__c"`
	EventEndDate   string  `json:"Event_End_Date__c"`
}

type SFTimingData struct {
	EndTime   string `json:"End_Time__c"`
	StartTime string `json:"Start_Time__c"`
	EndDate   string `json:"End_Date__c"`
	StartDate string `json:"Start_Date__c"`
	EventId   string `json:"Event__c"`
	Id        string `json:"Id,omitempty"`
}
