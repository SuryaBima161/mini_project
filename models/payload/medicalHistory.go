package payload

type (
	CreateMedicalHistoryRequest struct {
		Pendonor_id         uint   `json:"pendonor_id" from:"pendonor_id" validate:"required"`
		Tanggal_Pemeriksaan string `json:"tanggal_pemeriksaan" form:"tanggal_pemeriksaan" validate:"required"`
		Hasil_Pemeriksaan   string `json:"hasil_pemeriksaan" form:"hasil_pemeriksaan" validate:"required"`
		Golongan_Darah      string `json:"golongan_darah" form:"golongan_darah" validate:"required"`
	}

	CreatMedicalHistoryRespone struct {
		Pendonor_id         uint   `json:"pendonor_id" `
		Tanggal_Pemeriksaan string `json:"tanggal_pemeriksaan" `
		Hasil_Pemeriksaan   string `json:"hasil_pemeriksaan" `
		Golongan_Darah      string `json:"golongan_darah" `
	}
	GetMedicalHistoryResponse struct {
		Pendonor_id         uint   `json:"pendonor_id" `
		Tanggal_Pemeriksaan string `json:"tanggal_pemeriksaan" `
		Hasil_Pemeriksaan   string `json:"hasil_pemeriksaan" `
		Golongan_Darah      string `json:"golongan_darah" `
	}
)
