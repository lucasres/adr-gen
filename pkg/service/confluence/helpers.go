package confluence

import (
	"fmt"
	"net/http"
)

func assertResponseStatusCode(res *http.Response, expectedStatusCode int) error {
	if res.StatusCode != expectedStatusCode {
		return fmt.Errorf(
			"confluence API don't return expected HTTP Status Code %d. Returned: [%d] - \"%s\"",
			expectedStatusCode,
			res.StatusCode,
			res.Body,
		)
	}

	return nil
}
