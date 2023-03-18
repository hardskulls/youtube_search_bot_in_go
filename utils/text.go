package utils

import (
	"fmt"
	"strings"
)

/*
Source https://core.telegram.org/bots/api#formatting-options

<b>bold</b>, <strong>bold</strong>
<i>italic</i>, <em>italic</em>
<u>underline</u>, <ins>underline</ins>
<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
<span class="tg-spoiler">spoiler</span>, <tg-spoiler>spoiler</tg-spoiler>
<b>bold <i>italic bold <s>italic bold strikethrough <span class="tg-spoiler">italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>
<a href="http://www.example.com/">inline URL</a>
<a href="tg://user?id=123456789">inline mention of a user</a>
<code>inline fixed-width code</code>
<pre>pre-formatted fixed-width code block</pre>
<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>
*/

type HTMLizer struct{}

// !! Should be used BEFORE (!) applying any HTML tags, as it will break them !!
func (*HTMLizer) ReplaceUnsupported(text string) string {
	return strings.ReplaceAll(text, "<", "&lt;")
}

// Makes text bold. Uses '<b>...</b>' HTML tag.
func (*HTMLizer) Bold(text string) string {
	return fmt.Sprintf("<b>%v</b>", text)
}

// Makes text bold. Uses '<strong>...</strong>' HTML tag.
func (*HTMLizer) StrongBold(text string) string {
	return fmt.Sprintf("<strong>%v</strong>", text)
}

// Makes text italic. Uses '<i>...</i>' HTML tag.
func (*HTMLizer) Italic(text string) string {
	return fmt.Sprintf("<i>%v</i>", text)
}

// Makes text italic. Uses '<em>...</em>' HTML tag.
func (*HTMLizer) StrongItalic(text string) string {
	return fmt.Sprintf("<em>%v</em>", text)
}

// Makes text underline. Uses '<u>...</u>' HTML tag.
func (*HTMLizer) Underline(text string) string {
	return fmt.Sprintf("<u>%v</u>", text)
}

// Makes text underline. Uses '<ins>...</ins>' HTML tag.
func (*HTMLizer) StrongUnderline(text string) string {
	return fmt.Sprintf("<ins>%v</ins>", text)
}

// Makes text strikethrough. Uses '<s>...</s>' HTML tag.
func (*HTMLizer) Strikethrough(text string) string {
	return fmt.Sprintf("<s>%v</s>", text)
}

// Makes text strikethrough. Uses '<strike>...</strike>' HTML tag.
func (*HTMLizer) StrongStrikethrough(text string) string {
	return fmt.Sprintf("<strike>%v</strike>", text)
}

// Makes text strikethrough. Uses '<del>...</del>' HTML tag.
func (*HTMLizer) AlterStrikethrough(text string) string {
	return fmt.Sprintf("<del>%v</del>", text)
}

// Makes text a telegram spoiler. Uses '<span class="tg-spoiler">...</span>' HTML tag.
func (*HTMLizer) Spoiler(text string) string {
	return fmt.Sprintf(`<span class="tg-spoiler">%v</span>`, text)
}

// Makes text a telegram spoiler. Uses '<tg-spoiler>...</tg-spoiler>' HTML tag.
func (*HTMLizer) StrongSpoiler(text string) string {
	return fmt.Sprintf("<tg-spoiler>%v</tg-spoiler>", text)
}

// Makes text inlined URL. Uses '<a href=...>...</a>' HTML tag.
func (*HTMLizer) InlineURL(linkText string, url string) string {
	return fmt.Sprintf(`<a href="%v">%v</a>`, url, linkText)
}

// Makes text inlined user mention. Uses '<a href="tg://user?id=...">...</a>' HTML tag.
func (*HTMLizer) InlineUserMention(mentionText string, userId string) string {
	return fmt.Sprintf(`<a href="tg://user?id=%v">%v</a>`, userId, mentionText)
}

// Makes text a code block. Uses '<code>...</code>' HTML tag.
func (*HTMLizer) CodeBlock(code string) string {
	return fmt.Sprintf("<code>%v</code>", code)
}

// Makes text a code block. Uses '<pre>...</pre>' HTML tag.
func (*HTMLizer) StrongCodeBlock(code string) string {
	return fmt.Sprintf("<pre>%v</pre>", code)
}

// Makes text a code block. Uses '<pre><code class="language-...">...</code></pre>' HTML tag.
func (*HTMLizer) AlterCodeBlock(lang string, code string) string {
	return fmt.Sprintf(`<pre><code class="language-%v">%v</code></pre>`, lang, code)
}

/*
func (*HTMLizer) Striketrough(text string) string {
	return fmt.Sprintf("<b>%v</b>", text)
}

type HTMLize interface {
	Bold(text string) string
	StrongBold(text string) string
	Italic(text string) string
	StrongItalic(text string) string
	Striketrough(text string) string
} */
