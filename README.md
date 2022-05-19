# merchant-account-management

# Author: chungtran2801@gmail.com

# Setup and run
- This app is running locally.
- To run app please set up your postgres database with this information: 
    + Username: `postgres`
    + Password: `chungtm`
    + Host: `localhost:5432`
    + Database: `postgres`
    + Driver name: `postgres`
    
# Test:
- Run test coverage (coverage > 95%): `go test -coverprofile=profile.out ./...`
- View result: `go tool cover -html=profile.out`
- Generate mocks: `mockery --name=<interface_name>`

# Api


1. Merchant account
   1. Create
      + Url: `/v1/merchant/account?action=new`
      + Method: `post`
      + Request:
        ```json
        {
          "code": "",
          "name": "",
          "user_name": "",
          "password": ""
        }
        ```
      + Response:
        ```json
        {
          "status": "",
          "error": {
            "domain": "",
            "code": 0,
            "message": ""
          }
        }
        ```
   2. Read
      + Url: `v1/merchant/account?action=get&id=<merchant_id>`
      + Method: `post`
      + Response:
        ```json
        {
          "status": "",
          "error": {
            "domain": "",
            "code": 0,
            "message": ""
          },
          "data" : {
            "id": "",
            "merchant_code": "",
            "name": "",
            "is_active": ""
          }
        }
        ```
   3. Update
      + Url: `/v1/merchant/account?action=update`
      + Method: `post`
      + Request:
        ```json
        {
          "merchant_id": "",
          "name": "",
          "password": ""
        }
        ```
      + Response:
        ```json
        {
          "status": "",
          "error": {
            "domain": "",
            "code": 0,
            "message": ""
          }
        }
        ```
   4. Delete
      + Url: `/v1/merchant/account?action=delete&id=<merchant_id>`
      + Method: `post`
      + Response:
        ```json
        {
          "status": "",
          "error": {
            "domain": "",
            "code": 0,
            "message": ""
          }
        }
        ```
  2. Member of account
      1. Create
          + Url: `/v1/merchant/account?action=new`
          + Method: `post`
          + Request:
            ```json
            {
              "email": "",
              "merchant_id": "",
              "name": "",
              "address": "",
              "phone": ""
            }
            ```
          + Response:
            ```json
            {
              "status": "",
              "error": {
                "domain": "",
                "code": 0,
                "message": ""
              }
            }
            ```
      2. Read
          + Url: `v1/merchant/account?source=member&action=get&email=<member_email>`
          + Method: `post`
          + Response:
            ```json
            {
              "status": "",
              "error": {
                "domain": "",
                "code": 0,
                "message": ""
              },
              "data": {
                "merchant_id": "",
                "email": "",
                "name": "",
                "address": "",
                "phone": ""
              }
            }
            ```
      3. Update
          + Url: `/v1/merchant/account?action=update`
          + Method: `post`
          + Request:
            ```json
            {
              "email": "",
              "name": "",
              "address": "",
              "phone": ""
            }
            ```
          + Response:
            ```json
            {
              "status": "",
              "error": {
                "domain": "",
                "code": 0,
                "message": ""
              }
            }
            ```
      4. Delete
          + Url: `/v1/merchant/account?source=member&action=delete&email=<member_email>`
          + Method: `post`
          + Response:
            ```json
            {
              "status": "",
              "error": {
                "domain": "",
                "code": 0,
                "message": ""
              }
            }
            ```
      5. Read by offset and limit
          + Url: `v1/merchant/account?source=member&action=gets&merchant_id=<merchant_id>&offset=2&limit=2`
          + Method: `post`
          + Response:
            ```json
            {
              "status": "",
              "error": {
                "domain": "",
                "code": 0,
                "message": ""
              },
              "data": [
                  {
                    "merchant_id": "",
                    "email": "",
                    "name": "",
                    "address": "",
                    "phone": ""
                   }
                ]
            }
            ```
3. Response model explanation 
      1. Domain is `merchant_management`
      2. Response status can be `success` or `failed`
      3. Errors
        
          |Code|Message|
          |:---|---|
          | -1| router is nil|
          | -2| request is nil|
          | -3| Code is too long|
          | -4| internal error|
          | -5| internal error|
          | -6| password is empty|
          | -7| merchant code is empty|
          | -8| internal error|
          | -9| internal error|
          | -10| internal error|
          | -11| internal error|
          | -12| email existed|
          | -13| check existence failed|
          | -14| merchant Code existed|
          | -15| email is empty|
          | -16| internal error|
          | -17| internal error|
          | -18| internal error|
          | -19| merchant name is empty|
          | -20| username is empty|
          | -21| item not found|
          | -22| parameter is invalid|
          | -23| merchant id is empty|

# Upcoming parts
- Add logger
- Add tracer