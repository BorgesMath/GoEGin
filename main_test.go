package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"encoding/json"

	ct "github.com/BorgesMath/GoEGin/controllers"
	"github.com/BorgesMath/GoEGin/database"
	"github.com/BorgesMath/GoEGin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetUpDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaMetaMock() {

	metaTeste := models.Meta{
		Nome:  "Meta Exemplo",
		Check: false,
		Tempo: 120,
		Passos: []models.Passo{
			{
				Nome:      "Passo 1",
				Check:     false,
				Tempo:     30,
				Descricao: "Este é o primeiro passo",
			},
			{
				Nome:      "Passo 2",
				Check:     true,
				Tempo:     45,
				Descricao: "Este é o segundo passo",
			},
		},
	}

	database.DB.Create(&metaTeste)
	ID = int(metaTeste.ID)

}

func DeletaMetaMock() {
	var metaTeste models.Meta
	database.DB.Delete(&metaTeste, ID)
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

func TestListaMetas(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaMetaMock()
	defer DeletaMetaMock()

	r := SetUpDasRotasDeTeste()
	r.GET("/metas", ct.ExibeTodasMetas)
	req, _ := http.NewRequest("GET", "/metas", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	fmt.Println(resp.Body)
}

func TestBuscaMetaPorID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaMetaMock()
	defer DeletaMetaMock()
	r := SetUpDasRotasDeTeste()
	r.GET("/metas/:id", ct.AchaMetaPorID)

	caminho := fmt.Sprintf("/metas/%d", ID)

	req, _ := http.NewRequest("GET", caminho, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var metaMock models.Meta
	json.Unmarshal(resp.Body.Bytes(), &metaMock)
	fmt.Println(metaMock.Nome)

	assert.Equal(t, "Meta Exemplo", metaMock.Nome)

}

func DeletaMetaHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaMetaMock()
	r := SetUpDasRotasDeTeste()
	r.DELETE("/metas/:id", ct.DeletaMeta)

	pathBusca := "/metas/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", pathBusca, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}
