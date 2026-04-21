// audit:allow-oversize reason=bootstrap-wiring
package bootstrap

import (
	"context"

	"pos-go/internal/config"
	authhttp "pos-go/internal/modules/auth/transport/http"
	authusecase "pos-go/internal/modules/auth/usecase"
	systemhttp "pos-go/internal/modules/system/transport/http"
	googleoidc "pos-go/internal/platform/google"
	"pos-go/internal/platform/postgres"
	"pos-go/internal/platform/state/memory"
	jwtissuer "pos-go/internal/platform/token/jwt"
	httpmw "pos-go/internal/transport/http/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type App struct {
	Echo *echo.Echo
	DB   *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(httpmw.RequestID)
	e.Use(httpmw.Recover)

	api := e.Group("/api")

	healthHandler := systemhttp.NewHealthHandler(pool)
	healthHandler.Register(api)

	if cfg.Auth.Google.IsConfigured() {
		oidcProvider, err := googleoidc.NewOIDC(ctx, googleoidc.OIDCConfig{
			Issuer:       cfg.Auth.Google.Issuer,
			ClientID:     cfg.Auth.Google.ClientID,
			ClientSecret: cfg.Auth.Google.ClientSecret,
			RedirectURL:  cfg.Auth.Google.RedirectURL,
		})
		if err != nil {
			pool.Close()
			return nil, err
		}

		tokenIssuer, err := jwtissuer.NewHMACIssuer(
			cfg.Auth.JWT.Issuer,
			cfg.Auth.JWT.Aud,
			cfg.Auth.JWT.Kid,
			cfg.Auth.JWT.Secret,
			cfg.Auth.JWT.TTL,
		)
		if err != nil {
			pool.Close()
			return nil, err
		}

		tokenVerifier, err := jwtissuer.NewHMACVerifier(
			cfg.Auth.JWT.Issuer,
			cfg.Auth.JWT.Aud,
			cfg.Auth.JWT.Secret,
		)
		if err != nil {
			pool.Close()
			return nil, err
		}

		stateStore := memory.NewAuthStateStore()
		accountRepo := postgres.NewAccountIdentityRepository(pool)
		sessionStore := postgres.NewSessionStore(pool)
		sessionStatusChecker := postgres.NewSessionStatusChecker(pool)
		sessionRevoker := postgres.NewSessionRevoker(pool)
		refreshRepo := postgres.NewRefreshSessionRepository(pool)
		transactor := postgres.NewTransactor(pool)
		roleAssigner := postgres.NewAccountRoleAssigner(pool)
		roleRemover := postgres.NewAccountRoleRemover(pool)
		principalResolver := postgres.NewPrincipalResolver(pool)

		googleFlow := authusecase.NewGoogleFlow(
			oidcProvider,
			stateStore,
			accountRepo,
			sessionStore,
			tokenIssuer,
			transactor,
			cfg.Auth.StateTTL,
			cfg.Auth.SessionTTL,
		).WithRoleAssigner(roleAssigner)

		authGroup := api.Group("/auth")

		googleHandler := authhttp.NewGoogleHandler(
			authhttp.NewGoogleFlowAdapter(googleFlow),
			cfg.Auth.Google.RedirectURL,
		)
		googleHandler.Register(authGroup)

		refreshUsecase := authusecase.NewRefreshToken(
			refreshRepo,
			tokenIssuer,
			cfg.Auth.SessionTTL,
		)
		refreshHandler := authhttp.NewRefreshHandler(refreshUsecase)
		refreshHandler.Register(authGroup)

		meHandler := systemhttp.NewMeHandler()

		meGroup := api.Group("")
		meGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		meGroup.Use(httpmw.RequirePermission("profile.self.read"))
		meHandler.Register(meGroup)

		authzGroup := api.Group("/authz")
		authzGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		authzGroup.Use(httpmw.RequirePermission("profile.self.read"))
		meHandler.Register(authzGroup)

		logoutGroup := api.Group("/auth")
		logoutGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		logoutGroup.Use(httpmw.RequirePermission("auth.session.logout"))

		logoutUsecase := authusecase.NewLogoutCurrentSession(sessionRevoker)
		logoutHandler := authhttp.NewLogoutHandler(logoutUsecase)
		logoutHandler.Register(logoutGroup)

		adminGroup := api.Group("/admin")
		adminGroup.Use(httpmw.RequireAuth(tokenVerifier, principalResolver, sessionStatusChecker))
		adminGroup.Use(httpmw.RequirePermission("account.role.assign"))

		assignAccountRoleUsecase := authusecase.NewAssignAccountRole(roleAssigner)
		removeAccountRoleUsecase := authusecase.NewRemoveAccountRole(roleRemover)
		accountRoleHandler := authhttp.NewAccountRoleHandler(
			assignAccountRoleUsecase,
			removeAccountRoleUsecase,
		)
		accountRoleHandler.Register(adminGroup)
	}

	return &App{
		Echo: e,
		DB:   pool,
	}, nil
}
