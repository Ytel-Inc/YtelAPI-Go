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
 * Type definition for NumberType2Enum enum
 */
type NumberType2Enum int

/**
 * Value collection for NumberType2Enum enum
 */
const (
    NumberType2_ALL NumberType2Enum = 1 + iota
    NumberType2_VOICE
    NumberType2_SMS
)

func (r NumberType2Enum) MarshalJSON() ([]byte, error) { 
    s := NumberType2EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *NumberType2Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  NumberType2EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts NumberType2Enum to its string representation
 */
func NumberType2EnumToValue(numberType2Enum NumberType2Enum) string {
    switch numberType2Enum {
        case NumberType2_ALL:
    		return "all"		
        case NumberType2_VOICE:
    		return "voice"		
        case NumberType2_SMS:
    		return "sms"		
        default:
        	return "all"
    }
}

/**
 * Converts NumberType2Enum Array to its string Array representation
*/
func NumberType2EnumArrayToValue(numberType2Enum []NumberType2Enum) []string {
    convArray := make([]string,len( numberType2Enum))
    for i:=0; i<len(numberType2Enum);i++ {
        convArray[i] = NumberType2EnumToValue(numberType2Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func NumberType2EnumFromValue(value string) NumberType2Enum {
    switch value {
        case "all":
            return NumberType2_ALL
        case "voice":
            return NumberType2_VOICE
        case "sms":
            return NumberType2_SMS
        default:
            return NumberType2_ALL
    }
}
