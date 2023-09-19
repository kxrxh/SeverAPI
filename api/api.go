package api

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/goccy/go-json"
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

// Init initializes the API core.
//
// It takes a boolean parameter `hasLogger` which indicates whether the API core should have a logger or not.
// It does not return any value.
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

// initRoutes initializes the routes for the API.
//
// It sets up the root API route and creates various sub-routes for different functionalities.
// The function does not take any parameters and does not return any values.
func initRoutes() {
	// Root api route
	api := apiCore.app.Group("/api")

	// Route with authentication
	auth := api.Group("/auth")
	auth.Get("/conductor/", conductorLogin)
	auth.Get("/manager/", managerLogin)

	core := api.Group("/core")
	// core.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{
	// 		JWTAlg: jwtware.RS256,
	// 		Key:    apiCore.privateKey.Public(),
	// 	},
	// }))

	addManagerRoutes(core.Group("/manager"))

	addConductorRoutes(core.Group("/conductor"))

	// Dev routes. It doesn't use JWT. Just DEV TOKEN
	dev := api.Group("/dev")
	dev.Get("/monitor/:key", monitorRoute)
}

// Listen listens for incoming requests on the specified port.
//
// port: a string representing the port number to listen on.
// This function does not return any value.
func Listen(port string) {
	log.Printf("YOUR DEV API KEY: %s\nUSE IT TO ACCESS API MONITORING\n", apiCore.devkey)
	apiCore.app.Listen(":" + port)
}

// loadPrivateKey loads the private key used for JWT authentication.
//
// It reads the key file "jwtRS256.key" and parses it into an RSA private key.
// If there is an error reading the file or parsing the key, it logs a fatal
// error and terminates the program.
//
// Returns the parsed RSA private key.
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
