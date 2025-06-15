output "resource_group_name" {
  description = "The name of the Azure resource group"
  value       = azurerm_resource_group.rg.name
}

output "vm_name" {
  description = "The name of the virtual machine"
  value       = azurerm_linux_virtual_machine.webserver.name
}

output "nic_name" {
  description = "The name of the network interface card"
  value       = azurerm_network_interface.webserver.name
}

output "public_ip" {
  description = "The public IP address of the virtual machine"
  value       = azurerm_public_ip.webserver.ip_address
}
