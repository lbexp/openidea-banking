# openidea-banking

This is a mini Banking app API for the 3rd project submission in OpenIdea ProjectSprint initiated by [@nandapagi](https://twitter.com/nandapagi).

This project repository includes:
- Banking app API
- Database Migration
- Prometheus & Grafana setup

## Requirements

- Go v1.22.1
- PostgreSQL
- Docker

## Installation

### How to run locally (dev mode)

- Set Env Variables:
  - DB_NAME
  - DB_USERNAME
  - DB_PASSWORD
  - DB_HOST
  - DB_PORT
  - JWT_SECRET
  - BCRYPT_SALT
  - S3_ID
  - S3_SECRET_KEY
  - S3_BUCKET_NAME
  - S3_REGION
  - DB_PARAMS

- Migrate database
  ```
  Make migrateup
  ```

- Run the API service
  ```
  Make run
  ```

## List of APIs

### Authentication & Authorization

- POST /v1/user/register
  ```js
  // Example request
  {
    "email": "email",
    "name": "namadepan namabelakang",
    "password": ""
  }
  ```

  ```js
  // Example response
  {
    "message": "User registered successfully"
    "data": {
      "email": "email@email.com", 
      "name": "namadepan namabelakang", 
      "accessToken": "qwertyuiopasdfghjklzxcvbnm"
    }
  }
  ```
- POST /v1/user/login
  ```js
  // Example request
  {
    "email": "email",
    "password": ""
  }
  ```

  ```js
  // Example response
  {
    "message": "User logged successfully"
    "data": {
      "email": "email@email.com",
      "name": "namadepan namabelakang", 
      "accessToken": "qwertyuiopasdfghjklzxcvbnm"
    }
  }
  ```

### Balance

- POST /v1/balance
  ```js
  // Example request
  {
    "senderBankAccountNumber": "",
    "senderBankName": "",
    "addedBalance": 1,
    "currency":"USD",
    "transferProofImg": ""
  }
  ```
- GET /v1/balance
  ```js
  // Example response
  {
    "message": "success",
    "data": [
      {
        "balance": 1,
        "currency": "USD"
      },
      {
        "balance": 1,
        "currency": "IDR"
      },
    ]
  }
  ```
- GET /v1/balance/history
  ```js
  // Example response
  {
    "message": "success",
    "data": [
      {
        "transactionId":"",
        "balance":1,
        "currency":"",
        "transferProofImg": "",
        "createdAt": 1582605077000,
        "source": {
          "bankAccountNumber":"",
          "bankName":""
        }
      }
    ],
    "meta": {
      "limit":10,
      "offset":0,
      "total":100
    }
  }
  ```

### Transaction

- POST /v1/transaction
  ```js
  // Example request
  {
    "recipientBankAccountNumber": "",
	  "recipientBankName": "",
	  "fromCurrency":"",
	  "balances":1
  }
  ```

### Image upload

- POST /v1/image
  ```js
  // Example request
  {
    "message":"File uploaded sucessfully",
    "data" : {
      "imageUrl":"https://awss3.d87801e9-fcfc-42a8-963b-fe86d895b51a.jpeg"
    }
  }
  ```
