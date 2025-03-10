# **Golang Todo API**  

A simple RESTful API for managing Todo tasks.

## **Installation & Setup**  

### **1. Clone the Repository**  

```sh
git clone https://github.com/Ayobamidele/todo-api.git
cd todo-api
```

### **2. Install Dependencies**  

Ensure you have **Go** installed (version 1.18+). Then run:  

```sh
go mod tidy
```

### **3. Set Up Environment Variables**  

Create a `.env` file in the project root and add:  

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=todo_db
DB_PORT=5432
DB_SSLMODE=disable
```

### **4. Create the PostgreSQL Database**  

Ensure you create a PostgreSQL database is running and yoy add the connector to the .env file

### **5. Run the Application**  

```sh
go run main.go
```

The server will start on **<http://localhost:8080>** 🚀  

---

## **API Endpoints**  

### **1. Create a Todo**  

**POST** `/todos`  

```json
{
  "title": "Learn Go",
  "description": "Build a simple API",
  "completed": false
}
```

### **2. Get All Todos**  

**GET** `/todos`  

```json
[

    {
        "title": "Lorem",
        "description": "lorempixel",
        "completed": false
    },
    {
        "title": "Ipdum",
        "description": "ipdump",
        "completed": false
    }
]
```

### **3. Get a Todo by ID**  

**GET** `/todos/{id}`  

```json
{
  "title": "",
  "description": "",
  "completed": false
}
```

### **4. Update a Todo**  

**PUT** `/todos/{id}`  

```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "completed": true
}
```

### **5. Delete a Todo**  

**DELETE** `/todos/{id}`  

```json
{
  "message": "Todo deleted"
}
```

---

## **Future Enhancements**  

🔹 Add JWT authentication  
🔹 Deploy with Docker  
