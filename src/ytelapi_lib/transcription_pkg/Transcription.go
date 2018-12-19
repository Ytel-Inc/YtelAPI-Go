/*
 * ytelapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package transcription_pkg

import "ytelapi_lib/models_pkg"
import "ytelapi_lib/configuration_pkg"

/*
 * Interface for the TRANSCRIPTION_IMPL
 */
type TRANSCRIPTION interface {
    CreateTranscribeAudioURL (string) (string, error)

    CreateListTranscriptions (*int64, *int64, models_pkg.StatusEnum, *string) (string, error)

    CreateViewTranscription (string) (string, error)

    CreateTranscribeRecording (string) (string, error)
}

/*
 * Factory for the TRANSCRIPTION interaface returning TRANSCRIPTION_IMPL
 */
func NewTRANSCRIPTION(config configuration_pkg.CONFIGURATION) *TRANSCRIPTION_IMPL {
    client := new(TRANSCRIPTION_IMPL)
    client.config = config
    return client
}