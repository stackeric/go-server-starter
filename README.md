
## Development

1. Clone the (forked) repo.
2. Then, run

  ```sh
  $ go get github.com/codegangsta/gin
  ```


3. Run

  ```sh
  $ go get && go install && PORT=7000 DEBUG=* gin -p 9000 -a 7000 -i run # or run make dev
  ```

  Then visit `localhost:7000`

4. `MONGODB_URL` and `PORT` can be configured by setting the env variable.

  ```sh
  $ export MONGODB_URL=mongodb://
  $ export PORT=7000
  ```

