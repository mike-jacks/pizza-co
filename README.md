# Pizza-Co - Learning Go, Gorm, and gRPC

## Overview

This project was an exercise in learning and teaching myself Go, Gorm, gRPC and Hex design architecture.

This project focused around building a backend api for a pizza ordering service.

This micro-services api has 4 services (order-management-service, inventory-service, payment-processing-service, and order-processing-service). At the moment, only the first two services are implemented, with inventory-service being connected to a Postgres DB. The other two services each have a dummy "success" response that are always returned.

The program is ran by either running my 'send_request_to_order_management_server' go program, and configuring the number of requests you want to submit, or you can use a gRPC client to make the request, such as Kreya on Mac OS.

The starting endpoint/grpc request is the 'order_management_service' endpoint.

There are kubernetes deployment scripts along with Dockerfiles and publicly hosted images on [hub.docker.com](https://hub.docker.com/u/mikejacks). See docker images here:

`docker pull mikejacks/pizza-co-servers-inventory-service`

`docker pull mikejacks/pizza-co-servers-order-management-service`

## Installation and running application

1. Clone the repository

   ```bash
   git clone https://github.com/mike-jacks/pizza-co.git
   ```

2. Change directory to the `pizza-co` directory

   ```bash
   cd pizza-co
   ```

3. Create a Postgresql database and fill out the '.env.example' to match settings for Postgres. Make sure to remove the '.example' extension

4. Once your database is launched and running you can start the two go microservices

   ```bash
   go run order_management_service/main.go
   go run inventory_service/main.go
   ```

5. Assuming you are running 'order_management_service' on localhost port 9000, you will send gRPC requests to it through HTTP/2. You can either use a dedicated app such as Kreya to communicate to the api through gRPC or the provided go app 'send_request_to_order_management_server'.

## How it works

Here is a sample request someone could make to the PlaceOrder gRPC protobuf:

```json
{
  "customerInfo": {
    "id": "b51b6a14-7b5c-4d0e-a023-f71d2d918fbb",
    "firstName": "Alan",
    "lastName": "Smithy",
    "emailAddress": "sample@email.com",
    "deliveryAddress": {
      "houseNumber": "555",
      "streetName": "Fake Street",
      "aptNumber": "",
      "city": "Fake City",
      "state": "UT",
      "zipCode": "84791"
    },
    "phoneNumber": {
      "number": "435-555-1234",
      "type": "WORK"
    }
  },
  "pizzas": [
    {
      "toppings": ["PEPPERONI", "ONIONS"],
      "size": "EXTRA_LARGE",
      "crustType": "NEW_YORK",
      "extraOptions": ["EXTRA_CHEESE", "BBQ_SAUCE"],
      "quantity": 2
    }
  ],
  "paymentMethod": {
    "paymentType": "CREDIT_CARD",
    "paymentTimeframe": "PREPAID",
    "totalOrderAmount": "23.45"
  }
}
```

Here is the sample stream responses:

```json
// Response 1:
{
  "orderId": "ORDER24804568",
  "status": "RECEIVED",
  "message": "Your order has been received"
}

// Response 2:
{
  "orderId": "ORDER24804568",
  "status": "CHECKING_INVENTORY",
  "message": "Currently Checking Inventory....standby..."
}

// Response 3:
{
  "orderId": "ORDER24804568",
  "status": "CHECKING_INVENTORY",
  "message": "Your order includes the following toppings:\n2x PEPPERONI, 2x ONIONS,.\n\nYour order includes the following crust types:\n2x NEW_YORK.\n\nYour order includes the following crust sizes:\n2x EXTRA_LARGE. All items in your order are in inventory and available to order!\n"
}

// Response 4:
{
  "orderId": "ORDER24804568",
  "status": "PROCESSING_PAYMENT",
  "message": "Your payment successfully went through. Thank you!"
}

// Response 5:
{
  "orderId": "ORDER24804568",
  "status": "PROCESSING_ORDER",
  "message": "Order Processing complete!"
}

// Response 6:
{
  "orderId": "ORDER24804568",
  "status": "COMPLETE",
  "message": "Order is ready for pickup or delivery."
}
```

Example logs from the 'order-management-service' server:

```bash
[order-management-service-c847cdf89-kwv9c] 2024/08/14 07:13:09 Received order for customer: Alan Smithy
[order-management-service-c847cdf89-kwv9c] 2024/08/14 07:13:13 Checking inventory
[order-management-service-c847cdf89-kwv9c] 2024/08/14 07:13:26 Begin Processing Payment
[order-management-service-c847cdf89-kwv9c] 2024/08/14 07:13:30 Begin processiong order.
[order-management-service-c847cdf89-kwv9c] 2024/08/14 07:13:32 Order Process complete, sending final message
```

Example logs from the 'inventory-service' server:

```bash
[inventory-service-594ddbf8fc-jx9pm] 2024/08/14 07:11:31 Checking Inventory...
[inventory-service-594ddbf8fc-jx9pm] 2024/08/14 07:11:34 Order requesting toppings: [PEPPERONI ONIONS]
[inventory-service-594ddbf8fc-jx9pm] 2024/08/14 07:11:35 Order requesting crust types: [NEW_YORK]
[inventory-service-594ddbf8fc-jx9pm] 2024/08/14 07:11:36 Order requesting crust sizes: [EXTRA_LARGE]
[inventory-service-594ddbf8fc-jx9pm] 2024/08/14 07:11:37 Inventory Check Complete!
```

## Demo

![Pizza-Co-Demo](images/Pizza-co-demo.gif)
