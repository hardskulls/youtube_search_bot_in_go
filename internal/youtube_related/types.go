package youtube_related

import "google.golang.org/api/youtube/v3"

type SubscriptionAlias youtube.Subscription

func (s *SubscriptionAlias) Title() string {
	return s.Snippet.Title
}

func (s *SubscriptionAlias) Description() string {
	return s.Snippet.Description
}

func (s *SubscriptionAlias) Link() string {
	return "https://youtube.com/channel/" + s.Snippet.ResourceId.ChannelId
}

func (s *SubscriptionAlias) Date() string {
	return s.Snippet.PublishedAt
}

type PlaylistAlias youtube.Playlist

func (p *PlaylistAlias) Title() string {
	return p.Snippet.Title
}

func (p *PlaylistAlias) Description() string {
	return p.Snippet.Description
}

func (p *PlaylistAlias) Link() string {
	return "https://youtube.com/playlist?list=" + p.Id
}

func (p *PlaylistAlias) Date() string {
	return p.Snippet.PublishedAt
}
