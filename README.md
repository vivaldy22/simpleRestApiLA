# Simple REST API
I made this project for Link Aja Assignment<br>
This project was made using:
 1. <b>Go</b> (main language of this project)
 2. <b>MySQL</b> (main database)
 3. <b>sqlite3</b> (temporary database for unit testing)

## Requirements
 1. Go (This was made in go1.15.5)
 2. MySQL / MariaDB

## Preparation
 1. Import the database provided in database/db.sql file
 2. **Important** If you're executing with docker, you don't need to change anything in .env file. You can just skip step 3 and 4
 3. Inside .env file, there are configurations needed for the api to run,
  change them according to your system. The most important variable that 
  you have to change is:  
       - DB_USER
       - DB_PASSWORD
       - API_HOST (optional)
       - API_PORT (optional) 
 4. (Optional) Install ```make``` program for your system

## Execution with Docker
 1. Launch your terminal
 2. Go to the directory of this project
 3. Type ```docker-compose up --build```

## Execution without Docker
 1. Launch your terminal
 2. Go to the directory of this project
 3. Type ```make dev``` in the terminal. If your terminal says make command not found, 
 make sure to install it first on your system
 4. If everything is right, the output should be something like this
     ```
        $ make dev
        go run main.go
        2020/12/05 05:36:36 [MySQL] Database connected, simple_rest_la
        2020/12/05 05:36:36 Starting Web Server at localhost port: 8080
        Available endpoints:
        /account/{account_number} [GET]
        /account/{from_account_number}/transfer [POST]
     ``` <br>

# Endpoints
**Show Balance**
----
  Returns json data about a single account.

* **URL**

  `/account/{account_number}`

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `account_number=[string]`

* **Data Params**

  None

* **Success Response:**
    * **Code:** 200 <br />
    * **Content:**
  ``` json
    {
      "account_number": "555001",
      "customer_number": "1001",
      "balance": "10000"
    }
  ```
 
* **Error Response:**
    * **Code:** 404 <br />
    * **Content:**
  ``` json
      {
          "message": "Data not found",
          "error_code": 404
      }
  ```
  
**Transfer Money**
----
  Transfer money from an account to another account.

* **URL**

  `/account/{from_account_number}/transfer`

* **Method:**

  `POST`
  
*  **URL Params**

   **Required:**
 
   `from_account_number=[string]`

* **Data Params**

  ``` json
      {
          "to_account_number": "555002",
          "amount": "100"
      }
    ```

* **Success Response:**
    * **Code:** 201 <br />
    * **Content:**
    None.
 
* **Error Response:**
    * **Code:** 400 <br />
    * **Content:**
    ``` json
      {
          "message": "Decoding json failed",
          "error_code": 400
      }
    ```
  OR
    * There are multiple error message with code 404, this is one of them
    * **Code:** 404 <br />
    * **Content:**
    ``` json
        {
            "message": "Your Balance is not enough",
            "error_code": 404
        }
    ```
