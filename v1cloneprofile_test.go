// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/clone-global/cloneforce-go"
	"github.com/clone-global/cloneforce-go/internal/testutil"
	"github.com/clone-global/cloneforce-go/option"
)

func TestV1CloneProfileGet(t *testing.T) {
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
	_, err := client.V1.Clones.Profile.Get(context.TODO(), "cloneId")
	if err != nil {
		var apierr *cloneforce.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1CloneProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Clones.Profile.Update(
		context.TODO(),
		"cloneId",
		cloneforce.V1CloneProfileUpdateParams{
			AppearanceDesc:     cloneforce.String("appearanceDesc"),
			Birthdate:          cloneforce.Time(time.Now()),
			CareerDesc:         cloneforce.String("careerDesc"),
			ChildhoodDesc:      cloneforce.String("childhoodDesc"),
			CurrentCity:        cloneforce.String("currentCity"),
			EducationDesc:      cloneforce.String("educationDesc"),
			FamilyDesc:         cloneforce.String("familyDesc"),
			FinancesDesc:       cloneforce.String("financesDesc"),
			Gender:             cloneforce.String("gender"),
			Hometown:           cloneforce.String("hometown"),
			InterestsDesc:      cloneforce.String("interestsDesc"),
			IsEnabled:          cloneforce.Bool(true),
			Language:           cloneforce.String("language"),
			LifestyleDesc:      cloneforce.String("lifestyleDesc"),
			MaritalStatus:      cloneforce.String("maritalStatus"),
			Name:               cloneforce.String("name"),
			Nationality:        cloneforce.String("nationality"),
			Occupation:         cloneforce.String("occupation"),
			PrinciplesDesc:     cloneforce.String("principlesDesc"),
			Races:              []string{"string"},
			ScreenName:         cloneforce.String("screenName"),
			SocialCircleDesc:   cloneforce.String("socialCircleDesc"),
			SocialHobbiesDesc:  cloneforce.String("socialHobbiesDesc"),
			SportsDesc:         cloneforce.String("sportsDesc"),
			TravelDesc:         cloneforce.String("travelDesc"),
			UnusualHobbiesDesc: cloneforce.String("unusualHobbiesDesc"),
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
