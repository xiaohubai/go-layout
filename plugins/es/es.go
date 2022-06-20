package es

import (
	"github.com/olivere/elastic/v7"
	"github.com/xiaohubai/go-layout/configs/global"
)

//Es elastcisearch组件
func Init() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL(global.Cfg.Es.Path),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("elastic", "1qaz!QAZ"),
	)
	if err != nil {
		panic(err)
	}
	return client
}
