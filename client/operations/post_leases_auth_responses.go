// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLeasesAuthReader is a Reader for the PostLeasesAuth structure.
type PostLeasesAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLeasesAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLeasesAuthCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostLeasesAuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLeasesAuthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLeasesAuthNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLeasesAuthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLeasesAuthCreated creates a PostLeasesAuthCreated with default headers values
func NewPostLeasesAuthCreated() *PostLeasesAuthCreated {
	return &PostLeasesAuthCreated{}
}

/*PostLeasesAuthCreated handles this case with default header values.

PostLeasesAuthCreated post leases auth created
*/
type PostLeasesAuthCreated struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string

	Payload *PostLeasesAuthCreatedBody
}

func (o *PostLeasesAuthCreated) Error() string {
	return fmt.Sprintf("[POST /leases/auth][%d] postLeasesAuthCreated  %+v", 201, o.Payload)
}

func (o *PostLeasesAuthCreated) GetPayload() *PostLeasesAuthCreatedBody {
	return o.Payload
}

func (o *PostLeasesAuthCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	o.Payload = new(PostLeasesAuthCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLeasesAuthUnauthorized creates a PostLeasesAuthUnauthorized with default headers values
func NewPostLeasesAuthUnauthorized() *PostLeasesAuthUnauthorized {
	return &PostLeasesAuthUnauthorized{}
}

/*PostLeasesAuthUnauthorized handles this case with default header values.

Unauthorized
*/
type PostLeasesAuthUnauthorized struct {
}

func (o *PostLeasesAuthUnauthorized) Error() string {
	return fmt.Sprintf("[POST /leases/auth][%d] postLeasesAuthUnauthorized ", 401)
}

func (o *PostLeasesAuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesAuthForbidden creates a PostLeasesAuthForbidden with default headers values
func NewPostLeasesAuthForbidden() *PostLeasesAuthForbidden {
	return &PostLeasesAuthForbidden{}
}

/*PostLeasesAuthForbidden handles this case with default header values.

Failed to retrieve lease authentication
*/
type PostLeasesAuthForbidden struct {
}

func (o *PostLeasesAuthForbidden) Error() string {
	return fmt.Sprintf("[POST /leases/auth][%d] postLeasesAuthForbidden ", 403)
}

func (o *PostLeasesAuthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesAuthNotFound creates a PostLeasesAuthNotFound with default headers values
func NewPostLeasesAuthNotFound() *PostLeasesAuthNotFound {
	return &PostLeasesAuthNotFound{}
}

/*PostLeasesAuthNotFound handles this case with default header values.

No active lease exists for the requesting user
*/
type PostLeasesAuthNotFound struct {
}

func (o *PostLeasesAuthNotFound) Error() string {
	return fmt.Sprintf("[POST /leases/auth][%d] postLeasesAuthNotFound ", 404)
}

func (o *PostLeasesAuthNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLeasesAuthInternalServerError creates a PostLeasesAuthInternalServerError with default headers values
func NewPostLeasesAuthInternalServerError() *PostLeasesAuthInternalServerError {
	return &PostLeasesAuthInternalServerError{}
}

/*PostLeasesAuthInternalServerError handles this case with default header values.

Server failure
*/
type PostLeasesAuthInternalServerError struct {
}

func (o *PostLeasesAuthInternalServerError) Error() string {
	return fmt.Sprintf("[POST /leases/auth][%d] postLeasesAuthInternalServerError ", 500)
}

func (o *PostLeasesAuthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostLeasesAuthCreatedBody Lease Authentication
swagger:model PostLeasesAuthCreatedBody
*/
type PostLeasesAuthCreatedBody struct {

	// Access Key ID for access to the AWS API
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// URL to access the AWS Console
	ConsoleURL string `json:"consoleUrl,omitempty"`

	// expires on
	ExpiresOn float64 `json:"expiresOn,omitempty"`

	// Secret Access Key for access to the AWS API
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// Session Token for access to the AWS API
	SessionToken string `json:"sessionToken,omitempty"`
}

// Validate validates this post leases auth created body
func (o *PostLeasesAuthCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLeasesAuthCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLeasesAuthCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostLeasesAuthCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
