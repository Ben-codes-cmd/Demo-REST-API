# GO REST API

## Overview
This project demonstrates the functionality of a REST API from the client and server side. 

Upon execution of the program, an HTTP server is created on the localhost accessible via port 8081. This server supports a variety of endpoints for different API requests. 

The program utilizes go routines to concurrently handle requests and provide a "client side" menu to easily interface with a variety of requests that are built and sent programmatically.
Alternatively, one may opt to make requests with an external application such as insomnia or postman. 

As a proof of concept, I built this API to mimic a generic membership system that you might see in a subscription based business such as a gym or streaming service. 

## Supported Requests
This API demonstration supports a variety of operations including CRUD (Create, Read, Update, Delete). Supported functions include:

Base URL: localhost:8081

| Endpoint      | Verb   | Function |
| :------------ | :----- | :-------                                                                                        |
| /             | GET    | The root of the endpoint, currently no response                                                 |
| /member       | POST   | Create a new member (json attributes: string Id, string Name, string Age, string MembershipFee) |
| /member       | SECRET | Displays a secret message to the terminal - Not accessible via command line interface           |
| /member/{id}  | GET    | Get a json object representing the member of the specified id                                   |
| /member/{id}  | DELETE | Delete a member from the member list with the specified id                                      |
| /member/{id}  | PUT    | Update a member from the member list with the specified id                                      |
| /members      | GET    | Get a list of json objects representing each member in the member list                          |
| /members/{id} | PUT    | Deduplicate members sharing the same ID (retains first occurence)                               |

## Future Additions
- Add a UI
- Add data persistence/customized query parameters with SQL database
- Containerize application with docker/docker-compose
