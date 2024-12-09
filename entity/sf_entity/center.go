// this represents the location entity in salesforce
package entity

type Center struct {
	Object    string       `json:"object"`
	Value     Center_value `json:"value"`
	Operation string       `json:"operation"`
}

type Center_value struct {
	Ext_id             string      `json:"Id" gorm:"column:ext_id"`
	Tenant_id          int         `json:"Tenant_id" gorm:"column:tenant_id"`
	Ext_name           string      `json:"Name" gorm:"column:ext_name"`
	Address            Location    `json:"address" gorm:"column:address"`
	Geo_Location       GeoLocation `json:"geolocation" gorm:"column:geo_location"`
	Capacity           int         `json:"Max_Capacity__c" gorm:"column:capacity"`
	Mode               string      `json:"Center_Mode__c" gorm:"column:mode"`
	Webpage            string      `json:"Center_URL__c" gorm:"column:webpage"`
	Is_national_center bool        `json:"Is_National_Center__c" gorm:"column:is_national_center"`
	Is_enabled         bool        `json:"Is_enable__c" gorm:"column:is_enabled"`
	Created_at         string      `json:"CreatedDate" gorm:"column:created_at"`
	Updated_at         string      `json:"UpdatedDate" gorm:"column:updated_at"`
}

func (*Center_value) TableName() string {
	return "center"
}

func (*Center) NewCenter(Ext_id string,
	Tenant_id int,
	Ext_name string,
	Address Location,
	Geo_Location GeoLocation,
	Capacity int,
	Mode string,
	Webpage string,
	Is_national_center bool,
	Is_enabled bool,
	Created_at string,
	Updated_at string) *Center_value {
	return &Center_value{Ext_id: Ext_id, Tenant_id: Tenant_id, Ext_name: Ext_name, Address: Address, Geo_Location: Geo_Location, Capacity: Capacity, Mode: Mode, Webpage: Webpage, Is_national_center: Is_national_center, Is_enabled: Is_enabled, Created_at: Created_at, Updated_at: Updated_at}
}
