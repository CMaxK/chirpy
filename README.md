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


## Testing Endpoints:

### 1. User creation, login, auth, chirp posting, get chirps:
Creation:
![image](https://github.com/CMaxK/chirpy/assets/71667581/431e89d5-d62a-46e3-9645-4dcc4952e295)
Login:
![image](https://github.com/CMaxK/chirpy/assets/71667581/6e80a28b-06c8-49f7-a8e4-0b8a8597d35d)
Post chirp using auth token:
![image](https://github.com/CMaxK/chirpy/assets/71667581/7dd427a6-82b2-43bc-8ce5-042cdf65ef9f)
![image](https://github.com/CMaxK/chirpy/assets/71667581/84958697-5ec2-41dd-a0d0-7aef88f092e4)
![image](https://github.com/CMaxK/chirpy/assets/71667581/427fe560-1b24-4cce-a9e3-1f5ca89adbb0)
Get Chrips:
![image](https://github.com/CMaxK/chirpy/assets/71667581/27b83494-41f4-47fd-99b5-4235fdaf8164)

### 2. Chirps Delete:
![image](https://github.com/CMaxK/chirpy/assets/71667581/d0f36bc1-4c6c-44bb-ac23-d7e97388e11c)

### 3. Webhook for premium members:
Creation:
![image](https://github.com/CMaxK/chirpy/assets/71667581/a3ce59bd-a1a8-4f10-9288-a975ec0e040a)
Upgrade using API Key:
![image](https://github.com/CMaxK/chirpy/assets/71667581/0b754063-7f10-4729-8b16-6cec94a5999f)
![image](https://github.com/CMaxK/chirpy/assets/71667581/b5b2352d-525a-4cbe-813c-72756fbf0278)
Premium feature (is_chirpy_red) activated:
![image](https://github.com/CMaxK/chirpy/assets/71667581/73c397ef-9469-47bb-9061-342414cf6a4c)


### Further features such as:
* not allowing duplicate user creation
* Only user can delete their own tweets
* Admin panel with auto-increment visitor count
* Admin revocation of user access
* API health-check endpoint
* Reading chirps by chirp_id
* If a user should not be allowed to perform a certain task...they won't be able to without correct auth tokens etc.





