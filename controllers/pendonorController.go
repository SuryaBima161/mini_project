package controllers

import (
	"mini_project/models"
	"mini_project/models/payload"
	"mini_project/usecase"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

// get user by id
func GetPendonorId(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetPendonor(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get pendonor",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

// create user
func CreatePendonorController(c echo.Context) error {
	payload := payload.CreatPendonorRequest{}
	c.Bind(&payload)
	//validate request body
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed create pendonor",
			"errorDescription": err,
		})
	}
	respone, err := usecase.CreatePendonor(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed create pendonor",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create pendonor",
		"data":    respone,
	})
}

func UpdatePendonorController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	pendonor := models.Pendonor{}
	c.Bind(&pendonor)
	pendonor.ID = uint(id)
	if err := usecase.UpdatePendonor(&pendonor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update pendonor",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf pendonor tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

// delete user
func DeletePendonorController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete pendonor",
			"errorDescription": err,
		})
	}

	err = usecase.DeletePendonor(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete pendonor",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pendonor",
	})

}
