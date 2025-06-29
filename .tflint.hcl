plugin "terraform" {
  enabled = true
  preset  = "all"  # changed from recommended to all
}

plugin "azurerm" {
  enabled = true
  version = "0.25.1"
  source  = "github.com/terraform-linters/tflint-ruleset-azurerm"
}
