package headers

const (
	_contentTypeHeader   = "Content-Type"
	_contentTypeJSON     = "application/json"
	_authorizationHeader = "Authorization"

	_accessControlAllowOriginHeader       = "Access-Control-Allow-Origin"
	_accessControlAllowMethodsHeader      = "Access-Control-Allow-Methods"
	_accessControlAllowHeaders            = "Access-Control-Allow-Headers"
	_accessControlAllowsCredentialsHeader = "Access-Control-Allow-Credentials"

	_allowedOrigins = "*" //TODO: как перейдем на сервер поменять разрешенные хосты
	_allowedMethods = "GET, POST, PUT, DELETE, OPTIONS"
	_allowedHeaders = "Authorization,Content-Type,Accept,Origin,User-Agent,DNT,Cache-Control,X-Mx-ReqToken,Keep-Alive,X-Requested-With,If-Modified-Since,x-referer"
)

func ContentTypeHeader() string {
	return _contentTypeHeader
}

func ContentTypeJSON() string {
	return _contentTypeJSON
}

func AuthorizationHeader() string {
	return _authorizationHeader
}

func AccessControlAllowOriginHeader() string {
	return _accessControlAllowOriginHeader
}

func AccessControlAllowMethodsHeader() string {
	return _accessControlAllowMethodsHeader
}

func AccessControlAllowHeaders() string {
	return _accessControlAllowHeaders
}

func AccessControlAllowCredentialsHeader() string {
	return _accessControlAllowsCredentialsHeader
}

func AllowedOrigins() string {
	return _allowedOrigins
}

func AllowedMethods() string {
	return _allowedMethods
}

func AllowedHeaders() string {
	return _allowedHeaders
}
