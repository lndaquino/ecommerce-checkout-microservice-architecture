## Simple ecommerce checkout using microservice architecture, Golang, RabbitMQ and Docker
&nbsp;

### How to run
Initialize docker and run "docker-compose up -d" in project root directory.
At each folder there is a go file. Just enter in each folder and run "go run filename.go".
&nbsp;

At checkout you just need to fill de Credit Card Number (Número), Cupom and click Comprar button to purchase
- **valid credit card number** = *1* (approved status - declined otherwise)
- **valid cupom** = *abc* (valid status - invalid otherwise)

&nbsp;
### How to test
- stop the coupon microservice
- try a purchase
- watch queue retrials in payment microservice terminal 
- restart the coupon microservice and the process will be concluded