/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.11.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateCallRecording'
type CreateCallRecordingParams struct {
	//
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	RecordingStatusCallback *string `json:"RecordingStatusCallback,omitempty"`
	//
	RecordingStatusCallbackEvent *[]string `json:"RecordingStatusCallbackEvent,omitempty"`
}

func (params *CreateCallRecordingParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateCallRecordingParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateCallRecordingParams) SetPathAccountSid(PathAccountSid string) *CreateCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateCallRecordingParams) SetRecordingStatusCallback(RecordingStatusCallback string) *CreateCallRecordingParams {
	params.RecordingStatusCallback = &RecordingStatusCallback
	return params
}
func (params *CreateCallRecordingParams) SetRecordingStatusCallbackEvent(RecordingStatusCallbackEvent []string) *CreateCallRecordingParams {
	params.RecordingStatusCallbackEvent = &RecordingStatusCallbackEvent
	return params
}

func (c *ApiService) CreateCallRecording(CallSid string, params *CreateCallRecordingParams) (*TestResponseObject, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.RecordingStatusCallback != nil {
		data.Set("RecordingStatusCallback", *params.RecordingStatusCallback)
	}
	if params != nil && params.RecordingStatusCallbackEvent != nil {
		for _, item := range *params.RecordingStatusCallbackEvent {
			data.Add("RecordingStatusCallbackEvent", item)
		}
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TestResponseObject{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteCallRecording'
type DeleteCallRecordingParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteCallRecordingParams) SetPathAccountSid(PathAccountSid string) *DeleteCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) DeleteCallRecording(CallSid string, TestInteger int, params *DeleteCallRecordingParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{TestInteger}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"TestInteger"+"}", fmt.Sprint(TestInteger), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchCallRecording'
type FetchCallRecordingParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallRecordingParams) SetPathAccountSid(PathAccountSid string) *FetchCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchCallRecording(CallSid string, TestInteger int, params *FetchCallRecordingParams) (*TestResponseObject, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{TestInteger}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"TestInteger"+"}", fmt.Sprint(TestInteger), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TestResponseObject{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCallRecording'
type ListCallRecordingParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	//
	DateTest *string `json:"Date.Test,omitempty"`
	//
	DateCreatedBefore *time.Time `json:"DateCreated&lt;,omitempty"`
	//
	DateCreatedAfter *time.Time `json:"DateCreated&gt;,omitempty"`
	//
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCallRecordingParams) SetPathAccountSid(PathAccountSid string) *ListCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListCallRecordingParams) SetDateCreated(DateCreated time.Time) *ListCallRecordingParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *ListCallRecordingParams) SetDateTest(DateTest string) *ListCallRecordingParams {
	params.DateTest = &DateTest
	return params
}
func (params *ListCallRecordingParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListCallRecordingParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListCallRecordingParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListCallRecordingParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListCallRecordingParams) SetPageSize(PageSize int) *ListCallRecordingParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCallRecordingParams) SetLimit(Limit int) *ListCallRecordingParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of CallRecording records from the API. Request is executed immediately.
func (c *ApiService) PageCallRecording(CallSid string, params *ListCallRecordingParams, pageToken, pageNumber string) (*ListCallRecordingResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateTest != nil {
		data.Set("Date.Test", fmt.Sprint(*params.DateTest))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreated<", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreated>", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists CallRecording records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCallRecording(CallSid string, params *ListCallRecordingParams) ([]TestResponseObject, error) {
	response, err := c.StreamCallRecording(CallSid, params)
	if err != nil {
		return nil, err
	}

	records := make([]TestResponseObject, 0)

	for record := range response {
		records = append(records, record)
	}

	return records, err
}

// Streams CallRecording records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCallRecording(CallSid string, params *ListCallRecordingParams) (chan TestResponseObject, error) {
	if params == nil {
		params = &ListCallRecordingParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCallRecording(CallSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 1
	//set buffer size of the channel to 1
	channel := make(chan TestResponseObject, 1)

	go func() {
		for response != nil {
			responseRecords := response.Recordings
			for item := range responseRecords {
				channel <- responseRecords[item]
				curRecord += 1
				if params.Limit != nil && *params.Limit < curRecord {
					close(channel)
					return
				}
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, c.getNextListCallRecordingResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCallRecordingResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCallRecordingResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCallRecordingResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateCallRecording'
type UpdateCallRecordingParams struct {
	//
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	//
	PauseBehavior *string `json:"PauseBehavior,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateCallRecordingParams) SetPathAccountSid(PathAccountSid string) *UpdateCallRecordingParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateCallRecordingParams) SetPauseBehavior(PauseBehavior string) *UpdateCallRecordingParams {
	params.PauseBehavior = &PauseBehavior
	return params
}
func (params *UpdateCallRecordingParams) SetStatus(Status string) *UpdateCallRecordingParams {
	params.Status = &Status
	return params
}

func (c *ApiService) UpdateCallRecording(CallSid string, TestInteger int, params *UpdateCallRecordingParams) (*TestResponseObject, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{TestInteger}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)
	path = strings.Replace(path, "{"+"TestInteger"+"}", fmt.Sprint(TestInteger), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PauseBehavior != nil {
		data.Set("PauseBehavior", *params.PauseBehavior)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TestResponseObject{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
