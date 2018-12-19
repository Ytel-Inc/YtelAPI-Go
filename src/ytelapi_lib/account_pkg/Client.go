/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package account_pkg


import(
	"fmt"
	"github.com/apimatic/unirest-go"
	"ytelapi_lib/apihelper_pkg"
	"ytelapi_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type ACCOUNT_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Retrieve information regarding your Ytel account by a specific date. The response object will contain data such as account status, balance, and account usage totals.
 * @param    string        date     parameter: Required
 * @return	Returns the string response from the API call
 */
func (me *ACCOUNT_IMPL) CreateViewAccount (
            date string) (string, error) {
        //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/accounts/viewaccount.json"

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

        "Date" : date,

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

