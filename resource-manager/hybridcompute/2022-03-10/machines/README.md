
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/machines` Documentation

The `machines` SDK allows for interaction with the Azure Resource Manager Service `hybridcompute` (API Version `2022-03-10`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/machines"
```


### Client Initialization

```go
client := machines.NewMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MachinesClient.MachinesCreateOrUpdate`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue")

payload := machines.Machine{
	// ...
}


read, err := client.MachinesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinesClient.MachinesDelete`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue")

read, err := client.MachinesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinesClient.MachinesGet`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue")

read, err := client.MachinesGet(ctx, id, machines.DefaultMachinesGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinesClient.MachinesListByResourceGroup`

```go
ctx := context.TODO()
id := machines.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.MachinesListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.MachinesListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachinesClient.MachinesListBySubscription`

```go
ctx := context.TODO()
id := machines.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.MachinesListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.MachinesListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachinesClient.MachinesUpdate`

```go
ctx := context.TODO()
id := machines.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineValue")

payload := machines.MachineUpdate{
	// ...
}


read, err := client.MachinesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
