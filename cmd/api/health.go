package main

import (
	"net/http"
	"github.com/0x03ff/web_go/utils"
)

func (app *application) checkServerHandler(w http.ResponseWriter, r *http.Request){
	utils.WriteJSON(w , http.StatusOK,map[string]string{
    "message": "The service is online",
	})
	
}