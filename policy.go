package gconfig

type Policy int

//go:generate go run github.com/alvaroloes/enumer -type=Policy -linecomment
const (
	PolicyAllow           Policy = iota + 1 // allow
	PolicyRequireApproval                   // requireApproval
	PolicyRequireReason                     // requireReason
)
