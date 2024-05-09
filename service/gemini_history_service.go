package service

import (
	"fubuki-go/dto/request"
	"fubuki-go/model"
	repository "fubuki-go/repository_interface"
)

type GeminiHistoryService struct {
	Repository repository.GeminiHistoryRepositoryInterface
}

func NewGeminiHistoryService(Repository repository.GeminiHistoryRepositoryInterface) *GeminiHistoryService {
	return &GeminiHistoryService{Repository: Repository}
}

func (srv *GeminiHistoryService) CreateHistoryData(historyData *request.GeminiHistory) error {

	err := srv.Repository.Create(&model.History{
		UserQuestion: historyData.UserQuestion,
		ModelAnswer:  historyData.ModelAnswer,
	})
	return err
}

func (srv *GeminiHistoryService) CreateManyHistoryData(historiesData *[]request.GeminiHistory) error {

	var histories []model.History
	for _, data := range *historiesData {
		histories = append(histories, model.History{
			UserQuestion: data.UserQuestion,
			ModelAnswer:  data.ModelAnswer,
		})
	}
	err := srv.Repository.CreateMany(&histories)
	return err
}

func (srv *GeminiHistoryService) GetAllHistoryData() *[]model.History {
	results := srv.Repository.GetAll()

	return &results
}

func (srv *GeminiHistoryService) UpdateHistoryData(historyData *request.GeminiHistory) error {
	err := srv.Repository.Update(&model.History{
		UserQuestion: historyData.UserQuestion,
		ModelAnswer:  historyData.ModelAnswer,
	})

	return err
}

func (srv *GeminiHistoryService) DeleteHistoryData(id string) error {
	err := srv.Repository.Delete(id)
	return err
}
