# Mitramas Attendance Application

Mitramas' need is to create an attendance platform that can handle Mitramas employee absences. In this platform, employees (users) can check in (absent) and check out (absent). Employees can also add the activities they do every day. In addition, employees can also view their absence history and activity history. This project contains the API needed by Mitramas to perform attendance.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

- Clone the git repository:

```
git clone https://github.com/Assyatier21/mig-technical-test.git
cd mig-technical-test
```

- Install Dependencies

```
go mod tidy
```

- Create `config` folder in root path, then create a file `connection.go` in that folder with the following contents:

```
package config

const (
	User     = "YOUR_USERNAME_HERE"
	Password = "YOUR_PASSWORD_HERE"
	Host     = "localhost"
	Port     = "5432"
	Database = "YOUR_DATABASE_HERE"
	Schema   = "YOUR_SCHEMA_HERE"
	Sslmode  = "disable"
)
```

- Migrate Database From `mig-technical-test/tools/dump-mig-attendance-xxx.sql`

### Running

```
go run cmd/main.go
```

### API Endpoints Documentation

Details on the use of how this API works can be accessed via the following [link](https://documenter.getpostman.com/view/21912170/2s8Z72Tqqx).

- `/login`: auth for login
- `/logout`: logout user and clear session
- `/api/v1/activity`: get activity details between date given
- `/api/v1/activity`: (method: POST) add or edit activity in attendance
- `/api/v1/activity/delete`: delete activity from attendance specific using attendance_id
- `/api/v1/check-in`: set user checked-in on system
- `/api/v1/check-out`: set user checked-out from system
- `/api/v1/attendance/history`: get attendance history of user

We can test the endpoint using the collection (recommend to using Postman) located in : `mig-technical-test/tools`.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/mig-technical-test/blob/master/LICENSE) file for details.
