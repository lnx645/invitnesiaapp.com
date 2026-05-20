package config

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v3"
)

type AppCtx = fiber.Ctx
type structValidator struct {
	validate *validator.Validate
}

// validate
func (c *structValidator) Validate(out any) error {
	return c.validate.Struct(out)
}

func FiberConfig(viewsDir embed.FS) *fiber.Config {
	sub, err := fs.Sub(viewsDir, "views")
	if err != nil {
		panic("Failed to embed views folder: " + err.Error())
	}
	engine := html.NewFileSystem(http.FS(sub), ".html")
	return &fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
		Views:           engine,
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
		StrictRouting:   true,
		ServerHeader:    "Invitnesia Server",
		AppName:         "Invitnesia Api",
	}
}
