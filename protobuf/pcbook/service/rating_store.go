package service

import "sync"

type Rating struct {
	Count uint32
	Sum   float64
}

type RatingStore interface {
	Add(laptopID string, score float64) (*Rating, error)
}

type InMemoryRatingScore struct {
	mutex  sync.RWMutex
	rating map[string]*Rating
}

func (store *InMemoryRatingScore) Add(laptopID string, score float64) (*Rating, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	rating := store.rating[laptopID]
	if rating == nil {
		rating = &Rating{
			Count: 1,
			Sum:   score,
		}
	} else {
		rating.Count++
		rating.Sum += score
	}

	store.rating[laptopID] = rating
	return rating, nil
}

func NewInMemoryRatingStore() *InMemoryRatingScore {
	return &InMemoryRatingScore{
		rating: make(map[string]*Rating),
	}
}
