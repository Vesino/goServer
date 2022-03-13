package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMidleWare(HandleHome, CheckAuth(), Loggin()))
	server.Handle("POST", "/user/create", server.AddMidleWare(UserPostRequest, CheckAuth(), Loggin()))
	server.Listen()
}
