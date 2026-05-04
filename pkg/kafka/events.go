package kafka

import "time"

const TopicMovieSearched = "movie.searched"

// SearchEvent represents a movie search event published to Kafka.
type SearchEvent struct {
	Filters      SearchFilters `json:"filters"`
	ResultsCount int           `json:"resultsCount"`
	Timestamp    time.Time     `json:"timestamp"`
}

// SearchFilters represents the filters used in a movie search.
type SearchFilters struct {
	DirectorIDs []int    `json:"directorIds,omitempty"`
	GenreIDs    []int    `json:"genreIds,omitempty"`
	MinYear     *int     `json:"minYear,omitempty"`
	MaxYear     *int     `json:"maxYear,omitempty"`
	MinDuration *int     `json:"minDuration,omitempty"`
	MaxDuration *int     `json:"maxDuration,omitempty"`
	MinRating   *float64 `json:"minRating,omitempty"`
	MaxRating   *float64 `json:"maxRating,omitempty"`
}
