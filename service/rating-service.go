package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

type RatingService interface {
	All(allDTO dto.AllRatingsDTO) ([]entity.Rating, int64)
	InsertRating(r dto.RatingCreateDTO) entity.Rating
	GetRatingsByProductID(info dto.AllRatingsDTO) ([]entity.Rating, int64)
}

type ratingService struct {
	ratingRepository repository.RatingRepository
}

func NewRatingService(ratingRepo repository.RatingRepository) RatingService {
	return &ratingService{
		ratingRepository: ratingRepo,
	}
}

func (service *ratingService) All(allDTO dto.AllRatingsDTO) ([]entity.Rating, int64) {
	return service.ratingRepository.All(allDTO)
}

func (service *ratingService) InsertRating(r dto.RatingCreateDTO) entity.Rating {
	rating := entity.Rating{}
	err := smapping.FillStruct(&rating, smapping.MapFields(&r))
	if err != nil {
		log.Println("Rating Mapping Failed")
	}
	res := service.ratingRepository.InsertRating(rating)
	return res
}

func (service *ratingService) GetRatingsByProductID(info dto.AllRatingsDTO) ([]entity.Rating, int64) {
	return service.ratingRepository.GetRatingsByProductID(info)
}
