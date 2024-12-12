/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

type Config struct {
	Version  int      `json:"version"`
	Timezone []string `json:"timezone"`
	Auth     Auth     `json:"auth"`
}

type Auth struct {
	ClientID     string   `json:"clientID"`
	ClientSecret string   `json:"clientSecret"`
	Domain       string   `json:"domain"`
	Region       string   `json:"region"`
	UserPoolID   string   `json:"userPoolID"`
	URL          string   `json:"url"`
	OAuthScopes  []string `json:"oauthScopes"`
}
