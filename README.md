## Get Subordinates

### How to run
1. Install [Golang](https://golang.org/dl/)
2. Clone [repo](https://github.com/sandyleo26/getsubordinates)
`git clone https://github.com/sandyleo26/getsubordinates`
3. build & run
`make start`


### How to test
1. Run unit tests
`make test`

2. Run interactive tests (needs curl or postman)
For example, here's the curl command

```shell
# set roles
curl -X POST http://localhost:8080/roles -d "[
  {
      \"Id\": 1,
      \"Name\": \"System Administrator\",
      \"Parent\": 0
  },
  {
      \"Id\": 2,
      \"Name\": \"Location Manager\",
      \"Parent\": 1
  },
  {
      \"Id\": 3,
      \"Name\": \"Supervisor\",
      \"Parent\": 2
  },
  {
      \"Id\": 4,
      \"Name\": \"Employee\",
      \"Parent\": 3
  },
  {
      \"Id\": 5,
      \"Name\": \"Trainer\",
      \"Parent\": 3
  }
]"

# set users
curl -X POST http://localhost:8080/users -d "[
  {
    \"Id\":   1,
    \"Name\": \"Adam Admin\",
    \"Role\": 1
  },
  {
    \"Id\":   2,
    \"Name\": \"Emily Employee\",
    \"Role\": 4
  },
  {
    \"Id\":   3,
    \"Name\": \"Sam Supervisor\",
    \"Role\": 3
  },
  {
    \"Id\":   4,
    \"Name\": \"Mary Manager\",
    \"Role\": 2
  },
  {
    \"Id\":   5,
    \"Name\": \"Steve Trainer\",
    \"Role\": 5
  }
]"

# query
curl http://localhost:8080/subordinates/3

```
