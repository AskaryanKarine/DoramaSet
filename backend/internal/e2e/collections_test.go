//go:build e2e

package e2e

import (
	"DoramaSet/internal/handler/apiserver/DTO"
	"context"
	"github.com/ozontech/cute"
	"github.com/ozontech/cute/asserts/headers"
	"github.com/ozontech/cute/asserts/json"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test_Collection(t *testing.T) {
	tokens := make(map[string][]string)

	baseURI := "http://localhost:8080"
	cute.NewTestBuilder().
		Title("Просмотр чужих списков и добавление в избранное").
		Tags("collection").
		CreateStep("Register sub user").
		RequestBuilder(
			cute.WithURI(baseURI+"/auth/registration"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(DTO.Auth{
				Login:    "user1",
				Password: "user123456789",
				Email:    "user1@mail.com",
			}),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertHeaders(
			headers.Present("Set-Cookie"),
		).
		After(func(response *http.Response, errors []error) error {
			res := response.Header["Set-Cookie"]
			// res = strings.Split(res[0], ";")
			tokens["user1"] = res
			return nil
		}).
		NextTest().
		CreateStep("Create new public list").
		BeforeExecute(func(request *http.Request) error {
			log.Println(tokens["user1"])
			return nil
		}).
		RequestBuilder(
			cute.WithHeaders(map[string][]string{
				"Cookie": tokens["user1"],
			}),
			cute.WithURI(baseURI+"/list/"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(DTO.List{
				Name:        "test public collection",
				Description: "test public collection",
				Type:        "public",
			}),
		).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Present("data"),
			json.Equal("$.name", "test public collection"),
			json.Equal("$.description", "test public collection"),
			json.Equal("$.creator_name", "user1"),
		).
		NextTest().
		CreateStep("Add dorama in collection").
		RequestBuilder(
			cute.WithHeaders(map[string][]string{
				"Cookie": tokens["user1"],
			}),
			cute.WithURI(baseURI+"/list/1"),
			cute.WithMethod(http.MethodPost),
			cute.WithMarshalBody(DTO.Id{
				Id: 1,
			}),
		).
		ExpectStatus(http.StatusOK).
		ExecuteTest(context.Background(), t)
}
