package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

// 定义生产者
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// 定义路由处理函数
func addUser(c *gin.Context) {
	name := c.Query("name")
	if len(name) <= 0 {
		c.JSON(200, gin.H{
			"status": "缺少参数name",
		})
	}
	// 设置Kafka数据格式
	body := map[string]interface{}{"name": name}
	jsonBody, _ := json.Marshal(body)
	msg := kafka.Message{
		Key:   []byte("User"),
		Value: []byte(jsonBody),
	}
	// 数据写入Kafka
	err := kafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"status": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "请求成功",
		})
	}
}

func main() {
	r := gin.Default()
	kafkaWriter = getKafkaWriter("localhost:9092", "MyGin")
	defer kafkaWriter.Close()
	r.GET("/", addUser)
	r.Run(":8000")
}
