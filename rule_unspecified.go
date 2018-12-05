// +build !linux

package netlink

// RuleAdd adds a rule to the system.
// Equivalent to: ip rule add
func RuleAdd(rule *Rule) error {
	return ErrNotImplemented
}

// RuleDel deletes a rule from the system.
// Equivalent to: ip rule del
func RuleDel(rule *Rule) error {
	return ErrNotImplemented
}

// RuleList lists rules in the system.
// Equivalent to: ip rule list
func RuleList(family int) ([]Rule, error) {
	return nil, ErrNotImplemented
}
