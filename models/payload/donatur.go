package payload

type (
	CreateDonaturRequest struct {
		UserID       uint   `json:"user_id" form:"user_id" validate:"required"`
		Name         string `json:"name" form:"name" validate:"required"`
		JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required"`
		Tanggallahir string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
	}

	CreateDonaturRespone struct {
		UserID       uint   `json:"user_id"`
		Name         string `json:"name"`
		JenisKelamin string `json:"jenis_kelamin"`
		Tanggallahir string `json:"tanggal_lahir"`
	}
	GetDonaturResponse struct {
		UserID       uint   `json:"user_id"`
		Name         string `json:"name"`
		JenisKelamin string `json:"jenis_kelamin"`
		Tanggallahir string `json:"tanggal_lahir"`
	}
)
