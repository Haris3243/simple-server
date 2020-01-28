# Simple APIServer in Golang

This is the first basic example of a **server** in which a single API is defined which says the hello by obtaining value from environment variable **MYNAME** and concatenates it with **Hello**. for example if value of MYNAME in environment variable is Haris then when a call is made to this API it returns **Hello Haris**

to start with this basic implementation of a server in golang lets follow these steps after cloning the repository
## Compilation 
- Go to the folder where you have clone this repository.
- open terminal and run the following command 
    ```	
        go build ./src/main.go

    ```
This will generates an executable in folder by name ***"main"***

##  Building Docker Image
to build docker image of the server application run the following command:
    ```
        docker build -t server:latest .
    
    ```
Above command will create and image and you can verify it by getting your image list from command :
    ```
        docker images
    
    ```	
