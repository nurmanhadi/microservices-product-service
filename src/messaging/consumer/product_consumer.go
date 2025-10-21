package consumer

import (
	"encoding/json"
	"product-service/pkg/env"
	"product-service/src/dto"
	"product-service/src/internal/service"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type ProductConsumer interface {
	QueueProduct()
}

type productConsumer struct {
	logger         *logrus.Logger
	ch             *amqp091.Channel
	productService service.ProductService
}

func NewProductConsumer(logger *logrus.Logger, ch *amqp091.Channel, productService service.ProductService) ProductConsumer {
	return &productConsumer{
		ch:             ch,
		logger:         logger,
		productService: productService,
	}
}
func (c *productConsumer) QueueProduct() {
	msgs, err := c.ch.Consume(
		env.CONF.Broker.Queue.Product,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.logger.WithError(err).Error("failed to consume")
	}
	go func() {
		for d := range msgs {
			var datas []dto.OrderConsumerResponse
			err := json.Unmarshal(d.Body, &datas)
			if err != nil {
				c.logger.WithError(err).Error("failed to unmarshal json")
			}
			err = c.productService.UpdateProductBulkQuantityByID(datas)
			if err != nil {
				c.logger.WithError(err).Error("failed to update bulk product quantity by id")
			}
		}
	}()
}
