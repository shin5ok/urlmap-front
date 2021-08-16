package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "urlmap-front/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var g = gin.Default()

func initDefaultResponse() map[string]interface{} {
	return map[string]interface{}{
		"Status":  "fail",
		"Message": "",
	}
}

func main() {
	host := os.Getenv("URLMAP_HOST")
	if host == "" {
		host = "localhost:8080"
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}
	client := pb.NewRedirectionClient(conn)

	g.GET("/Ping", func(c *gin.Context) {
		log.Println("/Ping")
		body := initDefaultResponse()
		body["Status"] = "ok"
		body["Message"] = "Pong"
		c.JSON(http.StatusOK, body)
	})

	g.GET("/info/:u", func(c *gin.Context) {
		user := c.Param("u")
		log.Printf("/info/%s\n", user)
		body := initDefaultResponse()
		fmt.Println(body)

		u := &pb.User{User: user}
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
		path := c.Param("p")
		log.Printf("/get/%s\n", path)
		body := initDefaultResponse()

		rpath := &pb.RedirectPath{Path: path}

		if res, err := client.GetOrgByPath(context.TODO(), rpath); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, body)
		} else {
			body["Status"] = "ok"
			body["Org"] = res.GetOrg()
			c.JSON(http.StatusOK, body)
		}
	})

	g.POST("/register", func(c *gin.Context) {
		log.Println("/register")
		data := &pb.RedirectData{}
		err := c.Bind(&data.Redirect)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, err)
		}

		log.Println(data.Redirect)

		body := initDefaultResponse()
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
