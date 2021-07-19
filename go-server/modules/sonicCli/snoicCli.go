package sonicCli

import (
	"github.com/expectedsh/go-sonic/sonic"
)

var Ingester sonic.Ingestable
var Search sonic.Searchable

func InitSonicCli() {
	in, err := sonic.NewIngester("localhost", 11491, "SecretPassword")
	if err != nil {
		println(err)
	}
	Ingester = in

	search, err := sonic.NewSearch("localhost", 11491, "SecretPassword")
	if err != nil {
		println(err)
	}
	Search = search

}
