package youtube_related

type SearchedItem interface {
	Title() string
	Description() string
	Link() string
	Date() string
}
