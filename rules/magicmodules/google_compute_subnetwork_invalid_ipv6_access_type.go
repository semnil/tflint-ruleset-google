// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleComputeSubnetworkInvalidIpv6AccessTypeRule checks the pattern is valid
type GoogleComputeSubnetworkInvalidIpv6AccessTypeRule struct {
	resourceType  string
	attributeName string
}

// NewGoogleComputeSubnetworkInvalidIpv6AccessTypeRule returns new rule with default attributes
func NewGoogleComputeSubnetworkInvalidIpv6AccessTypeRule() *GoogleComputeSubnetworkInvalidIpv6AccessTypeRule {
	return &GoogleComputeSubnetworkInvalidIpv6AccessTypeRule{
		resourceType:  "google_compute_subnetwork",
		attributeName: "ipv6_access_type",
	}
}

// Name returns the rule name
func (r *GoogleComputeSubnetworkInvalidIpv6AccessTypeRule) Name() string {
	return "google_compute_subnetwork_invalid_ipv6_access_type"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleComputeSubnetworkInvalidIpv6AccessTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleComputeSubnetworkInvalidIpv6AccessTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleComputeSubnetworkInvalidIpv6AccessTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleComputeSubnetworkInvalidIpv6AccessTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"EXTERNAL", ""}, false)

		return runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssueOnExpr(r, err.Error(), attribute.Expr)
			}
			return nil
		})
	})
}
