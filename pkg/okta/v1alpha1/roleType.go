package gcoktav1alpha1

type RoleType int

//go:generate go run github.com/alvaroloes/enumer -type=RoleType -linecomment
const (
	RoleTypeOktaGroup RoleType = iota + 1 // OKTA_GROUP
)
