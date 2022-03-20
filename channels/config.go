package channels

// Type is to save the channel types
type Type struct {
	Gotify   bool
	Matrix   bool
	Telegram bool
	Discord  bool
	Trello   bool
	Zendesk  bool
	OsTicket bool
	SMTP     bool
	Webhook  bool
}
