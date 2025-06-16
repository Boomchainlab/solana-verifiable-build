// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agutoken_test

import (
	"context"
	"os"
	"testing"

	"github.com/Boomchainlab/solana-verifiable-build"
	"github.com/Boomchainlab/solana-verifiable-build/internal/testutil"
	"github.com/Boomchainlab/solana-verifiable-build/option"
)

func TestUsage(t *testing.T) {
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
	order, err := client.Store.Order.New(context.TODO(), agutoken.StoreOrderNewParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", order.ID)
}
