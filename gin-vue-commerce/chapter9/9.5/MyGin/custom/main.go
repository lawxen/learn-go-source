package main

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 定义消费者
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MaxBytes:       10e6, // 10MB 设置数据接收大小
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

// 定义消费方法
func Consumer(id int) {
	defer wg.Done()
	kafkaGroupID := "consumer-" + string(id)
	reader := getKafkaReader("localhost:9092", "MyGin", kafkaGroupID)
	defer reader.Close()
	fmt.Printf("%d 开始消费 ... !!\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("%d 消费错误 %v", err)
			continue
		}
		fmt.Printf("消费者：%d，订阅：%v，分区：%v，偏移位置：%v，时间：%d，数据：%s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), m.Key, m.Value)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Consumer(i)
		time.Sleep(10 * time.Second)
	}
	wg.Wait()
}
