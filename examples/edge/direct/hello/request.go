package hello

type SayHelloRequest struct {
	Name string `form:"name"`
}

type StreamHelloRequest struct {
	Name string `form:"name"`
}
