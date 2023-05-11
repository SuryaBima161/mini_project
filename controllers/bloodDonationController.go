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
func GetBloodDonationController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetBloodDonation(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get blood donation",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

// create user
func CreateBloodDonationController(c echo.Context) error {
	payload := payload.CreateBloodDonationRequest{}
	c.Bind(&payload)
	//validate request body
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed create PemberianDarah",
			"errorDescription": err,
		})
	}
	respone, err := usecase.CreateBloodDonation(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed create PemberianDarah",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create PemberianDarah",
		"data":    respone,
	})
}

func UpdateBloodDonationController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	bloodDonation := models.BloodDonation{}
	c.Bind(&bloodDonation)
	bloodDonation.ID = uint(id)
	if err := usecase.UpdateBloodDonation(&bloodDonation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update Blood Donation",
			"errorDescription": err,
			"errorMessage":     "Sorry Can't Update Blood Donation",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Blood Donation",
	})
}

// delete user
func DeleteBloodDonationController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete Blood Donation",
			"errorDescription": err,
		})
	}

	err = usecase.DeleteBloodDonation(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete PemberianDarah",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete PemberianDarah",
	})

}
