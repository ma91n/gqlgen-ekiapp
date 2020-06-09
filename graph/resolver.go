package graph

import (
	"context"
	"database/sql"
	"errors"
	"github.com/laqiiz/gqlgen-ekiapp/graph/model"
	"github.com/laqiiz/gqlgen-ekiapp/models"
	_ "github.com/lib/pq"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var db *sql.DB

func init() {
	conn, err := sql.Open("postgres", "user=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = conn
}

func (r *Resolver) getStationByName(ctx context.Context, name *string) ([]*model.Station, error) {
	sts, err := models.StationByNamesByStationName(db, *name)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.Station, 0, len(sts))
	for _, v := range sts {
		resp = append(resp, &model.Station{
			StationCd:   v.StationCd,
			StationName: v.StationName,
			LineName:    &v.LineName,
			Address:     &v.Address,
		})
	}
	return resp, nil
}

func (r *Resolver) getStationByCD(ctx context.Context, stationCd *int) (*model.Station, error) {
	stations, err := models.StationByCDsByStationCD(db, *stationCd)
	if err != nil {
		return nil, err
	}
	if len(stations) == 0 {
		return nil, errors.New("not found")
	}
	first := stations[0]

	return &model.Station{
		StationCd:   first.StationCd,
		StationName: first.StationName,
		LineName:    &first.LineName,
		Address:     &first.Address,
	}, nil
}

func (r *Resolver) transferStation(ctx context.Context, obj *model.Station) ([]*model.Station, error) {
	stationCd := obj.StationCd

	records, err := models.TransfersByStationCD(db, stationCd)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.Station, 0, len(records))
	for _, v := range records {
		if v.TransferStationName == "" {
			continue
		}
		resp = append(resp, &model.Station{
			StationCd:   v.TransferStationCd,
			StationName: v.TransferStationName,
			LineName:    &v.TransferLineName,
			Address:     &v.TransferAddress,
		})
	}

	return resp, nil
}

func (r *Resolver) beforeStation(ctx context.Context, obj *model.Station) (*model.Station, error) {
	stationCD := obj.StationCd
	records, err := models.BeforesByStationCD(db, stationCD)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	if records[0].BeforeStationName == "" {
		return nil, nil
	}

	return &model.Station{
		StationCd:   records[0].BeforeStationCd,
		StationName: records[0].BeforeStationName,
		LineName:    &records[0].LineName,
		Address:     &records[0].BeforeStationAddress,
	}, nil
}

func (r *Resolver) afterStation(ctx context.Context, obj *model.Station) (*model.Station, error) {
	stationCD := obj.StationCd
	records, err := models.AftersByStationCD(db, stationCD)
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	if records[0].AfterStationName == "" {
		return nil, nil
	}

	return &model.Station{
		StationCd:   records[0].AfterStationCd,
		StationName: records[0].AfterStationName,
		LineName:    &records[0].LineName,
		Address:     &records[0].AfterStationAddress,
	}, nil

}
