package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>ちゃっと</title>
				</head>
				<body>
					チャットしましょう！！
				</body>
			</html>
		`))
	})

	// WebサーバーStart
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
