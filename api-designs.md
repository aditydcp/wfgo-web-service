# API Endpoints Design

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