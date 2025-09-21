package model

type NotificationType struct {
	Token     string
	Topic     string
	Condition string
	Tokens    []string
}
