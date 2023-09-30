package errs

import "net/http"

type Code struct {
	Code        uint16 `json:"code"`
	Description string `json:"description"`
	Status      int
}

var (
	// System
	InternalServerError = Code{Code: 100, Description: "Internal server error", Status: http.StatusInternalServerError}
	InvalidJSON         = Code{Code: 101, Description: "Invalid JSON body", Status: http.StatusBadRequest}
	InvalidHashsum      = Code{Code: 102, Description: "Invalid Hash", Status: http.StatusBadRequest}
	InvalidUrlParam     = Code{Code: 103, Description: "Invalid Url param", Status: http.StatusBadRequest}

	// Auth codes
	AccessTokenHasExpired       = Code{Code: 1001, Description: "Invalid credentials, try again.", Status: http.StatusUnauthorized}
	AccessTokenInvalidFormat    = Code{Code: 1002, Description: "Access token has invalid format.", Status: http.StatusUnauthorized}
	AccessTokenInvalidSignature = Code{Code: 1003, Description: "Access token has invalid signature.", Status: http.StatusUnauthorized}
	IncorrectLoginOrPassword    = Code{Code: 1004, Description: "Invalid login credentials", Status: http.StatusUnauthorized}

	// User codes
	UserNotFound      = Code{Code: 1501, Description: "User not found, make sure user id is correct."}
	UserAlreadyExists = Code{Code: 1502, Description: "User already exists.", Status: http.StatusUnauthorized}

	// Monitor codes
	InvalidProtocol = Code{Code: 2001, Description: "The wrong protocol", Status: http.StatusBadRequest}
)
