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
 * Type definition for FileformatEnum enum
 */
type FileformatEnum int

/**
 * Value collection for FileformatEnum enum
 */
const (
    Fileformat_MP3 FileformatEnum = 1 + iota
    Fileformat_WAV
)

func (r FileformatEnum) MarshalJSON() ([]byte, error) { 
    s := FileformatEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *FileformatEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  FileformatEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts FileformatEnum to its string representation
 */
func FileformatEnumToValue(fileformatEnum FileformatEnum) string {
    switch fileformatEnum {
        case Fileformat_MP3:
    		return "mp3"		
        case Fileformat_WAV:
    		return "wav"		
        default:
        	return "mp3"
    }
}

/**
 * Converts FileformatEnum Array to its string Array representation
*/
func FileformatEnumArrayToValue(fileformatEnum []FileformatEnum) []string {
    convArray := make([]string,len( fileformatEnum))
    for i:=0; i<len(fileformatEnum);i++ {
        convArray[i] = FileformatEnumToValue(fileformatEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func FileformatEnumFromValue(value string) FileformatEnum {
    switch value {
        case "mp3":
            return Fileformat_MP3
        case "wav":
            return Fileformat_WAV
        default:
            return Fileformat_MP3
    }
}
