package user

import (
	"encoding/json"
	"testing"
)

func TestJsonUnmarshalFullUser(t *testing.T) {
	testName := "Case #1. All values are set"
	user := &User{}
	data :=
		`{
			"id": 1,
			"is_bot": false,
			"first_name": "Test",
			"last_name": "Pest",
			"username": "testpest",
			"language_code": "en",
			"is_premium": true,
			"added_to_attachment_menu": true,
			"can_join_groups": false,
			"can_read_all_group_messages": false,
			"supports_inline_queries": true
		 }`

	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		t.Fatalf("%v. Can't unmarshal json. %s.", testName, err.Error())
	}

	if user.ID != 1 {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, 1, user.ID)
	}

	if user.IsBot != false {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, false, user.IsBot)
	}

	if user.FirstName != "Test" {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, "Test", user.FirstName)
	}

	if *user.LastName != "Pest" {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, "Pest", *user.LastName)
	}

	if *user.Username != "testpest" {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, "testpest", *user.Username)
	}

	if *user.LanguageCode != "en" {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, "en", *user.LanguageCode)
	}

	if *user.IsPremium != true {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, true, *user.IsPremium)
	}

	if *user.AddedToAttachmentMenu != true {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, true, *user.AddedToAttachmentMenu)
	}

	if *user.CanJoinGroups != false {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, false, *user.CanJoinGroups)
	}

	if *user.CanReadAllGroupMessages != false {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, false, *user.CanReadAllGroupMessages)
	}

	if *user.SupportsInlineQueries != true {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, true, *user.SupportsInlineQueries)
	}
}

func TestJsonUnmarshalPartialUser(t *testing.T) {
	testName := "Case #2. Partial values are set"
	user := &User{}
	data :=
		`{
			"id": 1,
			"is_bot": false,
			"first_name": "Test",
			"last_name": "Pest",
			"username": "testpest",
			"added_to_attachment_menu": true,
			"can_join_groups": false,
			"supports_inline_queries": true
		 }`

	err := json.Unmarshal([]byte(data), user)
	if err != nil {
		t.Fatalf("%v. Can't unmarshal json. %s.", testName, err.Error())
	}

	if user.ID != 1 {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, 1, user.ID)
	}

	if user.IsBot != false {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, false, user.IsBot)
	}

	if user.FirstName != "Test" {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, "Test", user.FirstName)
	}

	if *user.LastName != "Pest" {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, "Pest", *user.LastName)
	}

	if *user.Username != "testpest" {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, "testpest", *user.Username)
	}

	if user.LanguageCode != nil {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, nil, user.LanguageCode)
	}

	if user.IsPremium != nil {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, nil, user.IsPremium)
	}

	if *user.AddedToAttachmentMenu != true {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, true, *user.AddedToAttachmentMenu)
	}

	if *user.CanJoinGroups != false {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, false, *user.CanJoinGroups)
	}

	if user.CanReadAllGroupMessages != nil {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, nil, user.CanReadAllGroupMessages)
	}

	if *user.SupportsInlineQueries != true {
		t.Fatalf("%v. Expected: %v. Actual %v.", testName, true, *user.SupportsInlineQueries)
	}
}
