# SpaceVoyagers Microservice

## Description
A Go microservice for managing different exoplanets. Supports adding, listing, updating, retrieving, deleting exoplanets, and fuel estimation for trips.

## Requirements
- Go 1.16 or later
- Docker (for containerization)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Abhishekagarwal1597/spaceVoyagers
    cd spaceVoyagers
    ```

2. Install dependencies:
    ```sh
    go mod download
    ```

## Running the Service

### Locally

1. Build the application:
    ```sh
    go build -o main .
    ```

2. Run the application:
    ```sh
    ./main
    ```

### Using Docker

1. Build the Docker image:
    ```sh
    docker build -t exoplanet-service .
    ```

2. Run the Docker container:
    ```sh
    docker run -p 8080:8080 exoplanet-service
    ```

## API Endpoints

### Add an Exoplanet

- **URL**: `/addExoplanet`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "name": "string",
    "description": "string",
    "distance": 100,
    "radius": 1.5,
    "mass": 5.0,   // Optional, only for Terrestrial type
    "type": "GasGiant"
  }
