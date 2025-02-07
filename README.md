# Geo Conversion Service

This project is a simple web application that acts as a conversion service from any georeferenced file to a MBTile pack. It uses Golang and the `ogr2ogr` command-line tool for the conversion process. The converted files are stored in an S3 bucket, and users can download them using pre-signed URLs.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- Golang (version 1.16 or later)
- GDAL (Geospatial Data Abstraction Library) with `ogr2ogr` tool
- AWS CLI (for configuring AWS credentials)

## Building the Project

To build the project, follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/geo-conversion-service.git
   cd geo-conversion-service
   ```

2. Install the required dependencies:

   ```sh
   go mod tidy
   ```

3. Build the project:

   ```sh
   go build -o geo-conversion-service cmd/main.go
   ```

## Running the Project

To run the project, follow these steps:

1. Ensure you have configured your AWS credentials using the AWS CLI:

   ```sh
   aws configure
   ```

2. Start the web server:

   ```sh
   ./geo-conversion-service
   ```

3. Open your web browser and navigate to `http://localhost:8080` to access the file upload page.

## Testing the Project

To test the project, follow these steps:

1. Run the unit tests:

   ```sh
   go test ./...
   ```

2. Check the test coverage:

   ```sh
   go test -cover ./...
   ```

3. Review the test results and ensure all tests pass successfully.
