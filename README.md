# wfgo-web-service
Simple CRUD RESTful Backend App using Go

# Getting Started

The configuration used in this app are as follows:

- Go version 1.21.4

For storage, this app is using MongoDB Cloud storage. There is no need to run it, you just need the URI to access. Access is open to anyone until 21 November 2023.

# What are inside?

This app features

- Go language
- Gin web framework
- MongoDB for storage

This app simulates an aquafarm management application with 2 entities: Farms and Ponds. 1 farm can have multiple ponds but 1 pond can only be registered to 1 farm. A standard CRUD operations for both entity is present with some additional notes.

- `POST` requests will be denied when attempting to create duplicate entry
- `PUT` requests will update existing entry, and will create a new one instead if specified entry does not yet exist.
- `DELETE` requests soft deletes a specified entry by moving it to a 'recycle bin' collection.
- `GET` requests will show all entries and can also be provided with an ID to specifically fetch an entry.

API Design (`api-designs.md`) and [documentation](https://documenter.getpostman.com/view/17672653/2s9YXmZ1ku) are also provided in this repo.

# How to Start

1. Clone this project using `$ git clone https://github.com/aditydcp/wfgo-web-service.git` and make sure to use the `main` or `dev-mongo` branch.

2. Setup your environment variables `.env` file. Make sure there are `MONGODB_URI`, `SERVER_PORT` and `BASE_URL`.

3. Open your terminal / command prompt and install the dependencies using `$ go get .`.

4. To run the app, use `$ go run .`.

5. The app is now running on `http://localhost:3000`. The port can be changed by changing the `SERVER_PORT` value in your `.env` file.

# API Address

As also mentioned in `api-designs.md`, here are the API addresses.

## /farms

### POST - /
Create a new farm entity. Deny duplicate entry.

### PUT - /:id
Update existing farm or create if specified farm does not exist yet.

### DELETE - /:id
Soft delete existing farm. Error if farm does not exist yet.

### GET - /
Get a list of all existing farm.

### GET (ID) - /:id
Get info about a farm with specified id.

## /ponds

### POST - /
Create a new pond entity. Deny duplicate entry.

### PUT - /:id
Update existing pond or create if specified pond does not exist yet.

### DELETE - /:id
Soft delete existing pond. Error if pond does not exist yet.

### GET - /
Get a list of all existing pond.

### GET (ID) - /:id
Get info about a pond with specified id.

## /statistics

### GET - /
Get info of each endpoints call count.

# Limitation

Unfortunately, this app is not yet final and still suffers some bugs and limitations.

- Enforcement of the relationship between the 2 entities has not been fully implemented. The intended behavior is that when creating a Farm, the app won't allow if the Farm uses a Pond that is already taken by other Farm.
- Statistics were able to record unique agents accessing each endpoints, however it does not add to the call count.
- `GET` requests on `/statistics` result on an `500` Internal Server Error code.
- `GET` requests on `/recycled/farms` and `/recycled/ponds` are unimplemented yet.
- Unit tests and other tests have not been conducted programmatically.