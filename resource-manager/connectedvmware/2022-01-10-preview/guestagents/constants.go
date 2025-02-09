package guestagents

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningAction string

const (
	ProvisioningActionInstall   ProvisioningAction = "install"
	ProvisioningActionRepair    ProvisioningAction = "repair"
	ProvisioningActionUninstall ProvisioningAction = "uninstall"
)

func PossibleValuesForProvisioningAction() []string {
	return []string{
		string(ProvisioningActionInstall),
		string(ProvisioningActionRepair),
		string(ProvisioningActionUninstall),
	}
}

func parseProvisioningAction(input string) (*ProvisioningAction, error) {
	vals := map[string]ProvisioningAction{
		"install":   ProvisioningActionInstall,
		"repair":    ProvisioningActionRepair,
		"uninstall": ProvisioningActionUninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningAction(input)
	return &out, nil
}
