# user-management

## deploy

1. create file log at "/var/log/efishery-test/efishery-test.log"
  ```bash
sudo mkdir /var/log/efishery-test && sudo chmod -R 777 /var/log/efishery-test
```
2. run file with command
  ```bash
// auth app
go run cmd/app/efishery-test/auth.go
// fetch app
go run cmd/app/efishery-test/fetch.go
```
3. Run mysql migration from sql/init.sql
4. visit url http://localhost:8081 (auth app) & http://localhost:8082 (fetch app)

5. auth app:
   1. register
       - url : http://localhost:8081/register
       - method : POST
       - payload json:
       ```bash
       {
           "username" : "bana1234",
           "phone" : "0811111111"
           "role" : "admin"
       }
       ```
   2. login:
       - url : http://localhost:8081/login
       - method : POST
       - payload json:
      ```bash
      {
          "phone" : "0811111111",
          "password" : "XVlB"
      }
      ```  
   3. auth token - claim data:
       - url : http://localhost:8081/auth
       - method : GET
       - headers
      ```bash
      {
          "Authorization" : "Bearer {token}"
      }
      ```

5. fetch app:
    1. fetch product
        - url : http://localhost:8082/fetch-product
        - method : GET
        - headers:
        ```bash
        {
            "Authorization" : "Bearer {token}"
        }
        ```
    2. fetch product compiled
        - url : http://localhost:8082/fetch-product-compiled
        - method : GET
        - headers:
       ```bash
       {
           "Authorization" : "Bearer {token}"
       }
       ```  
    3. auth token - claim data:
        - url : http://localhost:8082/auth
        - method : GET
        - headers
       ```bash
       {
           "Authorization" : "Bearer {token}"
       }
       ```