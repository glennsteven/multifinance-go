package presentations

type ConsumerRequest struct {
	NIK           string  `json:"nik" form:"nik"`
	FullName      string  `json:"full_name" form:"full_name"`
	LegalName     string  `json:"legal_name" form:"legal_name"`
	Pob           string  `json:"pob" form:"pob"`
	Dob           string  `json:"dob" form:"dob"`
	Salary        float64 `json:"salary" form:"salary"`
	ImageIdentity *File   `json:"image_identity" form:"image_identity"`
	ImageSelfie   *File   `json:"image_selfie" form:"image_selfie"`
}
