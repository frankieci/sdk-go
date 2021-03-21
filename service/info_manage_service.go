package service

import (
	"github.com/pkg/errors"
	"gitlab.com/sdk-go/factory"
	"gitlab.com/sdk-go/library/logger"
	"gitlab.com/sdk-go/wire"
)

type InfoManageService struct {
}

func NewInfoManageService() *InfoManageService {
	return &InfoManageService{}
}

func (svc *InfoManageService) GetInfo(iType *wire.InfoType) (interface{}, error) {
	collectorFactory := factory.NewCollectorFactory()
	collectorFactory.SetCollectorType(iType.IType)
	collector := collectorFactory.CreateCollector()
	data, err := collector.Collect()
	if err != nil {
		logger.Error(errors.WithStack(err))
		return nil, err
	}
	return data, nil
}
