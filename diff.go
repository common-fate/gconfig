package gconfig

import (
	"fmt"
)

type Changes struct {
	DeleteUsers  []string
	AddUsers     []string
	UpdateUsers  []UpdateUser
	DeleteAdmins []string
	AddAdmins    []string
}

// Empty returns true if no changes need to be made
func (c Changes) Empty() bool {
	return (len(c.DeleteUsers) == 0 &&
		len(c.AddUsers) == 0 &&
		len(c.UpdateUsers) == 0 &&
		len(c.DeleteAdmins) == 0 &&
		len(c.AddAdmins) == 0)
}

type UpdateUser struct {
	Email string
	// whether to make the user an admin or not
	Admin bool
}

type ErrNoInPlaceUpdates struct {
	Type  string
	ID    string
	Field string
	Old   string
	New   string
}

func (e *ErrNoInPlaceUpdates) Error() string {
	return fmt.Sprintf("Granted doesn't yet support in-place updates for %s %s (field %s changed from %s to %s)", e.Type, e.ID, e.Field, e.Old, e.New)
}

func (c *Config) ChangesFrom(old Config) (Changes, error) {
	var ch Changes

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
