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
 * Type definition for Direction1Enum enum
 */
type Direction1Enum int

/**
 * Value collection for Direction1Enum enum
 */
const (
    Direction1_IN Direction1Enum = 1 + iota
    Direction1_OUT
    Direction1_BOTH
)

func (r Direction1Enum) MarshalJSON() ([]byte, error) { 
    s := Direction1EnumToValue(r)
    return json.Marshal(s) 
} 

func (r *Direction1Enum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  Direction1EnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts Direction1Enum to its string representation
 */
func Direction1EnumToValue(direction1Enum Direction1Enum) string {
    switch direction1Enum {
        case Direction1_IN:
    		return "in"		
        case Direction1_OUT:
    		return "out"		
        case Direction1_BOTH:
    		return "both"		
        default:
        	return "in"
    }
}

/**
 * Converts Direction1Enum Array to its string Array representation
*/
func Direction1EnumArrayToValue(direction1Enum []Direction1Enum) []string {
    convArray := make([]string,len( direction1Enum))
    for i:=0; i<len(direction1Enum);i++ {
        convArray[i] = Direction1EnumToValue(direction1Enum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func Direction1EnumFromValue(value string) Direction1Enum {
    switch value {
        case "in":
            return Direction1_IN
        case "out":
            return Direction1_OUT
        case "both":
            return Direction1_BOTH
        default:
            return Direction1_IN
    }
}