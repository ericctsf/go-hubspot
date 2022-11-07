// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package go_hubspot

import (
	"sync"
)

// Ensure, that IHubspotFileAPIMock does implement IHubspotFileAPI.
// If this is not the case, regenerate this file with moq.
var _ IHubspotFileAPI = &IHubspotFileAPIMock{}

// IHubspotFileAPIMock is a mock implementation of IHubspotFileAPI.
//
// 	func TestSomethingThatUsesIHubspotFileAPI(t *testing.T) {
//
// 		// make and configure a mocked IHubspotFileAPI
// 		mockedIHubspotFileAPI := &IHubspotFileAPIMock{
// 			UploadFileFunc: func(file []byte, folderPath string, fileName string) (string, error) {
// 				panic("mock out the UploadFile method")
// 			},
// 		}
//
// 		// use mockedIHubspotFileAPI in code that requires IHubspotFileAPI
// 		// and then make assertions.
//
// 	}
type IHubspotFileAPIMock struct {
	// UploadFileFunc mocks the UploadFile method.
	UploadFileFunc func(file []byte, folderPath string, fileName string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// UploadFile holds details about calls to the UploadFile method.
		UploadFile []struct {
			// File is the file argument value.
			File []byte
			// FolderPath is the folderPath argument value.
			FolderPath string
			// FileName is the fileName argument value.
			FileName string
		}
	}
	lockUploadFile sync.RWMutex
}

// UploadFile calls UploadFileFunc.
func (mock *IHubspotFileAPIMock) UploadFile(file []byte, folderPath string, fileName string) (string, error) {
	if mock.UploadFileFunc == nil {
		panic("IHubspotFileAPIMock.UploadFileFunc: method is nil but IHubspotFileAPI.UploadFile was just called")
	}
	callInfo := struct {
		File       []byte
		FolderPath string
		FileName   string
	}{
		File:       file,
		FolderPath: folderPath,
		FileName:   fileName,
	}
	mock.lockUploadFile.Lock()
	mock.calls.UploadFile = append(mock.calls.UploadFile, callInfo)
	mock.lockUploadFile.Unlock()
	return mock.UploadFileFunc(file, folderPath, fileName)
}

// UploadFileCalls gets all the calls that were made to UploadFile.
// Check the length with:
//     len(mockedIHubspotFileAPI.UploadFileCalls())
func (mock *IHubspotFileAPIMock) UploadFileCalls() []struct {
	File       []byte
	FolderPath string
	FileName   string
} {
	var calls []struct {
		File       []byte
		FolderPath string
		FileName   string
	}
	mock.lockUploadFile.RLock()
	calls = mock.calls.UploadFile
	mock.lockUploadFile.RUnlock()
	return calls
}
