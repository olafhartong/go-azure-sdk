
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/guestagents` Documentation

The `guestagents` SDK allows for interaction with the Azure Resource Manager Service `connectedvmware` (API Version `2022-01-10-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/guestagents"
```


### Client Initialization

```go
client := guestagents.NewGuestAgentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GuestAgentsClient.Create`

```go
ctx := context.TODO()
id := guestagents.NewGuestAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "nameValue")

payload := guestagents.GuestAgent{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `GuestAgentsClient.Delete`

```go
ctx := context.TODO()
id := guestagents.NewGuestAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "nameValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `GuestAgentsClient.Get`

```go
ctx := context.TODO()
id := guestagents.NewGuestAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue", "nameValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GuestAgentsClient.ListByVm`

```go
ctx := context.TODO()
id := guestagents.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineValue")

// alternatively `client.ListByVm(ctx, id)` can be used to do batched pagination
items, err := client.ListByVmComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
