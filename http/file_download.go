package http

import (
	"bytes"
	"fmt"
	"github.com/signmem/httpfileserver/g"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func httpFileDownLoad() {
	http.HandleFunc("/api/v1/download",
		func(w http.ResponseWriter, r *http.Request) {

			clientIP, err := g.GetClientIP(r)
			if err != nil {
				clientIP = "None"
			}

			fileInfo, err := g.HTTPCheckContent(r)

			if err != nil {
				msg := fmt.Sprintf("httpFileDownLoad() client: %s, %s", clientIP, err)
				g.Logger.Error(msg)
				http.Error(w, msg , http.StatusInternalServerError)

				return
			}

			fileInfo.FSClient = clientIP

			//  http body check finish

			g.Logger.Debugf("httpFileDownLoad() client: %s, going download file: %s", clientIP, )

			file := g.Config().DownloadDir + "/" + fileInfo.FSPath

			fileName := file[strings.LastIndex(file,"/")+1:]

			downloadBytes, err := ioutil.ReadFile(file)
			if err != nil {
				msg := fmt.Sprintf("httpFileDownLoad() client: %s, error: ReadFile file %s error", clientIP, fileInfo.FSPath)
				g.Logger.Error(msg)
				http.Error(w, msg , http.StatusInternalServerError)
				return
			}

			mime := http.DetectContentType(downloadBytes)

			fileSize := len(string(downloadBytes))

			w.Header().Set("Content-Type", mime)
			w.Header().Set("Content-Disposition", "attachment; filename=" + fileName + "")
			w.Header().Set("Expires", "0")
			w.Header().Set("Content-Transfer-Encoding", "binary")
			w.Header().Set("Content-Length", strconv.Itoa(fileSize))
			w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

			http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))
			return
		})
}