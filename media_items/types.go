package media_items

// MediaItem represents of a media item (such as a photo or video) in Google Photos.
// See: https://developers.google.com/photos/library/reference/rest/v1/mediaItems
type MediaItem struct {
	ID            string
	Description   string
	ProductURL    string
	BaseURL       string
	MimeType      string
	MediaMetadata MediaMetadata
	Filename      string
}

// MediaMetadata represents metadata for a media item.
// See: https://developers.google.com/photos/library/reference/rest/v1/mediaItems
type MediaMetadata struct {
	CreationTime string
	Width        string
	Height       string

	// Photo: Metadata for a photo media type.
	Photo *Photo `json:"photo,omitempty"`
	// Video: Metadata for a video media type.
	Video *Video `json:"video,omitempty"`
}

type Photo struct {
	// ApertureFNumber: Apeture f number of the photo.
	ApertureFNumber float64 `json:"apertureFNumber,omitempty"`

	// CameraMake: Brand of the camera which took the photo.
	CameraMake string `json:"cameraMake,omitempty"`

	// CameraModel: Model of the camera which took the photo.
	CameraModel string `json:"cameraModel,omitempty"`

	// ExposureTime: Exposure time of the photo.
	ExposureTime string `json:"exposureTime,omitempty"`

	// FocalLength: Focal length of the photo.
	FocalLength float64 `json:"focalLength,omitempty"`

	// IsoEquivalent: ISO of the photo.
	IsoEquivalent int64 `json:"isoEquivalent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApertureFNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ApertureFNumber") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

type Video struct {
	// CameraMake: Brand of the camera which took the video.
	CameraMake string `json:"cameraMake,omitempty"`

	// CameraModel: Model of the camera which took the video.
	CameraModel string `json:"cameraModel,omitempty"`

	// Fps: Frame rate of the video.
	Fps float64 `json:"fps,omitempty"`

	// Status: Processing status of the video.
	//
	// Possible values:
	//   "UNSPECIFIED" - Video processing status is unknown.
	//   "PROCESSING" - Video is currently being processed. The user will
	// see an icon for this
	// video in the Google Photos app, however, it will not be playable yet.
	//   "READY" - Video is now ready for viewing.
	//   "FAILED" - Something has gone wrong and the video has failed to
	// process.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CameraMake") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CameraMake") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// SimpleMediaItem represents a simple media item to be created in Google Photos via an upload token.
// See: https://developers.google.com/photos/library/reference/rest/v1/mediaItems/batchCreate
type SimpleMediaItem struct {
	UploadToken string
	FileName    string
}
