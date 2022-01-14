package dispenser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mdeforest/PiPup/server/internal/app/dispenser/router"
)

func ListenAndServe() {
	r := router.Router()

	fmt.Println(("Starting server on the port 8080..."))

	log.Fatal(http.ListenAndServe(":8080", r))
}
