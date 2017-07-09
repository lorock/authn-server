package accounts

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/keratin/authn-server/api"
	"github.com/keratin/authn-server/services"
)

func patchAccountLock(app *api.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			panic(err)
		}

		err = services.AccountLocker(app.AccountStore, id)
		if err != nil {
			if fe, ok := err.(services.FieldErrors); ok {
				api.WriteJson(w, http.StatusNotFound, api.ServiceErrors{fe})
				return
			} else {
				panic(err)
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}
