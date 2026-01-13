## Database Connection

To connect to the database using `psql`:

```bash
psql -h localhost -U postgres -d exercises
```

Alternatively, using Docker Compose:

```bash
docker compose exec -it db psql -U postgres -d exercises
```

Password: `password`


## Chapter I Generating mock data

```CREATE TABLE trades(
      	id bigint,
      	buyer_id integer,
	symbol text,
	order_quantity integer,
	bid_price numeric(5,2),
	order_time timestamp
)
```





Commands

\d

show tables


