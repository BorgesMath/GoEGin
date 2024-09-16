package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	ct "github.com/BorgesMath/GoEGin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestQVerificaStatusCodeInicioParam(t *testing.T) {
	r := SetUpDasRotasDeTeste()
	r.GET("/:nome", ct.Inicio)
	req, _ := http.NewRequest("GET", "/nome", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Devem ser iguais")
	mockDaResp := `{"API":"Eai nome Bom dia"}`
	respBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, mockDaResp, string(respBody))
}
