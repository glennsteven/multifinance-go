package consts

const (
	RegistrationImageSelfieMaxSize   memory = 5 << 20 // 5 MB
	RegistrationImageIdentityMaxSize memory = 5 << 20 // 5 MB
)

var (
	ImageSelfieContentType   = []string{"image/jpg", "image/jpeg", "image/png"}
	ImageIdentityContentType = []string{"image/jpg", "image/jpeg", "image/png"}
)

var (
	MimeTypesJPEG = "image/jpeg"
	MimeTypesPNG  = "image/png"
)

var (
	MimeTypesAble = []string{
		MimeTypesJPEG,
		MimeTypesPNG,
	}
)

type memory int
