/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package subaccount_pkg

import "ytelapi_lib/models_pkg"
import "ytelapi_lib/configuration_pkg"

/*
 * Interface for the SUBACCOUNT_IMPL
 */
type SUBACCOUNT interface {
    CreateSubaccount (string, string, string, string, string) (string, error)

    CreateDeleteSubaccount (string, models_pkg.MergeNumberEnum) (string, error)

    CreateToggleSubaccountStatus (string, models_pkg.ActivateEnum) (string, error)
}

/*
 * Factory for the SUBACCOUNT interaface returning SUBACCOUNT_IMPL
 */
func NewSUBACCOUNT(config configuration_pkg.CONFIGURATION) *SUBACCOUNT_IMPL {
    client := new(SUBACCOUNT_IMPL)
    client.config = config
    return client
}