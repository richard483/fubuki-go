package impl

import (
	"fubuki-go/dto/request"
	"fubuki-go/model"
	"fubuki-go/repository"
)

type GeminiHistoryService struct {
	Repository repository.HistoryRepositoryInterface
}

func NewGeminiHistoryService(Repository repository.HistoryRepositoryInterface) *GeminiHistoryService {
	return &GeminiHistoryService{Repository: Repository}
}

func (srv *GeminiHistoryService) CreateHistoryData(historyData *request.History) error {

	err := srv.Repository.Create(&model.History{
		UserQuestion: historyData.UserQuestion,
		ModelAnswer:  historyData.ModelAnswer,
		ModelSource:  historyData.ModelSource,
	})
	return err
}

func (srv *GeminiHistoryService) CreateManyHistoryData(historiesData *[]request.History) error {

	var histories []model.History
	for _, data := range *historiesData {
		histories = append(histories, model.History{
			UserQuestion: data.UserQuestion,
			ModelAnswer:  data.ModelAnswer,
			ModelSource:  data.ModelSource,
		})
	}
	err := srv.Repository.CreateMany(&histories)
	return err
}

func (srv *GeminiHistoryService) GetAllHistoryDataByModelSource(modelSource string) *[]model.History {
	results := srv.Repository.GetAllByModelSource(modelSource)

	return &results
}

func (srv *GeminiHistoryService) UpdateHistoryData(historyData *request.UpdateHistory) error {
	err := srv.Repository.Update(&model.History{
		ID:           historyData.ID,
		UserQuestion: historyData.UserQuestion,
		ModelAnswer:  historyData.ModelAnswer,
	})

	return err
}

func (srv *GeminiHistoryService) DeleteHistoryData(id string) error {
	err := srv.Repository.Delete(id)
	return err
}
