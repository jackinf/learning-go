package common

func StartUp() {
	initConfig()
	initKeys()
	createDbSession()
	addIndexes()
}