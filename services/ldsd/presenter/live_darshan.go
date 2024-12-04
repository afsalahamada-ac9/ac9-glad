/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

// Package presenter defines data structures for the API responses
package presenter

type LiveDarshanConfig struct {
	Zoom ZoomInfo `json:"zoom"`
}

type ZoomInfo struct {
	Signature   string `json:"signature"`
	DisplayName string `json:"displayName"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details int    `json:"details"`
	TraceID int    `json:"traceID"`
}

type LiveDarshan struct {
	ID         int64  `json:"id"`
	Date       string `json:"date"`
	StartTime  string `json:"startTime"`
	MeetingID  string `json:"meetingID"`
	Password   string `json:"password"`
	MeetingURL string `json:"meetingURL"`
}
