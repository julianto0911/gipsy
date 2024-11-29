# Clean Architecture

##Brief Introduction
Pardon for my bad english as i am not a native english speaker.
Here i would like to share my piece of code for clean architecture with golang.

My essential understanding about the concept ofclean architecture, especially in the context of golang : 
1. Each parts should be independently able to be tested, replaceable, and maintainable.
2. Each layer should be indirectly access by other layer, in which we use interface to achieve it.

I understand that clean code approach sometimes overwhelming for people that try to understand it.
Therefore here i will try to explain it as simple and easy as possible.

Here i will use my own terms, in accordance with the clean architecture. Some of you who mayhave experience on MVC, 
may find some of my terms are similar to MVC,or the way it works.



## Foundation 
Standard Layers :
1. Adaptor/Handler
    - This layer is responsible for handling the request and response.
    - It is the entry point of the application.
    - It is the only layer that can access the external world (http request, db, etc).

2. Use Case/Controller
    - This layer is responsible for processing the request and response.
    - It is the only layer that can access the internal world (business logic).
    - It is the only layer that can access the external world (http request, db, etc).

3. Data/Repository/Model
    - This layer is responsible for handling the data.
    - It is the only layer that can access the database.
    - It is the only layer that can access the external world (http request, db, etc).
    - This layer is responsible for handling the data.

## Folder Structure : 
```
project
├── 📂 cmd
│   └── 📂 server
├── 📂 internal
│ ├── 📂 adaptor
│ ├── 📂 data
│ ├── 📂 usecase
│ ├── 📂 wire
├── 📂 pkg
│ ├── 📂 utils
│ ├── 📂 middleware
│ └── 📂 response
|-- .env
|-- go.mod
|-- go.sum
```