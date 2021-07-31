package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "urlmap-front/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var g = gin.Default()

func main() {
	host := flag.String("host", "localhost:8080", "host to connect")
	conn, err := grpc.Dial(*host, grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}
	client := pb.NewRedirectionClient(conn)

	g.GET("/info/:u", func(c *gin.Context) {
		u := &pb.User{User: c.Param("u")}
		res, err := client.GetInfoByUser(context.TODO(), u)
		if err != nil {
			log.Println(err)
			msg := map[string]string{"error": err.Error()}
			j, _ := json.Marshal(msg)
			c.JSON(http.StatusBadRequest, j)
		}
		c.JSON(http.StatusOK, res)
	})

	PortString := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PortString == ":" {
		PortString = ":8080"
	}
	g.Run(PortString)

}
