
# github.com/pocketGod/Basic_Auth_in_GO

Basic Authentication API Service


1. Set the `SECRET_KEY` environment variable:

```bash
set SECRET_KEY=first_time_with_GO_auThenTicaTion_secreT_key_
```

2. To run the service, execute the following command:

```bash
go run main.go
```

## Usage

### Token Issue

- **Method:** POST
- **URL:** `http://localhost:<PORT>/tokenIssue`
- **Body:**

```json
{
  "username": "string",
  "password": "string"
}
```

- **Response:**

```json
{
  "token": "<YOUR_JWT_TOKEN>"
}
```

- **Description:** This endpoint receives user credentials and validates them against the local `users.json` file (temporarily). If the credentials are valid, it returns a JWT token.

### Token Validate

- **Method:** GET
- **URL:** `http://localhost:<PORT>/tokenValidate?token=<YOUR_JWT_TOKEN>`
- **Response:**

```json
{
  "valid": boolean
}
```