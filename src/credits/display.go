package credits

import (
	"github.com/cruffinoni/EpiSchedule/src/environment"
)

func DisplayCreditsInfo(env environment.Environment) {
	env.Log(environment.VerboseMedium, "-> Below is the minimum of each step you can be in, based on the number of credits you currently have:\n")
	env.Logf(environment.VerboseMedium, environment.ColorRed+"	- Red: "+environment.ColorReset+
		"minimum amount before being in a deficit of credits: "+
		environment.ColorRed+"%v\n", env.User.Credits.Minimum)
	env.Logf(environment.VerboseMedium, environment.ColorYellow+"	- Yellow: "+environment.ColorReset+
		"minimum amount aimed at the moment of they year: "+
		environment.ColorYellow+"%v\n", env.User.Credits.Aimed)
	env.Logf(environment.VerboseMedium, environment.ColorGreen+"	- Green: "+environment.ColorReset+
		"minimum amount you should have before the end of the year: "+
		environment.ColorGreen+"%v\n", env.User.Credits.Objective)
}
