package replicationalertsettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ReplicationAlertSettingId{}

// ReplicationAlertSettingId is a struct representing the Resource ID for a Replication Alert Setting
type ReplicationAlertSettingId struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	AlertSettingName  string
}

// NewReplicationAlertSettingID returns a new ReplicationAlertSettingId struct
func NewReplicationAlertSettingID(subscriptionId string, resourceGroupName string, resourceName string, alertSettingName string) ReplicationAlertSettingId {
	return ReplicationAlertSettingId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ResourceName:      resourceName,
		AlertSettingName:  alertSettingName,
	}
}

// ParseReplicationAlertSettingID parses 'input' into a ReplicationAlertSettingId
func ParseReplicationAlertSettingID(input string) (*ReplicationAlertSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationAlertSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationAlertSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.AlertSettingName, ok = parsed.Parsed["alertSettingName"]; !ok {
		return nil, fmt.Errorf("the segment 'alertSettingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseReplicationAlertSettingIDInsensitively parses 'input' case-insensitively into a ReplicationAlertSettingId
// note: this method should only be used for API response data and not user input
func ParseReplicationAlertSettingIDInsensitively(input string) (*ReplicationAlertSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationAlertSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationAlertSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.AlertSettingName, ok = parsed.Parsed["alertSettingName"]; !ok {
		return nil, fmt.Errorf("the segment 'alertSettingName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateReplicationAlertSettingID checks that 'input' can be parsed as a Replication Alert Setting ID
func ValidateReplicationAlertSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationAlertSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Alert Setting ID
func (id ReplicationAlertSettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationAlertSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ResourceName, id.AlertSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Alert Setting ID
func (id ReplicationAlertSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticReplicationAlertSettings", "replicationAlertSettings", "replicationAlertSettings"),
		resourceids.UserSpecifiedSegment("alertSettingName", "alertSettingValue"),
	}
}

// String returns a human-readable description of this Replication Alert Setting ID
func (id ReplicationAlertSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Alert Setting Name: %q", id.AlertSettingName),
	}
	return fmt.Sprintf("Replication Alert Setting (%s)", strings.Join(components, "\n"))
}