package entity

type Course struct {
	Object    string       `json:"object"`
	Value     Course_value `json:"value"`
	Operation string       `json:"operation"`
}

type Course_value struct {
	Url           string   `json:"url" gorm:"url"`
	Max_attendees int      `json:"Max_attendees__c" gorm:"max_attendees"`
	Address       Location `json:"Address" gorm:"address"`
	Tenant_id     int      `json:"Tenant_id" gorm:"tenant_id"`
	Ext_id        string   `json:"Id" gorm:"ext_id"`
	Name          string   `json:"Name" gorm:"name"`
	Timezone      string   `json:"Timezone__c" gorm:"timezone"`
	Mode          string   `json:"Mode" gorm:"mode"`
	Center_id     int      `json:"Center_id" gorm:"center_id"`
	Status        string   `json:"Status__c" gorm:"status"`
	Created_at    string   `json:"CreatedDate" gorm:"created_at"`
	Num_attendees int      `json:"Number_Of_Students__c" gorm:"num_attendees"`
	Product_id    int      `json:"product_id" gorm:"product_id"`
	Updated_at    string   `json:"LastModifiedDate" gorm:"updated_at"`
	Notes         string   `json:"Notes__c" gorm:"notes"`
	Short_url     string   `json:"Short_url" gorm:"short_url"`
}

func (*Course_value) TableName() string {
	return "course"
}

func (*Course) NewCourse(Url string,
	Max_attendees int,
	Address Location,
	Tenant_id int,
	Ext_id string,
	Name string,
	Timezone string,
	Mode string,
	Center_id int,
	Status string,
	Created_at string,
	Num_attendees int,
	Product_id int,
	Updated_at string,
	Notes string,
	Short_url string,
) *Course_value {
	return &Course_value{Url: Url, Max_attendees: Max_attendees, Address: Address, Tenant_id: Tenant_id, Ext_id: Ext_id, Name: Name, Timezone: Timezone, Mode: Mode, Center_id: Center_id, Status: Status, Created_at: Created_at, Num_attendees: Num_attendees, Product_id: Product_id, Updated_at: Updated_at, Notes: Notes, Short_url: Short_url}
}
