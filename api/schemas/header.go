package schemas

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

// LinkHeader will add a "Link" header to the schema provided in the http response.
// e.g. `Link: <http://www.example.com/api/v1/schema/schema.json#/properties/loginParams>; rel="describedby"`
func LinkHeader(c *gin.Context, reference string) {
	baseURL := strings.TrimSuffix(os.Getenv("BASE_URL"), "/")
	path := strings.TrimPrefix(reference, "/")
	link := fmt.Sprintf(`<%s/api/v1/schema/%s>; rel="describedby"`, baseURL, path)
	c.Header("Link", link)
}
