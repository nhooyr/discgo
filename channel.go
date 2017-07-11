package discgo

import (
	"time"

	"encoding/json"

	"path"

	"net/url"
	"strconv"

	"bytes"
	"io"
	"mime/multipart"
)

// Channel represents a channel in Discord.
// If Recipient is set this is a DM Channel otherwise this is a Guild Channel.
// Guild Channel represents an isolated set of users and messages within a Guild.
// DM Channel represent a one-to-one conversation between two users, outside of the
// scope of guilds.
type Channel struct {
	ID                   string       `json:"id"`
	GuildID              string       `json:"guild_id"`
	Name                 string       `json:"name"`
	Type                 string       `json:"type"`
	Position             int          `json:"position"`
	IsPrivate            bool         `json:"is_private"`
	PermissionOverwrites []*Overwrite `json:"permission_overwrites"`
	Topic                string       `json:"topic"`
	Recipient            *User        `json:"recipient"`
	LastMessageID        string       `json:"last_message_id"`
	Bitrate              int          `json:"bitrate"`
	UserLimit            int          `json:"user_limit"`
}

// Message represents a message sent in a channel within Discord.
// The author object follows the structure of the user object, but
// is only a valid user in the case where the message is generated
// by a user or bot user. If the message is generated by a webhook,
// the author object corresponds to the webhook's id, username, and avatar.
// You can tell if a message is generated by a webhook by checking for the
// webhook_id on the message object.
type Message struct {
	ID              string        `json:"id"`
	ChannelID       string        `json:"channel_id"`
	Author          *User         `json:"author"`
	Content         string        `json:"content"`
	Timestamp       *time.Time    `json:"timestamp"`
	EditedTimestamp *time.Time    `json:"edited_timestamp"`
	TTS             bool          `json:"tts"`
	MentionEveryone bool          `json:"mention_everyone"`
	Mentions        []*User       `json:"mentions"`
	MentionRoles    []string      `json:"mention_roles"`
	Attachments     []*Attachment `json:"attachments"`
	Embeds          []*Embed      `json:"embeds"`
	Reactions       []*Reaction   `json:"reactions"`
	Nonce           string        `json:"nonce"`
	Pinned          bool          `json:"pinned"`
	WebhookID       string        `json:"webhook_id"`
}

type File struct {
	Name    string
	Content io.Reader
}

type Reaction struct {
	Count int
	Me    bool
	Emoji *ReactionEmoji
}

type ReactionEmoji struct {
	ID   *string // nullable
	Name string
}

type Overwrite struct {
	ID    string
	Type  string
	Allow int
	Deny  int
}

