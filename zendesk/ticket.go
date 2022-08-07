package zendesk

import (
	"encoding/json"
	"net/url"
	"time"
)

// CreateTicketBody is to structure the body data
type CreateTicketBody struct {
	Ticket CreateTicketBodyTicket `json:"ticket"`
}

type CreateTicketBodyTicket struct {
	Comment  CreateTicketBodyComment `json:"comment"`
	Priority string                  `json:"priority"`
	Subject  string                  `json:"subject"`
}

type CreateTicketBodyComment struct {
	Body string `json:"body"`
}

// CreateTicketReturn is to decode the json data
type CreateTicketReturn struct {
	Ticket struct {
		Url        string      `json:"url"`
		Id         int         `json:"id"`
		ExternalId interface{} `json:"external_id"`
		Via        struct {
			Channel string `json:"channel"`
			Source  struct {
				From struct {
				} `json:"from"`
				To struct {
				} `json:"to"`
				Rel interface{} `json:"rel"`
			} `json:"source"`
		} `json:"via"`
		CreatedAt           time.Time     `json:"created_at"`
		UpdatedAt           time.Time     `json:"updated_at"`
		Type                interface{}   `json:"type"`
		Subject             string        `json:"subject"`
		RawSubject          string        `json:"raw_subject"`
		Description         string        `json:"description"`
		Priority            string        `json:"priority"`
		Status              string        `json:"status"`
		Recipient           interface{}   `json:"recipient"`
		RequesterId         int64         `json:"requester_id"`
		SubmitterId         int64         `json:"submitter_id"`
		AssigneeId          interface{}   `json:"assignee_id"`
		OrganizationId      int64         `json:"organization_id"`
		GroupId             int64         `json:"group_id"`
		CollaboratorIds     []interface{} `json:"collaborator_ids"`
		FollowerIds         []interface{} `json:"follower_ids"`
		EmailCcIds          []interface{} `json:"email_cc_ids"`
		ForumTopicId        interface{}   `json:"forum_topic_id"`
		ProblemId           interface{}   `json:"problem_id"`
		HasIncidents        bool          `json:"has_incidents"`
		IsPublic            bool          `json:"is_public"`
		DueAt               interface{}   `json:"due_at"`
		Tags                []interface{} `json:"tags"`
		CustomFields        []interface{} `json:"custom_fields"`
		SatisfactionRating  interface{}   `json:"satisfaction_rating"`
		SharingAgreementIds []interface{} `json:"sharing_agreement_ids"`
		Fields              []interface{} `json:"fields"`
		FollowupIds         []interface{} `json:"followup_ids"`
		BrandId             int64         `json:"brand_id"`
		AllowChannelback    bool          `json:"allow_channelback"`
		AllowAttachments    bool          `json:"allow_attachments"`
	} `json:"ticket"`
	Audit struct {
		Id        int64     `json:"id"`
		TicketId  int       `json:"ticket_id"`
		CreatedAt time.Time `json:"created_at"`
		AuthorId  int64     `json:"author_id"`
		Metadata  struct {
			System struct {
				Client    string  `json:"client"`
				IpAddress string  `json:"ip_address"`
				Location  string  `json:"location"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"system"`
			Custom struct {
			} `json:"custom"`
		} `json:"metadata"`
		Events []struct {
			Id          int64         `json:"id"`
			Type        string        `json:"type"`
			AuthorId    int64         `json:"author_id,omitempty"`
			Body        string        `json:"body,omitempty"`
			HtmlBody    string        `json:"html_body,omitempty"`
			PlainBody   string        `json:"plain_body,omitempty"`
			Public      bool          `json:"public,omitempty"`
			Attachments []interface{} `json:"attachments,omitempty"`
			AuditId     int64         `json:"audit_id,omitempty"`
			Value       *string       `json:"value,omitempty"`
			FieldName   string        `json:"field_name,omitempty"`
			Via         struct {
				Channel string `json:"channel"`
				Source  struct {
					From struct {
						Deleted bool   `json:"deleted"`
						Title   string `json:"title"`
						Id      int64  `json:"id"`
					} `json:"from"`
					Rel string `json:"rel"`
				} `json:"source"`
			} `json:"via,omitempty"`
			Subject    string  `json:"subject,omitempty"`
			Recipients []int64 `json:"recipients,omitempty"`
		} `json:"events"`
		Via struct {
			Channel string `json:"channel"`
			Source  struct {
				From struct {
				} `json:"from"`
				To struct {
					Name    string `json:"name"`
					Address string `json:"address"`
				} `json:"to"`
				Rel interface{} `json:"rel"`
			} `json:"source"`
		} `json:"via"`
	} `json:"audit"`
}

// CreateTicket is to create a ticket in zendesk
func CreateTicket(body CreateTicketBody, r Request) (CreateTicketReturn, error) {

	address, err := url.JoinPath(r.BaseUrl, "api", "v2", "tickets")
	if err != nil {
		return CreateTicketReturn{}, err
	}

	convert, err := json.Marshal(body)
	if err != nil {
		return CreateTicketReturn{}, err
	}

	c := Config{
		Url:    address,
		Method: "POST",
		Body:   convert,
	}

	response, err := c.Send(r)
	if err != nil {
		return CreateTicketReturn{}, err
	}

	defer response.Body.Close()

	var decode CreateTicketReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateTicketReturn{}, err
	}

	return decode, nil

}
