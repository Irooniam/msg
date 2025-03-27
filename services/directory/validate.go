package directory

import (
	"errors"
	"fmt"
)

/*
*
this map is what all action field is
validated against
*
*/
var ACTIONS = map[string]string{
	"REGISTER-DEALER": "DR",
	"DEALER-EVENT":    "DE",
}

var RACTIONS = map[string]string{
	"DR": "REGISTER-DEALER",
	"DE": "DEALER-EVENT",
}

func TranslateAction(action []byte) ([]byte, error) {
	sact := string(action)
	if _, ok := RACTIONS[sact]; !ok {
		return []byte{}, errors.New(fmt.Sprintf("Cant find action %s in RACTIONS map", action))
	}

	return []byte(RACTIONS[sact]), nil
}
