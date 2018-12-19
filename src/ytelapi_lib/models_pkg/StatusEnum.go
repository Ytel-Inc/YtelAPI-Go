/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for StatusEnum enum
 */
type StatusEnum int

/**
 * Value collection for StatusEnum enum
 */
const (
    Status_INPROGRESS StatusEnum = 1 + iota
    Status_SUCCESS
    Status_FAILURE
)

func (r StatusEnum) MarshalJSON() ([]byte, error) { 
    s := StatusEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *StatusEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  StatusEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts StatusEnum to its string representation
 */
func StatusEnumToValue(statusEnum StatusEnum) string {
    switch statusEnum {
        case Status_INPROGRESS:
    		return "inprogress"		
        case Status_SUCCESS:
    		return "success"		
        case Status_FAILURE:
    		return "failure"		
        default:
        	return "inprogress"
    }
}

/**
 * Converts StatusEnum Array to its string Array representation
*/
func StatusEnumArrayToValue(statusEnum []StatusEnum) []string {
    convArray := make([]string,len( statusEnum))
    for i:=0; i<len(statusEnum);i++ {
        convArray[i] = StatusEnumToValue(statusEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func StatusEnumFromValue(value string) StatusEnum {
    switch value {
        case "inprogress":
            return Status_INPROGRESS
        case "success":
            return Status_SUCCESS
        case "failure":
            return Status_FAILURE
        default:
            return Status_INPROGRESS
    }
}