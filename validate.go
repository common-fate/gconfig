package gconfig

import (
	"fmt"
	"strings"

	gconfigv1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	"github.com/hashicorp/go-multierror"
)

type ErrInvalidAWSAccount struct {
	Account string
}

func (e *ErrInvalidAWSAccount) Error() string {
	return fmt.Sprintf("account %s is not a valid AWS account: must be 12 characters long", e.Account)
}

// This method should be used to validate users and groups
func (c *Config) Validate() error {
	var errs *multierror.Error

	providerMap := make(map[string]*gconfigv1alpha1.Provider)
	accountMap := make(map[string]*gconfigv1alpha1.Account)
	for _, p := range c.GetProviders() {
		if _, ok := providerMap[p.Id]; ok {
			// this should never happen as we have server-side validation to avoid duplicate provider IDs.
			return fmt.Errorf("duplicate provider %s", p.Id)
		}
		providerMap[p.Id] = p

		for _, acc := range p.Accounts {
			collectAccount(acc, accountMap)
		}
	}

	// group members must be defined in the "admins" or "users"
	userMap := make(map[string]bool)
	for _, a := range c.Admins {
		userMap[a.Email] = true
	}
	for _, u := range c.Users {
		userMap[u.Email] = true
	}

	groupMap := make(map[string]bool)

	for _, g := range c.Groups {
		_, ok := groupMap[g.ID]
		if ok {
			err := fmt.Errorf("duplicate group ID %s", g.ID)
			err = printLintError(g, err)
			errs = multierror.Append(errs, err)
		}

		for _, m := range g.Members {
			_, ok := userMap[m.Email]
			if !ok {
				err := fmt.Errorf("%s must be defined as a user or an admin", m.Email)
				err = printLintError(m, err)
				errs = multierror.Append(errs, err)
			}
		}
		groupMap[g.ID] = true
	}

	if errs.ErrorOrNil() != nil {
		errs.ErrorFormat = func(all []error) string {
			var errStrs []string
			for _, e := range all {
				errStrs = append(errStrs, e.Error())
			}
			return strings.Join(errStrs, "\n")
		}
		return errs
	} else {
		return nil
	}
}

func collectAccount(a *gconfigv1alpha1.Account, accountMap map[string]*gconfigv1alpha1.Account) {
	accountMap[a.Id] = a
	for _, child := range a.Children {
		collectAccount(child, accountMap)
	}
}

// filePositioners can get their positions in a YAML file.
// all of our config nodes are filePositioners
type filePositioner interface {
	filePosition() *FilePosition
}

// printLintError adds file position information to the error message if it exists.
func printLintError(p filePositioner, err error) error {
	pos := p.filePosition()
	if pos == nil {
		return err
	}
	return fmt.Errorf("%s:%d:%d: %s", pos.Filename, pos.Line, pos.Col, err)
}
