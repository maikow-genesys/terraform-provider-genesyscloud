package outbound_ruleset

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	gcloud "terraform-provider-genesyscloud/genesyscloud"
)

// dataSourceOutboundRulesetRead retrieves by name the id in question
func dataSourceOutboundRulesetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkConfig := meta.(*gcloud.ProviderMeta).ClientConfig
	proxy := newOutboundRulesetProxy(sdkConfig)

	name := d.Get("name").(string)

	return gcloud.WithRetries(ctx, 15*time.Second, func() *resource.RetryError {
		rulesetId, retryable, err := proxy.getOutboundRulesetIdByName(ctx, name)

		if err != nil && !retryable {
			return resource.NonRetryableError(fmt.Errorf("Error ruleset %s: %s", name, err))
		}

		if retryable {
			return resource.RetryableError(fmt.Errorf("No ruleset found with name %s", name))
		}

		d.SetId(rulesetId)
		return nil
	})
}
