package v1

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/plugins/email"
	"github.com/xiaohubai/go-layout/plugins/html"
)

func Warn(topic string) {
	partitionList, _ := global.KafkaConsumer.Partitions(topic)
	for partition := range partitionList {
		pc, _ := global.KafkaConsumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)

		defer pc.AsyncClose()
		for msg := range pc.Messages() {
			var w model.Warn
			_ = json.Unmarshal(msg.Value, &w)
			htmlText := "<html><body><p> {{.Type}}告警：" +
				"<p>● 时间：<td>{{.Date}}</td></p>" +
				"<p>● TraceId：<td>{{.TraceId}}</td></p>" +
				"</body></html>"

			if snedData, err := html.Format(htmlText, w); err == nil {
				_ = email.Send("1124938791@qq.com", "系统告警", snedData)
			}
		}
	}
}
