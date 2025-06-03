package main

import ( "log"
  "fmt"

  "cpru-api/internal/infrastructure/config"
	"cpru-api/internal/interface/graphql"
	"cpru-api/graph/generated"
	"cpru-api/internal/application/user"
	"cpru-api/internal/application/role"
	"cpru-api/internal/application/permission"
  "cpru-api/internal/application/rolepermission"
  "cpru-api/internal/application/userpermission"
  "cpru-api/internal/application/userrole"
  "cpru-api/internal/application/auth" 
	"cpru-api/internal/infrastructure/db"
  "cpru-api/internal/interface/http/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
  // "github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
  cfg := config.LoadConfig()
  dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
  )

	client, err := db.NewEntClient(dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Repository
	userRepo := db.NewUserRepository(client)
	roleRepo := db.NewRoleRepository(client)
	permissionRepo := db.NewPermissionRepository(client)
  rolepermissionRepo := db.NewRolePermissionRepository(client)
  userpermissionRepo := db.NewUserPermissionRepository(client)
  userroleRepo := db.NewUserRoleRepository(client)

	// Service
  authService := auth.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiresIn)
	userService := user.NewUserService(userRepo)
	roleService := role.NewRoleService(roleRepo)
	permissionService := permission.NewPermissionService(permissionRepo)
  rolepermissionService := rolepermission.NewRolePermissionService(rolepermissionRepo)
  userpermissionService := userpermission.NewUserPermissionService(userpermissionRepo)
  userroleService := userrole.NewUserRoleService(userroleRepo)


	// Resolver Root
	resolver := &graphql.Resolver{
    AuthService: authService,
		UserService:  userService,	
		RoleService:  roleService,
		PermissionService: permissionService,
    RolePermissionService: rolepermissionService,
    UserPermissionService: userpermissionService,
    UserRoleService: userroleService,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	e := echo.New()
  // e.Use(middleware.Logger())
  // e.Use(middleware.Recover())
	
  e.POST("/query", echo.WrapHandler(srv), middleware.JWTMiddleware(authService))
  e.GET("/", echo.WrapHandler(playground.Handler("GraphQL Playground", "/query"))) 
//   playgroundHandler := playground.HandlerWithHeaders(
//     "GraphQL Playground",           // title
//     "/query",                       // GraphQL endpoint path
//     map[string]string{              // HTTP headers ที่จะฝัง (header คงที่)
//         "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDg5MzU3MjcsImlhdCI6MTc0ODg0OTMyNywic3ViIjo1fQ.Yv_1RKw_CcxuFPC8FZpYuFkawIqvyIKbHr1mxF7hUGc",
//     },
//     nil,  // wsHeaders ถ้าไม่ใช้ WebSocket ให้ใส่ nil หรือ map[string]string{}
// )
//   e.GET("/", echo.WrapHandler(playgroundHandler))

  addr := fmt.Sprintf(":%s", cfg.Port)
  log.Printf("🚀 Starting server at %s ...", addr)
  if err := e.Start(addr); err != nil {
    log.Fatalf("failed to start server: %v", err)
  }	
}
