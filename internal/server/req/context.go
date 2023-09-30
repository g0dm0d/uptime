package req

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Ctx struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Ctx) ParseJSON(v interface{}) error {
	if c.Request.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("expected content-type %s, got %s", "application/json", c.Request.Header.Get("Content-Type"))
	}

	return json.NewDecoder(c.Request.Body).Decode(v)
}

// JSON Return client json
func (c *Ctx) JSON(v interface{}) error {
	c.Writer.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(c.Writer).Encode(v)
}

func (c *Ctx) BearerToken() string {
	authorization := c.Request.Header.Get("authorization")

	authorizationSplitted := strings.Split(authorization, " ")
	return authorizationSplitted[1]
}
