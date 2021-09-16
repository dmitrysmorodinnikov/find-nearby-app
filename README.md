# find-nearby-app

---

find-nearby-app is a system that, given coordinates, radius and limit, finds the nearest locations (e.g. vehicles/drivers/stores) withing the range defined by the radius. The system consists of backend, frontend and storage.

## How to run
You will need Docker https://docs.docker.com/get-docker/ and Docker Compose https://docs.docker.com/compose/install to run the solution.

1. Go to the Backend folder:

```
cd find-nearby-backend
```

2. Start the Backend Web Server and Postgres DB. DB will be automatically created, initialised and 1000 Singapore locations will be inserted:
```
make docker.start
```

3. Check that the server is up and running:
```
curl localhost:8080/ping
```
The server should have responded with `pong`.

4. Go to Frontend folder:
```
cd ../find-nearby-frontend
``` 

5. Start Frontend:
```
make
```

The app should be available at `localhost:3003`

6. To test the app, put the coordinates, radius and limit and press a "Find Vehicles" button. If the parameters are valid, the found locations will be shown on the map. Each marker is clickable - the vehicle info is shown when marker is clicked. Sample values: latitude=1.3261128155725437, longitude=103.69055657791628 radius=1000, limit=10. Check the results. Increase the radius to 2000 - there will be more results. Overall, 1000 vehicles were uniformly distributed across Singapore as a seed data.

##Backend: technical details
1. Backend is a Web Service written in Go. Go version required: >=1.13.

2. Two endpoints are supported:
  * GET '/ping'
  * GET '/locations/find?latitude=:latitude&longitude:=longitude&radius:=radius&limit=:limit

3. The system is covered by unit and integration tests. To run the tests locally (Go needs to be installed):

`cd find-nearby-backend`  
`make copy-config`  
`make build`  
`make db.docker-start`  
`make db.setup`  
`make test`


In case there are any problems with dependencies, run `go mod tidy`.

4. Postgres with [Postgis extension](https://postgis.net/) is used as a storage. Postgis provides support for spatial and geographic objects and manages location quieries efficiently (Postgis uses [R-Tree based indexes](https://postgis.net/workshops/postgis-intro/indexing.html)).
5. DB migrations are versioned, so that every change in a db schema can be tracked and rolled back.

6. Singapore locations were pre-generated with the help of [this awesome package](https://github.com/AleNegrini/PyCristoforo). Overall, 1000 locations were generated which gives a pretty high chance that some location will be found if a random coordinate within Singapore is provided as an input and a big enough radius.

7. The architecture of the backend is a standard layered architecture: Handler -> Usecase (business logic) -> repository -> underlying storage. Each layer relies on interfaces as dependencies (standard Dependency Injection) which facilitates proper testing and makes it easy to extend the system easily.

8. The model currently consists of only one entity - `Location`. However, it can be easily extended as new requirements emerge (e.g. we can add `Vehicle`, `City` etc).



##Frontend: technical details

1. Frontend is a basic React app. To run it locally (docker instructions were given at the beginning of this doc):  
`cd find-nearby-frontend`  
`npm install --silent`  
`npm install react-scripts@3.4.1 -g --silent`  
`npm start`

2. [Mapbox](https://www.mapbox.com/) is used to render the map and locations. It requires a token which is stored in `.env.local`. It's a free demo token, so no security concerns.

3. The found locations are displayed as red markers, while the "user input coordinates" is marked as a yellow marker.


