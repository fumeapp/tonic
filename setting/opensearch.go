package setting

type OpensearchSetting struct {
	Connect  string
	Address  string
	Signed   string
	Username string
	Password string
}

var Opensearch = &OpensearchSetting{}

func OpensearchSetup() *OpensearchSetting {

	Opensearch.Connect = env("OS_CONNECT", "false")
	Opensearch.Address = env("OS_HOST", "http://localhost:9200")
	Opensearch.Signed = env("OS_SIGNED", "false")
	Opensearch.Username = env("OS_USERNAME", "admin")
	Opensearch.Password = env("OS_PASSWORD", "admin")

	return Opensearch
}
