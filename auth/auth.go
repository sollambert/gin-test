package auth

import (
	"github.com/authboss/authboss"
	"github.com/authboss/defaults"
	"github.com/authboss/logout"
	"github.com/gorilla/sessions"
)

func configureAuthboss() authboss.Authboss {
	// Create Authboss instance
	ab := authboss.New()

	// Set up session storage
	store := sessions.NewCookieStore([]byte("your-secret-key"))
	ab.Config.Storage.SessionState = defaults.NewCookieStorer(store)

	// Set up user authentication options
	ab.Config.AuthenticateAfterRegister = true
	ab.Config.Modules.LogoutMethod = logout.RedirectMethod
	ab.Config.Modules.RecoverMethod = defaults.RecoverLoginRedirect
	ab.Config.Modules.RegisterMethod = defaults.RegisterLoginRedirect
	ab.Config.Modules.LoginMethod = defaults.LoginLoginRedirect

	// Set up cookie settings
	ab.Config.Cookie.Domain = "your-domain.com"
	ab.Config.Cookie.Name = "authcookie"
	ab.Config.Cookie.Path = "/"
	ab.Config.Cookie.Secure = false
	ab.Config.Cookie.HttpOnly = true
	ab.Config.Cookie.Expire = 3600

	// Set up password options
	ab.Config.Modules.BCryptCost = 10

	return ab
}