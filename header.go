package beatles

const (
	// Standar HTTP Headers
	HeaderContentType    = "Content-Type"
	HeaderAccept         = "Accept"
	HeaderAuthorization  = "Authorization"
	HeaderUserAgent      = "User-Agent"
	HeaderCacheControl   = "Cache-Control"
	HeaderAcceptEncoding = "Accept-Encoding"
	HeaderAcceptLanguage = "Accept-Language"

	// Custom Headers (X-Header)
	HeaderXRequestID      = "X-Request-ID"
	HeaderXCorrelationID  = "X-Correlation-ID"
	HeaderXRealIP         = "X-Real-IP"
	HeaderXForwardedFor   = "X-Forwarded-For"
	HeaderXForwardedProto = "X-Forwarded-Proto"
	HeaderXDeviceID       = "X-Device-ID"
	HeaderXPlatform       = "X-Platform"
	HeaderXVersion        = "X-Version"
	HeaderXSignature      = "X-Signature"

	// Security Headers
	HeaderXFrameOptions           = "X-Frame-Options"
	HeaderXXSSProtection          = "X-XSS-Protection"
	HeaderXContentTypeOptions     = "X-Content-Type-Options"
	HeaderStrictTransportSecurity = "Strict-Transport-Security"
	HeaderContentSecurityPolicy   = "Content-Security-Policy"
	HeaderReferrerPolicy          = "Referrer-Policy"

	// CORS Headers
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
)
