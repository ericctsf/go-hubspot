// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package go_hubspot

import (
	"bytes"
	"sync"
)

// Ensure, that IHubspotCRMAPIMock does implement IHubspotCRMAPI.
// If this is not the case, regenerate this file with moq.
var _ IHubspotCRMAPI = &IHubspotCRMAPIMock{}

// IHubspotCRMAPIMock is a mock implementation of IHubspotCRMAPI.
//
//	func TestSomethingThatUsesIHubspotCRMAPI(t *testing.T) {
//
//		// make and configure a mocked IHubspotCRMAPI
//		mockedIHubspotCRMAPI := &IHubspotCRMAPIMock{
//			GetCompanyForContactFunc: func(contactID string) (string, error) {
//				panic("mock out the GetCompanyForContact method")
//			},
//			GetDealForCompanyFunc: func(companyID string) (string, error) {
//				panic("mock out the GetDealForCompany method")
//			},
//			SearchCompaniesFunc: func(filterMap map[string]string, properties []string) ([]HubSpotSearchResult, error) {
//				panic("mock out the SearchCompanies method")
//			},
//			SearchContactsFunc: func(filterMap map[string]string, properties []string) ([]HubSpotSearchResult, error) {
//				panic("mock out the SearchContacts method")
//			},
//			UpdateCompanyFunc: func(companyID string, jsonPayload *bytes.Buffer) error {
//				panic("mock out the UpdateCompany method")
//			},
//		}
//
//		// use mockedIHubspotCRMAPI in code that requires IHubspotCRMAPI
//		// and then make assertions.
//
//	}
type IHubspotCRMAPIMock struct {
	// GetCompanyForContactFunc mocks the GetCompanyForContact method.
	GetCompanyForContactFunc func(contactID string) (string, error)

	// GetDealForCompanyFunc mocks the GetDealForCompany method.
	GetDealForCompanyFunc func(companyID string) (string, error)

	// SearchCompaniesFunc mocks the SearchCompanies method.
	SearchCompaniesFunc func(filterMap map[string]string, properties []string) ([]HubSpotSearchResult, error)

	// SearchContactsFunc mocks the SearchContacts method.
	SearchContactsFunc func(filterMap map[string]string, properties []string) ([]HubSpotSearchResult, error)

	// UpdateCompanyFunc mocks the UpdateCompany method.
	UpdateCompanyFunc func(companyID string, jsonPayload *bytes.Buffer) error

	// calls tracks calls to the methods.
	calls struct {
		// GetCompanyForContact holds details about calls to the GetCompanyForContact method.
		GetCompanyForContact []struct {
			// ContactID is the contactID argument value.
			ContactID string
		}
		// GetDealForCompany holds details about calls to the GetDealForCompany method.
		GetDealForCompany []struct {
			// CompanyID is the companyID argument value.
			CompanyID string
		}
		// SearchCompanies holds details about calls to the SearchCompanies method.
		SearchCompanies []struct {
			// FilterMap is the filterMap argument value.
			FilterMap map[string]string
			// Properties is the properties argument value.
			Properties []string
		}
		// SearchContacts holds details about calls to the SearchContacts method.
		SearchContacts []struct {
			// FilterMap is the filterMap argument value.
			FilterMap map[string]string
			// Properties is the properties argument value.
			Properties []string
		}
		// UpdateCompany holds details about calls to the UpdateCompany method.
		UpdateCompany []struct {
			// CompanyID is the companyID argument value.
			CompanyID string
			// JsonPayload is the jsonPayload argument value.
			JsonPayload *bytes.Buffer
		}
	}
	lockGetCompanyForContact sync.RWMutex
	lockGetDealForCompany    sync.RWMutex
	lockSearchCompanies      sync.RWMutex
	lockSearchContacts       sync.RWMutex
	lockUpdateCompany        sync.RWMutex
}

// GetCompanyForContact calls GetCompanyForContactFunc.
func (mock *IHubspotCRMAPIMock) GetCompanyForContact(contactID string) (string, error) {
	if mock.GetCompanyForContactFunc == nil {
		panic("IHubspotCRMAPIMock.GetCompanyForContactFunc: method is nil but IHubspotCRMAPI.GetCompanyForContact was just called")
	}
	callInfo := struct {
		ContactID string
	}{
		ContactID: contactID,
	}
	mock.lockGetCompanyForContact.Lock()
	mock.calls.GetCompanyForContact = append(mock.calls.GetCompanyForContact, callInfo)
	mock.lockGetCompanyForContact.Unlock()
	return mock.GetCompanyForContactFunc(contactID)
}

// GetCompanyForContactCalls gets all the calls that were made to GetCompanyForContact.
// Check the length with:
//
//	len(mockedIHubspotCRMAPI.GetCompanyForContactCalls())
func (mock *IHubspotCRMAPIMock) GetCompanyForContactCalls() []struct {
	ContactID string
} {
	var calls []struct {
		ContactID string
	}
	mock.lockGetCompanyForContact.RLock()
	calls = mock.calls.GetCompanyForContact
	mock.lockGetCompanyForContact.RUnlock()
	return calls
}

