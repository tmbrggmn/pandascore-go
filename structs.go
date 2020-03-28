package pandascore

// Series represents an instance of a league event.
//
// See Also
//
// https://developers.pandascore.co/doc/#section/Introduction/Events-hierarchy
type Series struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
}

// Represents an error coming directly from the PandaScore API (eg. no or invalid access token).
type PandaScoreError struct {
	Message string `json:"error"`
}

func (pse *PandaScoreError) Error() string {
	return "PandaScore error: " + pse.Message
}
