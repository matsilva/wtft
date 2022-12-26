package file

// Signature is a file signature
type Signature struct {
	Name                string   `json:"name"`
	HeaderSeq           []byte   `json:"header_seq"`
	Ext                 string   `json:"ext"`
	MimeType            string   `json:"mime_type"`
	SupportedExtensions []string `json:"supported_extensions"`
}

// Signatures is a list of file signatures
var Signatures = []Signature{
	//image file types
	{
		Name:                "PNG",
		HeaderSeq:           []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
		Ext:                 "png",
		MimeType:            "image/png",
		SupportedExtensions: []string{"png"},
	},
	{
		Name:                "JPEG",
		HeaderSeq:           []byte{0xFF, 0xD8, 0xFF},
		Ext:                 "jpg",
		MimeType:            "image/jpeg",
		SupportedExtensions: []string{"jpg", "jpeg"},
	},
	{
		Name:                "GIF",
		HeaderSeq:           []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61},
		Ext:                 "gif",
		MimeType:            "image/gif",
		SupportedExtensions: []string{"gif"},
	},
	{
		Name:                "BMP",
		HeaderSeq:           []byte{0x42, 0x4D},
		Ext:                 "bmp",
		MimeType:            "image/bmp",
		SupportedExtensions: []string{"bmp"},
	},
	{
		Name:                "WEBP",
		HeaderSeq:           []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x45, 0x42, 0x50},
		Ext:                 "webp",
		MimeType:            "image/webp",
		SupportedExtensions: []string{"webp"},
	},
	{
		Name:                "TIFF",
		HeaderSeq:           []byte{0x4D, 0x4D, 0x00, 0x2A},
		Ext:                 "tiff",
		MimeType:            "image/tiff",
		SupportedExtensions: []string{"tiff", "tif"},
	},
	//video file types
	{
		Name:                "AVI",
		HeaderSeq:           []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x41, 0x56, 0x49, 0x20},
		Ext:                 "avi",
		MimeType:            "video/x-msvideo",
		SupportedExtensions: []string{"avi"},
	},
	{
		Name:                "FLV",
		HeaderSeq:           []byte{0x46, 0x4C, 0x56, 0x01},
		Ext:                 "flv",
		MimeType:            "video/x-flv",
		SupportedExtensions: []string{"flv"},
	},
	{
		Name:                "MKV",
		HeaderSeq:           []byte{0x1A, 0x45, 0xDF, 0xA3},
		Ext:                 "mkv",
		MimeType:            "video/x-matroska",
		SupportedExtensions: []string{"mkv"},
	},
	{
		Name:                "MPG",
		HeaderSeq:           []byte{0x00, 0x00, 0x01, 0xBA},
		Ext:                 "mpg",
		MimeType:            "video/mpeg",
		SupportedExtensions: []string{"mpg"},
	},
	{
		Name:                "MP4",
		HeaderSeq:           []byte{0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70, 0x6D, 0x70, 0x34, 0x32},
		Ext:                 "mp4",
		MimeType:            "video/mp4",
		SupportedExtensions: []string{"mp4"},
	},
	{
		Name:                "WEBM",
		HeaderSeq:           []byte{0x1A, 0x45, 0xDF, 0xA3},
		Ext:                 "webm",
		MimeType:            "video/webm",
		SupportedExtensions: []string{"webm"},
	},
	{
		Name:                "OGG",
		HeaderSeq:           []byte{0x4F, 0x67, 0x67, 0x53},
		Ext:                 "ogg",
		MimeType:            "video/ogg",
		SupportedExtensions: []string{"ogg"},
	},
	{
		Name:                "OGV",
		HeaderSeq:           []byte{0x4F, 0x67, 0x67, 0x53},
		Ext:                 "ogv",
		MimeType:            "video/ogg",
		SupportedExtensions: []string{"ogv"},
	},
	{
		Name:                "OGA",
		HeaderSeq:           []byte{0x4F, 0x67, 0x67, 0x53},
		Ext:                 "oga",
		MimeType:            "video/ogg",
		SupportedExtensions: []string{"oga"},
	},
	{
		Name:                "OGX",
		HeaderSeq:           []byte{0x4F, 0x67, 0x67, 0x53},
		Ext:                 "ogx",
		MimeType:            "video/ogg",
		SupportedExtensions: []string{"ogx"},
	},
	{
		Name:                "M4V",
		HeaderSeq:           []byte{0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70, 0x6D, 0x70, 0x34, 0x32},
		Ext:                 "m4v",
		MimeType:            "video/mp4",
		SupportedExtensions: []string{"m4v"},
	},
	{
		Name:                "M4A",
		HeaderSeq:           []byte{0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70, 0x6D, 0x70, 0x34, 0x32},
		Ext:                 "m4a",
		MimeType:            "video/mp4",
		SupportedExtensions: []string{"m4a"},
	},
	{
		Name:                "MOV",
		HeaderSeq:           []byte{0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70, 0x6D, 0x70, 0x34, 0x32},
		Ext:                 "mov",
		MimeType:            "video/mp4",
		SupportedExtensions: []string{"mov"},
	},
	//audio file types
	{
		Name:                "MP3",
		HeaderSeq:           []byte{0x49, 0x44, 0x33},
		Ext:                 "mp3",
		MimeType:            "audio/mpeg",
		SupportedExtensions: []string{"mp3"},
	},
	{
		Name:                "WAV",
		HeaderSeq:           []byte{0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45},
		Ext:                 "wav",
		MimeType:            "audio/wav",
		SupportedExtensions: []string{"wav"},
	},
	{
		Name:                "FLAC",
		HeaderSeq:           []byte{0x66, 0x4C, 0x61, 0x43},
		Ext:                 "flac",
		MimeType:            "audio/flac",
		SupportedExtensions: []string{"flac"},
	},
	{
		Name:                "M4A",
		HeaderSeq:           []byte{0x4D, 0x34, 0x41, 0x20},
		Ext:                 "m4a",
		MimeType:            "audio/m4a",
		SupportedExtensions: []string{"m4a"},
	},
	{
		Name:                "M4B",
		HeaderSeq:           []byte{0x4D, 0x34, 0x41, 0x20},
		Ext:                 "m4b",
		MimeType:            "audio/m4b",
		SupportedExtensions: []string{"m4b"},
	},
	{
		Name:                "M4P",
		HeaderSeq:           []byte{0x4D, 0x34, 0x41, 0x20},
		Ext:                 "m4p",
		MimeType:            "audio/m4p",
		SupportedExtensions: []string{"m4p"},
	},
	{
		Name:                "M4R",
		HeaderSeq:           []byte{0x4D, 0x34, 0x41, 0x20},
		Ext:                 "m4r",
		MimeType:            "audio/m4r",
		SupportedExtensions: []string{"m4r"},
	},
	{
		Name:                "M4V",
		HeaderSeq:           []byte{0x4D, 0x34, 0x56, 0x20},
		Ext:                 "m4v",
		MimeType:            "audio/m4v",
		SupportedExtensions: []string{"m4v"},
	},
	{
		Name:                "3GP",
		HeaderSeq:           []byte{0x66, 0x74, 0x79, 0x70, 0x33, 0x67, 0x70},
		Ext:                 "3gp",
		MimeType:            "audio/3gp",
		SupportedExtensions: []string{"3gp"},
	},
	//compressed file types
	{
		Name:                "ZIP",
		HeaderSeq:           []byte{0x50, 0x4B, 0x03, 0x04},
		Ext:                 "zip",
		MimeType:            "application/zip",
		SupportedExtensions: []string{"zip"},
	},
	{
		Name:                "RAR",
		HeaderSeq:           []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07, 0x00},
		Ext:                 "rar",
		MimeType:            "application/x-rar-compressed",
		SupportedExtensions: []string{"rar"},
	},
	{
		Name:                "7Z",
		HeaderSeq:           []byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C},
		Ext:                 "7z",
		MimeType:            "application/x-7z-compressed",
		SupportedExtensions: []string{"7z"},
	},
	{
		Name:                "GZ",
		HeaderSeq:           []byte{0x1F, 0x8B, 0x08},
		Ext:                 "gz",
		MimeType:            "application/gzip",
		SupportedExtensions: []string{"gz"},
	},
	{
		Name:                "TAR",
		HeaderSeq:           []byte{0x75, 0x73, 0x74, 0x61, 0x72},
		Ext:                 "tar",
		MimeType:            "application/x-tar",
		SupportedExtensions: []string{"tar"},
	},
	{
		Name:                "BZ2",
		HeaderSeq:           []byte{0x42, 0x5A, 0x68},
		Ext:                 "bz2",
		MimeType:            "application/x-bzip2",
		SupportedExtensions: []string{"bz2"},
	},
	{
		Name:                "XZ",
		HeaderSeq:           []byte{0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00},
		Ext:                 "xz",
		MimeType:            "application/x-xz",
		SupportedExtensions: []string{"xz"},
	},
}
