# SingleStore Golang Template

This repo serves as an example of using Golang to connect and query a
[SingleStore][ssdb] database.

## Dependencies

* Make sure you have [Go][golang] 1.13 or above installed
* Install [Docker][docker] and ensure that the current user can run containers

## Example Usage

```console
foo@bar$ make start-singlestore
2020-12-17 02:45:20.354243 Initializing MemSQL Cluster in a Box
2020-12-17 02:45:20.354278 Creating...
2020-12-17 02:45:20.592035 Done.
2020-12-17 02:45:20.592061 Configuring...
2020-12-17 02:45:21.035554 Done.
2020-12-17 02:45:21.035579 Bootstrapping...
2020-12-17 02:45:26.066715 Done.
2020-12-17 02:45:26.066747 Configuring Toolbox...
2020-12-17 02:45:26.070704 Done.

Successful initialization!

To start the cluster:
    docker start (CONTAINER_NAME)

To stop the cluster (must be started):
    docker stop (CONTAINER_NAME)

To remove the cluster (all data will be deleted):
    docker rm (CONTAINER_NAME)

SingleStore DB started

foo@bar$ make
go build main.go
./main
2020/12/16 18:45:29 connecting to SingleStore
2020/12/16 18:45:29 creating example database
2020/12/16 18:45:33 creating example table
2020/12/16 18:45:33 inserting sample data
2020/12/16 18:45:33 Row 1: (2, world)
2020/12/16 18:45:33 Row 2: (1, hello)
2020/12/16 18:45:33 Row 3: (3, this is an example)

foo@bar$ make stop-singlestore
SingleStore DB stopped
```

[ssdb]: https://www.singlestore.com/
[s2ms]: https://portal.singlestore.com/
[docker]: https://docs.docker.com/get-docker/
[golang]: https://golang.org/doc/install
