This is a sample (incomplete) implementation of the Leaky Bucket Algorithm for API Rate limiting

## Installation

1. Clone the repo
   ```
   git clone https://github.com/Bhupesh-V/ratelimiter-demo
   ```
2. Start the server
   ```
   go run main.go
   ```

Try visiting the dev server at localhost:8080 and making quick requests. Play with capacity in `ratelimiter.CreateUserBucket()`


### TODO

- Fix support for concurrent requests
- Add benchmarks
- Better interface to choose between the 4 algorithms for rate-limiting
- Pick user-ids, tokens from env or some kind of configuration
- Interface for reading from a global distributed cache
- Interface for making requests/user count configurable