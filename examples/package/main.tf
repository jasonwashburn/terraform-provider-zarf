terraform {
  required_providers {
    zarf = {
      source = "jasonwashburn/zarf"
    }
  }
}

provider "zarf" {}

data "zarf_package" "example" {
  source = "ghcr.io/zarf-dev/packages/dos-games:1.3.0"
}

output "package" {
  value = data.zarf_package.example
}
