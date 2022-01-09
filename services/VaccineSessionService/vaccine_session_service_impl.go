package VaccineSessionService

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"vaccine-app-be/drivers/records"
	"vaccine-app-be/drivers/repository/VaccineRepository"
	"vaccine-app-be/drivers/repository/VaccineSessionRepository"
)

type VaccineSessionServiceImpl struct {
	vaccineSessionRepository VaccineSessionRepository.VaccineSessionRepository
	vaccineRepository        VaccineRepository.VaccineRepository
}

func NewSessionService(vaccineSessionRepository VaccineSessionRepository.VaccineSessionRepository, vaccineRepository VaccineRepository.VaccineRepository) VaccineSessionService {
	return &VaccineSessionServiceImpl{vaccineSessionRepository: vaccineSessionRepository, vaccineRepository: vaccineRepository}
}

func (service *VaccineSessionServiceImpl) CreateSession(ctx context.Context, domain VaccineSession) (VaccineSession, error) {
	err := domain.ValidateRequest()
	if err != nil {
		return VaccineSession{}, err
	}

	entitySession := records.VaccineSession{}
	copier.Copy(&entitySession, domain)
	vaccineData, err := service.vaccineRepository.FindVaccineById(ctx, domain.VaccineId)
	if err != nil {
		return VaccineSession{}, err
	}

	if domain.Quota > vaccineData.Stock {
		return VaccineSession{}, errors.New("vaccine stocks could not meet demand")
	} else {
		data, err := service.vaccineSessionRepository.Create(ctx, entitySession)
		if err != nil {
			return VaccineSession{}, err
		}

		//kurangi jumlah stock di DB untuk data realtime.
		entityVaccine := records.Vaccine{}
		entityVaccine.Stock = vaccineData.Stock - data.Quota
		_ , err = service.vaccineRepository.Update(ctx, vaccineData.HealthFacilitatorId, domain.VaccineId, entityVaccine)
		if err != nil {
			return VaccineSession{}, err
		}
		response := VaccineSession{}
		copier.Copy(&response, &data)

		return response, nil
	}
}

func (service *VaccineSessionServiceImpl) GetSessionById(ctx context.Context, id int) (VaccineSession, error) {
	data, err := service.vaccineSessionRepository.FindById(ctx, id)
	if err != nil {
		return VaccineSession{}, err
	}
	response := VaccineSession{}
	copier.Copy(&response, &data)

	return response, nil
}

func (service *VaccineSessionServiceImpl) GetSessionOwnedByHf(ctx context.Context, hfid int) ([]VaccineSession, error) {
	data, err := service.vaccineSessionRepository.FindSessionOwnedByHf(ctx, hfid)
	if err != nil {
		return nil, err
	}
	var response []VaccineSession
	copier.Copy(&response, &data)

	return response, nil
}

func (service *VaccineSessionServiceImpl) DeleteSession(ctx context.Context, hfid, id int) (string, error) {
	byId, err := service.vaccineSessionRepository.FindById(ctx, id)
	if err != nil {
		return "", err
	}
	if byId.SessionType == "" {
		return "", errors.New("data not found")
	}

	if byId.HealthFacilitatorId == hfid {
		_, err = service.vaccineSessionRepository.Delete(ctx, id, hfid)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("this session doesn't belongs to you")
	}

	return "success delete data", nil
}

func (service *VaccineSessionServiceImpl) UpdateSession(ctx context.Context, hfid, id int, domain VaccineSession) (VaccineSession, error) {
	err := domain.ValidateRequest()
	if err != nil {
		return VaccineSession{}, err
	}
	byId, err := service.vaccineSessionRepository.FindById(ctx, id)
	if err != nil {
		return VaccineSession{}, err
	}

	if byId.HealthFacilitatorId == hfid {
		entitySession := records.VaccineSession{}
		copier.Copy(&entitySession, domain)
		update, err := service.vaccineSessionRepository.Update(ctx, id, hfid, entitySession)
		if err != nil {
			return VaccineSession{}, err
		}
		update.Id = id
		update.HealthFacilitatorId = hfid
		response := VaccineSession{}
		copier.Copy(&response, &update)
		return response, nil
	} else {
		return VaccineSession{}, errors.New("cannot update this session, doesn't belong to yours")
	}
}

func (service *VaccineSessionServiceImpl) GetAllVaccineSession(ctx context.Context) ([]VaccineSession, error) {
	session, err := service.vaccineSessionRepository.GetAllVaccineSession(ctx)
	if err != nil {
		return nil, err
	}

	var response []VaccineSession
	copier.Copy(&response, &session)

	return response, nil
}
