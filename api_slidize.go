/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Smallize">
 *   Copyright (c) 2024 Slidize for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */
package slidizecloud

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// SlidizeApiService SlidizeApi service
type SlidizeApiService service

type ApiConvertRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	documents []*os.File
	format ExportFormat
}

func (r ApiConvertRequest) Documents(documents []*os.File) ApiConvertRequest {
	r.documents = documents
	return r
}

func (r ApiConvertRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ConvertExecute(r)
}

/*
Convert Converts files provided in the request body into target format and returns conversion result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param format Output file format.
 @return ApiConvertRequest
*/
func (a *SlidizeApiService) Convert(format ExportFormat) ApiConvertRequest {
	ctx := context.Background();
	return ApiConvertRequest{
		ApiService: a,
		ctx: ctx,
		format: format,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) ConvertExecute(r ApiConvertRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.Convert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/convert/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", url.PathEscape(parameterValueToString(r.format, "format")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.documents == nil {
		return localVarReturnValue, nil, reportError("documents is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentsLocalVarFormFileName string
	var documentsLocalVarFileName     string
	var documentsLocalVarFileBytes    []byte

	documentsLocalVarFormFileName = "documents"


	documentsLocalVarFile := r.documents

	if documentsLocalVarFile != nil {
		for _, file := range documentsLocalVarFile {
			fbs, _ := io.ReadAll(file)
			documentsLocalVarFileBytes = fbs
			documentsLocalVarFileName = file.Name()
			file.Close()
			formFiles = append(formFiles, formFile{fileBytes: documentsLocalVarFileBytes, fileName: documentsLocalVarFileName, formFileName: documentsLocalVarFormFileName})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConvertToVideoRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	document *os.File
	options *VideoOptions
}

func (r ApiConvertToVideoRequest) Document(document *os.File) ApiConvertToVideoRequest {
	r.document = document
	return r
}

func (r ApiConvertToVideoRequest) Options(options VideoOptions) ApiConvertToVideoRequest {
	r.options = &options
	return r
}

func (r ApiConvertToVideoRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ConvertToVideoExecute(r)
}

/*
ConvertToVideo Converts file provided in the request body into video.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConvertToVideoRequest
*/
func (a *SlidizeApiService) ConvertToVideo() ApiConvertToVideoRequest {
	ctx := context.Background();
	return ApiConvertToVideoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) ConvertToVideoExecute(r ApiConvertToVideoRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.ConvertToVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/video"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.document == nil {
		return localVarReturnValue, nil, reportError("document is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"


	documentLocalVarFile := r.document

	if documentLocalVarFile != nil {
		fbs, _ := io.ReadAll(documentLocalVarFile)

		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	}


	if r.options != nil {
		paramJson, err := parameterToJson(*r.options)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("options", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiImageWatermarkRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	documents []*os.File
	image *os.File
	options *ImageWatermarkOptions
}

func (r ApiImageWatermarkRequest) Documents(documents []*os.File) ApiImageWatermarkRequest {
	r.documents = documents
	return r
}

func (r ApiImageWatermarkRequest) Image(image *os.File) ApiImageWatermarkRequest {
	r.image = image
	return r
}

func (r ApiImageWatermarkRequest) Options(options ImageWatermarkOptions) ApiImageWatermarkRequest {
	r.options = &options
	return r
}

func (r ApiImageWatermarkRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ImageWatermarkExecute(r)
}

/*
ImageWatermark Adds image watermarks and return result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImageWatermarkRequest
*/
func (a *SlidizeApiService) ImageWatermark() ApiImageWatermarkRequest {
	ctx := context.Background();
	return ApiImageWatermarkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) ImageWatermarkExecute(r ApiImageWatermarkRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.ImageWatermark")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/watermark/image"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.documents == nil {
		return localVarReturnValue, nil, reportError("documents is required and must be specified")
	}
	if r.image == nil {
		return localVarReturnValue, nil, reportError("image is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentsLocalVarFormFileName string
	var documentsLocalVarFileName     string
	var documentsLocalVarFileBytes    []byte

	documentsLocalVarFormFileName = "documents"


	documentsLocalVarFile := r.documents

	if documentsLocalVarFile != nil {
		for _, file := range documentsLocalVarFile {
			fbs, _ := io.ReadAll(file)
			documentsLocalVarFileBytes = fbs
			documentsLocalVarFileName = file.Name()
			file.Close()
			formFiles = append(formFiles, formFile{fileBytes: documentsLocalVarFileBytes, fileName: documentsLocalVarFileName, formFileName: documentsLocalVarFormFileName})
		}
	}


	var imageLocalVarFormFileName string
	var imageLocalVarFileName     string
	var imageLocalVarFileBytes    []byte

	imageLocalVarFormFileName = "image"


	imageLocalVarFile := r.image

	if imageLocalVarFile != nil {
		fbs, _ := io.ReadAll(imageLocalVarFile)

		imageLocalVarFileBytes = fbs
		imageLocalVarFileName = imageLocalVarFile.Name()
		imageLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: imageLocalVarFileBytes, fileName: imageLocalVarFileName, formFileName: imageLocalVarFormFileName})
	}


	if r.options != nil {
		paramJson, err := parameterToJson(*r.options)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("options", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMergeRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	documents []*os.File
	format ExportFormat
	options *MergeOptions
}

func (r ApiMergeRequest) Documents(documents []*os.File) ApiMergeRequest {
	r.documents = documents
	return r
}

func (r ApiMergeRequest) Options(options MergeOptions) ApiMergeRequest {
	r.options = &options
	return r
}

func (r ApiMergeRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.MergeExecute(r)
}

/*
Merge Merges files provided in the request and saves the merge result into target format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param format Output file format.
 @return ApiMergeRequest
*/
func (a *SlidizeApiService) Merge(format ExportFormat) ApiMergeRequest {
	ctx := context.Background();
	return ApiMergeRequest{
		ApiService: a,
		ctx: ctx,
		format: format,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) MergeExecute(r ApiMergeRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.Merge")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merge/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", url.PathEscape(parameterValueToString(r.format, "format")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.documents == nil {
		return localVarReturnValue, nil, reportError("documents is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentsLocalVarFormFileName string
	var documentsLocalVarFileName     string
	var documentsLocalVarFileBytes    []byte

	documentsLocalVarFormFileName = "documents"


	documentsLocalVarFile := r.documents

	if documentsLocalVarFile != nil {
		for _, file := range documentsLocalVarFile {
			fbs, _ := io.ReadAll(file)
			documentsLocalVarFileBytes = fbs
			documentsLocalVarFileName = file.Name()
			file.Close()
			formFiles = append(formFiles, formFile{fileBytes: documentsLocalVarFileBytes, fileName: documentsLocalVarFileName, formFileName: documentsLocalVarFormFileName})
		}
	}


	if r.options != nil {
		paramJson, err := parameterToJson(*r.options)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("options", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiProtectRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	document *os.File
	options *ProtectionOptions
}

func (r ApiProtectRequest) Document(document *os.File) ApiProtectRequest {
	r.document = document
	return r
}

func (r ApiProtectRequest) Options(options ProtectionOptions) ApiProtectRequest {
	r.options = &options
	return r
}

func (r ApiProtectRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ProtectExecute(r)
}

/*
Protect Protects presentation with specified password and returns result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProtectRequest
*/
func (a *SlidizeApiService) Protect() ApiProtectRequest {
	ctx := context.Background();
	return ApiProtectRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) ProtectExecute(r ApiProtectRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.Protect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/lock"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.document == nil {
		return localVarReturnValue, nil, reportError("document is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"


	documentLocalVarFile := r.document

	if documentLocalVarFile != nil {
		fbs, _ := io.ReadAll(documentLocalVarFile)

		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	}


	if r.options != nil {
		paramJson, err := parameterToJson(*r.options)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("options", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveAnnotationsRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	document *os.File
}

func (r ApiRemoveAnnotationsRequest) Document(document *os.File) ApiRemoveAnnotationsRequest {
	r.document = document
	return r
}

func (r ApiRemoveAnnotationsRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.RemoveAnnotationsExecute(r)
}

/*
RemoveAnnotations Remove annotations from presentation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveAnnotationsRequest
*/
func (a *SlidizeApiService) RemoveAnnotations() ApiRemoveAnnotationsRequest {
	ctx := context.Background();
	return ApiRemoveAnnotationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) RemoveAnnotationsExecute(r ApiRemoveAnnotationsRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.RemoveAnnotations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/removeAnnotations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.document == nil {
		return localVarReturnValue, nil, reportError("document is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"


	documentLocalVarFile := r.document

	if documentLocalVarFile != nil {
		fbs, _ := io.ReadAll(documentLocalVarFile)

		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveMacrosRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	document *os.File
}

func (r ApiRemoveMacrosRequest) Document(document *os.File) ApiRemoveMacrosRequest {
	r.document = document
	return r
}

func (r ApiRemoveMacrosRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.RemoveMacrosExecute(r)
}

/*
RemoveMacros Remove macros from presentation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveMacrosRequest
*/
func (a *SlidizeApiService) RemoveMacros() ApiRemoveMacrosRequest {
	ctx := context.Background();
	return ApiRemoveMacrosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) RemoveMacrosExecute(r ApiRemoveMacrosRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.RemoveMacros")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/removeMacros"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.document == nil {
		return localVarReturnValue, nil, reportError("document is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"


	documentLocalVarFile := r.document

	if documentLocalVarFile != nil {
		fbs, _ := io.ReadAll(documentLocalVarFile)

		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReplaceTextRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	documents []*os.File
	options *ReplaceTextOptions
}

func (r ApiReplaceTextRequest) Documents(documents []*os.File) ApiReplaceTextRequest {
	r.documents = documents
	return r
}

func (r ApiReplaceTextRequest) Options(options ReplaceTextOptions) ApiReplaceTextRequest {
	r.options = &options
	return r
}

func (r ApiReplaceTextRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ReplaceTextExecute(r)
}

/*
ReplaceText Replace text in presentation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReplaceTextRequest
*/
func (a *SlidizeApiService) ReplaceText() ApiReplaceTextRequest {
	ctx := context.Background();
	return ApiReplaceTextRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) ReplaceTextExecute(r ApiReplaceTextRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.ReplaceText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/replaceText"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.documents == nil {
		return localVarReturnValue, nil, reportError("documents is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentsLocalVarFormFileName string
	var documentsLocalVarFileName     string
	var documentsLocalVarFileBytes    []byte

	documentsLocalVarFormFileName = "documents"


	documentsLocalVarFile := r.documents

	if documentsLocalVarFile != nil {
		for _, file := range documentsLocalVarFile {
			fbs, _ := io.ReadAll(file)
			documentsLocalVarFileBytes = fbs
			documentsLocalVarFileName = file.Name()
			file.Close()
			formFiles = append(formFiles, formFile{fileBytes: documentsLocalVarFileBytes, fileName: documentsLocalVarFileName, formFileName: documentsLocalVarFormFileName})
		}
	}


	if r.options != nil {
		paramJson, err := parameterToJson(*r.options)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("options", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSplitRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	document *os.File
	format ExportFormat
	options *SplitOptions
}

func (r ApiSplitRequest) Document(document *os.File) ApiSplitRequest {
	r.document = document
	return r
}

func (r ApiSplitRequest) Options(options SplitOptions) ApiSplitRequest {
	r.options = &options
	return r
}

func (r ApiSplitRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.SplitExecute(r)
}

/*
Split Splits presentation according to the specified slides range and saves result into target format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param format Output file format.
 @return ApiSplitRequest
*/
func (a *SlidizeApiService) Split(format ExportFormat) ApiSplitRequest {
	ctx := context.Background();
	return ApiSplitRequest{
		ApiService: a,
		ctx: ctx,
		format: format,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) SplitExecute(r ApiSplitRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.Split")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/split/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", url.PathEscape(parameterValueToString(r.format, "format")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.document == nil {
		return localVarReturnValue, nil, reportError("document is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"


	documentLocalVarFile := r.document

	if documentLocalVarFile != nil {
		fbs, _ := io.ReadAll(documentLocalVarFile)

		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	}


	if r.options != nil {
		paramJson, err := parameterToJson(*r.options)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("options", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTextWatermarkRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	documents []*os.File
	options *TextWatermarkOptions
}

func (r ApiTextWatermarkRequest) Documents(documents []*os.File) ApiTextWatermarkRequest {
	r.documents = documents
	return r
}

func (r ApiTextWatermarkRequest) Options(options TextWatermarkOptions) ApiTextWatermarkRequest {
	r.options = &options
	return r
}

func (r ApiTextWatermarkRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.TextWatermarkExecute(r)
}

/*
TextWatermark Adds text watermarks and return result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTextWatermarkRequest
*/
func (a *SlidizeApiService) TextWatermark() ApiTextWatermarkRequest {
	ctx := context.Background();
	return ApiTextWatermarkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) TextWatermarkExecute(r ApiTextWatermarkRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.TextWatermark")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/watermark/text"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.documents == nil {
		return localVarReturnValue, nil, reportError("documents is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}


	var documentsLocalVarFormFileName string
	var documentsLocalVarFileName     string
	var documentsLocalVarFileBytes    []byte

	documentsLocalVarFormFileName = "documents"


	documentsLocalVarFile := r.documents

	if documentsLocalVarFile != nil {
		for _, file := range documentsLocalVarFile {
			fbs, _ := io.ReadAll(file)
			documentsLocalVarFileBytes = fbs
			documentsLocalVarFileName = file.Name()
			file.Close()
			formFiles = append(formFiles, formFile{fileBytes: documentsLocalVarFileBytes, fileName: documentsLocalVarFileName, formFileName: documentsLocalVarFormFileName})
		}
	}


	if r.options != nil {
		paramJson, err := parameterToJson(*r.options)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("options", paramJson)
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUnprotectRequest struct {
	ctx context.Context
	ApiService *SlidizeApiService
	document *os.File
	password *string
}

func (r ApiUnprotectRequest) Document(document *os.File) ApiUnprotectRequest {
	r.document = document
	return r
}

// Password to remove.
func (r ApiUnprotectRequest) Password(password string) ApiUnprotectRequest {
	r.password = &password
	return r
}

func (r ApiUnprotectRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.UnprotectExecute(r)
}

/*
Unprotect Removes password from the presentation and returns result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnprotectRequest
*/
func (a *SlidizeApiService) Unprotect() ApiUnprotectRequest {
	ctx := context.Background();
	return ApiUnprotectRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SlidizeApiService) UnprotectExecute(r ApiUnprotectRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlidizeApiService.Unprotect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/unlock"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.document == nil {
		return localVarReturnValue, nil, reportError("document is required and must be specified")
	}
	if r.password == nil {
		return localVarReturnValue, nil, reportError("password is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "password", r.password, "")


	var documentLocalVarFormFileName string
	var documentLocalVarFileName     string
	var documentLocalVarFileBytes    []byte

	documentLocalVarFormFileName = "document"


	documentLocalVarFile := r.document

	if documentLocalVarFile != nil {
		fbs, _ := io.ReadAll(documentLocalVarFile)

		documentLocalVarFileBytes = fbs
		documentLocalVarFileName = documentLocalVarFile.Name()
		documentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: documentLocalVarFileBytes, fileName: documentLocalVarFileName, formFileName: documentLocalVarFormFileName})
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
