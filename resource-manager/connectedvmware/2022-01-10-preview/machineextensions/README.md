
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/machineextensions` Documentation

The `machineextensions` SDK allows for interaction with the Azure Resource Manager Service `connectedvmware` (API Version `2022-01-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/machineextensions"
```


### Client Initialization

```go
client := machineextensions.NewMachineExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MachineExtensionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := machineextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "nameValue", "extensionValue")

payload := machineextensions.MachineExtension{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MachineExtensionsClient.Delete`

```go
ctx := context.TODO()
id := machineextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "nameValue", "extensionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MachineExtensionsClient.Get`

```go
ctx := context.TODO()
id := machineextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "nameValue", "extensionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachineExtensionsClient.List`

```go
ctx := context.TODO()
id := machineextensions.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue")

// alternatively `client.List(ctx, id, machineextensions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, machineextensions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachineExtensionsClient.Update`

```go
ctx := context.TODO()
id := machineextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "nameValue", "extensionValue")

payload := machineextensions.MachineExtensionUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
