package presentations

type ConsumerRequest struct {
	NIK           string  `json:"nik" form:"nik" validate:"required"`
	FullName      string  `json:"full_name" form:"full_name" validate:"required"`
	LegalName     string  `json:"legal_name" form:"legal_name" validate:"required"`
	Pob           string  `json:"pob" form:"pob" validate:"required"`
	Dob           string  `json:"dob" form:"dob" validate:"required"`
	Salary        float64 `json:"salary" form:"salary" validate:"required"`
	ImageIdentity *File   `json:"image_identity" form:"image_identity" validate:"required"`
	ImageSelfie   *File   `json:"image_selfie" form:"image_selfie" validate:"required"`
}
