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

### Using `generate_series`

The `generate_series` function is essential for generating large sets of mock data. It produces a set of values from a start point to an end point, with an optional step.

**Syntax:**
`generate_series(start, stop, step)`

*   **start**: The value to start the series at.
*   **stop**: The value to stop the series (inclusive).
*   **step**: (Optional) The amount to increment by each step. Defaults to 1.

**Examples:**

1.  **Generate a sequence of numbers:**
    ```sql
    SELECT * FROM generate_series(1, 5);
    ```
    *Output: 1, 2, 3, 4, 5*

2.  **Generate numbers with a step:**
    ```sql
    SELECT * FROM generate_series(1, 10, 2);
    ```
    *Output: 1, 3, 5, 7, 9*

3.  **Generate a series of timestamps:**
    ```sql
    SELECT generate_series(
        '2026-01-01 00:00'::timestamp,
        '2026-01-01 12:00'::timestamp,
        '3 hours'
    );
    ```
    *Output: 2026-01-01 00:00, 2026-01-01 03:00, 2026-01-01 06:00, 2026-01-01 09:00, 2026-01-01 12:00*

### Generating Trades Data

You can combine `generate_series` with `random()` to populate the `trades` table with thousands of rows of realistic-looking data.

```sql
INSERT INTO trades (id, buyer_id, symbol, order_quantity, bid_price, order_time)
SELECT 
    id,
    floor(random() * 100 + 1)::int,                          -- Random buyer_id between 1 and 100
    (ARRAY['AAPL', 'GOOGL', 'MSFT', 'AMZN', 'TSLA'])[floor(random() * 5 + 1)], -- Random symbol
    floor(random() * 1000 + 1)::int,                        -- Random quantity between 1 and 1000
    (random() * 500 + 10)::numeric(5,2),                    -- Random price between 10.00 and 510.00
    now() - (random() * interval '30 days')                 -- Random time in the last 30 days
FROM generate_series(1, 1000) AS id;
```






## Useful psql Commands

Here are some commonly used `psql` commands to help you navigate and inspect your database.

| Command | Description | Example |
| :--- | :--- | :--- |
| `\d` | List relations (tables, views, sequences) or describe a specific relation. | `\d` (list all)<br>`\d trades` (describe 'trades' table) |
| `\dt` | List all tables in the current database. | `\dt` |
| `\l` | List all databases on the server. | `\l` |
| `\c <dbname>` | Connect to a different database. | `\c postgres` |
| `\du` | List all roles (users) and their attributes. | `\du` |
| `\q` | Quit the `psql` shell. | `\q` |



