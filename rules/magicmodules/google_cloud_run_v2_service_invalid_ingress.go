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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// GoogleCloudRunV2ServiceInvalidIngressRule checks the pattern is valid
type GoogleCloudRunV2ServiceInvalidIngressRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewGoogleCloudRunV2ServiceInvalidIngressRule returns new rule with default attributes
func NewGoogleCloudRunV2ServiceInvalidIngressRule() *GoogleCloudRunV2ServiceInvalidIngressRule {
	return &GoogleCloudRunV2ServiceInvalidIngressRule{
		resourceType:  "google_cloud_run_v2_service",
		attributeName: "ingress",
	}
}

// Name returns the rule name
func (r *GoogleCloudRunV2ServiceInvalidIngressRule) Name() string {
	return "google_cloud_run_v2_service_invalid_ingress"
}

// Enabled returns whether the rule is enabled by default
func (r *GoogleCloudRunV2ServiceInvalidIngressRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *GoogleCloudRunV2ServiceInvalidIngressRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *GoogleCloudRunV2ServiceInvalidIngressRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *GoogleCloudRunV2ServiceInvalidIngressRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		validateFunc := validation.StringInSlice([]string{"INGRESS_TRAFFIC_ALL", "INGRESS_TRAFFIC_INTERNAL_ONLY", "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER", ""}, false)

		err = runner.EnsureNoError(err, func() error {
			_, errors := validateFunc(val, r.attributeName)
			for _, err := range errors {
				runner.EmitIssue(r, err.Error(), attribute.Expr.Range())
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
