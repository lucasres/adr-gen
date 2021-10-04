package helpers

import "fmt"

// PrintAndExitIfGetFlagReturnError handle erros returned by get flag method of cobra package.
// These methods of the cobra package only return an error if there is a problem between trying to retrieve the value and setting the flag.
func PrintAndExitIfGetFlagReturnError(flagName string, err error) {
	if err != nil {
		PrintErrorAndExit(
			fmt.Errorf("can't retrive \"%s\" flag: %w", flagName, err),
		)
	}
}
