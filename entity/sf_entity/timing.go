package entity

type Timing struct {
	Object    string       `json:"object"`
	Operation string       `json:"operation"`
	Value     Timing_value `json:"value"`
}

type Timing_value struct {
	Course_id   int    `json:"Course_id" gorm:"column:course_id"` // todo: map with id from course
	Ext_id      string `json:"Id" gorm:"column:ext_id"`
	Course_date string `json:"Start_Date__c" gorm:"column:course_date"`
	Start_time  string `json:"Start_Time__c" gorm:"column:start_time"`
	End_time    string `json:"End_Time__c" gorm:"column:end_time"`
	Updated_at  string `json:"LastModifiedDate" gorm:"column:updated_at"`
	Created_at  string `json:"CreatedDate" gorm:"column:created_at"`
}

func (*Timing_value) TableName() string {
	return "course_timing"
}

func (*Timing) NewTiming(Course_id int,
	Ext_id string,
	Course_date string,
	Start_time string,
	End_time string,
	Updated_at string,
	Created_at string,
) *Timing_value {
	return &Timing_value{Course_id,
		Ext_id,
		Course_date,
		Start_time,
		End_time,
		Updated_at,
		Created_at,
	}
}
