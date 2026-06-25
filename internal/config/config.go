package config

type Config struct {
	Difficulty int    // Difficulty
	HddLimit   int    // Max size of lua program in roboto
	LuaAccess  string // Trusted commands of LUA
	CoreSpeed  int    // Speed of program in roboto
	OSspeed    int    // Speed of viruses and ROBOTOOS(Not roboto)
}

func SetDifficulty(level int) {
	switch level {
	case 1: // Easy
		GameConfig = Config{
			Difficulty: 1,
			HddLimit:   1024,
			LuaAccess:  "Basic",
			CoreSpeed:  50,
			OSspeed:    10,
		}
	case 2: // Normal
		GameConfig = Config{
			Difficulty: 2,
			HddLimit:   512,
			LuaAccess:  "Limited",
			CoreSpeed:  25,
			OSspeed:    20,
		}
	case 3: // Hard
		GameConfig = Config{
			Difficulty: 3,
			HddLimit:   256,
			LuaAccess:  "Minimal",
			CoreSpeed:  10,
			OSspeed:    40,
		}
	}
}

var GameConfig Config
