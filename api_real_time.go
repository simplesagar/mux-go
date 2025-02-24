// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type RealTimeApiService service

type GetRealtimeBreakdownParams struct {
	Dimension      string
	Timestamp      int32
	Filters        []string
	OrderBy        string
	OrderDirection string
}

// GetRealtimeBreakdown optionally accepts the APIOption of WithParams(*GetRealtimeBreakdownParams).
func (a *RealTimeApiService) GetRealtimeBreakdown(rEALTIMEMETRICID string, opts ...APIOption) (GetRealTimeBreakdownResponse, error) {
	var (
		localVarAPIOptions   = new(APIOptions)
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetRealTimeBreakdownResponse
	)

	for _, opt := range opts {
		opt(localVarAPIOptions)
	}

	localVarOptionals, ok := localVarAPIOptions.params.(*GetRealtimeBreakdownParams)
	if localVarAPIOptions.params != nil && !ok {
		return localVarReturnValue, reportError("provided params were not of type *GetRealtimeBreakdownParams")
	}

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/realtime/metrics/{REALTIME_METRIC_ID}/breakdown"
	localVarPath = strings.Replace(localVarPath, "{"+"REALTIME_METRIC_ID"+"}", fmt.Sprintf("%v", rEALTIMEMETRICID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && isSet(localVarOptionals.Dimension) {
		localVarQueryParams.Add("dimension", parameterToString(localVarOptionals.Dimension, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.Timestamp) {
		localVarQueryParams.Add("timestamp", parameterToString(localVarOptionals.Timestamp, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.Filters) {
		// This will "always work" for Mux's use case, since we always treat collections in query params as "multi" types.
		// The first version of this code checked the collectionFormat, but that's just wasted CPU cycles right now.
		for _, v := range localVarOptionals.Filters {
			localVarQueryParams.Add("filters[]", v)
		}
	}
	if localVarOptionals != nil && isSet(localVarOptionals.OrderBy) {
		localVarQueryParams.Add("order_by", parameterToString(localVarOptionals.OrderBy, ""))
	}
	if localVarOptionals != nil && isSet(localVarOptionals.OrderDirection) {
		localVarQueryParams.Add("order_direction", parameterToString(localVarOptionals.OrderDirection, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	// Check for common HTTP error status codes
	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return localVarReturnValue, err
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

type GetRealtimeHistogramTimeseriesParams struct {
	Filters []string
}

// GetRealtimeHistogramTimeseries optionally accepts the APIOption of WithParams(*GetRealtimeHistogramTimeseriesParams).
func (a *RealTimeApiService) GetRealtimeHistogramTimeseries(rEALTIMEHISTOGRAMMETRICID string, opts ...APIOption) (GetRealTimeHistogramTimeseriesResponse, error) {
	var (
		localVarAPIOptions   = new(APIOptions)
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetRealTimeHistogramTimeseriesResponse
	)

	for _, opt := range opts {
		opt(localVarAPIOptions)
	}

	localVarOptionals, ok := localVarAPIOptions.params.(*GetRealtimeHistogramTimeseriesParams)
	if localVarAPIOptions.params != nil && !ok {
		return localVarReturnValue, reportError("provided params were not of type *GetRealtimeHistogramTimeseriesParams")
	}

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/realtime/metrics/{REALTIME_HISTOGRAM_METRIC_ID}/histogram-timeseries"
	localVarPath = strings.Replace(localVarPath, "{"+"REALTIME_HISTOGRAM_METRIC_ID"+"}", fmt.Sprintf("%v", rEALTIMEHISTOGRAMMETRICID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && isSet(localVarOptionals.Filters) {
		// This will "always work" for Mux's use case, since we always treat collections in query params as "multi" types.
		// The first version of this code checked the collectionFormat, but that's just wasted CPU cycles right now.
		for _, v := range localVarOptionals.Filters {
			localVarQueryParams.Add("filters[]", v)
		}
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	// Check for common HTTP error status codes
	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return localVarReturnValue, err
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

type GetRealtimeTimeseriesParams struct {
	Filters []string
}

// GetRealtimeTimeseries optionally accepts the APIOption of WithParams(*GetRealtimeTimeseriesParams).
func (a *RealTimeApiService) GetRealtimeTimeseries(rEALTIMEMETRICID string, opts ...APIOption) (GetRealTimeTimeseriesResponse, error) {
	var (
		localVarAPIOptions   = new(APIOptions)
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetRealTimeTimeseriesResponse
	)

	for _, opt := range opts {
		opt(localVarAPIOptions)
	}

	localVarOptionals, ok := localVarAPIOptions.params.(*GetRealtimeTimeseriesParams)
	if localVarAPIOptions.params != nil && !ok {
		return localVarReturnValue, reportError("provided params were not of type *GetRealtimeTimeseriesParams")
	}

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/realtime/metrics/{REALTIME_METRIC_ID}/timeseries"
	localVarPath = strings.Replace(localVarPath, "{"+"REALTIME_METRIC_ID"+"}", fmt.Sprintf("%v", rEALTIMEMETRICID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && isSet(localVarOptionals.Filters) {
		// This will "always work" for Mux's use case, since we always treat collections in query params as "multi" types.
		// The first version of this code checked the collectionFormat, but that's just wasted CPU cycles right now.
		for _, v := range localVarOptionals.Filters {
			localVarQueryParams.Add("filters[]", v)
		}
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	// Check for common HTTP error status codes
	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return localVarReturnValue, err
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

func (a *RealTimeApiService) ListRealtimeDimensions(opts ...APIOption) (ListRealTimeDimensionsResponse, error) {
	var (
		localVarAPIOptions   = new(APIOptions)
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListRealTimeDimensionsResponse
	)

	for _, opt := range opts {
		opt(localVarAPIOptions)
	}

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/realtime/dimensions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	// Check for common HTTP error status codes
	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return localVarReturnValue, err
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

func (a *RealTimeApiService) ListRealtimeMetrics(opts ...APIOption) (ListRealTimeMetricsResponse, error) {
	var (
		localVarAPIOptions   = new(APIOptions)
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListRealTimeMetricsResponse
	)

	for _, opt := range opts {
		opt(localVarAPIOptions)
	}

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/realtime/metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := a.client.prepareRequest(localVarAPIOptions, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	// Check for common HTTP error status codes
	err = CheckForHttpError(localVarHttpResponse.StatusCode, localVarBody)
	if err != nil {
		return localVarReturnValue, err
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}
