package slugspace

import (
	"net/http"
			"encoding/json"
	"github.com/colbyleiske/slugspace/slugspaceapi/core/database"
	"github.com/colbyleiske/slugspace/slugspaceapi/models"
	"fmt"
)

//This will get gated by some sort of encryption eventually. Can't let anyone just make requests here
//Make a test for this at some point. Whether or not its private is TBD
func (s *Store) PostRegisterAppInstance() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		var payload database.JWTPayload

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payload)
		if err != nil {
			fmt.Println("Decoding payload issue")
			w.WriteHeader(http.StatusInternalServerError) //please add real logging to this asap
			return
		}

		tokenString, err := s.DAL().CreateJWT(&payload)
		if err != nil {
			if err.Error() == "Could not generate JWT" {
				fmt.Println("JWT generating issue")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err.Error() == "Insufficient claims" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if err.Error() == "Could not generate HMAC Key" {
				fmt.Println("HMAC Key issue")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.RegistrationResponse{ JWT: tokenString })
	})
}