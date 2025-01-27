# Home Assignment API


# ***Note***
what i observed was that the  result(response Json) when we request from the endpoint according to the test the distance here is 177 , i am not sure if this is right or wrong, because i am getting 0 and that is probably correct.
according to the algorithm
```bash
{
"total_price": 1190,
"small_order_surcharge": 0,
"cart_value": 1000,
"delivery": {
"fee": 190,
"distance": 177
}
}
```
the distance here is 177 , i am not sure if this is right or wrong, because i am getting 0 and that is probably correct. 



DOPC fetches essential venue-related data from the **Home Assignment API**. The following endpoints provide static and dynamic data for a venue:

 - **Static Information:** `/v1/venues/<VENUE SLUG>/static`
 - **Dynamic Information:** `/v1/venues/<VENUE SLUG>/dynamic`

### Relevant Fields:

#### Static Endpoint:
- `venue_raw -> location -> coordinates`: Venue location (longitude, latitude).

#### Dynamic Endpoint:
- `venue_raw -> delivery_specs -> order_minimum_no_surcharge`: Minimum cart value to avoid small order surcharge.
- `venue_raw -> delivery_specs -> delivery_pricing -> base_price`: Base price for the delivery fee.
- `venue_raw -> delivery_specs -> delivery_pricing -> distance_ranges`: Rules for calculating delivery fee based on distance.

---

# How to Run the Application

### Prerequisites:
- Go 1.19 or later installed on your system.
- Install the **Gin Web Framework**.

### Install Dependencies:

```bash
go get -u github.com/gin-gonic/gin
```
### Build the Application:
Navigate to the cmd/server directory and build the Go binary:
```bash
go build -o wolt ./cmd/server
```
### Run the Application:
Run the application, which will start a server on ` http://localhost:8000:`
```bash
./wolt
```
You can now test the `/api/v1/delivery-order-price` endpoint using the example reques

### Notes
- Ensure you follow the snake_case convention in query parameters and response fields.
- All money-related values are in the lowest denomination of the local currency (e.g., cents for Euro countries, øre for Sweden, and yen for Japan).
- This project is intended for learning purposes and should be deployed with appropriate optimizations for a production-grade system.




