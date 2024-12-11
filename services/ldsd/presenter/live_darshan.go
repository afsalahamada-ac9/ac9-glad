/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

// Package presenter defines data structures for the API responses
package presenter

import (
	"ac9/glad/pkg/id"
	"ac9/glad/services/ldsd/entity"
	"net/url"
	"strings"
)

// LiveDarshanConfig contains the zoom service configuration for live darshan
type LiveDarshanConfig struct {
	Zoom ZoomInfo `json:"zoom"`
}

// ZoomInfo contains the zoom information
type ZoomInfo struct {
	Signature   string `json:"signature"`
	DisplayName string `json:"displayName"`
}

// TODO: Move to a common location
// ErrorResponse sent after API handling failure
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details int    `json:"details"`
	TraceID int    `json:"traceID"`
}

// LiveDarshanReq object is sent to create the live darshan
type LiveDarshanReq struct {
	Date       string `json:"date"`
	StartTime  string `json:"startTime"`
	MeetingID  string `json:"meetingID,omitempty"`
	Password   string `json:"password,omitempty"`
	MeetingURL string `json:"meetingURL,omitempty"`
}

// LiveDarshanResponse object is sent in create live darshan response
type LiveDarshanResponse struct {
	ID id.ID `json:"id"`
}

// LiveDarshan contains the complete live darshan entity
type LiveDarshan struct {
	LiveDarshanReq
	LiveDarshanResponse
}

// FromEntityLiveDarshan creates live darshan response from entity
func (ld *LiveDarshan) FromEntityLiveDarshan(e *entity.LiveDarshan) error {
	ld.ID = e.ID
	ld.MeetingURL = e.MeetingURL
	ld.Date = e.Date.Format("2006-01-02")
	ld.StartTime = e.StartTime.Format("15:04:00")

	u, _ := url.Parse(e.MeetingURL)

	// Get zoom meeting id via last path segment
	segments := strings.Split(strings.Trim(u.Path, "/"), "/")
	ld.MeetingID = segments[len(segments)-1]

	// Get (encrypted) password from query parameter
	// Note: There is no way to decrypt the zoom password
	ld.Password = u.Query().Get("pwd")

	return nil
}
