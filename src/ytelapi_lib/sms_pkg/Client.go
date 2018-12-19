/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package sms_pkg


import(
	"fmt"
	"github.com/apimatic/unirest-go"
	"ytelapi_lib/apihelper_pkg"
	"ytelapi_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type SMS_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Retrieve a list of Outbound SMS message objects.
 * @param    *int64         page         parameter: Optional
 * @param    *int64         pageSize     parameter: Optional
 * @param    *string        from         parameter: Optional
 * @param    *string        to           parameter: Optional
 * @param    *string        dateSent     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) CreateListSMS (
            page *int64,
            pageSize *int64,
            from *string,
            to *string,
            dateSent *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/listsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Page" : page,
        "PageSize" : pageSize,
        "From" : from,
        "To" : to,
        "DateSent" : dateSent,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a list of Inbound SMS message objects.
 * @param    *int64         page         parameter: Optional
 * @param    *int64         pageSize     parameter: Optional
 * @param    *string        from         parameter: Optional
 * @param    *string        to           parameter: Optional
 * @param    *string        dateSent     parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) CreateListInboundSMS (
            page *int64,
            pageSize *int64,
            from *string,
            to *string,
            dateSent *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/getinboundsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "Page" : page,
        "PageSize" : pageSize,
        "From" : from,
        "To" : to,
        "DateSent" : dateSent,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Send an SMS from a Ytel number
 * @param    string         from                      parameter: Required
 * @param    string         to                        parameter: Required
 * @param    string         body                      parameter: Required
 * @param    *string        method                    parameter: Optional
 * @param    *string        messageStatusCallback     parameter: Optional
 * @param    *bool          smartsms                  parameter: Optional
 * @param    *bool          deliveryStatus            parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) CreateSendSMS (
            from string,
            to string,
            body string,
            method *string,
            messageStatusCallback *string,
            smartsms *bool,
            deliveryStatus *bool) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/sendsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "From" : from,
        "To" : to,
        "Body" : body,
        "Method" : method,
        "MessageStatusCallback" : messageStatusCallback,
        "Smartsms" : smartsms,
        "DeliveryStatus" : deliveryStatus,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a single SMS message object by its SmsSid.
 * @param    string        messageSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) CreateViewSMS (
            messageSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/viewsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "MessageSid" : messageSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}

/**
 * Retrieve a single SMS message object with details by its SmsSid.
 * @param    string        messageSid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *SMS_IMPL) CreateViewSMSDetails (
            messageSid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/sms/viewdetailsms.json"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return "", err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //form parameters
    parameters := map[string]interface{} {

        "MessageSid" : messageSid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return "", err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return "", err
    }

    //returning the response
    return _response.Body, nil

}
