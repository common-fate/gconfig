package gconfig

import (
	"fmt"
)

type Changes struct {
	DeleteProviders []string
	AddProviders    []Provider
	DeleteUsers     []string
	AddUsers        []string
	UpdateUsers     []UpdateUser
	DeleteAdmins    []string
	AddAdmins       []string
}

type UpdateUser struct {
	Email string
	// whether to make the user an admin or not
	Admin bool
}

type ErrNoInPlaceUpdates struct {
	Type string
	ID   string
}

func (e *ErrNoInPlaceUpdates) Error() string {
	return fmt.Sprintf("Granted doesn't yet support in-place updates for %s %s", e.Type, e.ID)
}

func (c *Config) ChangesFrom(old Config) (Changes, error) {
	var ch Changes

	oldProvidersToDelete := make(map[string]Provider)
	newProvidersToAdd := make(map[string]Provider)

	for _, p := range old.Providers {
		oldProvidersToDelete[p.ID] = p
	}

	for _, p := range c.Providers {
		if old, ok := oldProvidersToDelete[p.ID]; ok {
			// resource is common between old and new

			if ptrstr(old.BastionAccountID) != ptrstr(p.BastionAccountID) {
				return Changes{}, &ErrNoInPlaceUpdates{Type: "provider", ID: p.ID}
			}
			if ptrstr(old.InstanceARN) != ptrstr(p.InstanceARN) {
				return Changes{}, &ErrNoInPlaceUpdates{Type: "provider", ID: p.ID}
			}
			if ptrstr(old.IdentityStoreID) != ptrstr(p.IdentityStoreID) {
				return Changes{}, &ErrNoInPlaceUpdates{Type: "provider", ID: p.ID}
			}

			delete(oldProvidersToDelete, p.ID)
		} else {
			// resource is new
			newProvidersToAdd[p.ID] = p
			ch.AddProviders = append(ch.AddProviders, p)
		}
	}

	for id := range oldProvidersToDelete {
		ch.DeleteProviders = append(ch.DeleteProviders, id)
	}

	// user diffing
	// the full pool of users is both admins and users.
	// a user could:
	// 1. add a new user
	// 2. delete a user
	// 3. update a user (change them from a user to an admin, or vice versa)
	type userDetails struct {
		IsAdmin bool
	}

	oldUsersToDelete := make(map[string]userDetails)
	allNewUsers := make(map[string]userDetails)

	for _, u := range old.Users {
		oldUsersToDelete[u.Email] = userDetails{IsAdmin: false}
	}

	for _, u := range old.Admins {
		oldUsersToDelete[u.Email] = userDetails{IsAdmin: true}
	}

	// combine admins and users into one list
	for _, u := range c.Users {
		allNewUsers[u.Email] = userDetails{IsAdmin: false}
	}
	for _, u := range c.Admins {
		allNewUsers[u.Email] = userDetails{IsAdmin: true}
	}

	for email, new := range allNewUsers {
		if old, ok := oldUsersToDelete[email]; ok {
			// user is common between old and new, so don't delete them
			if old.IsAdmin != new.IsAdmin {
				// user has been updated
				ch.UpdateUsers = append(ch.UpdateUsers, UpdateUser{
					Email: email,
					Admin: new.IsAdmin,
				})
			}
			delete(oldUsersToDelete, email)

		} else {
			// user is new
			if new.IsAdmin {
				ch.AddAdmins = append(ch.AddAdmins, email)
			} else {
				ch.AddUsers = append(ch.AddUsers, email)
			}
		}
	}

	for email, u := range oldUsersToDelete {
		if u.IsAdmin {
			ch.DeleteAdmins = append(ch.DeleteAdmins, email)
		} else {
			ch.DeleteUsers = append(ch.DeleteUsers, email)
		}
	}

	return ch, nil
}

// ptrstr converts a string pointer to a string. It returns an empty string "" if the pointer is nil.
func ptrstr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
