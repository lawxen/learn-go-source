package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
)

func main() {
	// 连接Elasticsearch
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200/",
		},
	}
	es, _ := elasticsearch.NewClient(cfg)
	// 创建索引
	es.Indices.Create("my_go")
	// 在索引里创建文档
	document := struct {
		Name string `json:"name"`
	}{
		"golang-elasticsearch",
	}
	data, _ := json.Marshal(document)
	dr, _ := es.Index("my_go", bytes.NewReader(data))
	fmt.Printf("在索引里创建文档：%v\n", dr)
	// 获取文档ID
	var docData map[string]interface{}
	json.NewDecoder(dr.Body).Decode(&docData)
	did := docData["_id"].(string)
	// 查找文档的id
	sr, _ := es.Get("my_go", did)
	fmt.Printf("查找文档的id：%v\n", sr)
	// 按条件搜索文档
	query := `{ "query": { "match": {"name" : "golang"} } }`
	sr1, _ := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("my_go"),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithFrom(0),
		es.Search.WithSize(1000),
	)
	fmt.Printf("按条件搜索文档：%v\n", sr1)
	// 搜索全部文档
	sr2, _ := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("my_go"),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithFrom(0),
		es.Search.WithSize(1000),
	)
	fmt.Printf("搜索全部文档：%v\n", sr2)
	// 更新文档
	ur, _ := es.Update("my_go", did, strings.NewReader(`{"doc":{"name" : "gin"}} `))
	fmt.Printf("更新文档：%v\n", ur)
	// 删除文档
	es.Delete("my_go", did)
	// 删除索引
	es.Indices.Delete([]string{"my_go"})
}
