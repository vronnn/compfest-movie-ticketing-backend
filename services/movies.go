package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
)

type MovieService interface {
	CreateMovies(ctx context.Context, movieDTO dto.MovieCreateDTO) (entities.Movies, error)
	GetAllMovies(ctx context.Context) ([]entities.Movies, error)
	GetMovieByID(ctx context.Context, movieID uuid.UUID) (any, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(mr repository.MovieRepository) MovieService {
	return &movieService{
		movieRepository: mr,
	}
}

func (ms *movieService) CreateMovies(ctx context.Context, movieDTO dto.MovieCreateDTO) (entities.Movies, error) {
	var movie entities.Movies

	t, err := time.Parse("2006-01-02", movieDTO.ReleaseDate)
	if err != nil {
		fmt.Println("Error parsing time:", err)
    return entities.Movies{}, err
	}

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error load location:", err)
		return entities.Movies{}, err
	}

	movie.ReleaseDate = t.In(location)
	movie.Title = movieDTO.Title
	movie.Description = movieDTO.Description
	movie.AgeRating = movieDTO.AgeRating
	movie.TicketPrice = movieDTO.TicketPrice
	fmt.Println("movie.ReleaseDate:", movie.ReleaseDate)

	return ms.movieRepository.CreateMovies(ctx, movie)
}

func (ms *movieService) GetAllMovies(ctx context.Context) ([]entities.Movies, error) {
	return ms.movieRepository.GetAllMovies(ctx)
}

func (ms *movieService) GetMovieByID(ctx context.Context, movieID uuid.UUID) (any, error) {
	return ms.movieRepository.GetMovieByID(ctx, movieID)
}
