package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/WASENDERAUTO/backren"
)

func init() {
	functions.HTTP("Wasenderauto", wasenderauto_Users)
}

func wasenderauto_Users(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	w.Header().Set("Access-Control-Allow-Origin", "https://cyberoren.my.id")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, backren.InsertUserApp("PASETOPUBLICKEY", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, backren.GetuserByAdmin("PASETOPUBLICKEY", r))

}
