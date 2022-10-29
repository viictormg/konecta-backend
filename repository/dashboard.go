package repository

import (
	"context"
	"konecta/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) SaveInfo(info model.DashboardReponse) error {
	col := r.db.Collection("dashboard")

	_, err := col.InsertOne(context.Background(), info)

	return err
}

func (r Repository) GetInfoDashboardDBByID(id string) ([]model.DashboardReponse, error) {
	var poinst []model.DashboardReponse
	opt := options.Find()

	filter := bson.M{"id": id}

	rows, err := r.db.Collection("dashboard").Find(context.TODO(), filter, opt)
	if err != nil {
		return poinst, err
	}
	for rows.Next(context.TODO()) {
		var info model.DashboardReponse

		err = rows.Decode(&info)

		if err != nil {
			return poinst, err
		}

		poinst = append(poinst, info)
	}

	return poinst, nil
}
