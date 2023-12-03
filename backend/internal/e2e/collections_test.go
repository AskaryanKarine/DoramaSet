//go:build e2e

package e2e

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"context"
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/headers"
	"github.com/ozontech/cute/asserts/json"
	"net/http"
	"testing"
)

type AsserBody func(body []byte) error

type responseData struct {
	Data []DTO.List `json:"data"`
}

func assertPublicLists() AsserBody {
	return func(body []byte) error {
		if len(body) == 0 {
			return errors.New("response is empty")
		}

		var d responseData
		err := json2.Unmarshal(body, &d)
		if err != nil {
			return fmt.Errorf("unmarshal: %w", err)
		}
		if len(d.Data) != 1 {
			return errors.New("incorrect items in page")
		}

		list := d.Data[0]
		if list.Id != 2 {
			return fmt.Errorf("incorrect id collection, got %d, but expected 2", list.Id)
		}
		if list.Name != "test public collection" {
			return fmt.Errorf("incorrect name collection, got %s, but expected \"test public collection\"", list.Name)
		}
		if list.Description != "test public collection" {
			return fmt.Errorf("incorrect description collection, got %s, but expected \"test public collection\"",
				list.Description)
		}
		if list.Type != "public" {
			return fmt.Errorf("incorrect type collection, got %s, but expected \"public\"", list.Type)
		}
		if list.CreatorName != "user1" {
			return fmt.Errorf("incorrect creatorName collection, got %s, but expected \"user1\"", list.CreatorName)
		}
		if len(list.Doramas) != 1 {
			return errors.New("incorrect content len in collection")
		}
		if list.Doramas[0].Id != 1 {
			return fmt.Errorf("incorrect items in collection, got id elem %d, but expected 1", list.Doramas[0].Id)
		}

		return nil
	}
}

func Test_Collection(t *testing.T) {
	myHeaders := make(map[string][]string)
	baseURI := "http://localhost:8080"
	cute.NewTestBuilder().
		Title("Watching public collection and adding to favourite").
		Tags("collection").
		CreateStep("Registered user1").
		RequestBuilder(
			cute.WithURI(baseURI+"/auth/registration"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(DTO.Auth{
				Login:    "user1",
				Password: "user123456789",
				Email:    "user1@mail.com",
			}),
		).
		ExpectStatus(http.StatusOK).
		AssertHeaders(
			headers.Present("Set-Cookie"),
		).
		After(func(response *http.Response, errors []error) error {
			myHeaders["Cookie"] = response.Header["Set-Cookie"]
			return nil
		}).
		NextTest().
		CreateStep("Create user1's public list").
		RequestBuilder(
			cute.WithURI(baseURI+"/list/"),
			cute.WithMethod(http.MethodPost),
			cute.WithHeaders(myHeaders),
			cute.WithMarshalBody(DTO.List{
				Name:        "test public collection",
				Description: "test public collection",
				Type:        "public",
			}),
		).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Present("$.data"),
			json.Equal("$.data.name", "test public collection"),
			json.Equal("$.data.description", "test public collection"),
			json.Equal("$.data.creator_name", "user1"),
		).
		NextTest().
		CreateStep("Add dorama in user1's collection").
		RequestBuilder(
			cute.WithHeaders(myHeaders),
			cute.WithURI(baseURI+"/list/2"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(DTO.Id{
				Id: 1,
			}),
		).
		ExpectStatus(http.StatusOK).
		NextTest().
		CreateStep("Register user2").
		RequestBuilder(
			cute.WithURI(baseURI+"/auth/registration"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(DTO.Auth{
				Login:    "user2",
				Password: "user123456789",
				Email:    "user2@mail.com",
			}),
		).
		ExpectStatus(http.StatusOK).
		AssertHeaders(
			headers.Present("Set-Cookie"),
		).
		After(func(response *http.Response, errors []error) error {
			res := response.Header["Set-Cookie"]
			myHeaders["Cookie"] = res
			return nil
		}).
		NextTest().
		CreateStep("Watch public collection").
		RequestBuilder(
			cute.WithURI(baseURI+"/list/public"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Present("$.data"),
			cute.AssertBody(assertPublicLists()),
		).
		NextTest().
		CreateStep("Add in favourite user2").
		RequestBuilder(
			cute.WithURI(baseURI+"/user/favorite"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(DTO.Id{Id: 2}),
			cute.WithHeaders(myHeaders),
		).
		ExpectStatus(http.StatusOK).
		NextTest().
		CreateStep("Get favourite").
		RequestBuilder(
			cute.WithURI(baseURI+"/user/favorite"),
			cute.WithMethod(http.MethodGet),
			cute.WithHeaders(myHeaders),
		).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Present("$.data"),
			cute.AssertBody(assertPublicLists()),
		).
		ExecuteTest(context.Background(), t)
}
