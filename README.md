# chirpy

Web server that mimics a popular short-form-messaging platform backend with:
* CRUD operations for messages
* User auth with JWT
* Refresh tokens as extra security layer on top of of JWT
* User login/logout and relevant access permissions
* JSON object to mimic document store for messages and users

## To run:

`go build -o out && ./out`

## On re-run:

Either manually delete the previously created **database.json** file or run: `go build -o out && ./out --debug`
