package gconfig

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
)

type ErrInvalidAWSAccount struct {
	Account string
}

func (e *ErrInvalidAWSAccount) Error() string {
	return fmt.Sprintf("account %s is not a valid AWS account: must be 12 characters long", e.Account)
}

func (c *Config) Validate() error {
	var errs *multierror.Error

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

	// provider IDs must match
	providerMap := make(map[string]bool)
	for _, p := range c.Providers {
		_, ok := providerMap[p.ID]
		if ok {
			err := fmt.Errorf("duplicate provider ID %s", p.ID)
			err = printLintError(&p, err)
			errs = multierror.Append(errs, err)
		}

		if p.BastionAccountID != nil {
			err := mustBeAWSAccount(*p.BastionAccountID)
			if err != nil {
				err = printLintError(p, err)
				errs = multierror.Append(errs, err)
			}
		}

		providerMap[p.ID] = true
	}

	accountMap := make(map[string]bool)
	for _, a := range c.Accounts {
		accErrs := checkAccount(&a, accountMap, providerMap)
		errs = multierror.Append(errs, accErrs...)
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

func checkAccount(a *Account, accountMap map[string]bool, providerMap map[string]bool) []error {
	var errs []error

	_, ok := accountMap[a.ID]
	if ok {
		err := fmt.Errorf("duplicate account ID %s", a.ID)
		err = printLintError(a, err)
		errs = append(errs, err)
	}

	accountMap[a.ID] = true

	if a.parentId != nil && a.Provider != nil {
		err := fmt.Errorf("if accounts are grouped together only the top-level group may specify the provider, account %s is in a group but has a provider %s", a.ID, *a.Provider)
		err = printLintError(a, err)
		errs = append(errs, err)
	}

	if a.parentId == nil {
		if a.Provider == nil {
			err := fmt.Errorf("account %s must specify a provider", a.ID)
			err = printLintError(a, err)
			errs = append(errs, err)
		} else {
			_, ok := providerMap[*a.Provider]
			if !ok {
				err := fmt.Errorf("provider %s for account %s doesn't match any providers in your config", *a.Provider, a.ID)
				err = printLintError(a, err)
				errs = append(errs, err)
			}
		}
	}

	if a.AwsAccountID != nil {
		err := mustBeAWSAccount(*a.AwsAccountID)
		err = printLintError(a, err)
		if err != nil {
			errs = append(errs, err)
		}
	}

	for _, c := range a.Children {
		childErrs := checkAccount(&c, accountMap, providerMap)
		errs = append(errs, childErrs...)
	}

	return errs
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

func mustBeAWSAccount(acc string) error {
	if len(acc) != 12 {
		return &ErrInvalidAWSAccount{Account: acc}
	}
	return nil
}
