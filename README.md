## Simple ecommerce checkout using microservice architecture and Golang
&nbsp;

### How to run
At each folder there is a go file. Just enter in each folder and run "go run filename.go"
&nbsp;

At checkout you just need to fill de Credit Card Number (Número), Cupom and click Comprar button to purchase
- **valid credit card number** = *1* (approved status - declined otherwise)
- **valid cupom** = *abc* (valid status - invalid otherwise)

&nbsp;
### How to test retrial strategy
- stop the payment or coupon microservice
- try a purchase
- watch retrials in terminal 
- restart the service before 5 retrials and the process will be concluded