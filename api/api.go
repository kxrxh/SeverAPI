package api

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/goccy/go-json"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt"
	"sever.hack/api/utils"
)

type ApiCore struct {
	app        *fiber.App
	devkey     string
	privateKey *rsa.PrivateKey
}

var apiCore ApiCore

func Init(hasLogger bool) {
	apiCore = ApiCore{
		app: fiber.New(fiber.Config{
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
		devkey:     utils.RandStringBytesMaskImpr(64),
		privateKey: loadPrivateKey(),
	}
	// Initialize default config
	apiCore.app.Use(cors.New())
	apiCore.app.Use(compress.New())

	if hasLogger {
		apiCore.app.Use(logger.New())
	}

	// Initialize routes
	initRoutes()
}

func initRoutes() {
	// Root api route
	api := apiCore.app.Group("/api")

	auth := api.Group("/auth")
	auth.Get("/conductor/", conductorLogin)

	core := api.Group("/core")
	core.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    apiCore.privateKey.Public(),
		},
	}))

	dev := api.Group("/dev")
	dev.Get("/monitor/:key", monitorRoute)
}

func Listen(port string) {
	log.Printf("YOUR DEV API KEY: %s\nUSE IT TO ACCESS API MONITORING\n", apiCore.devkey)
	apiCore.app.Listen(":" + port)
}

func loadPrivateKey() *rsa.PrivateKey {
	keyFile, err := os.ReadFile("jwtRS256.key")
	if err != nil {
		log.Fatal("Failed to load private key:", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyFile)
	if err != nil {
		log.Fatal("Failed to parse private key:", err)
	}
	return privateKey
}