// GetDealForCompany calls GetDealForCompanyFunc.
func (mock *IHubspotCRMAPIMock) GetDealForCompany(companyID string) (string, error) {
	if mock.GetDealForCompanyFunc == nil {
		panic("IHubspotCRMAPIMock.GetDealForCompanyFunc: method is nil but IHubspotCRMAPI.GetDealForCompany was just called")
	}
	callInfo := struct {
		CompanyID string
	}{
		CompanyID: companyID,
	}
	mock.lockGetDealForCompany.Lock()
	mock.calls.GetDealForCompany = append(mock.calls.GetDealForCompany, callInfo)
	mock.lockGetDealForCompany.Unlock()
	return mock.GetDealForCompanyFunc(companyID)
}

// GetDealForCompanyCalls gets all the calls that were made to GetDealForCompany.
// Check the length with:
//
//	len(mockedIHubspotCRMAPI.GetDealForCompanyCalls())
func (mock *IHubspotCRMAPIMock) GetDealForCompanyCalls() []struct {
	CompanyID string
} {
	var calls []struct {
		CompanyID string
	}
	mock.lockGetDealForCompany.RLock()
	calls = mock.calls.GetDealForCompany
	mock.lockGetDealForCompany.RUnlock()
	return calls
}

// SearchCompanies calls SearchCompaniesFunc.
func (mock *IHubspotCRMAPIMock) SearchCompanies(filterMap map[string]string, properties []string) ([]HubSpotSearchResult, error) {
	if mock.SearchCompaniesFunc == nil {
		panic("IHubspotCRMAPIMock.SearchCompaniesFunc: method is nil but IHubspotCRMAPI.SearchCompanies was just called")
	}
	callInfo := struct {
		FilterMap  map[string]string
		Properties []string
	}{
		FilterMap:  filterMap,
		Properties: properties,
	}
	mock.lockSearchCompanies.Lock()
	mock.calls.SearchCompanies = append(mock.calls.SearchCompanies, callInfo)
	mock.lockSearchCompanies.Unlock()
	return mock.SearchCompaniesFunc(filterMap, properties)
}

// SearchCompaniesCalls gets all the calls that were made to SearchCompanies.
// Check the length with:
//
//	len(mockedIHubspotCRMAPI.SearchCompaniesCalls())
func (mock *IHubspotCRMAPIMock) SearchCompaniesCalls() []struct {
	FilterMap  map[string]string
	Properties []string
} {
	var calls []struct {
		FilterMap  map[string]string
		Properties []string
	}
	mock.lockSearchCompanies.RLock()
	calls = mock.calls.SearchCompanies
	mock.lockSearchCompanies.RUnlock()
	return calls
}

// SearchContacts calls SearchContactsFunc.
func (mock *IHubspotCRMAPIMock) SearchContacts(filterMap map[string]string, properties []string) ([]HubSpotSearchResult, error) {
	if mock.SearchContactsFunc == nil {
		panic("IHubspotCRMAPIMock.SearchContactsFunc: method is nil but IHubspotCRMAPI.SearchContacts was just called")
	}
	callInfo := struct {
		FilterMap  map[string]string
		Properties []string
	}{
		FilterMap:  filterMap,
		Properties: properties,
	}
	mock.lockSearchContacts.Lock()
	mock.calls.SearchContacts = append(mock.calls.SearchContacts, callInfo)
	mock.lockSearchContacts.Unlock()
	return mock.SearchContactsFunc(filterMap, properties)
}

// SearchContactsCalls gets all the calls that were made to SearchContacts.
// Check the length with:
//
//	len(mockedIHubspotCRMAPI.SearchContactsCalls())
func (mock *IHubspotCRMAPIMock) SearchContactsCalls() []struct {
	FilterMap  map[string]string
	Properties []string
} {
	var calls []struct {
		FilterMap  map[string]string
		Properties []string
	}
	mock.lockSearchContacts.RLock()
	calls = mock.calls.SearchContacts
	mock.lockSearchContacts.RUnlock()
	return calls
}

// UpdateCompany calls UpdateCompanyFunc.
func (mock *IHubspotCRMAPIMock) UpdateCompany(companyID string, jsonPayload *bytes.Buffer) error {
	if mock.UpdateCompanyFunc == nil {
		panic("IHubspotCRMAPIMock.UpdateCompanyFunc: method is nil but IHubspotCRMAPI.UpdateCompany was just called")
	}
	callInfo := struct {
		CompanyID   string
		JsonPayload *bytes.Buffer
	}{
		CompanyID:   companyID,
		JsonPayload: jsonPayload,
	}
	mock.lockUpdateCompany.Lock()
	mock.calls.UpdateCompany = append(mock.calls.UpdateCompany, callInfo)
	mock.lockUpdateCompany.Unlock()
	return mock.UpdateCompanyFunc(companyID, jsonPayload)
}

// UpdateCompanyCalls gets all the calls that were made to UpdateCompany.
// Check the length with:
//
//	len(mockedIHubspotCRMAPI.UpdateCompanyCalls())
func (mock *IHubspotCRMAPIMock) UpdateCompanyCalls() []struct {
	CompanyID   string
	JsonPayload *bytes.Buffer
} {
	var calls []struct {
		CompanyID   string
		JsonPayload *bytes.Buffer
	}
	mock.lockUpdateCompany.RLock()
	calls = mock.calls.UpdateCompany
	mock.lockUpdateCompany.RUnlock()
	return calls
}
