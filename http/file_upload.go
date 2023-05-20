package http

import (
	"fmt"
	"github.com/signmem/httpfileserver/g"
	"io"
	"net/http"
	"os"
)


func uploadHandler() {
	http.HandleFunc("/api/v1/fileupload",
		func(w http.ResponseWriter, r *http.Request) {

			/*
			if r.Method != "POST" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}


			clientIP, err := g.GetClientIP(r)
			if err != nil {
				clientIP = "None"
			}
			*/

			fileInfo, err := g.HTTPCheckContent(r)

			/*

			r.Body = http.MaxBytesReader(w, r.Body, g.Config().MaxUploadSize)

			if err := r.ParseMultipartForm(g.Config().MaxUploadSize); err != nil {
				msg := fmt.Sprintf("client %s, The uploaded file is too big. err %s", clientIP, err)
				http.Error(w, msg, http.StatusBadRequest)
				return
			}
			*/

			srcFile, handler, err := r.FormFile("fsname")

			if err != nil {
				msg := fmt.Sprintf("from file error %s", err)
				http.Error(w, msg, http.StatusBadRequest)
				return
			}

			defer srcFile.Close()

			_ = g.CheckAndCreateDir(fileInfo.FSPath)
			destName := g.Config().DownloadDir + "/" + fileInfo.FSPath + "/" + handler.Filename
			destFile, err := os.OpenFile(destName, os.O_WRONLY|os.O_CREATE, 0666)

			if err != nil {
				msg := fmt.Sprintf("openfile error %s", err)
				http.Error(w, msg, http.StatusInternalServerError)
				return
			}

			defer destFile.Close()
			_, err = io.Copy(destFile, srcFile)

			if err != nil {
				msg := fmt.Sprintf("copy file error %s", err)
				http.Error(w, msg, http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Upload successful")

		})
}




