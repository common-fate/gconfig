package gconfig

type Status int

//go:generate go run github.com/alvaroloes/enumer -type=Status -linecomment
const (
	StatusPending  Status = iota + 1 // PENDING
	StatusApproved                   // APPROVED
)
