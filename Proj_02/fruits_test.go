package examples

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gavv/httpexpect/v2"
)

func TestFriuts(T *testing.T){
	handler :=FruitsHandler()
	server :=httptest.NewServer(handler)
	defer server.Close()
	e := httpexpect.New(t,server.URL)
	e.GET("/fruits").
		Expect().
		Status(http.StatusOK).JSON().Array().Empty()

		orange :=map[string]interface{}{
			"weight":100,
		}
		e.PUT("/fruits/orange").WithJSON(orange).
		 	Expect().
			Status(http.StatusContent).NoContent()
		
		apple :=map[string]interface{}{
			"colors":[]interface{}{"green","red"},
			"weight":200,
		}
		e.PUT("/fruits/apple").WithJSON(apple).
		 	Expect().
			Status(http.StatusContent).NoContent()
		e.GET("/fruits").
			Expect().
			status(http.StatusOK).JSON().Array().ContainsOnly("orange","apple")	
		e.GET("/fruits/orange").
			Expect().
			status(http.StatusOK).JSON().Array().Equal(orange).NotEqual(apple)	
		e.GET("/fruits/orange").
			Expect().
			status(http.StatusOK).
			JSON().Object().ContainsKey("weight").ValueEqual("weight",100)
		obj :=e.GET("/fruits/apple").
			Expect().
			status(http.StatusOK).JSON().Object()

		obj.Key().ContainsOnly("weight","colors")	

		obj.Value("colors").Array().Element("green","red")
		obj.Value("colors").Array().Element(0).String().Equal("green")
		obj.Value("colors").Array().Element(1).String().Equal("red")

		obj.Value("weight").Number().Equal(200)
		e.GET("/fruits/melon").
		Expect().
			status(http.StatusNotFound).


}