# Clean Architecture

## Brief Introduction
Pardon for my bad english as i am not a native english speaker.
Here i would like to share my piece of code for clean architecture with golang.

My essential understanding about the concept ofclean architecture, especially in the context of golang : 
1. Each parts should be independently able to be tested, replaceable, and maintainable.
2. Each layer should be indirectly access by other layer, in which we use interface to achieve it.
3. Each implementation of component of each layer, should be solitary, non accessible from outside of its layer.

I understand that clean code approach sometimes overwhelming for people that try to understand it.
Therefore here i will try to explain it as simple and easy as possible.

Here i will use my own terms, in accordance with the clean architecture. Some of you who may have experience on MVC, 
may find some of my terms are similar to MVC,or the way it works.
My structure may be different from the traditional clean architecture, but i believe it is more flexible and easier to understand for me, or hopefully for you too.



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

## First Form (Part 1):
_establish definition for data and its interface_

I would like to introduce terms "entity".
You can think of entity as a table structure/definition that will be used in the application.
It's similar to model in MVC.

For entity is stored in folder `internal/data/entity`.

```
type eProduct struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

```

e stands for entity, followed by name of entity.

You may also make definition of actual data using tagging, i used gorm tag for this.

You may also define foreign key, entity relationship, etc. 

JSON tag is not used here as it is not designed to be use for response directly.

I also create method named `TableName()` in the entity, it is used to define the table name in the database.

You may ask why stay private? 
For component that is private, you cannot by mistake, use it directly from other layer/package. 
This rule will be apply to all components on any layer.


## First Form (Part 2):
_establish interface for entity_

It is useless if you create something which private/solitary, you need some connectors or introducer for outside of it's package. Here we use interface for this. 

Interface here means the face of the entity, for others to able to use it indirectly.

For interface is stored in folder `internal/data/entity`. 
I store it in the same folder as entity, as it is closely related to entity. 
Other structure split it to different folder, it is up to you.           

```
type IProduct interface {
	TableName() string
}
```

the prefix I means interface, followed by name of entity.

I use smallcase e for entity, and uppercase I for interface.
e stands for entity , i/I stands for it's interface.
Uppercase also means it is public component, accessible from outside of its package.

I specifically use this for in order to understand and separate it exactly so you don't get confused :D.


