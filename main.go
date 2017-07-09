package main

import "os"

func main() {
	a := App{}
	a.Initialize(os.Getenv("APP_ROOT_FOLDER_PATH"))
	defer a.Close()
	a.Run(":8080")
}
