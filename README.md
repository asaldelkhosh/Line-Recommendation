# :oncoming_taxi: Line Recommendation

Implementing a route recommendation system for ```Snappline!``` users. In this project
we simulate ```Snappline!``` models, after that we create a recommendation system in order
to suggest new routes based on users search results on map.

As you know, ```Snappline!``` is a new taxi application which provides shared cab in specific routes
in ```Tehran```. What we are going to do is to simulate the user data and application data, and build
a recommendation system to find out which routes should be add into the system.

## :mammoth: HTTP

In order to get data, run the http server with the following command:

```sh
cd http
go run main.go
```

### request

Make a get request to ```localhost:8080/data``` to get the following:

```json
{
  "routes": [
    {
      "id": 1,
      "start": {
        "ID": 10,
        "CreatedAt": "2023-05-09T11:00:12.181975+03:30",
        "UpdatedAt": "2023-05-09T11:00:12.181975+03:30",
        "DeletedAt": null,
        "x": 955,
        "y": -862
      },
      "stop": {
        "ID": 6,
        "CreatedAt": "2023-05-09T11:00:11.2225+03:30",
        "UpdatedAt": "2023-05-09T11:00:11.2225+03:30",
        "DeletedAt": null,
        "x": -584,
        "y": 872
      }
    },
    {
      "id": 2,
      "start": {
        "ID": 10,
        "CreatedAt": "2023-05-09T11:00:12.181975+03:30",
        "UpdatedAt": "2023-05-09T11:00:12.181975+03:30",
        "DeletedAt": null,
        "x": 955,
        "y": -862
      },
      "stop": {
        "ID": 4,
        "CreatedAt": "2023-05-09T11:00:08.386891+03:30",
        "UpdatedAt": "2023-05-09T11:00:08.386891+03:30",
        "DeletedAt": null,
        "x": 475,
        "y": -265
      }
    }
  ],
  "searches": [
    {
      "x1": -545,
      "y1": -850,
      "x2": -215,
      "y2": 471
    },
    {
      "x1": 877,
      "y1": 239,
      "x2": -594,
      "y2": -762
    },
    {
      "x1": 171,
      "y1": -629,
      "x2": 40,
      "y2": -741
    }
  ]
}
```

## :robot: Model
