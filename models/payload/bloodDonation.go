package payload

type (
	CreateBloodDonationRequest struct {
		DonaturID     uint   `json:"donatur_id" from:"donatur_id" validate:"required"`
		PenerimaID    uint   `json:"penerima_id" form:"penerima_id" validate:"required"`
		TanggalDonasi string `json:"tanggal_donasi" form:"tanggal_donasi" validate:"required"`
		Qty           string `json:"qty" form:"qty" validate:"required"`
	}

	CreateBloodDonationRespone struct {
		DonaturID     uint   `json:"donatur_id"`
		PenerimaID    uint   `json:"penerima_id"`
		TanggalDonasi string `json:"tanggal_donasi" `
		Qty           string `json:"qty" form:"qty"`
	}
	GetBloodDonationResponse struct {
		DonaturID     uint   `json:"donatur_id"`
		PenerimaID    uint   `json:"penerima_id"`
		TanggalDonasi string `json:"tanggal_donasi" `
		Qty           string `json:"qty" form:"qty"`
	}
)
