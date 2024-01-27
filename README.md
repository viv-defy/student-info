# student-info

This repo is an exercise as a part of one2n's [Site Reliability Engineering BootCamp](https://playbook.one2n.in/sre-bootcamp). Here, we build a CRUD REST API to perform the following operations

- Add a new student
- Get all students
- Get a student with an Id
- Update existing student information
- Delete a student record

## Local Setup

- Clone the repository
  ```
  git clone git@github.com:viv-defy/student-info.git
  ```
- Database Setup
  - You also need Postgres setup locally to run the code. You can follow the instructions given [here](https://www.postgresql.org/download/) to setup a local version of postgres
  - Login to postgres from terminal/command-line
  - Create a new db "students"
    ```
    create database "students"
    ```
  - Add a new user
    ```
    create user "gorm" with password 'gorm';
    ```
    You can change/customise the username and password. Make sure to use the same username and password in the .env file(mentioned later)
- Copy the contents of .env.example into .env file. Enter the appropriate values of the variables
- Run the make file
  ```
  make build
  ```
- Run the executable
  ```
  ./student_info
  ```
  You now have the server up and running!!!

## Using the APIs

You can call the API end points using curl, Postman or any other method of your choice. A sample collection of API calls has been added in the repo. You can follow the below instructions to use the the collection

- Download [Postman](https://www.postman.com/downloads/) onto your local system. In case you have VS code, you can download the the Postman extension.
- Create a free Postman account
- Import the Postman collection `students-info.json`, present in the root-dir of this repo

You can now use and modify the collection according to your needs.
