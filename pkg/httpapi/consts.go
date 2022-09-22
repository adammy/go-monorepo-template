package httpapi

const (
	// ContentTypeHeader is the "Content-Type" header key.
	ContentTypeHeader = "Content-Type"

	// ApplicationJSON is the "application/json" header value for ContentTypeHeader.
	ApplicationJSON = "application/json"

	// MultipartFormData is the "multipart/form-data" header value for ContentTypeHeader.
	// If checking for this value, use strings.HasPrefix since it typically has additional content after.
	MultipartFormData = "multipart/form-data"

	// ImagePNG is the "image/png" header value for ContentTypeHeader.
	ImagePNG = "image/png"

	// ImageJPG is the "image/jpg" header value for ContentTypeHeader.
	ImageJPG = "image/jpg"

	// ImageJPEG is the "image/jpeg" header value for ContentTypeHeader.
	ImageJPEG = "image/jpeg"

	// TraceLogLevel denotes to show trace and all log levels below.
	TraceLogLevel LogLevel = "trace"

	// DebugLogLevel denotes to show debug and all log levels below.
	DebugLogLevel LogLevel = "debug"

	// InfoLogLevel denotes to show info and all log levels below.
	InfoLogLevel LogLevel = "info"

	// WarnLogLevel denotes to show warn and all log levels below.
	WarnLogLevel LogLevel = "warn"

	// ErrorLogLevel denotes to show error and all log levels below.
	ErrorLogLevel LogLevel = "error"

	// CriticalLogLevel denotes to show critical logs only.
	CriticalLogLevel LogLevel = "critical"

	// EnvFlag defines the incoming flag to determine the Server environment.
	EnvFlag = "env"

	// EnvFlagDefault defines the default value for EnvFlag.
	EnvFlagDefault = "local"

	// EnvFlagDescription provides a description for EnvFlag.
	EnvFlagDescription = "the environment for the application"

	// ConfigPath defines the path for getting Server config.
	ConfigPath = "./configs"
)
