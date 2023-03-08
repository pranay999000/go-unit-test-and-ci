package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRootHandler(t *testing.T) {
	mockResponse := `{"message":"Hello World!"}`

	r := SetUpRouter()
	r.GET("/", RootHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(resp))
	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestRootHandlerDocker(t *testing.T) {
// 	mockResponse := `{"message":"Hello World!"}`

// 	pool, err := dockertest.NewPool("")
// 	require.NoError(t, err, "could not connect to Docker")

// 	resource, err := pool.Run("unit-test", "multistage", []string{})
// 	require.NoError(t, err, "could not start container")

// 	t.Cleanup(func() {
// 		require.NoError(t, pool.Purge(resource), "failed to remove container")
// 	})

// 	var resp *http.Response

// 	err = pool.Retry(func() error {
// 		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetPort("8080/tcp"), "/"))
// 		if err != nil {
// 			t.Log("container not ready, waiting...")
// 			return err
// 		}
// 		return nil
// 	})
// 	require.NoError(t, err, "HTTP error")
// 	defer resp.Body.Close()

// 	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

// 	body, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err, "failed to read HTTP body")

// 	require.Equal(t, string(body), mockResponse)
// }