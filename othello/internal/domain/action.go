package domain

type Action int8

const (
	ActionUnknown Action = iota
	ActionPutDisk
	ActionGiveUp
	ActionPass
)
