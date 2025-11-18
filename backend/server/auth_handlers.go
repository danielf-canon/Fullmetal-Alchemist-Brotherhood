package server

import (
    "backend-alquimia/api"
    "backend-alquimia/models"
    "encoding/json"
    "net/http"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPassword(hashed string, plain string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
    return err == nil
}

func generateJWT(userID uint, rol string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "rol":     rol,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}


func (s *Server) HandleRegister(w http.ResponseWriter, r *http.Request) {
    var req api.RegisterRequest
    json.NewDecoder(r.Body).Decode(&req)

    existing, _ := s.UsuarioRepository.FindByEmail(req.Email)
    if existing != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("email already exists"))
        return
    }

    hash, _ := hashPassword(req.Password)

    user := &models.Usuario{
        Email:        req.Email,
        PasswordHash: hash,
        Rol:          "alquimista", 
    }

    user, _ = s.UsuarioRepository.Save(user)

    alquimista := &models.Alquimista{
        Nombre:       req.Name, 
        Edad:         18,
        Especialidad: "Desconocida",
        Rango:        "Aprendiz",
    }

    alquimista, _ = s.AlquimistaRepository.Save(alquimista)

    user.AlquimistaID = &alquimista.ID
    s.UsuarioRepository.Save(user)

    token, _ := generateJWT(user.ID, user.Rol)

    res, _ := json.Marshal(api.AuthResponse{Token: token})
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}


func (s *Server) HandleLogin(w http.ResponseWriter, r *http.Request) {
    var req api.LoginRequest
    json.NewDecoder(r.Body).Decode(&req)

    user, _ := s.UsuarioRepository.FindByEmail(req.Email)
    if user == nil || !checkPassword(user.PasswordHash, req.Password) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    token, _ := generateJWT(user.ID, user.Rol)

    res, _ := json.Marshal(api.AuthResponse{Token: token})
    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}
