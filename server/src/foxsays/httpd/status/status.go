package status

const (
	Continue           = 100
	SwitchingProtocols = 101
)

const (
	OK                   = 200
	Created              = 201
	Accepted             = 202
	NonAuthoritativeInfo = 203
	NoContent            = 204
	ResetContent         = 205
	PartialContent       = 206
)

const (
	MultipleChoices   = 300
	MovedPermanently  = 301
	Found             = 302
	SeeOther          = 303
	NotModified       = 304
	UseProxy          = 305
	TemporaryRedirect = 307
)

const (
	BadRequest                   = 400
	Unauthorized                 = 401
	PaymentRequired              = 402
	Forbidden                    = 403
	NotFound                     = 404
	MethodNotAllowed             = 405
	NotAcceptable                = 406
	ProxyAuthRequired            = 407
	RequestTimeout               = 408
	Conflict                     = 409
	Gone                         = 410
	LengthRequired               = 411
	PreconditionFailed           = 412
	RequestEntityTooLarge        = 413
	RequestURITooLong            = 414
	UnsupportedMediaType         = 415
	RequestedRangeNotSatisfiable = 416
	ExpectationFailed            = 417
	Teapot                       = 418
	UnprocessableEntity          = 422
)

const (
	InternalServerError     = 500
	NotImplemented          = 501
	BadGateway              = 502
	ServiceUnavailable      = 503
	GatewayTimeout          = 504
	HTTPVersionNotSupported = 505
)
