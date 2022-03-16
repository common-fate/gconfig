package gcoktav1alpha1

import (
	"crypto/sha256"
	"fmt"
)

type Changes struct {
	DeleteUsers  []string
	AddUsers     []string
	UpdateUsers  []UpdateUser
	DeleteAdmins []string
	AddAdmins    []string
	AddRoles     []string
	DeleteRoles  []string
	UpdateRoles  []UpdateRole
	// @TODO:
	// - AddGroups
	// - DeleteGroups
	// - UpdateGroups
}

// Empty returns true if no changes need to be made
func (c Changes) Empty() bool {
	return (len(c.DeleteUsers) == 0 &&
		len(c.AddUsers) == 0 &&
		len(c.UpdateUsers) == 0 &&
		len(c.DeleteAdmins) == 0 &&
		len(c.AddRoles) == 0 &&
		len(c.DeleteRoles) == 0 &&
		len(c.UpdateRoles) == 0 &&
		len(c.AddAdmins) == 0)
}

type UpdateUser struct {
	Email string
	// whether to make the user an admin or not
	Admin bool
}

type UpdateRole struct {
	ID string
	// String describing what field changed
	AlteredField []string // @TODO: potentially make into enum?

	AddRules    []AddRule
	DeleteRules []DeleteRule
}

type AddRule struct {
	Group      string
	Policy     string
	Breakglass bool
}
type DeleteRule struct {
	Group      string
	Policy     string
	Breakglass bool
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
	type ruleDetails struct {
		policy     string
		group      string
		Breakglass bool
	}

	oldUsersToDelete := make(map[string]userDetails)
	allNewUsers := make(map[string]userDetails)

	for _, u := range old.Admins {
		oldUsersToDelete[u.Email] = userDetails{IsAdmin: true}
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

	allNewRoles := make(map[string]Role)
	allPrevRoles := make(map[string]Role)
	// for _, u := range c.Roles {

	// 	append(ch.UpdateUsers, UpdateUser{
	// 				Email: email,
	// 				Admin: new.IsAdmin,
	// 			})
	// 	allNewUsers[u.Email] = userDetails{IsAdmin: true}
	// }
	// oldRoles  := make(map[string]*Role)

	for _, u := range c.Roles {
		allNewRoles[u.ID] = *u
		// allNewUsers[u.Email] = userDetails{IsAdmin: true}
	}
	for _, o := range old.Roles {
		allPrevRoles[o.ID] = *o
	}

	// for each role, check if it's been updated
	for id, new := range allNewRoles {
		// If theres a match in ID, that means the roll hasn't been deleted,
		// either updated or stayed the same
		if old, ok := allPrevRoles[id]; ok {
			// role is common between old and new, so don't delete them
			if old.ID == new.ID {
				// instantiate a new UpdateRole obj for eact role, only add it to the list
				// if the role has bene updated
				ruleUpdateObj := UpdateRole{
					ID:           id,
					AlteredField: []string{},
				}

				oldRuleCount := len(old.Rules)
				newRuleCount := len(new.Rules)

				// If there's a rule count difference
				if oldRuleCount != newRuleCount {
					ruleUpdateObj.AlteredField = append(ruleUpdateObj.AlteredField, "Rules")
				}

				//loop through old rules and hash the combination of policy+group
				//Make that the key of the map
				//Do the same with the new rules
				//loop through new rules and key the old rules with the hash of each -> if doesnt exist: create new rule
				//If we find a difference add altered field

				oldRules := make(map[[32]byte]ruleDetails)
				newRules := make(map[[32]byte]ruleDetails)

				for _, rule := range old.Rules {
					hash := sha256.Sum256([]byte(rule.Policy.Policy + rule.Group))
					oldRules[hash] = ruleDetails{group: rule.Group, policy: rule.Policy.Policy, Breakglass: rule.Breakglass}

				}

				for _, rule := range new.Rules {
					hash := sha256.Sum256([]byte(rule.Policy.Policy + rule.Group))
					newRules[hash] = ruleDetails{group: rule.Group, policy: rule.Policy.Policy, Breakglass: rule.Breakglass}

				}

				updatedRole := &UpdateRole{}

				for hash, new_rule := range newRules {
					//if we dont find it in the old hash then its new or edited
					if _, ok := oldRules[hash]; !ok {
						updatedRole = &UpdateRole{
							ID:           old.ID,
							AlteredField: append(ruleUpdateObj.AlteredField, "Rules"),
							AddRules:     append(updatedRole.AddRules, AddRule{Group: new_rule.group, Policy: new_rule.policy, Breakglass: new_rule.Breakglass}),
						}

					} else {
						//if we find the hash that means this rule has stayed the same
						delete(oldRules, hash)
					}

				}

				//add all the deleted rules
				//leftover old rules are deleted
				for _, rule := range oldRules {

					updatedRole = &UpdateRole{
						ID:           old.ID,
						AlteredField: append(ruleUpdateObj.AlteredField, "Rules"),
						AddRules:     updatedRole.AddRules,
						DeleteRules:  append(ruleUpdateObj.DeleteRules, DeleteRule{Group: rule.group, Policy: rule.policy, Breakglass: rule.Breakglass}),
					}
				}

				//set
				ruleUpdateObj = *updatedRole

				// @TODO: Support deep rule diffing...

				// If there's been any changes to the Role, add it to the UpdatedRoles list
				if len(ruleUpdateObj.AlteredField) > 0 {
					ch.UpdateRoles = append(ch.UpdateRoles, ruleUpdateObj)
				}

			}
			delete(allPrevRoles, id)

		} else {
			// role is new
			ch.AddRoles = append(ch.AddRoles, new.ID)
		}
	}

	for id, old := range allPrevRoles {
		// if old role is not in new, then it has been deleted
		if _, ok := allNewRoles[id]; !ok {
			ch.DeleteRoles = append(ch.DeleteRoles, old.ID)
		}
	}

	return ch, nil
}
