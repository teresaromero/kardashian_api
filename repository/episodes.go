package repository

import (
	"kardashian_api/custom_errors"
	"kardashian_api/database"
	"kardashian_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllEpisodes() ([]models.Episode, *custom_errors.HttpError) {

	ctx, cancel := database.Context()
	defer cancel()

	cursor, err := database.Use("episodes").Find(ctx, bson.D{})

	if err != nil {
		return nil, custom_errors.InternalServerError(err)
	}
	defer cursor.Close(ctx)

	var episodes []models.Episode
	err = cursor.All(ctx, &episodes)
	if err != nil {
		return nil, custom_errors.InternalServerError(err)
	}

	return episodes, nil
}

func GetEpisodeByNumber(num int) (interface{}, *custom_errors.HttpError) {

	ctx, cancel := database.Context()
	defer cancel()

	result := database.Use("episodes").FindOne(ctx, bson.D{primitive.E{Key: "episode_overall", Value: num}})

	var episode models.Episode
	err := result.Decode(&episode)
	if err != nil {
		return nil, custom_errors.BadRequest(err)
	}
	return episode, nil
}