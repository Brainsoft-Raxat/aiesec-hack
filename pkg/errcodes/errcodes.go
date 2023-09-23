package errcodes

import (
	"net/http"

	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
)

// TODO: errcodes here
var (
	IncorrectCredsError = apperror.NewErrorCode(1, http.StatusBadRequest, "Incorrect email or password")
	JerryNotFound       = apperror.NewErrorCode(2, http.StatusNotFound, "Jerry not found")
	CreateEventError    = apperror.NewErrorCode(3, http.StatusInternalServerError, "Unable to create event")
	InternalServerError = apperror.NewErrorCode(4, http.StatusInternalServerError, "Internal Server Error")
)
