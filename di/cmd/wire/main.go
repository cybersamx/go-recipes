package main

func main() {
	hs, err := InitializeServer()
	if err != nil {
		panic(err)
	}

	err = hs.Start()
	if err != nil {
		panic(err)
	}
}
