package command

import "flag"

func CreateNewCommand(name *string) *flag.FlagSet {
	return flag.NewFlagSet(*name, flag.ExitOnError)
}

func CreateImdbIdParameter(command *flag.FlagSet) *string {
	return command.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")
}
