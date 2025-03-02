package go_hubspot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func getMockCRMAPI(mockClient *IHTTPClientMock) HubspotCRMAPI {
	return HubspotCRMAPI{
		APIKey:     "api_key",
		httpClient: mockClient,
	}
}

var singleObjectResponse []byte = []byte(`
	{
		"total":1,
		"results":[
			{
				"id":"123id",
				"properties":{
					"company_number":"11762819"
				},
				"createdAt":"anyDate",
				"updatedAt":"anyDate",
				"archived":false
			}
		]
	}
`)

var noObjectResponse []byte = []byte(`
	{}
`)

var multipleObjectResponse []byte = []byte(`
	{
		"total":2,
		"results":[
			{
				"id":"123id",
				"properties":{
					"company_number":"11762819"
				},
				"createdAt":"anyDate",
				"updatedAt":"anyDate",
				"archived":false
			},
			{
				"id":"456id",
				"properties":{
					"company_number":"87654321"
				},
				"createdAt":"anyDate",
				"updatedAt":"anyDate",
				"archived":false
			}
		]
	}
`)

func generateMock(t *testing.T, response []byte) IHTTPClientMock {
	return IHTTPClientMock{
		DoFunc: func(req *http.Request) (resp *http.Response, err error) {

			expectedBody := hubSpotSearchRequest{
				FilterGroups: []filterGroup{
					{
						Filters: []filter{
							{
								Value:        "example-application-id",
								PropertyName: "application_id",
								Operator:     "EQ",
							},
						},
					},
				},
				Properties: []string{
					"id",
					"company_number",
				},
			}

			url := fmt.Sprintf("%s", req.URL)

			var body []byte
			body, err = ioutil.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}

			var bodyStruct hubSpotSearchRequest
			err = json.Unmarshal(body, &bodyStruct)
			if err != nil {
				t.Errorf("Error unmarshalling hubspot search request: %s", err.Error())
			}

			if !reflect.DeepEqual(bodyStruct, expectedBody) {
				t.Errorf("Incorrect body, expected\n%s\ngot:\n%s", bodyStruct, expectedBody)
			}

			w := httptest.NewRecorder()
			expectedUrl := "https://api.hubapi.com/crm/v3/objects/objecttype/search"
			if url == expectedUrl {
				if req.Method != "POST" {
					t.Errorf("A request was made that should have been POST but was %s", req.Method)
				}

				w.WriteHeader(200)
				w.Write(response)
			} else {
				t.Errorf("Unexpected url, expected:\n%s\ngot:\n%s", expectedUrl, url)
			}
			return w.Result(), nil
		},
	}
}

// searchForApplicationIdTest runs a test on SearchHubSpot
// with a given response from the api and an expected result
func searchForApplicationIdTest(t *testing.T, response []byte, expectedIDs []HubSpotSearchResult) {
	mockHubSpotHTTPClient := generateMock(t, response)

	api := getMockCRMAPI(&mockHubSpotHTTPClient)
	gotResult, err := api.SearchHubSpot(
		"objecttype",
		map[string]string{
			"application_id": "example-application-id",
		},
		[]string{
			"id",
			"company_number",
		},
	)
	if err != nil {
		t.Errorf("SearchHubSpot failed; %s", err.Error())
	}

	if len(gotResult) != len(expectedIDs) {
		t.Errorf("Incorrect number of returned results, expected %d, got %d", len(expectedIDs), len(gotResult))
	}

	if len(mockHubSpotHTTPClient.DoCalls()) != 1 {
		t.Errorf("Expected 1 call to HubSpot API")
	}

	for i := range gotResult {
		if gotResult[i].Id != expectedIDs[i].Id {
			t.Errorf("SearchHubSpot returned incorrect IDs, expected:\n%s\ngot:\n%s\n", expectedIDs, gotResult)
		}
	}
}

func TestSearchHubSpot(t *testing.T) {
	searchForApplicationIdTest(t, noObjectResponse, []HubSpotSearchResult{})
	searchForApplicationIdTest(t, singleObjectResponse, []HubSpotSearchResult{
		{
			Id:           "123id",
			Properties:   map[string]string{"company_number": "11762819"},
			Associations: map[string]Associations{},
		},
	})
	searchForApplicationIdTest(t, multipleObjectResponse, []HubSpotSearchResult{
		{
			Id:           "123id",
			Properties:   map[string]string{"company_number": "11762819"},
			Associations: map[string]Associations{},
		},
		{
			Id:           "456id",
			Properties:   map[string]string{"company_number": "87654321"},
			Associations: map[string]Associations{},
		},
	})
}

func createCompanyForContactMock(t *testing.T, numberOfResults int) IHTTPClientMock {
	return IHTTPClientMock{
		DoFunc: func(req *http.Request) (resp *http.Response, err error) {

			url := fmt.Sprintf("%s", req.URL)

			w := httptest.NewRecorder()
			expectedUrl := "https://api.hubapi.com/crm/v3/objects/contacts/contactid/associations/company"
			if url == expectedUrl {

				if req.Method != "GET" {
					t.Errorf("A request was made that should have been GET but was %s", req.Method)
				}

				w.WriteHeader(200)

				assocs := make([]Association, numberOfResults)
				for i := 0; i < numberOfResults; i++ {
					assocs[i] = Association{fmt.Sprintf("id%d", i), "type"}
				}

				response := Associations{assocs}

				respBody, err := json.Marshal(response)
				if err != nil {
					t.Errorf("Failed to marshal respBody for company mock")
				}

				w.Write(respBody)
			} else {
				t.Errorf("Unexpected url, expected:\n%s\ngot:\n%s", expectedUrl, url)
			}
			return w.Result(), nil
		},
	}
}

