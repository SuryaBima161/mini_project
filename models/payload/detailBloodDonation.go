package payload

type (
	CreateDetailBloodDonationRequest struct {
		DonaturID        uint `json:"donatur_id" from:"donatur_id" validate:"required"`
		PenerimaID       uint `json:"penerima_id" form:"penerima_id" validate:"required"`
		MedicalHistoryID uint `json:"medical_history_id" form:"medical_history_id" validate:"required"`
	}

	CreateDetailDonationRespone struct {
		DonaturID        uint   `json:"donatur_id"`
		PenerimaID       uint   `json:"penerima_id"`
		MedicalHistoryID uint   `json:"medical_history_id"`
		TotalQty         string `json:"total_qty"`
	}
	GetDetailDonationResponse struct {
		DonaturID        uint   `json:"donatur_id"`
		PenerimaID       uint   `json:"penerima_id"`
		MedicalHistoryID uint   `json:"medical_history_id"`
		TotalQty         string `json:"total_qty"`
	}
)
