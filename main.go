package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shin5ok/shoutouthostnamegcp"
	pb "github.com/shin5ok/urlmap-front/pb"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	g          = gin.Default()
	version    = "0.21"
	portString = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
	shoutouthostnamegcp.SetSigHandler(os.Getenv("SLACK_URL"), os.Getenv("SLACK_CHANNEL"))

}

func initDefaultResponse() map[string]interface{} {
	return map[string]interface{}{
		"Status":  "fail",
		"Message": "",
	}
}

func Ping(c *gin.Context) {
	log.Info().Str("method", "Ping").Send()
	body := initDefaultResponse()
	body["Status"] = "ok"
	body["Message"] = "Pong"
	body["Version"] = version
	c.JSON(http.StatusOK, body)
}

func CreateRouter() *gin.Engine {

	host := os.Getenv("URLMAP_API")
	if host == "" {
		host = "urlmap-api:8080"
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		log.Error().Err(err)
	}
	client := pb.NewRedirectionClient(conn)

	g.GET("/", func(c *gin.Context) {
		body := initDefaultResponse()
		body["Status"] = "ok"
		log.Info().
			Str("path", "/").
			Str("status", body["Status"].(string))
		c.JSON(http.StatusOK, body)
	})

	g.GET("/Ping", Ping)

	g.GET("/info/:u", func(c *gin.Context) {
		user := c.Param("u")
		log.Info().Msgf("/info/%s", user)
		body := initDefaultResponse()
		fmt.Println(body)

		u := &pb.User{User: user}
		res, err := client.GetInfoByUser(context.TODO(), u)
		if err != nil {
			log.Error().Err(err)
			body["Message"] = err.Error()
			c.JSON(http.StatusBadRequest, body)
		}
		body["Status"] = "ok"
		body["Data"] = res
		c.JSON(http.StatusOK, body)
	})

	g.GET("/get/:p", func(c *gin.Context) {
		path := c.Param("p")
		log.Info().Msgf("/get/%s", path)
		body := initDefaultResponse()

		rpath := &pb.RedirectPath{Path: path}

		if res, err := client.GetOrgByPath(context.TODO(), rpath); err != nil {
			log.Error().Err(err)
			c.JSON(http.StatusInternalServerError, body)
		} else {
			body["Status"] = "ok"
			body["Org"] = res.GetOrg()
			c.JSON(http.StatusOK, body)
		}
	})

	g.POST("/user/:u", func(c *gin.Context) {
		data := &pb.User{}
		err := c.Bind(&data)

		if err != nil {
			log.Error().Err(err)
			c.JSON(http.StatusBadRequest, err)
		}

		body := initDefaultResponse()

		user := c.Param("u")
		log.Info().Msgf("/user/%s create or update", user)
		u := &pb.User{User: user, NotifyTo: data.NotifyTo}

		if res, err := client.SetUser(context.TODO(), u); err != nil {
			log.Info().Msgf("%+v", err)
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

		body := initDefaultResponse()

		user := c.Param("u")
		log.Info().Msgf("/user/%s delete and user's entries", user)
		u := &pb.User{User: user}

		if res, err := client.RemoveUser(context.TODO(), u); err != nil {
			log.Error().Err(err)
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
		log.Info().Msg("/register")
		data := &pb.RedirectData{}
		err := c.Bind(&data.Redirect)
		log.Debug().Msgf("%+v", data)

		if err != nil {
			log.Error().Err(err)
			c.JSON(http.StatusBadRequest, err)
		}

		log.Info().Msgf("%+v", data.Redirect)

		body := initDefaultResponse()
		if res, err := client.SetInfo(context.TODO(), data); err != nil {
			log.Error().Err(err)
			body["Message"] = err.Error()
			c.JSON(http.StatusInternalServerError, body)
		} else {
			log.Info().Msgf("%+v", res)
			body["Status"] = "ok"
			body["Data"] = res.GetOrg()
			c.JSON(http.StatusAccepted, body)
		}

	})

	g.GET("/listusers", func(c *gin.Context) {
		log.Info().Msg("/listusers\n")
		body := initDefaultResponse()

		if res, err := client.ListUsers(context.TODO(), &emptypb.Empty{}); err != nil {
			log.Error().Err(err)
			c.JSON(http.StatusInternalServerError, body)
		} else {
			body["Status"] = "ok"
			body["List"] = res.Users
			c.JSON(http.StatusOK, body)
		}
	})
	return g

}

func main() {
	// for prom agent ref: https://github.com/penglongli/gin-metrics
	forProm := gin.Default()

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.UseWithoutExposingEndpoint(g)
	m.Expose(forProm)

	go func() {
		forProm.Run(":18080")
	}()

	if portString == ":" {
		portString = ":8080"
	}

	g := CreateRouter()
	g.Run(portString)

}
