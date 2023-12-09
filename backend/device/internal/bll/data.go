package bll

import (
	"context"
	"ims-server/device/internal/dal/repo"
	"ims-server/device/internal/param"
	"ims-server/device/pkg/pack"

	egoerror "ims-server/pkg/error"
	"ims-server/pkg/util"
)

type dataService struct {
}

func NewDataService() *dataService {
	return &dataService{}
}

// TODO：以下全待测：@姚礼兴
func (s *dataService) GetDataByID(ctx context.Context, req *param.GetDataByID) (*param.GetDataByIDResponse, error) {
	id := req.ID
	sensorDate, err := repo.NewSensorRepo().Get(ctx, id)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}

	resp := pack.ToDataResponse(sensorDate)
	return &param.GetDataByIDResponse{
		DataResponse: resp,
	}, nil
}

func (s *dataService) MGetSensorByIDs(ctx context.Context, req *param.MGetDataByIDsRequest) (*param.MGetDataByIDsResponse, error) {
	res, err := repo.NewSensorRepo().MGet(ctx, req.IDs)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}

	resp := []param.GetDataByIDResponse{}
	for _, sensorDate := range res {
		info := pack.ToDataResponse(&sensorDate)
		resp = append(resp, param.GetDataByIDResponse{
			DataResponse: info,
		})
	}

	return &param.MGetDataByIDsResponse{
		List: resp,
	}, nil
}

func (s *dataService) UpdateMqDateByID(ctx context.Context, req *param.UpdateDataByIDRequest) (*param.UpdateDataByIDResponse, error) {
	_, err1 := repo.NewSensorRepo().Get(ctx, req.ID)
	if err1 != nil {
		return nil, egoerror.ErrNotFound
	}

	mqDateMap := util.RequestToSnakeMapWithIgnoreZeroValueAndIDKey(req)

	update, err2 := repo.NewSensorRepo().Update(ctx, req.ID, mqDateMap)
	if err2 != nil {
		return nil, egoerror.ErrInvalidParam
	}

	resp := pack.ToDataResponse(update)
	return &param.UpdateDataByIDResponse{
		DataResponse: resp,
	}, nil
}

func (s *dataService) DeleteMqDaterByID(ctx context.Context, req *param.DeleteDataByIDRequest) (*param.DeleteDataByIDResponse, error) {
	_, err := repo.NewSensorRepo().Get(ctx, req.ID)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}
	err = repo.NewSensorRepo().Delete(ctx, req.ID)
	if err != nil {
		return nil, egoerror.ErrNotFound
	}
	return &param.DeleteDataByIDResponse{}, nil
}
