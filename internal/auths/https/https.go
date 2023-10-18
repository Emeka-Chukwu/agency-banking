package auth_handler

import (
	domain "agency-banking/domain/auths"
	usecases_handler "agency-banking/internal/auths/usecases"
	"agency-banking/util"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func NewAuthHandler(ChiRouter *chi.Mux,
	usecase usecases_handler.Authusecase,
	Config util.Config) {
	handler := AuthsHandler{usecase: usecase, Config: Config, ChiRouter: ChiRouter}
	ChiRouter.Route("/auths", func(r chi.Router) {
		r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("fine and running perfectly "))
		})
		r.Post("/signup", handler.Register)
		r.Post("/login", handler.Login)
	})

}

type AuthsHandler struct {
	ChiRouter *chi.Mux
	usecase   usecases_handler.Authusecase
	Config    util.Config
}

// @Summary Login User
// @Description assign a valid token to user
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.LoginResp
// @Router /auths/login [post]
func (ah AuthsHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.Login
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}
	login, err := ah.usecase.Login(req)
	if err != nil {
		ah.Config.ErrorJSON(w, err)
	}
	ah.Config.WriteJSON(w, http.StatusOK, login)
}

// @Summary Register User
// @Description register and assign a valid token to user
// @ID get-user-by-id
// @Produce json
// // @Param id path int true "User ID"
// @Success 201 {object} domain.LoginResp
// @Router /auths/register [post]
func (ah AuthsHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return
	}
	login, err := ah.usecase.Register(req)
	if err != nil {
		ah.Config.ErrorJSON(w, err)
	}
	ah.Config.WriteJSON(w, http.StatusOK, login)
}
