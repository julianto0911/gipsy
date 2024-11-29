# Clean Architecture

## Brief Introduction
Pardon for my bad english as i am not a native english speaker.
Here i would like to share my piece of code for clean architecture with golang.

My essential understanding about the concept of clean architecture, especially in the context of golang : 
1. Each parts should be independently able to be tested, replaceable, and maintainable.
2. Each layer should be indirectly access by other layer, in which we use interface to achieve it. 

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
│ ├── 📂 repository
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
_establish definition for data/table_

I would like to introduce terms "entity".
You can think of entity as a table structure/definition that will be used in the application.
It's similar to model in MVC.

For entity is stored in folder `internal/repository`.

```
type EProduct struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

```

E stands for entity, followed by name of entity.

You may also make definition of actual data using tagging, i used gorm tag for this.

You may also define foreign key, entity relationship, etc. 

JSON tag is not used here as it is not designed to be use for response directly.

I also create method named `TableName()` in the entity, it is used to define the table name in the database.
```
type EProduct struct {
	ID   int64  `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (c *EProduct) TableName() string {
	return "products"
}

``` 


## First Form (Part 2):
_establish interface for entity_

After defining the entity, some source will interface it.
In our case, to maintain simplicity, we will skip it for the moment.

I will rather explain why we need interface, and how it works.

Interface enable us to replace the entity with other entity, without changing the code.
This is extremely useful especially when we do testing.

Example :
```
type productA struct{
    TableName() string
    Color string
}

type productB struct{
    TableName() string
    Size string
}

type IProduct interface{
    TableName() string
}

//we can return productA since productA implements IProduct interface
func returnA() IProduct{
    return &productA{}
}

//we can return productB since productB implements IProduct interface
func returnB() IProduct{
    return &productB{}
}

```

### First Form (Part 3):
_establish repository for entity_

In repository level, things get interesting.
```
type rProduct struct {
	db *gorm.DB
}
```

Here we have db property inside the repository. It is used to connect and access the database.
We can say that in repository, we do the actual database operation. 
Same like model in MVC, except in clean architecture, the operation and the object of the table are separated.
Smallcase r stands for repository, _i would like to make it private_.

Next, lets create method called `Create` inside the repository.

```
func (r *productRepo) Create(name string) (entity.EProduct, error) {
	product := EProduct{
        Name: name,
    }

	err := r.db.Create(&product).Error

	return &product, err
}
```
The Create method return an entity, and error.

You can add more method inside the repository, like `Get`, `Update`, `Delete`, etc, according to your need.

Ok, from here you already have a basic repository : 
1. a repository uses entity to get or push data into or from database, using database connection (db property).
2. a repository should contains only create,update,delete, get method, in the simplest way. It should not contain business logic.

## First Form (Part 4):
_establish interface for repository_

After defining the repository, we need to define the interface for it.

```
type RProduct interface {
	Create(name string) (*entity.EProduct, error)
}

```

Note that there is rProduct and RProduct. RProduct is the interface, and rProduct is the repository.

Then we create a function to return the repository, and inject the database connection into it.
```
// repository for product
func NewProductRepo(db *gorm.DB) RProduct {
	repo := &rProduct{
		db: db,
	}

	return repo
}
```

You will notice that we return an interface on the function.

Ok , before we continue, let's summarize what we got : 
1. entity : table definition
2. repository : actual operation of data, using database connection.
3. repository process and return entity.
4. repository's interface : to make the repository interchangeable on other level.
5. we create function that returns repository's interface, but what we return is the implementation of repository itself.

After First Form finish, you have the data/entity/repository layer at hand. We will continue to the usecase layer.


## Second Form (Part 1):
_establish usecase_

Simplified : Usecase is to Clean, as Controller is to MVC.

Ok, the usecase is the business, logic layer of the application.
What is business logic? 
Example : 
1. A Product must have at least 3 character name.
2. Product Brand ID must be defined first in Brand table.
3. After create product, send email notification to supervisor.
4. etc...

Here in usecase , you need to define the rules, and flow of the application.
We will do it in a very simple way, just to get you familiar with the concept.

```
type productUC struct {
	product repository.RProduct
}

