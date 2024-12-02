/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import "sudhagar/glad/entity"

// Address
type Address struct {
	Street1 string `json:"street,omitempty"`
	Street2 string `json:"street_2,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
}

// Date/time
type DateTime struct {
	Date      string `json:"date,omitempty"`      // Only date in YYYY-MM-DD format
	StartTime string `json:"startTime,omitempty"` // Only time in HH:MM:SS format (SS is optional, default 00)
	EndTime   string `json:"endTime,omitempty"`
}

func (a *Address) CopyFrom(sa entity.CourseAddress) {
	a.Street1 = sa.Street1
	a.Street2 = sa.Street2
	a.City = sa.City
	a.State = sa.State
	a.Zip = sa.Zip
	a.Country = sa.Country
}

func (dt *DateTime) CopyFrom(sdt entity.CourseDateTime) {
	dt.Date = sdt.Date
	dt.StartTime = sdt.StartTime
	dt.EndTime = sdt.EndTime
}
