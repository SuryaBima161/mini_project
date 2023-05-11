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
func GetDetailBloodDonationIdController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetDetailBloodDonation(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get detail blood Donation",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}
func GetDetailBloodDonationTotalQtyController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Invalid ID",
			"errorDescription": err.Error(),
		})
	}

	donationResp, totalQty, err := usecase.GetDetailBloodDonationWithTotalQty(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error getting detail blood donation",
			"errorDescription": err.Error(),
		})
	}

	response := map[string]interface{}{
		"status":    "success",
		"donation":  donationResp,
		"total_qty": totalQty,
	}

	return c.JSON(http.StatusOK, response)
}

// create user
func CreateDetailBloodDonationController(c echo.Context) error {
	payload := payload.CreateDetailBloodDonationRequest{}
	c.Bind(&payload)
	//validate request body
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed create detail blood Donation",
			"errorDescription": err,
		})
	}
	respone, err := usecase.CreateDetailBloodDonation(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":          "failed create detail blood Donation",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create detail blood Donation",
		"data":    respone,
	})
}

func UpdateDetailBloodDonationController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	detailDonation := models.DetailBloodDonation{}
	c.Bind(&detailDonation)
	detailDonation.ID = uint(id)
	if err := usecase.UpdateDetailBloodDonation(&detailDonation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update detail blood Donation",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf detail blood Donation tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update detail blood Donation",
	})
}

// delete user
func DeleteDetailBloodDonationController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete PemberianDarah",
			"errorDescription": err,
		})
	}

	err = usecase.DeleteDetailBloodDonation(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message":          "failed delete detail blood Donation",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete detail blood Donation",
	})

}
