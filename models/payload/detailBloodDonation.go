package payload

type (
	CreateDetailBloodDonationRequest struct {
		Pendonor_id    uint   `json:"pendonor_id" from:"pendonor_id" validate:"required"`
		Penerima_id    uint   `json:"penerima_id" form:"penerima_id" validate:"required"`
		Tanggal_Donasi string `json:"tanggal_donasi" form:"tanggal_donasi" validate:"required"`
		Qty            string `json:"qqty" form:"qqty" validate:"required"`
	}

	CreateDetailDonationRespone struct {
		Pendonor_id    uint   `json:"pendonor_id"`
		Penerima_id    uint   `json:"penerima_id"`
		Tanggal_Donasi string `json:"tanggal_donasi" `
		Qty            string `json:"qty" form:"qty"`
	}
	GetDetailDonationResponse struct {
		Pendonor_id    uint   `json:"pendonor_id"`
		Penerima_id    uint   `json:"penerima_id"`
		Tanggal_Donasi string `json:"tanggal_donasi" `
		Qty            string `json:"qty" form:"qty"`
	}
)
