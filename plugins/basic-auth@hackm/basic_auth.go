package plugin

import (
	"encoding/base64"
	"github.com/labstack/echo"
)

type (
	// BasicAuthConfig defines the config for HTTP basic auth middleware.
	BasicAuthConfig struct {
		// Validator is the function to validate basic auth credentials.
		Validator BasicAuthValidator
		// Interceptor is the function to success basic auth call until next.
		Interceptor BasicAuthInterceptor
	}

	// BasicAuthValidator defines a function to validate basic auth credentials.
	BasicAuthValidator func(string, string) bool
	// BasicAuthInterceptor defines a function to intercept basic auth successed.
	BasicAuthInterceptor func(echo.Context, string)
)

const (
	basic = "Basic"
)

// BasicAuth returns an HTTP basic auth middleware.
//
// For valid credentials it calls the next handler.
// For invalid credentials, it sends "401 - Unauthorized" response.
// For empty or invalid `Authorization` header, it sends "400 - Bad Request" response.
func BasicAuth(validator BasicAuthValidator, interceptor BasicAuthInterceptor) echo.MiddlewareFunc {
	return BasicAuthWithConfig(BasicAuthConfig{validator, interceptor})
}

// BasicAuthWithConfig returns an HTTP basic auth middleware from config.
// See `BasicAuth()`.
func BasicAuthWithConfig(config BasicAuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header().Get(echo.HeaderAuthorization)
			l := len(basic)

			if len(auth) > l+1 && auth[:l] == basic {
				b, err := base64.StdEncoding.DecodeString(auth[l+1:])
				if err != nil {
					return err
				}
				cred := string(b)
				for i := 0; i < len(cred); i++ {
					if cred[i] == ':' {
						// Verify credentials
						if config.Validator(cred[:i], cred[i+1:]) {
							// Call interceptor
							config.Interceptor(c, cred[:i])
							return next(c)
						}
					}
				}
			}
			// Need to return `401` for browsers to pop-up login box.
			c.Response().Header().Set(echo.HeaderWWWAuthenticate, basic+" realm=Restricted")
			return echo.ErrUnauthorized
		}
	}
}
