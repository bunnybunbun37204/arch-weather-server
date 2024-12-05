Here is a `README.md` for using the API based on your instructions:

# Weather API

This API allows you to get and post weather data to a PostgreSQL database. 

## Prerequisites

Before using this API, ensure you have the following installed:

1. **Go**: [Download Go](https://golang.org/dl/)
2. **Docker**: [Install Docker](https://www.docker.com/get-started)

## Setup

Follow the steps below to get started with the API:

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/bunnybunbun37204/arch-weather-server.git
cd weather-api
```

### 2. Create `.env` File

You need to create a `.env` file to store environment variables. You can get the `.env` file from my [Discord](https://discord.com).

### 3. Install Necessary Libraries

After setting up the `.env` file, install the necessary Go libraries:

```bash
go mod tidy
```

### 4. Start the Server

#### Using Docker

1. Build and run the Docker containers:

```bash
docker-compose up
```

#### Running the Server

2. Start the server by running the following Go command:

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`.

## API Documentation

### GET `/weather`

This endpoint retrieves the latest weather value from the database.

**Response Format:**

```json
{
    "ID": 1,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "Pressure": 1013.25,
    "RelativeHumidity": 85,
    "Temperature": 30.5,
    "WindDirection": 90,
    "WindSpeed": 5,
    "CHP1": 1.2,
    "DirectSun": 400,
    "GlobalSun": 800,
    "DiffuseSun": 200,
    "RainFall": 10.5,
    "AllDayIllumination": 1200,
    "PM25": 15
}
```

### POST `/weather`

This endpoint allows you to post new weather data to the database.

**Request Body Format:**

```json
{
    "Pressure": 1013.25,
    "RelativeHumidity": 85,
    "Temperature": 30.5,
    "WindDirection": 90,
    "WindSpeed": 5,
    "CHP1": 1.2,
    "DirectSun": 400,
    "GlobalSun": 800,
    "DiffuseSun": 200,
    "RainFall": 10.5,
    "AllDayIllumination": 1200,
    "PM25": 15
}
```

**Example POST Request (Using `curl`):**

```bash
curl -X POST http://localhost:8080/weather \
    -H "Content-Type: application/json" \
    -d '{
        "Pressure": 1013.25,
        "RelativeHumidity": 85,
        "Temperature": 30.5,
        "WindDirection": 90,
        "WindSpeed": 5,
        "CHP1": 1.2,
        "DirectSun": 400,
        "GlobalSun": 800,
        "DiffuseSun": 200,
        "RainFall": 10.5,
        "AllDayIllumination": 1200,
        "PM25": 15
    }'
```

## Troubleshooting

- If you encounter issues with Docker, ensure that Docker is properly installed and running.
- If you have issues with Go dependencies, make sure you have the correct versions of Go installed.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

### Key Sections:
1. **Prerequisites**: Instructions for installing Go and Docker.
2. **Setup**: Instructions for setting up the environment, creating the `.env` file, and running the server.
3. **API Documentation**: Descriptions of the `GET /weather` and `POST /weather` endpoints, with the respective response and request formats.
4. **Troubleshooting**: Provides basic troubleshooting information in case the user encounters issues.