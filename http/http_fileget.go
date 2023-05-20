package http

import (
	"fmt"
	"github.com/signmem/httpfileserver/g"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func httpFileGet() {
	http.HandleFunc("/api/v1/fileget",
		func(w http.ResponseWriter, r *http.Request) {

			clientIP, err := g.GetClientIP(r)
			if err != nil {
				clientIP = "None"
			}

			fileInfo, err := g.HTTPCheckContent(r)

			if err != nil {
				msg := fmt.Sprintf("httpFileGet() client: %s, %s", clientIP, err)
				g.Logger.Error(msg)
				http.Error(w, msg , http.StatusInternalServerError)

				return
			}

			fileInfo.FSClient = clientIP

			//  http body check finish

			g.Logger.Debugf("httpFileGet() client: %s, going download file: %s", clientIP, fileInfo.FSPath)

			file, err := os.Open(g.Config().DownloadDir + "/" + fileInfo.FSPath)

			if err != nil {
				msg := fmt.Sprintf("httpFileGet() client: %s, error: open file %s error", clientIP, fileInfo.FSPath)
				g.Logger.Error(msg)
				http.Error(w, msg , http.StatusInternalServerError)
				return
			}

			info, err := file.Stat()

			if err != nil {
				msg := fmt.Sprintf("httpFileGet() client: %s, error: stat() %s error", clientIP, fileInfo.FSPath)
				g.Logger.Error(msg)
				http.Error(w, msg , http.StatusInternalServerError)
				return
			}

			fileName := fileInfo.FSPath[strings.LastIndex(fileInfo.FSPath,"/")+1:]

			contentType := "application/octet-stream"
			w.Header().Set("Content-Disposition", "attachment; filename=" + fileName)
			w.Header().Set("Content-Type", contentType)
			w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))
			file.Seek(0, 0)
			io.Copy(w, file)
			file.Close()
			msg := fmt.Sprintf("httpFileGet() client: %s, file: %s download complete.", clientIP, fileInfo.FSPath)
			g.Logger.Debug(msg)
			return
		})
}