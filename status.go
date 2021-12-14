package gconfig

type ConfigStatus int

//go:generate go run github.com/alvaroloes/enumer -type=ConfigStatus -linecomment
const (
	ConfigStatusPending  ConfigStatus = iota + 1 // PENDING
	ConfigStatusApproved                         // APPROVED
)
