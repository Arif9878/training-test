### Install package

```bash
make tidy
```

### Running Answer of Question Number 2
    
1. Rename .env-example file to .env
2. Set environment variable API_KEY_OMDB
3. Run app
    ```bash
    make run-question-2
    ```

> To test log search make sure have run the log_search.sql in your mysql
> Comment code `db, _, err := sqlmock.New()` on question-2/cmd/main.go
> Uncomment mysql db setup on question-2/cmd/main.go and set environment variable 
> Run again question-2 
 
# Execute the call
$ curl localhost:8080/omdb?search=Batman&page1

# To test grpc service run
    ```bash
    make run-test-client-grpc
    ```
t
### Running Answer of Question Number 3

```bash
make run-question-3
```

## Running Answer of Question Number 4

```bash
make run-question-4
```

## For Test 

```bash
make test
```