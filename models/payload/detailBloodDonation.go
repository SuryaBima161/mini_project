package payload

type (
	CreateDetailBloodDonationRequest struct {
		Pendonor_id uint `json:"pendonor_id" from:"pendonor_id" validate:"required"`
		Penerima_id uint `json:"penerima_id" form:"penerima_id" validate:"required"`
		Riwayat_id  uint `json:"riwayat_id"`
	}

	CreateDetailDonationRespone struct {
		Pendonor_id uint   `json:"pendonor_id"`
		Penerima_id uint   `json:"penerima_id"`
		Riwayat_id  uint   `json:"riwayat_id"`
		Total_Qty   string `json:"total_qty"`
	}
	GetDetailDonationResponse struct {
		Pendonor_id uint   `json:"pendonor_id"`
		Penerima_id uint   `json:"penerima_id"`
		Riwayat_id  uint   `json:"riwayat_id"`
		Total_Qty   string `json:"total_qty"`
	}
)
