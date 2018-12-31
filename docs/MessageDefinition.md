# HTTP API definition

All of request/response format  are JSON

## passphrase /api/passphrase

This is for generating passphrase.

### GET /api/passphrase

- Request format:
	```json
	{
		"title": "string of this passphrase's title",
		"keyphrase": "keyphrase string",
		"algorithm": "generating passphrase algorithm, write hash string",
		"seed": "seed value of string to generate passphrase"
	}
	```

- Response:
	- HTTP 200 OK if server generate passphrase:
		```json
		{
			"result": "passphrase result string"
		}
		```
	- (Other status code if server failed to generate passphrase, not implement)

### Others

- POST/PUT/DELETE api/passphrase are Nothing

## passphraseInformation /api/passphraseInfo

This is for login user's passphrase information.  
User can store base of generating passphrase information.

### GET /api/passphraseInfo

- Request format:
	- body data is NULL
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server generate passphrase:
		```json
		{
			"value of title": {
				"algorithm": "generating passphrase algorithm, write hash string",
				"seed": "seed value of string to generate passphrase"
			},
			...
		}
		```
	- (Other status code if server failed to generate passphrase, not implement)

### POST/PUT /api/passphraseInfo

Create/Update passphraseInformation

- Request format:
	- body data is following:
		```json
		[
			"value of title": {
				"algorithm": "generating passphrase algorithm, write hash string",
				"seed": "seed value of string to generate passphrase"
			},
			...
		]
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if PUT request and there is no same title's data, and
			- if POST reqest and there is already created same title's data.

### DELETE /api/passphraseInfo

Delete passphraseInformation

- Request format:
	- body data is following:
		```json
		{
			"value of title",
			...
		}
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if there is no data.

## UserAccount /api/userAccount

This is for generating/updating/deleting login user account.

### GET /api/userAccount

- Nothing

### POST /api/userAccount

Create UserAccount

- Request format:
	- body data is following:
		```json
		{
			"username": "user name information",
			"passphrase": "passphrase string"
		}
		```

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if there is already used username.

### PUT /api/userAccount

Update UserAccount

- Request format:
	- body data is following:
		```json
		{
			"username": "user name information",
			"passphrase": "passphrase string"
		}
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if there is no user.

### DELETE /api/userAccount

Delete UserAccount

- Request format:
	- body data is following:
		```json
		{
			"username": "user name information",
		}
		```
	- User autorization is by Digest autorization.

- Response:
	- HTTP 200 OK if server success to update (body data is NULL).
	- Other status code if server failed to generate passphrase, not implement
		- 400 Bad request
			- if there is no user.

- Note:
	- All of passphrase information which is created by the user are deleted.