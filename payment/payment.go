package main

import (
	"log"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/wesleywillians/go-rabbitmq/queue"
	"github.com/streadway/amqp"
)

type Order struct {
	ID uuid.UUID
	Coupon string
	CcNumber string
}

type Result struct {
	Status string
}

func NewOrder() Order {
	return Order{ID: uuid.NewV4()}
}

const (
	InvalidCoupon = "invalid"
	ValidCoupon = "valid"
	ConnectionError = "connection error"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func main() {
	messageChannel := make(chan amqp.Delivery)
	
	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	for msg := range messageChannel {
		process(msg)
	}
}

func process(msg amqp.Delivery) {
	order := NewOrder()
	json.Unmarshal(msg.Body, &order)

	resultCoupon := makeHttpCall("http://localhost:9092", order.Coupon)

	switch resultCoupon.Status {
	case InvalidCoupon:
		log.Println("Order: ", order.ID, ": invalid coupon!")

	case ConnectionError:
		msg.Reject(false)
		log.Println("Order: ", order.ID, ": couldn´t process!")

	case ValidCoupon:
		log.Println("Order: ", order.ID, ": processed!")

	}
}

func makeHttpCall(urlMicroservice string, coupon string) Result {
	values := url.Values{}
	values.Add("coupon", coupon)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 0

	res, err := retryClient.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: ConnectionError}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Erro no processamento do resultado!")
	}

	result := Result{}
	json.Unmarshal(data, &result)

	return result
}
