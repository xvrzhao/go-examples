package http

import (
	"compress/gzip"
	"log"
	"net/http"
)

func RunServerEnableGzip(addr string) {
	serveMux.HandleFunc("/test/gzip", func(writer http.ResponseWriter, request *http.Request) {
		// 1. set header: content encoding and type, can't set one without the other
		writer.Header().Set("content-encoding", "gzip")
		writer.Header().Set("content-type", "text/plain; charset=utf-8")

		// 2. wrap the writer
		gzipWriter := gzip.NewWriter(writer)

		// 3. write data by wrapper
		_, err := gzipWriter.Write([]byte("hello gzip!"))
		if err != nil {
			http.Error(writer, "server error", 500)
			log.Printf("failed to write compressed stream: %v", err)
			return
		}

		// 4. remember to flush!!
		err = gzipWriter.Flush()
		if err != nil {
			http.Error(writer, "server error", 500)
			log.Printf("failed to flush compressed stream: %v", err)
			return
		}
	})

	err := http.ListenAndServe(addr, serveMux)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
}
