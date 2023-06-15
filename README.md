# MTA Hosting 

MTA Hosting is a service designed to uncover inefficient servers hosting only a few active Mail Transfer Agents (MTAs). It helps optimize MTA hosting for better resource utilization and cost efficiency.

## Features

- Retrieves hostnames having less or equal active IP addresses based on a configurable threshold.
- Configurable threshold using the `THRESHOLD` environment variable (default value: 1).
- Provides IP configuration data through a mock service.
- Includes unit and integration tests.

## Installation

To install and run the MTA Hosting Optimizer locally, follow these steps:

1. Clone the repository from GitLab:

   ```
   git clone https://github.com/akshay-singla/mta-hosting.git
   ```

2. Navigate to the project directory:
    ``` 
    cd mta-hosting-optimizer
    ```

3. Set up the required environment variables:
    * THRESHOLD: Specify the threshold value for active IP addresses (optional, default: 1).

4. Build and run the service:
    ``` 
    go build
    ./mta-hosting 
    ``` 
    
    ``` 
    go run main.go 
    ``` 

    The service will be available at http://localhost:8080.



# API Endpoint
### Retrieve hostnames with fewer or equal active IP addresses
 * URL: /hostnames
 * Method: GET

### Query Parameters
 * threshold (optional): Specify the threshold value for active IP addresses (overrides the environment variable)

### Response
The response will be a JSON array containing the hostnames with fewer or equal active IP addresses based on the threshold.

#### Example response:
``` 
["mta-prod-1", "mta-prod-3"] 
```

## Testing

To run the tests, execute the following command:

```
go test ./...
```

