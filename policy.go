package gconfig

type RulePolicy int

//go:generate go run github.com/alvaroloes/enumer -type=RulePolicy -linecomment
const (
	RulePolicyAllow           RulePolicy = iota + 1 // allow
	RulePolicyRequireApproval                       // requireApproval
	RulePolicyRequireReason                         // requireReason
)
