package srv

import (
	"encoding/json"
	"fmt"
	"github.com/ArkronHomeWork/goforfun/model"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UserService struct {
	userRepository *model.UserRepository
}

func GetUserService() *UserService {
	return &UserService{userRepository: model.GetUserRepository()}
}

func (service *UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if users, err := service.userRepository.GetAll(); err != nil {
		log.Printf("error %e while geting all users", err)
		http.Error(w, http.StatusText(400), 400)
		return
	} else {
		rawJson, err := json.Marshal(users)
		if err != nil {
			log.Printf("error %e while making json from %+v", err, users)
			http.Error(w, http.StatusText(500), 500)
			return
		} else {
			if _, err := fmt.Fprintf(w, string(rawJson)); err != nil {
				log.Printf("error %e while saving data %s to responce ", err, rawJson)
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}
	}
}

func (service *UserService) SaveNewUser(w http.ResponseWriter, r *http.Request) {
	data := new(model.UserData)
	rawData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("can not read %+v with error %e", rawData, err)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	err = data.ToObject(rawData)
	if err != nil {
		log.Printf("can not make object from %+v with error %e", rawData, err)
		http.Error(w, http.StatusText(400), 400)
		return
	}
	err = service.userRepository.Save(data)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func (service *UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(chi.URLParam(r, "id")); err != nil {
		log.Printf("can not parse %s to integer", chi.URLParam(r, "id"))
		http.Error(w, http.StatusText(400), 400)
		return
	} else {
		if user, err := service.userRepository.GetById(id); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		} else {
			if res, err := user.ToJson(); err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			} else {
				if _, err = w.Write(res); err != nil {
					http.Error(w, http.StatusText(500), 500)
				}
			}
		}
	}
}
