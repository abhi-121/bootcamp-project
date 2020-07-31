# Team2-Case-Study-1
Repository for Case Study 1

### To run this project, follow the below steps :
-  Clone this repo : `git clone github.com/shashijangra22/Team2-Case-Study-1`
-  Make sure u have an instance of dynamodb running.
-  Create the tables [if db is empty] : `make tables`
-  Generate Stub Code from proto files : `make protos`
-  Start the Server : `make server`

### To Perform Unit Testing :
- Customer API Testing: run `make c-tests`
- Restaurant API Testing: run `make r-tests`


### EndPoints exposed by the API to deal with Customer and Orders: 

- `POST /login`                 Sign in the Admin to use the API
- `POST /customer`              Add a new customer
- `GET /customers`              Get all customers
- `GET /customer/{cid}`         Get a particular customer  
- `GET /customer{cid}/orders`   Get all orders of a particular customer
- `POST /order`                 Add a new order
- `GET /order/{oid}`            Get a particular order
- `PUT /order`                  Update an existing order
- `PUT /order/item`             Add item to an order
- `DELETE /order{oid}`          Delete an order
- `DELETE /order/{oid}/{item}`  Delete an item from the order
- `POST /order/discount`        Apply discount coupon to an order [need to verify this]

### EndPoints exposed by the API to deal with Restaurants:
  
- `POST /restaurant`                                    Add a new Restaurant
- `GET /restaurants`                                    Get all restaurant
- `GET /restaurant/{rid}`                               Get a particular restaurant
- `POST /restaurant/item`                               Add a new item to a restaurant
- `PUT /restaurant/item`                                Update an item to a restaurant
- `DELETE /restaurant/{rid}/item`                       Delete an item from a restaurant
- `GET /restaurant/{rid}/items`                         Get all items from a restaurant
- `GET /restaurant/{rid}/item/{itemId}`                 Get particular item from a restaurant
- `GET /restaurant/{rid}/items?min={min}&max={max}`     Get items from a restaurant in a price range

### Data Model of the DB

Customer
----

```json
{
    "_id": "<string>",
    "name": "<string>",
    "phone": "<string>",
    "addr": "<string>"
}
```

Order
----

```json
{
    "_id": "<string>",
    "c_id": "<string>",
    "r_id": "<string>",
    "itemline": [
        {
            "name":"<string>",
            "price":"<string>"
        },
        {
            "name":"<string>",
            "price":"<string>"
        }
    ],
    "price":"<string>",
    "discount":"<string>"
}
```

Restaurant
----

```json
{
    "_id": "<string>",
    "name": "<string>",
    "online": "<string>",
    "menu": [
        {
            "name":"<string>",
            "price":"<string>"
        },
        {
            "name":"<string>",
            "price":"<string>"
        }
    ],
    "rating":"<string>",
    "category":"<string>"
}
```