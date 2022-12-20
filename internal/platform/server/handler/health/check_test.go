package health

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Check(t *testing.T) {
	r := fiber.New()
	r.Get("/health", CheckHandler())

	t.Run("it returns 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/health", nil)
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		res, _ := r.Test(req)
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