```
The same principle for productUC to stay private, it will be represent later by ProductUC interface.

Then we create function to receive the input of the product, name it InputProduct, stores in `usecase/product/parameter.go`
```
type InputProduct struct {
	Name string `json:"name"`
}
```

You can see we added json tag, later this InputProduct struct will be used for http request in adaptor layer.

Next, we create function to process the input, and return the output.
```
func (uc *productUC) Create(input InputProduct) (*repository.EProduct, error) {
	return uc.product.Create(input.Name)
}
```

The principle of Clean Architecture, you can use/pass/reference elements  from inside the layer but not the other way around.

## Second Form (Part 2):
_establish interface for usecase_

Next we will interface the productUC, and create function to return the interface.

```
type ProductUC interface {
	Create(input InputProduct) (*repository.EProduct, error)
}
```

```
func NewProductUseCase(product repository.RProduct) ProductUC {
	return &productUC{
		product: product,
	}
}
```

Here you may find resemblance of how usecase formed as in the repository layer. Yes it is practically the same.


Ok, before we continue, let's summarize what we got : 
1. entity : table definition
2. repository : actual operation of data, using database connection.
3. repository process and return entity.
4. repository's interface : to make the repository interchangeable on other level.
5. we create function that returns repository's interface, but what we return is the implementation of repository itself.
6. usecase : define the business logic, process the input, and return the output.
7. usecase's interface : to make the usecase interchangeable on other level.
8. we create function that returns usecase's interface, but what we return is the implementation of usecase itself.

After Second Form finish, you have the repository at first layer,usecase at secondlayer. We will continue to the adaptor layer.

## Third Form (Part 1):
_establish adaptor_

Adaptor is the layer that communicate with the external world, in this case, http request.
```
type productAdaptor struct {
	ucProduct usecase.ProductUC
}
```

We create function to receive input and pass it to usecase.
```
func (adp *productAdaptor) Create(c *gin.Context) {
	input := usecase.InputProduct{}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := adp.ucProduct.Create(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
```

Here i use gin framework, but you can use any framework that you want.      

## Third Form (Part 2):
_establish interface for adaptor_

```
type ProductAdaptor interface {
	Create(c *gin.Context)
}
```

Last we create function to return the adaptor's interface, and inject the usecase into it.
```
func NewProductAdaptor(ucProduct usecase.ProductUC) ProductAdaptor {
	return &productAdaptor{
		ucProduct: ucProduct,
	}
}
```

Ok, before we continue, let's summarize what we got : 
1. entity : table definition
2. repository : actual operation of data, using database connection.
3. repository process and return entity.
4. repository's interface : to make the repository interchangeable on other level.
5. we create function that returns repository's interface, but what we return is the implementation of repository itself.
6. usecase : define the business logic, process the input, and return the output.
7. usecase's interface : to make the usecase interchangeable on other level.
8. we create function that returns usecase's interface, but what we return is the implementation of usecase itself.
9. adaptor : communicate with the external world, in this case, http request.
10. adaptor's interface : to make the adaptor interchangeable on other level.
11. we create function that returns adaptor's interface, but what we return is the implementation of adaptor itself.

An addition to the codebase : 
1. 1 adaptor may have 1 or more usecase, and 1 usecase may be used by 1 or more adaptor.
2. 1 usecase may have 1 or more repository, and 1 repository may be used by 1 or more usecase.
3. 1 repository may have 1 or more entity.

After Third Form finish, you have the repository at first layer,usecase at secondlayer, and adaptor at third layer.

The last part is the wiring.

## Wiring
Each elements in layers doesn't know about each other, that's why in golang there are terms dependency injection.
In my own terms, i use word "wire" or "wiring" , just like what we do in electronics where we link 1 component to other through wiring.

```
func Wiring(db *gorm.DB) *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(
		gin.Recovery(),
		middleware.RequestID(),
	)

	api := router.Group("/api/v1")
	wireProduct(api, db)

	return router
}



func wireProduct(router *gin.RouterGroup, db *gorm.DB) {
	rProduct := repository.NewProductRepo(db)
	ucProduct := ucproduct.NewProductUseCase(rProduct)
	adpProduct := adaptor.NewProductAdaptor(ucProduct)

	router.POST("/create", adpProduct.Create)
}

```

Explanation : 
1. We initiate the database connection.
2. We create a function named `wireProduct` that will wire the product layer.
3. In `wireProduct` function, we create repository, usecase, and adaptor.
4. We tell gin to use `wireProduct` function when receiving `/api/v1/create` request.

For basic programming purpose, this clean approach may not seem necessary, often make things complicated.
But for moving forward to software engineering, clean approach is a must.

Next topic i will cover on how to use this clean architecture to simplify your life when doing unit testing.

If you have any input or suggestion, please feel free to contact me. 
Email : julianto@lumoshive.com

Thank you.