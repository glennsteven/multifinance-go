package resource

type ConsumerResource struct {
	Id             string  `json:"id"`
	IdentityNumber string  `json:"identity_number"`
	FullName       string  `json:"full_name"`
	LegalName      string  `json:"legal_name"`
	BirthPlace     string  `json:"birth_place"`
	BirthDate      string  `json:"birth_date"`
	Salary         float64 `json:"salary"`
	URL            URL     `json:"url"`
}

type URL struct {
	ImageIdentity string `json:"image_identity"`
	ImageSelfie   string `json:"image_selfie"`
}
