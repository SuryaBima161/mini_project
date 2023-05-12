package controllers

import (
	"mini_project/models"
	"mini_project/models/payload"
	"mini_project/usecase"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMedicalHistory(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetMedicalHistory(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get  Medical History",
			"errorDescription": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get Medical History",
		"users":  user,
	})
}

func CreateMedicalHistoryController(c echo.Context) error {
	payload := payload.CreateMedicalHistoryRequest{}
	c.Bind(&payload)
	//validate request body
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed create  Medical History",
			"errorDescription": err,
		})
	}
	respone, err := usecase.CreateMedicalHistory(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed create  Medical History",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create  Medical History",
		"data":    respone,
	})
}

func UpdateMedicalHistoryController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	medicalHistory := models.MedicalHistory{}
	c.Bind(&medicalHistory)
	medicalHistory.ID = uint(id)
	if err := usecase.UpdateMedicalHistory(&medicalHistory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update  Medical History",
			"errorDescription": err,
			"errorMessage":     "Sorry cant update Medical History now",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Medical History",
	})
}

// delete user
func DeleteMedicalHistoryController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete  Medical History",
			"errorDescription": err,
		})
	}

	err = usecase.DeleteMedicalHistory(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete  Medical History",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Medical History",
	})

}
