package setting

type OpensearchSetting struct {
	Address  string
	Username string
	Password string
}

var Opensearch = &OpensearchSetting{}

func OpensearchSetup() *OpensearchSetting {

	Opensearch.Address = env("OS_HOST", "http://localhost:9200")
	Opensearch.Username = env("OS_USERNAME", "admin")
	Opensearch.Password = env("OS_PASSWORD", "admin")

	return Opensearch
}
