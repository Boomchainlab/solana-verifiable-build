// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package agutoken_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Boomchainlab/solana-verifiable-build"
	"github.com/Boomchainlab/solana-verifiable-build/internal/testutil"
	"github.com/Boomchainlab/solana-verifiable-build/option"
)

func TestStoreListInventory(t *testing.T) {
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
	_, err := client.Store.ListInventory(context.TODO())
	if err != nil {
		var apierr *agutoken.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
