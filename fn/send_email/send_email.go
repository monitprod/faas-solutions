/*

 DON'T MOVE THIS FILE TO ANOTHER FOLDER THAT ISN'T THE ROOT FOLDER
 DOING THIS WILL HARM GetRootPath METHOD

*/
package send_email

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var rootPath string

func StartEnv() {
	envPath := GetRootPath() + "/.env"
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatal("Send email dot env not initialized:", err)
	}

	log.Println("Send Email Env Started!")
}

func GetRootPath() string {
	if rootPath != "" {
		return rootPath
	}

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)

	rootPath = basepath
	return rootPath
}
