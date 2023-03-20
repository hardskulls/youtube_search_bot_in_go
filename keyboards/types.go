package keyboards

type ListButton string

const (
	ListCancel         ListButton = "List/Cancel"
	ListSettings       ListButton = "List/ListSettings"
	ListResultLimit    ListButton = "List/ResultLimit"
	ListTargetOptions  ListButton = "List/TargetOptions"
	ListSortingOptions ListButton = "List/SortingOptions"
	ListExecute        ListButton = "List/Execute"
)

type SearchButton string

const (
	SearchCancel          SearchButton = "Search/Cancel"
	SearchSettings        SearchButton = "Search/ListSettings"
	SearchResultLimit     SearchButton = "Search/ResultLimit"
	SearchTextToSearch    SearchButton = "Search/SearchTextToSearch"
	SearchTargetOptions   SearchButton = "Search/TargetOptions"
	SearchSearchInOptions SearchButton = "Search/SearchSearchInOptions"
	SearchExecute         SearchButton = "Search/Execute"
)

// Defines what to search: subscription or playlist.
type SearchTarget string

const (
	SearchTargetSubscription SearchTarget = "SearchTarget/Subscription"
	SearchTargetPlaylist     SearchTarget = "SearchTarget/Playlist"
)

// Defines what to search: subscription or playlist.
type ListTarget string

const (
	ListTargetSubscription ListTarget = "ListTarget/Subscription"
	ListTargetPlaylist     ListTarget = "ListTarget/Playlist"
)

// Defines where to search user text: in title or in description.
type SearchIn string

const (
	SearchInTitle       SearchIn = "SearchIn/Title"
	SearchInDescription SearchIn = "SearchIn/Description"
)

type Sorting string

const (
	SortingDate         Sorting = "Sorting/Date"
	SortingAlphabetical Sorting = "Sorting/Alphabetical"
)
