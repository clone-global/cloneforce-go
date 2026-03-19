// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/cloneforce-go"
	"github.com/stainless-sdks/cloneforce-go/internal/testutil"
	"github.com/stainless-sdks/cloneforce-go/option"
)

func TestV1CloneFileNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloneforce.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Clones.Files.New(
		context.TODO(),
		"cloneId",
		cloneforce.V1CloneFileNewParams{
			URL:         "url",
			ContentType: cloneforce.String("contentType"),
			Name:        cloneforce.String("name"),
		},
	)
	if err != nil {
		var apierr *cloneforce.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CloneFileGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloneforce.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Clones.Files.Get(
		context.TODO(),
		"fileId",
		cloneforce.V1CloneFileGetParams{
			CloneID: "cloneId",
		},
	)
	if err != nil {
		var apierr *cloneforce.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CloneFileList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloneforce.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Clones.Files.List(context.TODO(), "cloneId")
	if err != nil {
		var apierr *cloneforce.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CloneFileDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloneforce.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Clones.Files.Delete(
		context.TODO(),
		"fileId",
		cloneforce.V1CloneFileDeleteParams{
			CloneID: "cloneId",
		},
	)
	if err != nil {
		var apierr *cloneforce.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
