// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/clone-global/cloneforce-go"
	"github.com/clone-global/cloneforce-go/internal/testutil"
	"github.com/clone-global/cloneforce-go/option"
)

func TestV1CloneSkillNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Clones.Skills.New(
		context.TODO(),
		"cloneId",
		cloneforce.V1CloneSkillNewParams{
			SkillID:  "skillId",
			IsActive: cloneforce.Bool(true),
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

func TestV1CloneSkillUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Clones.Skills.Update(
		context.TODO(),
		"skillName",
		cloneforce.V1CloneSkillUpdateParams{
			CloneID:  "cloneId",
			IsActive: cloneforce.Bool(true),
			Settings: []cloneforce.V1CloneSkillUpdateParamsSetting{{
				Name:         "name",
				ConnectionID: cloneforce.String("connectionId"),
				Value:        cloneforce.String("value"),
			}},
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

func TestV1CloneSkillListWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Clones.Skills.List(
		context.TODO(),
		"cloneId",
		cloneforce.V1CloneSkillListParams{
			Active: cloneforce.V1CloneSkillListParamsActiveTrue,
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

func TestV1CloneSkillDelete(t *testing.T) {
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
	_, err := client.V1.Clones.Skills.Delete(
		context.TODO(),
		"skillName",
		cloneforce.V1CloneSkillDeleteParams{
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
