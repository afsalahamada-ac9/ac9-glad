package entity

import "errors"

type Token struct {
	AuthToken string `json:"access_token"`
}

func (t *Token) Validate() error {
	if t.AuthToken == "" {
		return errors.New("invalid entity")
	}
	return nil
}
