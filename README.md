In week 4 I moved away from "Concrete Implementation" to "Interface-Based Design"
The ProductService no longer depends on a specific MongoDB struct. Instead, it depends on a ProductRepository interface.
I implemented a Global Security Gatekeeper using the Middleware pattern.
Before a request reaches the "Create" or "Delete" handlers, it must pass through the AuthMiddleware.
The middleware intercepts the HTTP request, extracts the X-Admin-Key from the header, and compares it against the ADMIN_KEY stored in the system's environment variables (.env).
If the keys do not match, the request is rejected with a 403 Forbidden status before any business logic is executed or any database resources are consumed
I added logic to ensure that:
Prices cannot be negative (preventing financial data corruption).
Product Names cannot be empty (ensuring data integrity in the catalog).
If a user tries to save a product with a negative price, the Service returns a custom error, stopping the process before it hits the database.
If MongoDB takes too long to respond, the context logic will automatically kill the request to prevent the server from locking up.
