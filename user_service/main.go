package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type RegisterUserResponse struct {
	UserID   int64  `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type User struct {
	ID       int64
	Email    string
	Password string
	Name     string
	Gender   string
	Phone    string
	Role     string
}

var (
	users      = make(map[int64]*User)
	usersMutex sync.Mutex
	idCounter  int64 = 1
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Simple check: email must be unique
	for _, u := range users {
		if u.Email == req.Email {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
	}

	user := &User{
		ID:       idCounter,
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Gender:   req.Gender,
		Phone:    req.Phone,
		Role:     req.Role,
	}
	users[idCounter] = user
	resp := RegisterUserResponse{
		UserID:   user.ID,
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Gender:   user.Gender,
		Phone:    user.Phone,
		Role:     user.Role,
	}
	idCounter++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/register", registerHandler)
	log.Println("Mock user_service running at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// curl -X POST http://localhost:8080/api/v1/shippers/   -H "Content-Type: application/json"   -d '{
//     "email": "duong@example.com",
//     "password": "supersecurepassword",
//     "name": "Thai Duong",
//     "gender": "gay",
//     "phone": "1234567890",
//     "vehicleType": "car",
//     "vehiclePlate": "30K-999.99"
//   }'