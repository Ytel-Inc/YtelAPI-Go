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
 * Type definition for GroupConfirmFileEnum enum
 */
type GroupConfirmFileEnum int

/**
 * Value collection for GroupConfirmFileEnum enum
 */
const (
    GroupConfirmFile_MP3 GroupConfirmFileEnum = 1 + iota
    GroupConfirmFile_WAV
)

func (r GroupConfirmFileEnum) MarshalJSON() ([]byte, error) { 
    s := GroupConfirmFileEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *GroupConfirmFileEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  GroupConfirmFileEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts GroupConfirmFileEnum to its string representation
 */
func GroupConfirmFileEnumToValue(groupConfirmFileEnum GroupConfirmFileEnum) string {
    switch groupConfirmFileEnum {
        case GroupConfirmFile_MP3:
    		return "mp3"		
        case GroupConfirmFile_WAV:
    		return "wav"		
        default:
        	return "mp3"
    }
}

/**
 * Converts GroupConfirmFileEnum Array to its string Array representation
*/
func GroupConfirmFileEnumArrayToValue(groupConfirmFileEnum []GroupConfirmFileEnum) []string {
    convArray := make([]string,len( groupConfirmFileEnum))
    for i:=0; i<len(groupConfirmFileEnum);i++ {
        convArray[i] = GroupConfirmFileEnumToValue(groupConfirmFileEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func GroupConfirmFileEnumFromValue(value string) GroupConfirmFileEnum {
    switch value {
        case "mp3":
            return GroupConfirmFile_MP3
        case "wav":
            return GroupConfirmFile_WAV
        default:
            return GroupConfirmFile_MP3
    }
}
