package payload

type (
	CreatPendonorRequest struct {
		User_id       uint   `json:"user_id" form:"user_id" validate:"required"`
		Name          string `json:"name" form:"name" validate:"required"`
		Jenis_Kelamin string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required"`
		Tanggal_lahir string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
	}

	CreatPendonorRespone struct {
		User_id       uint   `json:"user_id"`
		Name          string `json:"name"`
		Jenis_Kelamin string `json:"jenis_kelamin"`
		Tanggal_lahir string `json:"tanggal_lahir"`
	}
	GetPendonorResponse struct {
		User_id       uint   `json:"user_id"`
		Name          string `json:"name"`
		Jenis_Kelamin string `json:"jenis_kelamin"`
		Tanggal_lahir string `json:"tanggal_lahir"`
	}
)
