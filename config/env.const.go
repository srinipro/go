package config

//Env type direct to load the corresponding configuration file
//from the proposal made as JSON.
type Env string

//CurrentEnv to switch build stages before do the deployment
const CurrentEnv Env = Development

const (
	//Development direct to load dev.json on bootstrap of application
	Development Env = "DEV"
	//Quality direct to load qas.json on bootstrap of application
	Quality Env = "QAS"
	//Production direct to load prod.json on bootstrap of application
	Production Env = "PROD"
)