type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   *time.Time      `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Attachment struct {
	ID       string `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	Size     int    `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

// TODO create mention method on User, Channel, Role and Custom ReactionEmoji structs
// https://discordapp.com/developers/docs/resources/channel#message-formatting

const endpointChannels = "channels"

func endpointChannel(cID string) string {
	return path.Join(endpointChannels, cID)
}

func (c *Client) GetChannel(cID string) (ch *Channel, err error) {
	endpoint := endpointChannel(cID)
	req, err := c.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	body, err := c.do(req, endpoint, 0)
	if err != nil {
		return nil, err
	}
	return ch, json.Unmarshal(body, &ch)
}

type ParamsModifyChannel struct {
	Name      string `json:"name,omitempty"`
	Position  int    `json:"position,omitempty"`
	Topic     string `json:"topic,omitempty"`
	Bitrate   int    `json:"bitrate,omitempty"`
	UserLimit int    `json:"user_limit,omitempty"`
}

func (c *Client) ModifyChannel(cID string, chmp *ParamsModifyChannel) error {
	endpoint := endpointChannel(cID)
	req, err := c.newRequestJSON("PATCH", endpoint, chmp)
	if err != nil {
		return err
	}
	_, err = c.do(req, endpoint, 0)
	return err
}

func (c *Client) DeleteChannel(cID string) (ch *Channel, err error) {
	endpoint := endpointChannel(cID)
	req, err := c.newRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}
	body, err := c.do(req, endpoint, 0)
	if err != nil {
		return nil, err
	}
	return ch, json.Unmarshal(body, &ch)
}

func endpointChannelMessages(cID string) string {
	return path.Join(endpointChannel(cID), "messages")
}

type ParamsGetChannelMessages struct {
	AroundID string
	BeforeID string
	AfterID  string
	Limit    int
}

func (gcmsp *ParamsGetChannelMessages) RawQuery() string {
	v := make(url.Values)
	if gcmsp.AroundID != "" {
		v.Set("around", gcmsp.AroundID)
	}
	if gcmsp.BeforeID != "" {
		v.Set("before", gcmsp.BeforeID)
	}
	if gcmsp.AfterID != "" {
		v.Set("after", gcmsp.AfterID)
	}
	if gcmsp.Limit > 0 {
		v.Set("limit", strconv.Itoa(gcmsp.Limit))
	}
	return v.Encode()
}

func (c *Client) GetChannelMessages(cID string, pgcms *ParamsGetChannelMessages) (msgs []*Message, err error) {
	endpoint := endpointChannelMessages(cID)
	req, err := c.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	if pgcms != nil {
		req.URL.RawQuery = pgcms.RawQuery()
	}

	body, err := c.do(req, endpoint, 0)
	if err != nil {
		return nil, err
	}
	return msgs, json.Unmarshal(body, &msgs)
}

func endpointChannelMessage(cID, mID string) string {
	return path.Join(endpointChannelMessages(cID), mID)
}

func (c *Client) GetChannelMessage(cID, mID string) (m *Message, err error) {
	endpoint := endpointChannelMessage(cID, mID)
	req, err := c.newRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	rateLimitPath := endpointChannelMessage(cID, "{id}")
	body, err := c.do(req, rateLimitPath, 0)
	if err != nil {
		return nil, err
	}
	return m, json.Unmarshal(body, &m)
}

type ParamsCreateMessage struct {
	Content string `json:"content,omitempty"`
	Nonce   string `json:"nonce,omitempty"`
	TTS     bool   `json:"tts,omitempty"`
	File    *File  `json:"-"`
	Embed   *Embed `json:"embed,omitempty"`
}

func (c *Client) CreateMessage(cID string, cmp *ParamsCreateMessage) (m *Message, err error) {
	reqBody := &bytes.Buffer{}
	reqBodyWriter := multipart.NewWriter(reqBody)

	payloadJSON, err := json.Marshal(cmp)
	if err != nil {
		return nil, err
	}
	w, err := reqBodyWriter.CreateFormField("payload_json")
	if err != nil {
		return nil, err
	}
	_, err = w.Write(payloadJSON)
	if err != nil {
		return nil, err
	}

	if cmp.File != nil {
		w, err := reqBodyWriter.CreateFormFile("file", cmp.File.Name)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(w, cmp.File.Content)
		if err != nil {
			return nil, err
		}
	}

	err = reqBodyWriter.Close()
	if err != nil {
		return nil, err
	}

	endpoint := endpointChannelMessages(cID)
	req, err := c.newRequest("POST", endpoint, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", reqBodyWriter.FormDataContentType())

	body, err := c.do(req, endpoint, 0)
	if err != nil {
		return nil, err
	}
	return m, json.Unmarshal(body, &m)
}

func endpointReaction(cID, mID, emoji, uID string) string {
	return path.Join(endpointChannelMessage(cID, mID), "reactions", emoji, uID)
}

func (c *Client) CreateReaction(cID, mID, emoji string) error {
	endpoint := endpointReaction(cID, mID, emoji, "@me")
	req, err := c.newRequest("PUT", endpoint, nil)
	if err != nil {
		return err
	}
	rateLimitPath := endpointReaction(cID, "{id}", "{id}", "{id}")
	_, err = c.do(req, rateLimitPath, 0)
	return err
}

// uID = "@me" for your own reaction
func (c *Client) DeleteReaction(cID, mID, emoji, uID string) error {
	endpoint := endpointReaction(cID, mID, emoji, uID)
	req, err := c.newRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	rateLimitPath := endpointReaction(cID, "{id}", "{id}", "{id}")
	_, err = c.do(req, rateLimitPath, 0)
	return err
}

// uID = "@me" for your own reaction
func (c *Client) GetReactions(cID, mID, emoji string) error {
	endpoint := endpointReaction(cID, mID, emoji, "TODO")
	req, err := c.newRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	rateLimitPath := endpointReaction(cID, "{id}", "{id}", "{id}")
	_, err = c.do(req, rateLimitPath, 0)
	return err
}

