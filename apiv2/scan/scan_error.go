package scan

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/zelat/goharbor-client/apiv2/errors"
)

func handleSwaggerScanErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			return nil
		case http.StatusBadRequest:
			return &errors.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrForbidden{}
		case http.StatusConflict:
			return &errors.ErrConflict{}
		case http.StatusInternalServerError:
			return &errors.ErrInternalErrors{}
		}
	}
	return in
}
