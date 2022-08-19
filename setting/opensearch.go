package setting

type OpensearchSetting struct {
	Address  string
	Username string
	Password string
}

var Opensearch = &OpensearchSetting{}

func OpensearchSetup() *OpensearchSetting {

	Opensearch.Address = env("ES_HOST", "http://localhost:9200")
	Opensearch.Username = env("ES_USERNAME", "admin")
	Opensearch.Password = env("ES_PASSWORD", "admin")

	return Opensearch
}
