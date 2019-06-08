/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package usage_pkg

import "github.com/Ytel-Inc/YtelAPI-Go/src/ytelapi_lib/models_pkg"
import "github.com/Ytel-Inc/YtelAPI-Go/src/ytelapi_lib/configuration_pkg"

/*
 * Interface for the USAGE_IMPL
 */
type USAGE interface {
    CreateListUsage (models_pkg.ProductCodeEnum, *string, *string, *string) (string, error)
}

/*
 * Factory for the USAGE interaface returning USAGE_IMPL
 */
func NewUSAGE(config configuration_pkg.CONFIGURATION) *USAGE_IMPL {
    client := new(USAGE_IMPL)
    client.config = config
    return client
}
