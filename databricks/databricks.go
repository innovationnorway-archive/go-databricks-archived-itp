package databricks

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the Databricks API
	DefaultBaseURI = "/api/2.0"
	// HeaderXDatabricksOrgID is the Azure Databricks extension header of the
	// organization ID returned in the response.
	HeaderXDatabricksOrgID = "x-databricks-org-id"
)

// ExtractOrgID extracts the Databricks organization ID from the
// x-databricks-org-id header.
func ExtractOrgID(resp *http.Response) string {
	return autorest.ExtractHeaderValue(HeaderXDatabricksOrgID, resp)
}

// Error encapsulates the error response from a Databricks service.
type Error struct {
	autorest.DetailedError
	ErrorCode      string `json:"error_code"`
	Message        string `json:"message"`
	Response       *http.Response
	OrganizationID string
}

// Error returns a human-friendly error message from Databricks error.
func (e Error) Error() string {
	return fmt.Sprintf("ErrorCode=%q Message=%q", e.ErrorCode, e.Message)
}

// IsError returns true if the passed error is an Databricks error; false otherwise.
func IsError(e error) bool {
	_, ok := e.(*Error)
	return ok
}

// WithError returns a RespondDecorator that emits an
// Error by reading the response body.
func WithError() autorest.RespondDecorator {
	return func(r autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(resp *http.Response) error {
			err := r.Respond(resp)
			if err == nil {
				var e Error
				defer resp.Body.Close()

				encodedAs := autorest.EncodedAsJSON

				b, decodeErr := autorest.CopyAndDecode(encodedAs, resp.Body, &e)
				resp.Body = ioutil.NopCloser(&b)
				if decodeErr != nil {
					return fmt.Errorf("error response cannot be parsed: %q error: %v", b.String(), decodeErr)
				}

				if e.ErrorCode != "" && e.Message != "" {
					e.Response = resp
					e.OrganizationID = ExtractOrgID(resp)
					err = &e
				}
			}

			return err
		})
	}
}
