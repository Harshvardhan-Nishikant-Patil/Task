Clone the Repository:

git clone <repository_url>
Navigate to the Backend Directory:
 
cd e-commerce-platform/backend
Initialize Go Module:

go mod init e-commerce-platform
Install Dependencies: 

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
Run the Backend Server:

go run main.go
Frontend (Svelte):
Navigate to the Frontend Directory:

cd ../frontend
Install Node.js Dependencies:

npm install
Run the Frontend Development Server:

npm run dev
Access the Application:
Open your web browser and go to http://localhost:5000 to access the frontend of the e-commerce platform.

Additional Notes:
Make sure you have SQLite installed to run the backend application with the provided SQLite database file (test.db).
Adjust the database connection settings in main.go if you want to use a different database.
The backend server will be running on port 8080 by default. If you need to change it, modify the port in the main.go file.
Access the frontend application at http://localhost:5000 by default (or as specified in the console output).
