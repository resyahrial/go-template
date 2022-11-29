package route

func (r *R) InitV1Route() {
	v1 := r.engine.Group("/v1")
	users := v1.Group("/users")
	{
		users.POST("", wrapHandler(r.handler.CreateUser))
	}
}
