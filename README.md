# Intro
It is backend API for ehr_webapp written in Go.

# Get started
## Setup postgres sql database
- Install postgresql in your local computer without using docker 
- Connect to postgresql database via the bash terminal
```bash
# Login into postgres
$ psql postgres
# List the available databases
postgres-# \l
# Create database named 'patients'
CREATE DATABASE patients;
# Connect to 'patients' database
\c patients;
# Fill database with data by running each of the sql files in extras/
# Below is only an example
\i /user/home/patient_api/extras/patientIDs.sql;
```
- Update the postgres connection string in data/database.go file

## Run the api
- Load all the modules by executing the following in root directory 
```bash
$ go get .
```
- To Run the API. Execute the following in root directory 
```bash
$ go run .
```

## Testing api
Postman is used to test different API endpoint.

Since, the data is only present local. Postman App needs to be installed locally. 

 View: [API Documentation](https://winter-resonance-708968.postman.co/workspace/Baum-Team-Space~9b9975ce-b871-450a-886a-39856df32780/collection/24423207-b2292700-66a5-41a5-a184-5256203baab3?action=share&creator=24423207)