package dao

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/xiaohubai/go-layout/configs/global"
)

// InsertDoc 插入记录，自动创建id
func InsertDoc(indexName string, data interface{}) (err error) {
	_, err = global.Es.Index().Index(indexName).BodyJson(data).Do(context.TODO())
	return err
}

// InsertDocById 插入记录，根据id
func InsertDocById(indexName, id string, data interface{}) (err error) {
	_, err = global.Es.Index().Index(indexName).Id(id).BodyJson(data).Do(context.TODO())
	return err
}

// ExistsIndex 判断index是否存在
func ExistsIndex(indexName string) (bool, error) {
	exists, err := global.Es.IndexExists(indexName).Do(context.TODO())
	if err != nil {
		return false, err
	}
	if !exists {
		return false, err
	}
	return true, nil
}

// ExistsDocByQuery 判断doc是否存在，根据查询条件
func ExistsDocByQuery(indexName string, query elastic.Query) (bool, error) {
	//query=elastic.NewTermQuery("msgid.keyword", msgItem.MsgID)
	ret, err := global.Es.Search(indexName).Query(query).Size(0).Do(context.TODO())
	if err != nil {
		return false, err
	}
	if ret.TotalHits() <= 0 {
		return false, err
	}
	return true, nil
}

// SearchDocList 分页搜索数据，根据查询条件
/* query := elastic.NewBoolQuery()
if len(request.SearchDate)>0 {
	query.Must(elastic.NewTermQuery("date", request.SearchDate))
}
offset=(request.Page-1)*request.Size
ascending=false
*/
func SearchDocList(indexName string, query elastic.Query, offset, size int, sortField string, ascending bool) (
	*elastic.SearchResult, error) {

	op := global.Es.Search(indexName).Query(query)
	if len(sortField) > 0 {
		op = op.Sort(sortField, ascending)
	}
	ret, err := op.Size(size).From(offset).Do(context.TODO())
	return ret, err
}

// SearchDoc 搜索doc,根据查询条件
func SearchDoc(indexName string, query elastic.Query) (*elastic.SearchResult, error) {
	ret, err := global.Es.Search(indexName).Query(query).Do(context.TODO())
	return ret, err
}

// GetDocById 获取doc 根据id
func GetDocById(indexName, id string) (*elastic.GetResult, error) {
	ret, err := global.Es.Get().Index(indexName).Id(id).Do(context.TODO())
	return ret, err
}

// CreateIndex 创建index
func CreateIndex(indexName, mapping string) error {
	ret, err := global.Es.CreateIndex(indexName).BodyString(mapping).Do(context.TODO())
	if err != nil {
		return err
	}
	if !ret.Acknowledged {
		return fmt.Errorf("create index err")
	}
	return nil
}

// DeleteIndex 删除index
func DeleteIndex(indexName string) error {
	ret, err := global.Es.DeleteIndex(indexName).Do(context.TODO())
	if err != nil {
		return err
	}
	if !ret.Acknowledged {
		return fmt.Errorf("delete index err")
	}
	return nil
}

/* refresh:
"true":强制同步(慎用) 节点立即完成刷新(性能最差)
"false":异步刷新(默认值) 写入数据一段时间后可见(性能最好)
"wait_for":/等待同步 客户端等待, 直到节点自动完成刷新 */
// DeleteDocById 删除doc 根据id
func DeleteDocById(indexName, id, refresh string) error {
	ret, err := global.Es.Delete().Index(indexName).Id(id).Refresh(refresh).Do(context.TODO())
	if err != nil {
		return err
	}
	if ret.Result != "deleted" {
		return fmt.Errorf("delete id err")
	}
	return nil
}

// UpdateDocById 更新doc 根据id
func UpdateDocById(indexName, id, refresh string, data interface{}) error {
	ret, err := global.Es.Update().Index(indexName).Id(id).Doc(data).Refresh(refresh).Do(context.TODO())
	if err != nil {
		return err
	}
	if ret.Result != "updated" {
		return fmt.Errorf("delete id err")
	}
	return nil
}
