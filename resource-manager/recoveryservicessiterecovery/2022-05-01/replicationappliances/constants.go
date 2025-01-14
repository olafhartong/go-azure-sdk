package replicationappliances

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthErrorCustomerResolvability string

const (
	HealthErrorCustomerResolvabilityAllowed    HealthErrorCustomerResolvability = "Allowed"
	HealthErrorCustomerResolvabilityNotAllowed HealthErrorCustomerResolvability = "NotAllowed"
)

func PossibleValuesForHealthErrorCustomerResolvability() []string {
	return []string{
		string(HealthErrorCustomerResolvabilityAllowed),
		string(HealthErrorCustomerResolvabilityNotAllowed),
	}
}

func parseHealthErrorCustomerResolvability(input string) (*HealthErrorCustomerResolvability, error) {
	vals := map[string]HealthErrorCustomerResolvability{
		"allowed":    HealthErrorCustomerResolvabilityAllowed,
		"notallowed": HealthErrorCustomerResolvabilityNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthErrorCustomerResolvability(input)
	return &out, nil
}

type ProtectionHealth string

const (
	ProtectionHealthCritical ProtectionHealth = "Critical"
	ProtectionHealthNone     ProtectionHealth = "None"
	ProtectionHealthNormal   ProtectionHealth = "Normal"
	ProtectionHealthWarning  ProtectionHealth = "Warning"
)

func PossibleValuesForProtectionHealth() []string {
	return []string{
		string(ProtectionHealthCritical),
		string(ProtectionHealthNone),
		string(ProtectionHealthNormal),
		string(ProtectionHealthWarning),
	}
}

func parseProtectionHealth(input string) (*ProtectionHealth, error) {
	vals := map[string]ProtectionHealth{
		"critical": ProtectionHealthCritical,
		"none":     ProtectionHealthNone,
		"normal":   ProtectionHealthNormal,
		"warning":  ProtectionHealthWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionHealth(input)
	return &out, nil
}

type RcmComponentStatus string

const (
	RcmComponentStatusCritical RcmComponentStatus = "Critical"
	RcmComponentStatusHealthy  RcmComponentStatus = "Healthy"
	RcmComponentStatusUnknown  RcmComponentStatus = "Unknown"
	RcmComponentStatusWarning  RcmComponentStatus = "Warning"
)

func PossibleValuesForRcmComponentStatus() []string {
	return []string{
		string(RcmComponentStatusCritical),
		string(RcmComponentStatusHealthy),
		string(RcmComponentStatusUnknown),
		string(RcmComponentStatusWarning),
	}
}

func parseRcmComponentStatus(input string) (*RcmComponentStatus, error) {
	vals := map[string]RcmComponentStatus{
		"critical": RcmComponentStatusCritical,
		"healthy":  RcmComponentStatusHealthy,
		"unknown":  RcmComponentStatusUnknown,
		"warning":  RcmComponentStatusWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RcmComponentStatus(input)
	return &out, nil
}
