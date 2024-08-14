# Pizza-Co - Learning Go, Gorm, and gRPC

## Overview

This project was an exercise in learning and teaching myself Go, Gorm, gRPC and Hex design architecture.

This project focused around building a backend api for a pizza ordering service.

This micro-services api has 4 services (order_management_service, inventory_service, payment-processing-service, and order processing service). At the moment, on only the first two services are implmented, with inventory-service being connected to a Postgres DB. The other 2 services have dummy "success" data state that always returns successful.

The program is ran by either running my 'send_request_to_order_management_server' go program, and configure the number of requests you want to submit, or you can use a gRPC client to make the request, such as Kreya on mac.

The starting endpoint/grpc request is the 'order_management_service' end point.

AI also have kubernetes deployment scripts along with Dockerfiles and publicly hosted images on [hub.docker.com](https://hub.docker.com/u/mikejacks). See docker images here:

`docker pull mikejacks/pizza-co-servers-inventory-service`

`docker pull mikejacks/pizza-co-servers-order-management-service`

## Installation and running application

1. Clone the repository to a location on your hard drive  
   `git clone https://github.com/mike-jacks/pizza-co.git`

2. Change directory to the pizza-co directory  
   `cd pizza-co`

3. Create a Postgresql database and fill out the '.env.example' to match settings for Postgres. Make sure to remove the '.example' extension

4. Once your database is launched and running you can start the two go microservices  
   `go run order_management_service/main.go`  
   `go run inventory_service/main.go`

5. Assuming you are running 'order_management_service' on localhost port 9000, you will send gRPC requests to it through HTTP/2. You can either use a dedicated app, such as Kreya, to communicate to the api through gRPC.

## How it works

Here is a sample of a request someone could make to the PlaceOrder gRPC protobuf:

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
