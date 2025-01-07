## Blog Aggregator

RSS Feed Fetcher Project

### Background
Built to learn how RSS feeding works and utilizing a Postgres DB to store information.

### Built with:
- Go
- Postgres
- Goose (migrations)
- SQLC (SQL -> Go Generated Code)

### Instructions

1. **Have Go installed**: [Download Go](https://go.dev/dl/)

2. **Commands**:

   - Register a user:
     ```shell
     go run . register <name>
     ```

   - Add a feed:
     ```shell
     go run . addfeed "<feed_name>" "<feed_url>"
     ```

   - Aggregate feeds at a specific time interval:
     ```shell
     go run . agg <time_duration>  # e.g., "30s" or "1m"
     ```

   - Reset the database:
     ```shell
     go run . reset
     ```

### Full Command List:

- `register <name>`: Register a new user with the given name.
- `addfeed <feed_name> <feed_url>`: Add a new RSS feed to the aggregator.
- `agg <time_duration>`: Aggregate the feeds every specified time duration (e.g., `30s`, `1m`).
- `reset`: Resets the entire database.

---

### Notes:
- Make sure your Postgres database is set up and configured.
- Ensure that all necessary dependencies are installed before running the commands.
