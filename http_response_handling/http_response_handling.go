package http_response_handling

import (
	"fmt"
	"io"
	"net/http"
)

func HandleHttpRequestError(resp *http.Response, err error, httpStatusCode int) error {
	if err != nil {
		return err
	} else if resp.StatusCode != httpStatusCode {
		b := resp.Body
		defer b.Close()
		by, _ := io.ReadAll(b)
		return fmt.Errorf("%s: %s", resp.Status, string(by))
	}

	return nil
}
