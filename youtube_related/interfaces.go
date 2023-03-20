package youtube_related

// Anything that can be searched on user's YouTube channel.
type SearchedItem interface {
	Title() string
	Description() string
	Link() string
	Date() string
}
