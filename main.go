package main

import (
	"context"
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
var defaultResponse = map[string]interface{}{
	"Status": "fail",
}

func main() {
	host := flag.String("host", "localhost:8080", "host to connect")
	conn, err := grpc.Dial(*host, grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}
	client := pb.NewRedirectionClient(conn)

	g.GET("/ping", func(c *gin.Context) {
		body := map[string]string{
			"Message": "pong",
		}
		c.JSON(http.StatusOK, body)
	})

	g.GET("/info/:u", func(c *gin.Context) {
		body := defaultResponse

		u := &pb.User{User: c.Param("u")}
		res, err := client.GetInfoByUser(context.TODO(), u)
		if err != nil {
			log.Println(err)
			body["Message"] = err.Error()
			c.JSON(http.StatusBadRequest, body)
		}
		body["Status"] = "ok"
		body["Data"] = res
		c.JSON(http.StatusOK, body)
	})

	g.GET("/get/:p", func(c *gin.Context) {
		body := defaultResponse

		path := &pb.RedirectPath{Path: c.Param("p")}

		if res, err := client.GetOrgByPath(context.TODO(), path); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, body)
		} else {
			body["Status"] = "ok"
			body["Org"] = res.GetOrg()
			c.JSON(http.StatusOK, body)
		}
	})

	g.POST("/register", func(c *gin.Context) {
		data := &pb.RedirectData{}
		err := c.Bind(&data.Redirect)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, err)
		}

		log.Println(data.Redirect)

		body := defaultResponse
		if res, err := client.SetInfo(context.TODO(), data); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, body)
		} else {
			log.Println(res)
			body["Status"] = "ok"
			body["Data"] = fmt.Sprintf("%s", res.GetOrg())
			c.JSON(http.StatusAccepted, body)
		}

	})

	PortString := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PortString == ":" {
		PortString = ":8080"
	}
	g.Run(PortString)

}
