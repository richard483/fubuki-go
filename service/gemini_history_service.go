package service

import (
	"fubuki-go/dto/request"
	"fubuki-go/model"
	repository "fubuki-go/repository_interface"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GeminiHistoryService struct {
	Repository repository.GeminiHistoryRepositoryInterface
}

func NewGeminiHistoryService(Repository repository.GeminiHistoryRepositoryInterface) *GeminiHistoryService {
	return &GeminiHistoryService{Repository: Repository}
}

func (srv *GeminiHistoryService) CreateHistoryData(c *gin.Context) {
	var historyData request.CreateGeminiHistory
	if err := c.BindJSON(&historyData); err != nil {
		return
	}

	rowAffected := srv.Repository.Create(&model.History{
		UserQuestion: historyData.UserQuestion,
		ModelAnswer:  historyData.ModelAnswer,
	})

	c.IndentedJSON(http.StatusOK, rowAffected)
	return
}

func (srv *GeminiHistoryService) GetAllHistoryData(c *gin.Context) {
	results := srv.Repository.GetAll()

	c.IndentedJSON(http.StatusOK, results)
	return
}

func (srv *GeminiHistoryService) UpdateHistoryData(c *gin.Context) {
	var historyData request.UpdateGeminiHistory
	if err := c.BindJSON(&historyData); err != nil {
		log.Fatalln(err)
		return
	}

	rowAffected := srv.Repository.Update(&model.History{
		UserQuestion: historyData.UserQuestion,
		ModelAnswer:  historyData.ModelAnswer,
	})

	c.IndentedJSON(http.StatusOK, rowAffected)
	return
}

func (srv *GeminiHistoryService) DeleteHistoryData(c *gin.Context) {

	res, ok := c.GetQuery("id")

	if !ok {
		log.Fatalln("No 'id' parameter found on REST request")
		return
	}
	rowAffected := srv.Repository.Delete(res)

	c.IndentedJSON(http.StatusOK, rowAffected)
	return
}
