package database

import (
	"errors"
	"fmt"
	"github.com/ivanmeca/DesafioPismo/v2/internal/model"
	"gorm.io/gorm"
)

type IEventRepository interface {
	CreateLogEvent(e *model.LogEventMessage) (model.LogEventMessage, error)
	CreateMonitoringEvent(e *model.MonitoringEventMessage) (model.LogEventMessage, error)
	CreateUserOperationEvent(e *model.UserOperationEventMessage) (model.LogEventMessage, error)
}

type EventRepository struct {
	IEventRepository
}

func (r *Repository) CreateWorkerFunction(workerFunction model.WorkerFunction) (*model.InnerWorkerFunction, error) {

	fmt.Println("workerFunction ", workerFunction)

	err := r.db.Create(&workerFunction).Error
	if err != nil {
		return nil, err
	}

	return &workerFunction.InnerWorkerFunction, nil
}

func (r *Repository) GetWorkerFunctions() ([]model.WorkerFunctionRequest, error) {

	var workerFunction []model.WorkerFunctionRequest
	result := r.db.Where("deleted_at IS NULL ORDER BY id ASC").Find(&workerFunction)
	if result.Error != nil {
		//"error": "falha ao executar a consulta"
		return nil, result.Error
	}

	if workerFunction == nil {
		//"msg": "funções não encontradas"
		return nil, result.Error
	}

	if len(workerFunction) == 0 {
		//"funções não encontradas"
		return nil, result.Error
	}

	return workerFunction, nil

}

func (r *Repository) GetWorkerFunctionByID(id string) (*model.WorkerFunctionRequest, error) {

	var workerFunction model.WorkerFunctionRequest
	result := r.db.Where("deleted_at IS NULL").Find(&workerFunction, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	if result.Error != nil {
		//"falha ao executar a consulta "
		return nil, result.Error
	}

	return &workerFunction, nil
}

func (r *Repository) PutWorkerFunctionID(id string, workerFunctionUpdate model.WorkerFunction) (*model.WorkerFunction, error) {

	var workerFunction model.WorkerFunction
	result := r.db.First(&workerFunction, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	if result.Error != nil {
		//"falha ao executar a consulta "
		return nil, result.Error
	}

	// Atualização dos campos desejados
	workerFunction.InnerWorkerFunction = workerFunctionUpdate.InnerWorkerFunction
	workerFunction.UpdatedBy = workerFunctionUpdate.UpdatedBy

	err := r.db.Save(&workerFunction).Error
	if err != nil {
		// "cannot update function: "
		return nil, err
	}

	return &workerFunction, nil

}

func (r *Repository) DeleteWorkerFunction(id string, name string) error {

	var workerFunction model.WorkerFunction
	result := r.db.Delete(&workerFunction, id)
	if result.Error != nil {
		// "falha ao excluir função: "
		return result.Error
	} else {
		errExec := r.db.Exec("UPDATE worker_function SET deleted_by = ? WHERE id = ?;", name, id)
		if errExec.Error != nil {
			//"error": "falha ao registrar o usuario que executou a ação: "
			return errExec.Error
		}

	}

	if result.RowsAffected == 0 {
		// "função não encontrada"
		return errors.New("Not found")
	}

	return nil

}
