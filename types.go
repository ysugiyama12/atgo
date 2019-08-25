package atgo

// AtCoderRating struct
type AtCoderRating struct {
	ContestName       string
	ContestScreenName string
	EndTime           string
	InnerPerformance  int
	IsRated           bool
	NewRating         int
	OldRating         int
	Performance       int
	Place             int
}

type AtCoderUser struct {
	UserID string
	Rating int
	Color string
	Details []AtCoderRating
}