package gconfig

// Tests the the correct ordering of rules is observed
// func Test_RuleSelector(t *testing.T) {
// 	requireReason, err := structpb.NewStruct(map[string]interface{}{
// 		"reason": true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	allow, err := structpb.NewStruct(map[string]interface{}{
// 		"allow": true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	approval, err := structpb.NewStruct(map[string]interface{}{
// 		"approval": true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cert := &x509.Certificate{Subject: pkix.Name{OrganizationalUnit: []string{"developer", "tester"}}}
// 	rules := []*gconfigv1alpha1.Rule{{Policy: requireReason, Group: "developer"}, {Policy: allow, Group: "tester"}, {Policy: approval, Group: "reasonNeeders"}}
// 	rule, _ := RuleSelector(cert, rules)
// 	assert.Equal(t, rules[1], rule)

// 	// user does not have tester group so the developer rule is retured
// 	cert = &x509.Certificate{Subject: pkix.Name{OrganizationalUnit: []string{"developer"}}}
// 	rule, _ = RuleSelector(cert, rules)
// 	assert.Equal(t, rules[0], rule)
// }
// func Test_RuleSelector1(t *testing.T) {
// 	requireReason, err := structpb.NewStruct(map[string]interface{}{
// 		"reason": true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	allow, err := structpb.NewStruct(map[string]interface{}{
// 		"allow": true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	approval, err := structpb.NewStruct(map[string]interface{}{
// 		"approval": true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	cert := &x509.Certificate{Subject: pkix.Name{OrganizationalUnit: []string{"developer", "tester"}}}
// 	rules := []*gconfigv1alpha1.Rule{{Policy: requireReason, Group: "developer"}, {Policy: allow, Group: "tester"}, {Policy: approval, Group: "reasonNeeders"}}
// 	rule, _ := RuleSelector(cert, rules)
// 	assert.Equal(t, rules[1], rule)

// 	// user does not have tester group so the developer rule is retured
// 	cert = &x509.Certificate{Subject: pkix.Name{OrganizationalUnit: []string{"developer"}}}
// 	rule, _ = RuleSelector(cert, rules)
// 	assert.Equal(t, rules[0], rule)
// }
