// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agutoken_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/agu-token-go"
	"github.com/stainless-sdks/agu-token-go/internal/testutil"
	"github.com/stainless-sdks/agu-token-go/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agutoken.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.New(context.TODO(), agutoken.UserNewParams{
		User: agutoken.UserParam{
			ID:         agutoken.Int(10),
			Email:      agutoken.String("john@email.com"),
			FirstName:  agutoken.String("John"),
			LastName:   agutoken.String("James"),
			Password:   agutoken.String("12345"),
			Phone:      agutoken.String("12345"),
			Username:   agutoken.String("theUser"),
			UserStatus: agutoken.Int(1),
		},
	})
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agutoken.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.Get(context.TODO(), "username")
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agutoken.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.User.Update(
		context.TODO(),
		"username",
		agutoken.UserUpdateParams{
			User: agutoken.UserParam{
				ID:         agutoken.Int(10),
				Email:      agutoken.String("john@email.com"),
				FirstName:  agutoken.String("John"),
				LastName:   agutoken.String("James"),
				Password:   agutoken.String("12345"),
				Phone:      agutoken.String("12345"),
				Username:   agutoken.String("theUser"),
				UserStatus: agutoken.Int(1),
			},
		},
	)
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agutoken.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.User.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agutoken.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.NewWithList(context.TODO(), agutoken.UserNewWithListParams{
		Items: []agutoken.UserParam{{
			ID:         agutoken.Int(10),
			Email:      agutoken.String("john@email.com"),
			FirstName:  agutoken.String("John"),
			LastName:   agutoken.String("James"),
			Password:   agutoken.String("12345"),
			Phone:      agutoken.String("12345"),
			Username:   agutoken.String("theUser"),
			UserStatus: agutoken.Int(1),
		}},
	})
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agutoken.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.Login(context.TODO(), agutoken.UserLoginParams{
		Password: agutoken.String("password"),
		Username: agutoken.String("username"),
	})
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := agutoken.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.User.Logout(context.TODO())
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
