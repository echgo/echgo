// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package trello is used to define the configuration
// for the request, the functions so that the
// request can be built and processed.
package trello

import (
	"encoding/json"
	"net/url"
	"time"
)

// CreateCardReturn is to decode the json data.
type CreateCardReturn struct {
	Attachments []interface{} `json:"attachments"`
	Id          string        `json:"id"`
	Badges      struct {
		AttachmentsByType struct {
			Trello struct {
				Board int `json:"board"`
				Card  int `json:"card"`
			} `json:"trello"`
		} `json:"attachmentsByType"`
		Location              bool        `json:"location"`
		Votes                 int         `json:"votes"`
		ViewingMemberVoted    bool        `json:"viewingMemberVoted"`
		Subscribed            bool        `json:"subscribed"`
		Fogbugz               string      `json:"fogbugz"`
		CheckItems            int         `json:"checkItems"`
		CheckItemsChecked     int         `json:"checkItemsChecked"`
		CheckItemsEarliestDue interface{} `json:"checkItemsEarliestDue"`
		Comments              int         `json:"comments"`
		Attachments           int         `json:"attachments"`
		Description           bool        `json:"description"`
		Due                   interface{} `json:"due"`
		DueComplete           bool        `json:"dueComplete"`
		Start                 interface{} `json:"start"`
	} `json:"badges"`
	CheckItemStates  []interface{} `json:"checkItemStates"`
	Closed           bool          `json:"closed"`
	DueComplete      bool          `json:"dueComplete"`
	DateLastActivity time.Time     `json:"dateLastActivity"`
	Desc             string        `json:"desc"`
	DescData         struct {
		Emoji struct {
		} `json:"emoji"`
	} `json:"descData"`
	Due                   interface{}   `json:"due"`
	DueReminder           interface{}   `json:"dueReminder"`
	Email                 interface{}   `json:"email"`
	IdBoard               string        `json:"idBoard"`
	IdChecklists          []interface{} `json:"idChecklists"`
	IdList                string        `json:"idList"`
	IdMembers             []interface{} `json:"idMembers"`
	IdMembersVoted        []interface{} `json:"idMembersVoted"`
	IdShort               int           `json:"idShort"`
	IdAttachmentCover     interface{}   `json:"idAttachmentCover"`
	Labels                []interface{} `json:"labels"`
	IdLabels              []interface{} `json:"idLabels"`
	ManualCoverAttachment bool          `json:"manualCoverAttachment"`
	Name                  string        `json:"name"`
	Pos                   int           `json:"pos"`
	ShortLink             string        `json:"shortLink"`
	ShortUrl              string        `json:"shortUrl"`
	Start                 interface{}   `json:"start"`
	Subscribed            bool          `json:"subscribed"`
	Url                   string        `json:"url"`
	Cover                 struct {
		IdAttachment         interface{} `json:"idAttachment"`
		Color                interface{} `json:"color"`
		IdUploadedBackground interface{} `json:"idUploadedBackground"`
		Size                 string      `json:"size"`
		Brightness           string      `json:"brightness"`
		IdPlugin             interface{} `json:"idPlugin"`
	} `json:"cover"`
	IsTemplate bool          `json:"isTemplate"`
	CardRole   interface{}   `json:"cardRole"`
	Stickers   []interface{} `json:"stickers"`
	Limits     struct {
	} `json:"limits"`
}

// CreateCard is to create a card to trello.
func CreateCard(name, desc string, r Request) (CreateCardReturn, error) {

	address, err := url.JoinPath(baseUrl, "1", "cards")
	if err != nil {
		return CreateCardReturn{}, err
	}

	c := Config{
		Url:    address,
		Method: "POST",
	}

	parse, err := url.Parse(c.Url)
	if err != nil {
		return CreateCardReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return CreateCardReturn{}, err
	}

	newUrl.Add("idList", r.IdList)
	newUrl.Add("name", name)
	newUrl.Add("desc", desc)
	newUrl.Add("key", r.Key)
	newUrl.Add("token", r.Token)

	parse.RawQuery = newUrl.Encode()
	c.Url = parse.String()

	response, err := c.Send()
	if err != nil {
		return CreateCardReturn{}, err
	}

	defer response.Body.Close()

	var decode CreateCardReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateCardReturn{}, err
	}

	return decode, nil

}
