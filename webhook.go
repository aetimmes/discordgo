package discordgo

// Webhook stores the data for a webhook.
type Webhook struct {
	ID        string      `json:"id"`
	Type      WebhookType `json:"type"`
	GuildID   string      `json:"guild_id"`
	ChannelID string      `json:"channel_id"`
	User      *User       `json:"user"`
	Name      string      `json:"name"`
	Avatar    string      `json:"avatar"`
	Token     string      `json:"token"`

	// ApplicationID is the bot/OAuth2 application that created this webhook
	ApplicationID string `json:"application_id,omitempty"`

	SourceGuild   *Guild   `json:"source_guild"`
	SourceChannel *Channel `json:"source_channel"`
	URL           string   `json:"url"`
}

// WebhookType is the type of Webhook (see WebhookType* consts) in the Webhook struct
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
type WebhookType int

// Valid WebhookType values
const (
	WebhookTypeIncoming        WebhookType = 1
	WebhookTypeChannelFollower WebhookType = 2
	WebhookTypeApplication     WebhookType = 3
)

// WebhookParams is a struct for webhook params, used in the WebhookExecute command.
type WebhookParams struct {
	Content         string                  `json:"content,omitempty"`
	Username        string                  `json:"username,omitempty"`
	AvatarURL       string                  `json:"avatar_url,omitempty"`
	TTS             bool                    `json:"tts,omitempty"`
	Files           []*File                 `json:"-"`
	Components      []MessageComponent      `json:"components"`
	Embeds          []*MessageEmbed         `json:"embeds,omitempty"`
	AllowedMentions *MessageAllowedMentions `json:"allowed_mentions,omitempty"`
	// NOTE: Works only for followup messages.
	Flags    uint64 `json:"flags,omitempty"`
	ThreadID string `json:"thread_id,omitempty"`
}

// WebhookEdit stores data for editing of a webhook message.
type WebhookEdit struct {
	Content         string                  `json:"content,omitempty"`
	Components      []MessageComponent      `json:"components"`
	Embeds          []*MessageEmbed         `json:"embeds,omitempty"`
	Files           []*File                 `json:"-"`
	AllowedMentions *MessageAllowedMentions `json:"allowed_mentions,omitempty"`
}
