# Transaction Service

## Introduction
This is a basic REST service written in Golang, and uses PostgreSQL to fulfill product
the following requirements:
1. Create a transaction record, which can have a parent transaction as well.
2. Get transaction by ID/Type.
3. Get recursive sum of transactions referencing their parent ID.

## Usage
1. You need to create a `env.json` config file under `cloud/` directory. Refer env.json.sample for the JSON sample.
2. Run the server using `go run main.go` or build a binary using `go build -o transaction-service` and execute using `chmod +x transaction-service && ./transaction-service`
3. Generate a JWT using the signing key added in the `env.json` field `jwt_signing_key`. Use this key as a Bearer token for all transaction routes.
4. Create some dummy transactions using `PUT /transactionservice/transactions/:id` or `POST /transactionservice/transactions` route.
5. You can fetch a transaction using `GET /transactionservice/transactions/:id` route or get transaction IDs matching a type using `GET /transactionservice/types/:type` route.
6. You can fetch the sum of all child transactions' amounts using the `PUT /transactionservice/sum/:parent-id` route.
