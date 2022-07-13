# shared-bike
## How to start projects
### By docker-compose
1. Run command `docker-compose build`
1. Run command `docker-compose up -d frontend`
1. Waiting for all services started
1. Restart dev_api by running `docker restart dev_api` for migration of the first time
1. Wait for the project to start and access the frontend via `http://localhost:3000` and the API will serve on `http://localhost:8000`
### By machine environment
#### Start the API
1. Install MySQL by running the command `brew install mysql`
1. Start the MySQL `brew services start mysql`
1. Go to `api` folder and run the command `make install`
1. Copy `.env.sample` to `.env` file and change the `DB_CONNECTION_STRING` as your local config
1. Run DB migration command `goose -dir ./sql/migrations mysql $DB_CONNECTION_STRING up`
1. Run DB seeder command `goose -dir ./sql/migrations mysql $DB_CONNECTION_STRING up`
1. Run `make start` for starting the API service
1. Access the service swagger doc via `http://localhost:8000/swagger/index.html`

#### Start the Frontend
1. Go to `frontend` folder
1. Change the API config for the development environment in `public/config.json`
1. Run command `yarn install` for installing the dependencies
1. Run command `yarn start`
1. Access to the website via `http://localhost:3000`

## High-level solution
### Context

### Containers

### Components
## DB Diagram
## Detail API design

### User
#### Register
#### Login

### Bike
#### Get All Bikes
1. Sequence Diagram
    ```uml
    @startuml
    actor       User
    boundary    API
    control     BikesHandler
    control     BikeUseCase
    entity      BikeRepository
    entity      UserRepository
    database    Database
    User -> API : HTTPS
    API -> BikesHandler : check authenticate
    BikesHandler -> BikeUseCase
    BikeUseCase -> BikeRepository
    BikeRepository -> Database : Get all bikes
    BikeRepository <-- Database : Return []domain.Bike
    BikeUseCase <-- BikeRepository : Return domain.Bike
    BikeUseCase -> UserRepository
    UserRepository -> Database : Get users by IDs
    UserRepository <-- Database : []domain.User
    BikeUseCase <-- UserRepository : []domain.User
    BikesHandler <-- BikeUseCase: []domain.BikeDTO
    API <-- BikesHandler: []domain.BikeDTO
    User <-- API : JSON or Error
    @enduml
    ```
1. Params
    - No parameters
1. Response
    - Status 200  
        ```json
          [
            [
              {
                "id": 1,
                "lat": "50.119504",
                "long": "8.638137",
                "name": "henry",
                "nameOfRenter": "Bob",
                "status": "rented",
                "userId": 1,
              }
            ]
          ]
        ```
    - Status 500  
        `internal server error`
1. Response property
    - `id` is a unique id of bike
    - `lat` is latitude
    - `long` is longitude
    - `name` is the name of the bike
    - `nameOfRenter`(optional) is name of the renter
    - `userId` is renter id
#### Rent Bike
#### Return Bike

## Tech stacks
### Backend
1. Golang
1. Echo Framework for API
1. MySQL for DB
1. Testing by mockery and testify

### Frontend
1. React for frontend
1. Jest for testing


## Improvement
For API to get all bikes we'll improve it to get bikes only near locations of the users by using lat and long calculation
