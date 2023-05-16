package room

import (
	"encoding/json"
	"gather-go/package/room/room_models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func (h handler) Room(c *gin.Context) {
	id := c.Param("id")
	var room room_models.Room

	if err := h.DB.Preload("Owner").Where("id = ?", id).First(&room).Error; err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, &room)

}

func getToken(t *testing.T) {
	r := gin.Default()
	r.GET("/get-token", getRtcToken)

	// create a request object
	req, err := http.NewRequest("GET", "/rtc/test/pubisher/uid/9/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// call the handler function
	r.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// parse the JSON response body
	var resp map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	// check the token value
	if token, ok := resp["rtcToken"]; !ok || token != "my_token_value" {
		t.Errorf("handler returned unexpected token value: got %v, want %v", token, "my_token_value")

	}

}
