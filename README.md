# Event Flux

A Go-based application that interacts with both **Cassandra** and **ScyllaDB** to manage event data, specifically related to `fraudnetic_events`. This system allows for querying, filtering, and retrieving event data, using **Fiber** for handling HTTP requests and **gRPC** for RPC-based communication. Event Flux supports both databases and makes it easy to switch between them.

**DISCLAMER 1: I skipped unit testing for faster development process.**

## Requirements
- **Go** (version 1.22 or later)
- **Docker** and **Docker Compose** (for running Cassandra and ScyllaDB)
- **Postman** (for API testing)

## Build and Run the Containers
You can start the application along with Cassandra and ScyllaDB using **Docker Compose**:

```bash
docker-compose up --build
```

**NOTE: database containers are fully self-contained and do not require manual migrations or data insertion. All necessary data and schema setup are handled automatically when the containers start using these files: [load-data-cassandra.sh](scripts/load-data-cassandra.sh) and [load-data-scylla.sh](scripts/load-data-scylla.sh)**

This command will:
- Start **Cassandra** on port `9042`.
- Start **ScyllaDB** on port `9050`.
- Launch the **Event Flux** application on port `8080` (by default).

## Postman Collection
I provided a **Postman collection** that includes all the API endpoints you need to test the application. You can find the collection at:
[event-flux.postman_collection.json](event-flux.postman_collection.json)

## Endpoints

1. **GET `/events/:id`**  
   Retrieve an event by its `id`.  
   Example:  
   `GET http://0.0.0.0:8080/events/9cab3f76-331e-11ef-ae33-0242ac150002`

2. **GET `/events`**  
   Retrieve all events in the database.  
   Example:  
   `GET http://0.0.0.0:8080/events`

3. **POST `/events/filter`**  
   Filter events based on specific parameters like `event_name`, and `created_at`.  
   Example:
   `GET http://0.0.0.0:8080/events/filter/?start_date=2024-01-01 00:00:00&end_date=2024-12-31 00:00:00&event_name=registration`

## Environment Variables
The system can be configured using environment variables, defined in the `docker-compose.yml` file:

- **DB_DRIVER_TYPE**: Choose between `cassandra` or `scylla`.
- **CASSANDRA_HOST**: Hostname for Cassandra.
- **SCYLLA_HOST**: Hostname for ScyllaDB.
- **APP_HOST**: The host address where the app will run.
- **APP_PORT**: The port on which the app will run.

**DISCLAMER 2: This is a anti-pattern, ENVs should not be inside docker compose files.**

## Cassandra vs ScyllaDB

After doing some research, I found that both Cassandra and ScyllaDB are pretty similar when it comes to the basics. They’re both highly reliable, distributed NoSQL databases that are great for handling large amounts of data across many nodes. They even use the same Cassandra Query Language (CQL), which means the commands, queries, and interactions in this project are the same no matter which one you choose. From a development perspective, there’s really not much difference.

ScyllaDB was built with performance in mind. It's written in C++, which allows it to make better use of modern multi-core processors, unlike Cassandra, which is based on Java. Because of this, ScyllaDB can handle more requests per second with lower latency, especially when things get really busy. It’s designed to scale efficiently without the overhead that sometimes slows Cassandra down.

For this project, Event Flux, it really doesn’t matter much which one you use. The API endpoints, data handling, and everything else are the same whether you choose Cassandra or ScyllaDB. So, you can pick whichever database fits your needs—if you’re dealing with a high-traffic environment, ScyllaDB might be the better choice, but for most use cases, Cassandra will work just fine too.

I chose `PRIMARY KEY ((event_name), created_at) WITH CLUSTERING ORDER BY (created_at DESC)` because it allows us to efficiently query events by event type, which provides a good balance in terms of cardinality. If we used columns like `id`, `dates`, or `user_id` as the partition key, it would lead to too many partitions in the database since these columns have high cardinality (a large variety of unique values). On the other hand, columns like `incognito` and `processed` have very low cardinality, meaning they would lead to all the data being stored in just a few partitions, or even a single one, which could overload specific nodes.

By using `event_name`, which has moderate cardinality, we can distribute the data more evenly across the cluster, ensuring that no single node is overwhelmed and queries remain efficient. This also works well with time-based queries, as `created_at` is used as the clustering key to order the events within each partition by date, making it easy to retrieve recent events quickly.


