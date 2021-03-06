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
 * Type definition for NumberTypeEnum enum
 */
type NumberTypeEnum int

/**
 * Value collection for NumberTypeEnum enum
 */
const (
    NumberType_ALL NumberTypeEnum = 1 + iota
    NumberType_VOICE
    NumberType_SMS
)

func (r NumberTypeEnum) MarshalJSON() ([]byte, error) { 
    s := NumberTypeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *NumberTypeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  NumberTypeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts NumberTypeEnum to its string representation
 */
func NumberTypeEnumToValue(numberTypeEnum NumberTypeEnum) string {
    switch numberTypeEnum {
        case NumberType_ALL:
    		return "all"		
        case NumberType_VOICE:
    		return "voice"		
        case NumberType_SMS:
    		return "sms"		
        default:
        	return "all"
    }
}

/**
 * Converts NumberTypeEnum Array to its string Array representation
*/
func NumberTypeEnumArrayToValue(numberTypeEnum []NumberTypeEnum) []string {
    convArray := make([]string,len( numberTypeEnum))
    for i:=0; i<len(numberTypeEnum);i++ {
        convArray[i] = NumberTypeEnumToValue(numberTypeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NumberTypeEnumFromValue(value string) NumberTypeEnum {
    switch value {
        case "all":
            return NumberType_ALL
        case "voice":
            return NumberType_VOICE
        case "sms":
            return NumberType_SMS
        default:
            return NumberType_ALL
    }
}
