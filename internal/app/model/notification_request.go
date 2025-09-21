package model

import "errors"

type NotificationRequest struct {
	Token     string            `json:"token"`
	Topic     string            `json:"topic"`
	Condition string            `json:"condition"`
	Tokens    []string          `json:"tokens"`
	Title     string            `json:"title"`
	Body      string            `json:"body"`
	Data      map[string]string `json:"data"`
}

func (r NotificationRequest) Validate() error {
	count := 0

	if r.Token != "" {
		count++
	}
	if r.Topic != "" {
		count++
	}
	if r.Condition != "" {
		count++
	}

	if r.Tokens != nil {
		if len(r.Tokens) == 0 {
			return errors.New("tokens tidak boleh kosong")
		}
		count++
	}

	if count == 0 {
		return errors.New("harus ada minimal satu dari token/topic/condition/tokens")
	}
	if count > 1 {
		return errors.New("hanya boleh isi salah satu dari token/topic/condition/tokens")
	}

	return nil
}
