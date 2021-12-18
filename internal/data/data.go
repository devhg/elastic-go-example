package data

import (
	"fmt"
	"log"
	"os"

	"github.com/google/wire"
	"github.com/olivere/elastic/v7"

	"github.com/devhg/es/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewEsClient)

func NewEsClient(conf *conf.Config) *elastic.Client {
	url := fmt.Sprintf("http://%s", conf.Elastic.Addr)
	client, err := elastic.NewClient(
		elastic.SetSniff(false), // docker需要 https://blog.csdn.net/finghting321/article/details/105991741
		elastic.SetURL(url),     // elastic 服务地址
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ERR ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "ELASTIC INFO ", log.LstdFlags)),
	)
	if err != nil {
		log.Fatalln("Failed to create elastic client")
	}
	return client
}
