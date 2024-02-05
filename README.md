# chirpy

Web server that mimics "twitter" backend with:
* CRUD operations for messages
* User auth with JWT
* User login/logout

## To run:

go build -o out && ./out

## On re-run:

Either manually delete the previously created database.json file or run: go build -o out && ./out --debug
