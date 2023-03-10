package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ArthurQR98/challenge_fiber/kit/command/commandmocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	commandBus := new(commandmocks.Bus)
	commandBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("creating.CourseCommand"),
	).Return(nil)

	r := fiber.New()
	r.Post("/courses", CreateHandler(commandBus))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			Name:     "Demo Course",
			Duration: "10 months",
		}
		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		res, err := r.Test(req, -1)
		require.NoError(t, err)

		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Name:     "Demo Course",
			Duration: "10 months",
		}
		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
		require.NoError(t, err)

		res, _ := r.Test(req, -1)

		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("given a valid request with invalid id returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "ba57",
			Name:     "Demo Course",
			Duration: "10 months",
		}
		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		res, _ := r.Test(req)

		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
