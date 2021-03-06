// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

// SortByEnum Enum with underlying type: string
type SortByEnum string

// Set of constants representing the allowable values for SortByEnum
const (
	SortByTimeCreated SortByEnum = "timeCreated"
	SortByDisplayName SortByEnum = "displayName"
)

var mappingSortBy = map[string]SortByEnum{
	"timeCreated": SortByTimeCreated,
	"displayName": SortByDisplayName,
}

// GetSortByEnumValues Enumerates the set of values for SortByEnum
func GetSortByEnumValues() []SortByEnum {
	values := make([]SortByEnum, 0)
	for _, v := range mappingSortBy {
		values = append(values, v)
	}
	return values
}
