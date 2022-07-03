package geo

import (
	"net/http"

	"where2go/server"
	"where2go/server/models"
	"where2go/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/mmcloughlin/geohash"
	"go.uber.org/zap"
)

func MapHandler(c *gin.Context) {
	var topLeftGeoHash string
	var topRightGeoHash string
	var bottomLeftGeoHash string
	var bottomRightGeoHash string

	logger := server.Application.Logger

	topLeftLatLng := c.Query("topLeft")
	topRightLatLng := c.Query("topRight")
	bottomLeftLatLng := c.Query("bottomLeft")
	bottomRightLatLng := c.Query("bottomRight")

	lat, lng, err := utils.ParseLatLngQueryParam(topLeftLatLng)
	if err != nil {
		response := models.Success{
			Success: false,
			Message: "incorrect topLeft query param",
		}
		c.JSON(http.StatusBadRequest, response)

		logger.Warn(
			"incorrect topLeft query param", zap.Error(err))
		return
	}
	topLeftGeoHash = geohash.Encode(lat, lng)

	lat, lng, err = utils.ParseLatLngQueryParam(topRightLatLng)
	if err != nil {
		response := models.Success{
			Success: false,
			Message: "incorrect topRight query param",
		}
		c.JSON(http.StatusBadRequest, response)

		logger.Warn(
			"incorrect topRight query param", zap.Error(err))
		return
	}
	topRightGeoHash = geohash.Encode(lat, lng)

	lat, lng, err = utils.ParseLatLngQueryParam(bottomLeftLatLng)
	if err != nil {
		response := models.Success{
			Success: false,
			Message: "incorrect bottomLeft query param",
		}
		c.JSON(http.StatusBadRequest, response)

		logger.Warn(
			"incorrect bottomLeft query param", zap.Error(err))
		return
	}
	bottomLeftGeoHash = geohash.Encode(lat, lng)

	lat, lng, err = utils.ParseLatLngQueryParam(bottomRightLatLng)
	if err != nil {
		response := models.Success{
			Success: false,
			Message: "incorrect bottomRight query param",
		}
		c.JSON(http.StatusBadRequest, response)

		logger.Warn(
			"incorrect bottomRight query param", zap.Error(err))
		return
	}
	bottomRightGeoHash = geohash.Encode(lat, lng)

	logger.Info(
		"GeoHash details",
		zap.String("topLeft", topLeftGeoHash),
		zap.String("topRight", topRightGeoHash),
		zap.String("bottomLeft", bottomLeftGeoHash),
		zap.String("bottomRight", bottomRightGeoHash))

	response := models.Success{
		Success: true,
		Message: "success",
	}
	c.JSON(http.StatusOK, response)
}
