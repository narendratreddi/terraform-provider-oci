// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AvailableSoftwareSourceSummary A software source which can be added to a managed instance. Once a software source is added, packages from that software source can be installed on that managed instance.
type AvailableSoftwareSourceSummary struct {

	// unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User friendly name for the software source
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID for the parent software source, if there is one
	ParentId *string `mandatory:"false" json:"parentId"`

	// Display name of the parent software source, if there is one
	ParentName *string `mandatory:"false" json:"parentName"`
}

func (m AvailableSoftwareSourceSummary) String() string {
	return common.PointerString(m)
}
