// Package presenter defines data structures for the API responses
package presenter

type ZoomDetails struct {
	ID         int64  `json:"id"`
	Date       string `json:"date"`
	StartTime  string `json:"start_time"`
	MeetingID  string `json:"meeting_id"`
	Password   string `json:"password"`
	MeetingURL string `json:"meeting_url"`
}

type ZoomWrapper struct {
	Zoom ZoomDetails `json:"zoom"`
}

type MetadataWrapper struct {
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Zoom ZoomSignature `json:"zoom"`
}

type ZoomSignature struct {
	Signature string `json:"signature"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details int    `json:"details"`
	TraceID int    `json:"trace_id"`
}
