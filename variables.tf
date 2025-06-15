# Define config variables
variable "label_prefix" {
  description = "Prefix used to label all resources"
  type        = string
  default     = "lab6"
}

variable "region" {
  description = "Azure region to deploy the resources"
  type        = string
  default     = "canadacentral"
}

variable "admin_username" {
  type        = string
  default     = "azureadmin"
  description = "The username for the local user account on the VM."
}
