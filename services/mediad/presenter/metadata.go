/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

type Metadata struct {
	Version     int64  `json:"version"`
	LastUpdated string `json:"lastUpdated"`
	Total       int    `json:"total"`
	URL         string `json:"url"`
}

type MetadataResponse struct {
	Quote *Metadata `json:"quote,omitempty"`
	Media *Metadata `json:"media,omitempty"`
}
