# :oncoming_taxi: Line Recommendation System

![Version](https://img.shields.io/badge/Version-v0.1-blue)
![App](https://img.shields.io/badge/App-Recommendation_System-blue)
![Lang](https://img.shields.io/badge/Language-Python-blue)
![Context](https://img.shields.io/badge/Context-Data_Analysis-blue)
![Tests](https://img.shields.io/badge/Tests-Pass-blue)

Implementing a route recommendation system for ```Snappline!``` application. In this project
we simulate ```Snappline!``` routes and search models, after that we create a recommendation system in order
to suggest new routes based on users search results on map.

As you may not know, ```Snappline!``` is a new taxi application which provides shared cab in specific routes
in ```Tehran```. What we are going to do is to simulate the user input data and application data, and build
a recommendation system to find out which routes should be add into the system.

Since the stations are constant and we cannot add new stations, we take search results on the map and
by using the existing routes on ```Snappline!``` we decide which routes should be added into the system.

## :mammoth: HTTP

For simulating the application, we create an http server that generates fake data for us.
This http server works like ```Snappline!``` app. It stores stations, routes and user search results on map.
In order to setup the http server, you can run the following command (make sure to have golang installed on your system):

```sh
cd http
go run main.go
```

### request

No you can get all data by making a get request to ```localhost:8080/data```, result is same as the following:

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
    }
  ],
  "searches": [
    {
      "x1": -545,
      "y1": -850,
      "x2": -215,
      "y2": 471
    }
  ]
}
```

## :robot: Data

Out model uses ```pandas``` library in order to read the ```json``` data and convert it into a
dataframe. After that we use ```ast``` in order to convert the inner json objects into a directory.
In all first steps, we are trying to make our data clean and useable for processing. Our pseudo code is as follow:

```shell
[1] http request to get json data.
[2] convert json data to `pandas` dataframe.
[3] create `routes` and `search` dataframe.
[4] map the `search` entities to a station/node in `routes` dataframe. (since we cannot create stations)
[5] map the `search` source and destination to a new route in `routes` dataframe.
[6] count the number of new entities.
[7] sort them in decreasing order.
```

After we cleaned our data we create two dataframes for search and routes. We convert each entity to
```shapely.geometry.Point``` object to use the ```nearest_points``` method in order to apply our
algorithm on the data to map them into routes dataframe. Now we perform ```nearest_points``` to
map a search to route and then we use ```value_counts``` method in ```pandas``` dataframe in order
to get the results.

### execute

You can execute the script by running ```python main.py``` but make sure to have all of the libraries
in ```requirements.txt``` installed on an virtual environment. Result would be as follow:

```sh
POINT (805 683)-POINT (-408 834)      10
POINT (805 683)-POINT (-629 -733)     6
POINT (805 683)-POINT (475 -265)      4
                                     ..
POINT (-660 562)-POINT (-666 -109)    1
POINT (-245 -149)-POINT (805 683)     1
```
