package main

import (
	user_server "gomq/user/server"
	vote_server "gomq/vote/server"
)

func main() {

	go user_server.UserServer()
	vote_server.VoteServer()

}
