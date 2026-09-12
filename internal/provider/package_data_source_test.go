package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPackageDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `data "zarf_package" "test" {source = "ghcr.io/zarf-dev/packages/dos-games:1.3.0"}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.zarf_package.test", "metadata.name", "dos-games"),
					resource.TestCheckResourceAttr("data.zarf_package.test", "metadata.description", "Simple example to load classic DOS games into K8s in the airgap"),
					resource.TestCheckResourceAttr("data.zarf_package.test", "metadata.version", "1.3.0"),
					resource.TestCheckResourceAttr("data.zarf_package.test", "metadata.url", ""),
					resource.TestCheckResourceAttr("data.zarf_package.test", "source", "ghcr.io/zarf-dev/packages/dos-games:1.3.0"),
					resource.TestCheckResourceAttrSet("data.zarf_package.test", "metadata.architecture"),
					resource.TestCheckResourceAttrSet("data.zarf_package.test", "digest"),
				),
			},
		},
	})
}
