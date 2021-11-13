package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "urlmap-front/pb"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var g = gin.Default()

var version string = "0.12"

func initDefaultResponse() map[string]interface{} {
	return map[string]interface{}{
		"Status":  "fail",
		"Message": "",
	}
}

func Ping(c *gin.Context) {
	log.Println("/Ping")
	body := initDefaultResponse()
	body["Status"] = "ok"
	body["Message"] = "Pong"
	body["Version"] = version
	c.JSON(http.StatusOK, body)
}

func main() {
	var logger, _ = zap.NewProduction()
	defer logger.Sync()
	suger := logger.Sugar()

	host := os.Getenv("URLMAP_API")
	if host == "" {
		host = "urlmap-api:8080"
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		suger.Infow(err.Error())
		log.Println(err)
	}
	client := pb.NewRedirectionClient(conn)

	g.GET("/", func(c *gin.Context) {
		body := initDefaultResponse()
		body["Status"] = "ok"
		// suger logging which is simple and little slower than using logger directly
		suger.Infow("/", "status", body["Status"])
		// logger method which we should use when we need high performance
		logger.Info("/", zap.String("body_status", body["Status"].(string)))
		c.JSON(http.StatusOK, body)
	})

	g.GET("/Ping", Ping)

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

	g.POST("/user/:u", func(c *gin.Context) {
		log.Println("/user create or update")
		data := &pb.User{}
		err := c.Bind(&data)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, err)
		}

		body := initDefaultResponse()

		user := c.Param("u")
		u := &pb.User{User: user, NotifyTo: data.NotifyTo}

		if res, err := client.SetUser(context.TODO(), u); err != nil {
			log.Printf("%+v\n", err)
			body["Message"] = err
			c.JSON(http.StatusInternalServerError, body)
		} else {
			body["Status"] = "ok"
			body["Data"] = fmt.Sprintf("%s is Created or Updated", u.User)
			body["Message"] = res
		}
		c.JSON(http.StatusAccepted, body)
	})

	g.DELETE("/user/:u", func(c *gin.Context) {
		log.Println("/user delete and user's entries")

		body := initDefaultResponse()

		user := c.Param("u")
		u := &pb.User{User: user}

		if res, err := client.RemoveUser(context.TODO(), u); err != nil {
			log.Printf("%+v\n", err)
			body["Message"] = err
			c.JSON(http.StatusInternalServerError, body)
		} else {
			body["Status"] = "ok"
			body["Data"] = fmt.Sprintf("%s is Removed with entries", u.User)
			body["Message"] = res
		}
		c.JSON(http.StatusAccepted, body)
	})

	g.POST("/register", func(c *gin.Context) {
		log.Println("/register")
		data := &pb.RedirectData{}
		err := c.Bind(&data.Redirect)
		log.Printf("%+v\n", data)

		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, err)
		}

		log.Println(data.Redirect)

		body := initDefaultResponse()
		if res, err := client.SetInfo(context.TODO(), data); err != nil {
			log.Println(err)
			body["Message"] = err.Error()
			c.JSON(http.StatusInternalServerError, body)
		} else {
			log.Println(res)
			body["Status"] = "ok"
			body["Data"] = res.GetOrg()
			c.JSON(http.StatusAccepted, body)
		}

	})

	PortString := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PortString == ":" {
		PortString = ":8080"
	}

	g.GET("/listusers", func(c *gin.Context) {
		log.Printf("/listusers\n")
		body := initDefaultResponse()

		if res, err := client.ListUsers(context.TODO(), &emptypb.Empty{}); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, body)
		} else {
			body["Status"] = "ok"
			body["List"] = res.Users
			c.JSON(http.StatusOK, body)
		}
	})

	g.Run(PortString)

}
