Week 3 description:
There is transition from an in-memory storage system to a professional Layered Architecture with a persistent MongoDB database
Implemented MongoDB as the primary data store using the official Go Driver instead of handling all IDs manually.
Isolated all database queries (Find, Insert, Update, Delete) from the business logic.
Refactored HTTP handlers to focus strictly on request/response parsing.
Added a docker-compose.yml file to easily spin up a local MongoDB instance.
implemented Create, Read (List & Single), Update, and Delete endpoints.
Updated the Product model with BSON tags for MongoDB compatibility.
POST /products: creates a document in the inventory_db.products collection.
GET /products/list: retrieves all products from MongoDB.
