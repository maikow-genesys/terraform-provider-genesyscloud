package external_contacts_organization

import (
	"fmt"
	"strconv"
	"strings"
	"terraform-provider-genesyscloud/genesyscloud/provider"
	"terraform-provider-genesyscloud/genesyscloud/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mypurecloud/platform-client-sdk-go/v146/platformclientv2"
)

func TestAccResourceExternalContacts(t *testing.T) {
	var (
		resourceLabel     = "externalOrganizationAbc"
		name              = "ABCCorp"
		phoneDisplay      = "+1 321-700-1243"
		countryCode       = "US"
		address           = "1011 New Hope St"
		city              = "Norristown"
		state             = "PA"
		postalCode        = "19401"
		twitterId         = "twitterId"
		twitterName       = "twitterName"
		twitterScreenName = "twitterScreenname"
		symbol            = "ABC"
		exchange          = "NYSE"
		tags              = []string{
			strconv.Quote("news"),
			strconv.Quote("channel"),
		}
		symbolUpdate       = "CBA"
		twitterUpdatedName = "twitterSecond"
		externalsystemurl  = "https://externalsystemurl.com"
	)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { util.TestAccPreCheck(t) },
		ProviderFactories: provider.GetProviderFactories(providerResources, providerDataSources),
		Steps: []resource.TestStep{
			{
				// Create
				Config: GenerateBasicExternalOrganizationResource(
					resourceLabel,
					name,
					phoneDisplay, countryCode,
					address, city, state, postalCode, countryCode,
					twitterId, twitterName, twitterScreenName,
					symbol, exchange,
					tags,
					"",
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "name", name),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "phone_number.0.display", phoneDisplay),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "phone_number.0.country_code", countryCode),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.address1", address),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.city", city),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.state", state),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.postal_code", postalCode),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.country_code", countryCode),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "twitter.0.twitter_id", twitterId),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "twitter.0.name", twitterName),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "twitter.0.screen_name", twitterScreenName),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "tickers.0.symbol", symbol),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "tickers.0.exchange", exchange),
				),
			},

			{
				//Update
				Config: GenerateBasicExternalOrganizationResource(
					resourceLabel,
					name,
					phoneDisplay, countryCode,
					address, city, state, postalCode, countryCode,
					twitterId, twitterUpdatedName, twitterScreenName,
					symbolUpdate, exchange,
					tags,
					externalsystemurl,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "name", name),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "phone_number.0.display", phoneDisplay),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "phone_number.0.country_code", countryCode),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.address1", address),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.city", city),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.state", state),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.postal_code", postalCode),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "address.0.country_code", countryCode),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "twitter.0.twitter_id", twitterId),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "twitter.0.name", twitterUpdatedName),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "twitter.0.screen_name", twitterScreenName),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "tickers.0.symbol", symbolUpdate),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "tickers.0.exchange", exchange),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "tickers.0.exchange", exchange),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "tickers.0.exchange", exchange),
					resource.TestCheckResourceAttr("genesyscloud_externalcontacts_organization."+resourceLabel, "external_system_url", externalsystemurl),
				),
			},
			{
				// Import/Read
				ResourceName:      "genesyscloud_externalcontacts_organization." + resourceLabel,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
		CheckDestroy: testVerifyContactDestroyed,
	})
}

func GenerateBasicExternalOrganizationResource(resourceID string, name string,
	phoneDisplay, phoneCountrycode string,
	address, city, state, postalCode, countryCode string,
	twitterId, twitterName, twitterScreenName string,
	symbol, exchange string,
	tags []string,
	externalUrl string,
) string {
	return fmt.Sprintf(`resource "genesyscloud_externalcontacts_organization" "%s" {
        name = "%s"
        phone_number {
          display = "%s"
          country_code = "%s"
        }
        address {
          address1 = "%s"
          city = "%s"
          state = "%s"
          postal_code = "%s"
          country_code = "%s"
        }
        twitter {
          twitter_id = "%s"
          name = "%s"
          screen_name = "%s"
        }
		tickers{
			symbol = "%s"
			exchange = "%s"
		}
		tags = [%s]
		external_system_url = "%s"
    }
    `, resourceID, name,
		phoneDisplay, phoneCountrycode,
		address, city, state, postalCode, countryCode,
		twitterId, twitterName, twitterScreenName,
		symbol, exchange,
		strings.Join(tags, ", "),
		externalUrl)
}

func testVerifyContactDestroyed(state *terraform.State) error {
	externalAPI := platformclientv2.NewExternalContactsApi()
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "genesyscloud_externalcontacts_organization" {
			continue
		}

		externalContact, resp, err := externalAPI.GetExternalcontactsContact(rs.Primary.ID, nil)
		if externalContact != nil {
			return fmt.Errorf("External contact (%s) still exists", rs.Primary.ID)
		} else if util.IsStatus404(resp) {
			// External Contact not found as expected
			continue
		} else {
			// Unexpected error
			return fmt.Errorf("Unexpected error: %s", err)
		}
	}
	// Success. All contacts destroyed
	return nil
}
