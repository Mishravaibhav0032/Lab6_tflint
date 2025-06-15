package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Your Azure subscription ID
var subscriptionID string = "eed2085b-133e-4466-8ed6-6da86c403fc0"

func TestAzureLinuxVMCreation(t *testing.T) {
	// Define the prefix to match your Terraform variable
	labelPrefix := "lab6" 

	// Terraform options for Terratest
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"label_prefix": labelPrefix, 
		},
		EnvVars: map[string]string{
			"ARM_SUBSCRIPTION_ID": subscriptionID,
		},
	}

	// Clean up resources with `terraform destroy` after the test
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)

	// Get Terraform output values
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	// Validate the VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))

	// Validate NIC is attached to VM
	actualNicNames := azure.GetVirtualMachineNics(t, vmName, resourceGroupName, subscriptionID)
	assert.Equal(t, nicName, actualNicNames[0])

	// Validate correct image is used
	vmImage := azure.GetVirtualMachineImage(t, vmName, resourceGroupName, subscriptionID)
	expectedOSPublisher := "Canonical"
	expectedOSVersion := "22_04-lts-gen2"
	assert.Equal(t, expectedOSPublisher, vmImage.Publisher)
	assert.Equal(t, expectedOSVersion, vmImage.SKU)
}