func TestGetCompanyForContact(t *testing.T) {
	// Create an api that returns 1 company for "contactid"
	companyMock := createCompanyForContactMock(t, 1)
	api := getMockCRMAPI(&companyMock)

	companyId, err := api.GetCompanyForContact("contactid")
	if err != nil {
		t.Errorf("GetCompanyForContact returned an unexpected error: %s", err.Error())
	}

	if companyId != "id0" {
		t.Errorf("GetCompanyForContact returned incorrect id, expected:id0 got:%s", companyId)
	}

	if len(companyMock.DoCalls()) != 1 {
		t.Errorf("Expected 1 call to HubSpot API")
	}

	// Create an api that returns no companies for "contactid"
	companyMock = createCompanyForContactMock(t, 0)
	api = getMockCRMAPI(&companyMock)

	companyId, err = api.GetCompanyForContact("contactid")
	if err != nil {
		t.Errorf("GetCompanyForContact returned an unexpected error: %s", err.Error())
	}

	if companyId != "" {
		t.Errorf(`GetCompanyForContact did not return "" when it found no associated company`)
	}

	if len(companyMock.DoCalls()) != 1 {
		t.Errorf("Expected 1 call to HubSpot API")
	}

	// Create an api that returns multiple companies for "contactid"
	companyMock = createCompanyForContactMock(t, 2)
	api = getMockCRMAPI(&companyMock)

	companyId, err = api.GetCompanyForContact("contactid")

	expectedError := "There are multiple companies associated with contact 'contactid' there should only be one"
	if err.Error() != expectedError {
		t.Errorf(`GetCompanyForContact did not fail with correct error for multiple retrieved companies, expected:
			%s
			got:
			%s`, expectedError, err.Error())
	}

	if companyId != "" {
		t.Errorf(`GetCompanyForContact did not return "" when an error was thrown`)
	}

	if len(companyMock.DoCalls()) != 1 {
		t.Errorf("Expected 1 call to HubSpot API")
	}
}

func createDealForCompanyMock(t *testing.T, numberOfResults int) IHTTPClientMock {
	return IHTTPClientMock{
		DoFunc: func(req *http.Request) (resp *http.Response, err error) {

			url := fmt.Sprintf("%s", req.URL)

			w := httptest.NewRecorder()
			expectedUrl := "https://api.hubapi.com/crm/v3/objects/companies/companyid/associations/deal?limit=500"
			if url == expectedUrl {

				if req.Method != "GET" {
					t.Errorf("A request was made that should have been GET but was %s", req.Method)
				}

				w.WriteHeader(200)

				assocs := make([]Association, numberOfResults)
				for i := 0; i < numberOfResults; i++ {
					assocs[i] = Association{fmt.Sprintf("id%d", i), "type"}
				}

				response := Associations{assocs}

				respBody, err := json.Marshal(response)
				if err != nil {
					t.Errorf("Failed to marshal respBody for company mock")
				}

				w.Write(respBody)
			} else {
				t.Errorf("Unexpected url, expected:\n%s\ngot:\n%s", expectedUrl, url)
			}
			return w.Result(), nil
		},
	}
}

func TestGetDealForCompany(t *testing.T) {
	// Create an api that returns 1 deal for "companyid"
	dealMock := createDealForCompanyMock(t, 1)
	api := getMockCRMAPI(&dealMock)

	dealId, err := api.GetDealForCompany("companyid")
	if err != nil {
		t.Errorf("GetDealForCompany returned an unexpected error: %s", err.Error())
	}

	if dealId != "id0" {
		t.Errorf("GetDealForCompany returned incorrect id, expected:id0 got:%s", dealId)
	}

	if len(dealMock.DoCalls()) != 1 {
		t.Errorf("Expected 1 call to HubSpot API")
	}

	// Create an api that returns no companies for "companyid"
	dealMock = createDealForCompanyMock(t, 0)
	api = getMockCRMAPI(&dealMock)

	dealId, err = api.GetDealForCompany("companyid")
	if err != nil {
		t.Errorf("GetDealForCompany returned an unexpected error: %s", err.Error())
	}

	if dealId != "" {
		t.Errorf(`GetDealForCompany did not return "" when it found no associated deal`)
	}

	if len(dealMock.DoCalls()) != 1 {
		t.Errorf("Expected 1 call to HubSpot API")
	}

	// Create an api that returns multiple companies for "companyid"
	dealMock = createDealForCompanyMock(t, 2)
	api = getMockCRMAPI(&dealMock)

	dealId, err = api.GetDealForCompany("companyid")

	expectedError := "There are multiple deals associated with company 'companyid' there should only be one"
	if err.Error() != expectedError {
		t.Errorf(`GetDealForCompany did not fail with correct error for multiple retrieved deals, expected:
			%s
			got:
			%s`, expectedError, err.Error())
	}

	if dealId != "" {
		t.Errorf(`GetDealForCompany did not return "" when an error was thrown`)
	}

	if len(dealMock.DoCalls()) != 1 {
		t.Errorf("Expected 1 call to HubSpot API")
	}
}
