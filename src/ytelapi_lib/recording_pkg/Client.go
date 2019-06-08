/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package recording_pkg


import(
	"github.com/apimatic/unirest-go"
	"github.com/Ytel-Inc/YtelAPI-Go/src/ytelapi_lib/apihelper_pkg"
	"github.com/Ytel-Inc/YtelAPI-Go/src/ytelapi_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type RECORDING_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Remove a recording from your Ytel account.
 * @param    string        recordingsid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateDeleteRecording (
            recordingsid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/deleterecording.json"

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

        "recordingsid" : recordingsid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request, false);
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
 * Retrieve a list of recording objects.
 * @param    *int64         page            parameter: Optional
 * @param    *int64         pagesize        parameter: Optional
 * @param    *string        datecreated     parameter: Optional
 * @param    *string        callsid         parameter: Optional
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateListRecordings (
            page *int64,
            pagesize *int64,
            datecreated *string,
            callsid *string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/listrecording.json"

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

        "page" : page,
        "pagesize" : pagesize,
        "Datecreated" : datecreated,
        "callsid" : callsid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request, false);
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
 * Retrieve the recording of a call by its RecordingSid. This resource will return information regarding the call such as start time, end time, duration, and so forth.
 * @param    string        recordingsid     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *RECORDING_IMPL) CreateViewRecording (
            recordingsid string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/recording/viewrecording.json"

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

        "recordingsid" : recordingsid,

    }


    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, parameters, me.config.BasicAuthUserName(), me.config.BasicAuthPassword())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request, false);
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

