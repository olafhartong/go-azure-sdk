package product

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductState string

const (
	ProductStateNotPublished ProductState = "notPublished"
	ProductStatePublished    ProductState = "published"
)

func PossibleValuesForProductState() []string {
	return []string{
		string(ProductStateNotPublished),
		string(ProductStatePublished),
	}
}

func parseProductState(input string) (*ProductState, error) {
	vals := map[string]ProductState{
		"notpublished": ProductStateNotPublished,
		"published":    ProductStatePublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductState(input)
	return &out, nil
}
