// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agutoken_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Boomchainlab/solana-verifiable-build"
	"github.com/Boomchainlab/solana-verifiable-build/internal/testutil"
	"github.com/Boomchainlab/solana-verifiable-build/option"
	"github.com/Boomchainlab/solana-verifiable-build/shared"
)

func TestStoreOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Store.Order.New(context.TODO(), agutoken.StoreOrderNewParams{
		Order: shared.OrderParam{
			ID:       agutoken.Int(10),
			Complete: agutoken.Bool(true),
			PetID:    agutoken.Int(198772),
			Quantity: agutoken.Int(7),
			ShipDate: agutoken.Time(time.Now()),
			Status:   shared.OrderStatusApproved,
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

func TestStoreOrderGet(t *testing.T) {
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
	_, err := client.Store.Order.Get(context.TODO(), 0)
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreOrderDelete(t *testing.T) {
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
	err := client.Store.Order.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
