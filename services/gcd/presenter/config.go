package presenter

type Config struct {
	Version  int      `json:"version"`
	Timezone []string `json:"timezone"`
	Auth     Auth     `json:"auth"`
}

type Auth struct {
	ClientId     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	Domain       string `json:"domain"`
	Region       string `json:"region"`
	UserPoolId   string `json:"userPoolID"`
	Url          string `json:"url"`
}
