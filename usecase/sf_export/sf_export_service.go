package service

import (
	"ac9/glad/entity"
	infra "ac9/glad/ops/db"
	pkgentity "ac9/glad/pkg/id"
	util "ac9/glad/pkg/util"
	"ac9/glad/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SFExportService struct {
	courseRepo *repository.CoursePGSQL
	timingRepo *repository.TimingPGSQL
	sfEndpoint string
}

func NewSFExportService() (*SFExportService, error) {
	sqlDB, err := infra.GetDB()
	if err != nil {
		return nil, err
	}
	db, err := sqlDB.DB()
	if err != nil {
		return nil, err
	}
	return &SFExportService{
		courseRepo: repository.NewCoursePGSQL(db),
		timingRepo: repository.NewTimingPGSQL(db),
		sfEndpoint: "https://aol-dev--awspoc.sandbox.my.salesforce.com/services/apexrest/handleAolEvent",
	}, nil
}

func (s *SFExportService) ExportToSF(courseID pkgentity.ID) error {
	// Get course data
	course, err := s.courseRepo.Get(courseID)
	if err != nil {
		return fmt.Errorf("failed to get course: %w", err)
	}

	// Transform course data to SF format
	sfEvent := entity.SFEventData{
		ExtId:        *course.ExtID,
		NumStudents:  int(course.NumAttendees),
		MaxAttendees: int(course.MaxAttendees),
		Notes:        &course.Notes,
		Status:       string(course.Status),
	}

	if course.Address.Validate() == nil {
		sfEvent.Country = course.Address.Country
		sfEvent.City = course.Address.City
		sfEvent.State = course.Address.State
		sfEvent.ZipCode = course.Address.Zip
		sfEvent.StreetAddress2 = course.Address.Street2
		sfEvent.StreetAddress1 = &course.Address.Street1
	}

	// Create payload
	payload := []entity.SFPayload{
		{
			Object: "Event__c",
			Items: []entity.SFRecord{
				{
					Operation: "Insert", // or "Update" based on logic
					Value:     sfEvent,
				},
			},
		},
	}

	// Send to SF
	return s.sendToSF(payload)
}

func (s *SFExportService) sendToSF(payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	client := http.Client{}
	request, err := http.NewRequest("POST", s.sfEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("there was an error generating the request")
	}
	token, err := util.GenerateTokens()
	if err != nil {
		log.Println("error generating the tokens", err)
	}
	request.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return fmt.Errorf("failed to send to SF: %w", err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("error performing the request", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("SF returned non-200 status: %d", resp.StatusCode)
	}
	log.Println("response from SF:", resp, resp.Body)
	return nil
}
