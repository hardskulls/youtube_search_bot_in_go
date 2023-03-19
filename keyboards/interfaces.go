package keyboards

import "gopkg.in/telebot.v3"

type CreateBtn interface {
	CreateBtn() telebot.Btn
}

func (st SearchTarget) CreateBtn() telebot.Btn {
	b := telebot.Btn{Text: string(st), Data: string(st)}
	switch st {
	case SearchTargetSubscription:
		b.Text = "Subscription 🏷"
	case SearchTargetPlaylist:
		b.Text = "Playlist ⏯"
	}
	return b
}

func (lt ListTarget) CreateBtn() telebot.Btn {
	b := telebot.Btn{Text: string(lt), Data: string(lt)}
	switch lt {
	case ListTargetSubscription:
		b.Text = "Subscription 🏷"
	case ListTargetPlaylist:
		b.Text = "Playlist ⏯"
	}
	return b
}

func (s Sorting) CreateBtn() telebot.Btn {
	b := telebot.Btn{Text: string(s), Data: string(s)}
	switch s {
	case SortingDate:
		b.Text = "Date 🗓"
	case SortingAlphabetical:
		b.Text = "Alphabetical 🔠"
	}
	return b
}

func (si SearchIn) CreateBtn() telebot.Btn {
	b := telebot.Btn{Text: string(si), Data: string(si)}
	switch si {
	case SearchInTitle:
		b.Text = "Title 🔤"
	case SearchInDescription:
		b.Text = "Description 📃"
	}
	return b
}

func (lb ListButton) CreateBtn() telebot.Btn {
	b := telebot.Btn{Text: string(lb), Data: string(lb)}
	switch lb {
	case ListCancel:
		b.Text = "Cancel ❌"
	case ListSettings:
		b.Text = "Settings ⚙"
	case ListTargetOptions:
		b.Text = "Target 🎯"
	case ListSortingOptions:
		b.Text = "Sorting 🗃"
	case ListResultLimit:
		b.Text = "Result limit 📥"
	case ListExecute:
		b.Text = "Execute ✔"
	}
	return b
}

func (sb SearchButton) CreateBtn() telebot.Btn {
	b := telebot.Btn{Text: string(sb), Data: string(sb)}
	switch sb {
	case SearchCancel:
		b.Text = "Cancel ❌"
	case SearchSettings:
		b.Text = "Settings ⚙"
	case SearchTargetOptions:
		b.Text = "Target 🎯"
	case SearchSearchInOptions:
		b.Text = "Search in 🗃"
	case SearchResultLimit:
		b.Text = "Result limit 📥"
	case SearchExecute:
		b.Text = "Execute ✔"
	}
	return b
}

type CreateKb interface {
	CreateKb() telebot.ReplyMarkup
}

func (st SearchTarget) CreateKb() telebot.ReplyMarkup {
	return SearchSettings.CreateKB()
}

func (lt ListTarget) CreateKb() telebot.ReplyMarkup {
	return ListSettings.CreateKB()
}

func (si SearchIn) CreateKb() telebot.ReplyMarkup {
	return SearchSettings.CreateKB()
}

func (s Sorting) CreateKb() telebot.ReplyMarkup {
	return ListSettings.CreateKB()
}

func (lb ListButton) CreateKB() telebot.ReplyMarkup {
	replyMarkup := telebot.ReplyMarkup{}

	switch lb {
	case ListCancel:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 3)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(ListTargetOptions, ListSortingOptions)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(ListResultLimit)
		replyMarkup.InlineKeyboard[2] = CreateInlineButtonRow(ListExecute, ListCancel)
	case ListSettings:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 3)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(ListTargetOptions, ListSortingOptions)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(ListResultLimit)
		replyMarkup.InlineKeyboard[2] = CreateInlineButtonRow(ListExecute, ListCancel)
	case ListResultLimit:
		replyMarkup.Text("Send result limit")
	case ListTargetOptions:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 2)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(ListTargetSubscription, ListTargetPlaylist)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(ListCancel)
	case ListSortingOptions:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 2)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(SortingDate, SortingAlphabetical)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(ListCancel)
	case ListExecute:
	}

	return replyMarkup
}

func (sb SearchButton) CreateKB() telebot.ReplyMarkup {
	replyMarkup := telebot.ReplyMarkup{}

	switch sb {
	case SearchCancel:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 3)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(SearchTargetOptions, SearchSearchInOptions)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(SearchResultLimit, SearchTextToSearch)
		replyMarkup.InlineKeyboard[2] = CreateInlineButtonRow(SearchExecute, SearchCancel)
	case SearchSettings:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 3)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(SearchTargetOptions, SearchSearchInOptions)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(SearchResultLimit, SearchTextToSearch)
		replyMarkup.InlineKeyboard[2] = CreateInlineButtonRow(SearchExecute, SearchCancel)
	case SearchResultLimit:
		replyMarkup.Text("Send result limit")
	case SearchTargetOptions:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 2)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(SearchTargetSubscription, SearchTargetPlaylist)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(SearchCancel)
	case SearchSearchInOptions:
		replyMarkup.InlineKeyboard = make([][]telebot.InlineButton, 2)
		replyMarkup.InlineKeyboard[0] = CreateInlineButtonRow(SearchInTitle, SearchInDescription)
		replyMarkup.InlineKeyboard[1] = CreateInlineButtonRow(SearchCancel)
	case SearchExecute:
	}

	return replyMarkup
}

func CreateBtnRow(iFaces ...interface{ CreateBtn }) telebot.Row {
	var row telebot.Row
	for _, btn := range iFaces {
		row = append(row, btn.CreateBtn())
	}
	return row
}

func CreateInlineButtonRow(iFaces ...interface{ CreateBtn }) []telebot.InlineButton {
	var row []telebot.InlineButton
	for _, btn := range iFaces {
		row = append(row, *btn.CreateBtn().Inline())
	}
	return row
}
