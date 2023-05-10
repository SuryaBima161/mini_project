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
func GetRiwayatId(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetRiwayat(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get Riwayat",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

// create user
func CreateRiwayatController(c echo.Context) error {
	payload := payload.CreateRiwayatRequest{}
	c.Bind(&payload)
	//validate request body
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed create Riwayat",
			"errorDescription": err,
		})
	}
	respone, err := usecase.CreateRiwayat(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed create Riwayat",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create Riwayat",
		"data":    respone,
	})
}

func UpdateRiwayatController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	riwayat := models.Riwayat{}
	c.Bind(&riwayat)
	riwayat.ID = uint(id)
	if err := usecase.UpdateRiwayat(&riwayat); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update Riwayat",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf Riwayat tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

// delete user
func DeleteRiwayatController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete Riwayat",
			"errorDescription": err,
		})
	}

	err = usecase.DeleteRiwayat(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete Riwayat",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Riwayat",
	})

}
