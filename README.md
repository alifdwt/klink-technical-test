<h1 align="center">K-Link Technical Test</h1>

> **Employee Data Management**, built using the [Go programming language](https://golang.org), the [Gin](https://gin-gonic.com) web framework, and the [GORM](https://gorm.io) ORM for interacting with a [PostgreSQL](https://www.postgresql.org) database.

## 🧰 Installation

1. Clone the repository

```bash
git clone https://github.com/alifdwt/klink-technical-test.git
```

2. Install the dependencies

```bash
cd klink-technical-test
go mod download
```

3. Add your .env file

```sh
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_DRIVER=

JWT_SECRET=
TOKEN_SYMETRIC_KEY=
ACCESS_TOKEN_DURATION=

APP_TIMEZONE=
SERVER_ADDRESS=
```

4. Run make swag to generate swagger documentations

```bash
make swag
```

5. Run make run to start the server

```bash
make run
```

## 📖 Documentation

> You can access the documentation at http://SERVER_ADDRESS/docs/index.html

# K-Link Technical Test API

This is K-Link Technical Test API for Full Stack Developer position

## Version: 1.0

### Terms of service

http://swagger.io/terms/

**Contact information:**  
Alif Dewantara  
http://github.com/alifdwt  
aputradewantara@gmail.com

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### Security

**BearerAuth**

| apiKey | _API Key_     |
| ------ | ------------- |
| In     | header        |
| Name   | Authorization |

### /auth/login

#### POST

##### Summary:

Login to the application

##### Description:

Login a user

##### Parameters

| Name | Located in | Description | Required | Schema                                  |
| ---- | ---------- | ----------- | -------- | --------------------------------------- |
| data | body       | data        | Yes      | [auth.LoginRequest](#auth.LoginRequest) |

##### Responses

| Code | Description           | Schema                                            |
| ---- | --------------------- | ------------------------------------------------- |
| 200  | OK                    | [responses.Token](#responses.Token)               |
| 400  | Bad Request           | [responses.ErrorMessage](#responses.ErrorMessage) |
| 500  | Internal Server Error | [responses.ErrorMessage](#responses.ErrorMessage) |

### /auth/register

#### POST

##### Summary:

Register to the application

##### Description:

Register a new user

##### Parameters

| Name | Located in | Description | Required | Schema                                        |
| ---- | ---------- | ----------- | -------- | --------------------------------------------- |
| data | body       | data        | Yes      | [auth.RegisterRequest](#auth.RegisterRequest) |

##### Responses

| Code | Description           | Schema                                            |
| ---- | --------------------- | ------------------------------------------------- |
| 200  | OK                    | [responses.Token](#responses.Token)               |
| 400  | Bad Request           | [responses.ErrorMessage](#responses.ErrorMessage) |
| 500  | Internal Server Error | [responses.ErrorMessage](#responses.ErrorMessage) |

### /member/create

#### POST

##### Summary:

Create new member

##### Description:

Create new member

##### Parameters

| Name | Located in | Description | Required | Schema                                                    |
| ---- | ---------- | ----------- | -------- | --------------------------------------------------------- |
| data | body       | data        | Yes      | [member.CreateMemberRequest](#member.CreateMemberRequest) |

##### Responses

| Code | Description           | Schema                                                |
| ---- | --------------------- | ----------------------------------------------------- |
| 201  | Created               | [responses.MemberResponse](#responses.MemberResponse) |
| 400  | Bad Request           | [responses.ErrorMessage](#responses.ErrorMessage)     |
| 500  | Internal Server Error | [responses.ErrorMessage](#responses.ErrorMessage)     |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /member/showAll

#### GET

##### Summary:

Get all member

##### Description:

Get all member

##### Responses

| Code | Description           | Schema                                                    |
| ---- | --------------------- | --------------------------------------------------------- |
| 200  | OK                    | [ [responses.MemberResponse](#responses.MemberResponse) ] |
| 400  | Bad Request           | [responses.ErrorMessage](#responses.ErrorMessage)         |
| 500  | Internal Server Error | [responses.ErrorMessage](#responses.ErrorMessage)         |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### /member/showById/{id}

#### GET

##### Summary:

Get member by id

##### Description:

Get member by id

##### Parameters

| Name | Located in | Description | Required | Schema  |
| ---- | ---------- | ----------- | -------- | ------- |
| id   | path       | id          | Yes      | integer |

##### Responses

| Code | Description           | Schema                                                |
| ---- | --------------------- | ----------------------------------------------------- |
| 200  | OK                    | [responses.MemberResponse](#responses.MemberResponse) |
| 400  | Bad Request           | [responses.ErrorMessage](#responses.ErrorMessage)     |
| 500  | Internal Server Error | [responses.ErrorMessage](#responses.ErrorMessage)     |

##### Security

| Security Schema | Scopes |
| --------------- | ------ |
| BearerAuth      |        |

### Models

#### auth.LoginRequest

| Name     | Type   | Description | Required |
| -------- | ------ | ----------- | -------- |
| password | string |             | Yes      |
| username | string |             | Yes      |

#### auth.RegisterRequest

| Name       | Type   | Description | Required |
| ---------- | ------ | ----------- | -------- |
| admin_code | string |             | No       |
| password   | string |             | Yes      |
| username   | string |             | Yes      |

#### member.CreateMemberRequest

| Name          | Type   | Description | Required |
| ------------- | ------ | ----------- | -------- |
| birthdate     | string |             | Yes      |
| gender        | string |             | Yes      |
| level         | string |             | Yes      |
| nama_belakang | string |             | Yes      |
| nama_depan    | string |             | Yes      |

#### responses.ErrorMessage

| Name    | Type   | Description                        | Required |
| ------- | ------ | ---------------------------------- | -------- |
| message | string | StatusCode int `json:"statusCode"` | No       |

#### responses.MemberResponse

| Name          | Type    | Description | Required |
| ------------- | ------- | ----------- | -------- |
| birthdate     | string  |             | No       |
| gender_id     | integer |             | No       |
| id            | integer |             | No       |
| join_date     | string  |             | No       |
| level_id      | integer |             | No       |
| nama_belakang | string  |             | No       |
| nama_depan    | string  |             | No       |
| update_at     | string  |             | No       |
| user_id       | integer |             | No       |

#### responses.Token

| Name  | Type   | Description | Required |
| ----- | ------ | ----------- | -------- |
| token | string |             | No       |
