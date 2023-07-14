package repository

import (
	"context"
	"math/rand"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovies(ctx context.Context, movies entities.Movies) (entities.Movies, error)
	GetAllMovies(ctx context.Context) ([]entities.Movies, error)
	GetMovieByID(ctx context.Context, movieID uuid.UUID) (any, error)
	GenerateRandomTimeMovie() ([]entities.TimeMovie, error)
}

type movieRepository struct {
	connection *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{
		connection: db,
	}
}

func (mr *movieRepository) CreateMovies(ctx context.Context, movies entities.Movies) (entities.Movies, error) {
	if err := mr.connection.Create(&movies).Error; err != nil {
		return entities.Movies{}, err
	}
	return movies, nil
}

func (mr *movieRepository) GetAllMovies(ctx context.Context) ([]entities.Movies, error) {
	var movies []entities.Movies
	if err := mr.connection.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (mr *movieRepository) GetMovieByID(ctx context.Context, movieID uuid.UUID) (interface{}, error) {
	var data struct {
		Movies entities.Movies
		Studio []struct {
			ID        uint               
			Name      string            
			TimeMovie []entities.TimeMovie
		}
	}

	var studio []entities.Place

	if err := mr.connection.Where("id = ?", movieID).First(&data.Movies).Error; err != nil {
		return nil, err
	}

	if err := mr.connection.Find(&studio).Error; err != nil {
		return nil, err
	}

	for i := range studio {
		timeMovie, err := mr.GenerateRandomTimeMovie()
		if err != nil {
			return nil, err
		}
		data.Studio = append(data.Studio, struct {
			ID        uint               
			Name      string            
			TimeMovie []entities.TimeMovie
		}{
			ID:        uint(studio[i].ID), 
			Name:      studio[i].Name,
			TimeMovie: timeMovie,
		})
	}

	return data, nil
}

func (mr *movieRepository) GenerateRandomTimeMovie() ([]entities.TimeMovie, error) {
	var timeMovie []entities.TimeMovie

	if err := mr.connection.Where("type = ?", rand.Intn(5) + 1).Find(&timeMovie).Error; err != nil {
		return []entities.TimeMovie{}, err
	}

	return timeMovie, nil
}