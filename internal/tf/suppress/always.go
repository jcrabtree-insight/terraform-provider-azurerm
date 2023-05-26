package suppress

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Always is a SchemaDiffSuppressFunc that always returns true.
// It can be used for the DiffSuppressFunc field in schema values.
// When used, Terraform will ignore any changes to that schema value when a plan is created.
func Always(_, _, _ string, _ *schema.ResourceData) bool {
	return true
}
