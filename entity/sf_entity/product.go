package entity

type Product struct {
	Value     Product_value `json:"value"`
	Object    string        `json:"object"`
	Operation string        `json:"operation"`
}

type Product_value struct {
	Updated_at string `json:"LastModifiedDate" gorm:"column:updated_at"`
	Created_at string `json:"CreatedDate" gorm:"column:created_at"`
	//Is_deleted          bool   `json:"IsDeleted" gorm:"column:is_deleted"`
	Format              string `json:"Online_Or_In_Person__c" gorm:"column:format"`
	Max_Attendees       int32  `json:"Max_Attendees__c,omitempty" gorm:"column:max_attendees"`
	Listing_Visibity    string `json:"Listing_Visibity__c,omitempty" gorm:"column:visibility"` // todo: rename visibity to visibility -> spelling mistake in sf
	Event_Duration      int32  `json:"Event_Duration__c,omitempty" gorm:"column:duration_days"`
	Product             string `json:"Product__c,omitempty" gorm:"column:base_product_ext_id"`
	CType               string `json:"CType_Id__c" gorm:"column:ctype"`
	Title               string `json:"Title__c" gorm:"column:title"`
	Name                string `json:"name" gorm:"column:ext_name"`
	TenantID            int32  `json:"Tenant_id" gorm:"column:tenant_id"`
	ExtID               string `json:"Id" gorm:"column:ext_id"`
	Base_product_ext_id string `json:"base_product_ext_id" gorm:"base_product_ext_id"`
	Is_auto_approve     bool   `json:"Auto_Approve_Event__c" gorm:"column:is_auto_approve"`
}

func (*Product_value) TableName() string {
	return "product"
}

func (*Product) NewProduct(Updated_at string,
	Created_at string,
	/*Is_deleted bool,*/ // todo: is_deleted is not in the data model, update it and add this parameter back
	Format string,
	Max_Attendees int32,
	Listing_Visibity string,
	Event_Duration int32,
	Product string,
	CType string,
	Title string,
	Name string,
	TenantID int32,
	ExtID string,
	Base_product_ext_id string,
	Is_auto_approve bool) *Product_value {
	return &Product_value{Updated_at: Updated_at, Created_at: Created_at /*Is_deleted: Is_deleted,*/, Format: Format, Max_Attendees: Max_Attendees, Listing_Visibity: Listing_Visibity, Event_Duration: Event_Duration, Product: Product, CType: CType, Title: Title, Name: Name, TenantID: TenantID, ExtID: ExtID, Base_product_ext_id: Base_product_ext_id, Is_auto_approve: Is_auto_approve}
}
