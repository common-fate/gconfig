package gconfig

func (r Role) MatchingRules(groups []string) []Rule {
	var matched []Rule

	for _, r := range r.Rules {
		for _, g := range groups {
			if r.Group == g {
				matched = append(matched, r)
			}
		}
	}

	return matched
}
