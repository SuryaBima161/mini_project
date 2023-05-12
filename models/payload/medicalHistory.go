package payload

type (
	CreateMedicalHistoryRequest struct {
		DonaturID          uint   `json:"donatur_id" from:"donatur_id" validate:"required"`
		TanggalPemeriksaan string `json:"tanggal_pemeriksaan" form:"tanggal_pemeriksaan" validate:"required"`
		HasilPemeriksaan   string `json:"hasil_pemeriksaan" form:"hasil_pemeriksaan" validate:"required"`
		GolonganDarah      string `json:"golongan_darah" form:"golongan_darah" validate:"required"`
	}

	CreatMedicalHistoryRespone struct {
		DonaturID          uint   `json:"donatur_id" `
		TanggalPemeriksaan string `json:"tanggal_pemeriksaan" `
		HasilPemeriksaan   string `json:"hasil_pemeriksaan" `
		GolonganDarah      string `json:"golongan_darah" `
	}
	GetMedicalHistoryResponse struct {
		DonaturID          uint   `json:"donatur_id" `
		TanggalPemeriksaan string `json:"tanggal_pemeriksaan" `
		HasilPemeriksaan   string `json:"hasil_pemeriksaan" `
		GolonganDarah      string `json:"golongan_darah" `
	}
)
