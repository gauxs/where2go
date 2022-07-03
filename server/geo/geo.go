package geo

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MapHandler(c *gin.Context) {
	topLeft := c.Param("topLeft")
	// latTopRight := c.Param("latTopRight")
	// latBottomLeft := c.Param("latBottomLeft")
	// latBottomRight := c.Param("latBottomRight")

	coordinates := strings.Split(topLeft, ",")

	c.String(http.StatusOK, "Hello %s", coordinates[0])
}
