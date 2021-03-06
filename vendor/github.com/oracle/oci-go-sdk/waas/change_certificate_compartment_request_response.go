// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ChangeCertificateCompartmentRequest wrapper for the ChangeCertificateCompartment operation
type ChangeCertificateCompartmentRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SSL certificate used in the WAAS policy. This number is generated when the certificate is added to the policy.
	CertificateId *string `mandatory:"true" contributesTo:"path" name:"certificateId"`

	ChangeCertificateCompartmentDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the `PUT` or `DELETE` call for a resource, set the `if-match` parameter to the value of the etag from a previous `GET` or `POST` response for that resource. The resource will be updated or deleted only if the etag provided matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated before then due to conflicting operations
	// *Example:* If a resource has been deleted and purged from the system, then a retry of the original delete request may be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangeCertificateCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangeCertificateCompartmentRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangeCertificateCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ChangeCertificateCompartmentResponse wrapper for the ChangeCertificateCompartment operation
type ChangeCertificateCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ChangeCertificateCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangeCertificateCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
