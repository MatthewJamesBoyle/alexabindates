package main

import (
	"github.com/MatthewJamesBoyle/bindates/parser"
	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

// Handler is the lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
	return DispatchIntents(request), nil
}

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "allbins":
		response = handleResponse(ALLBINS)
	case "rubbish":
		response = handleResponse(RUBBISH)
	case "recycling":
		response = handleResponse(RECYCLING)
	case "food":
		response = handleResponse(FOOD)
	case alexa.HelpIntent:
		response = handleHelp()
	}

	return response
}

func handleHelp() alexa.Response {
	return alexa.NewSimpleResponse("Hello", "Hope you enjoy the app")
}

func handleResponse(collectionType int) alexa.Response {
	dates := parser.Parse(parser.Config{
		os.Getenv("URL"),
		os.Getenv("HOUSE_NUMBER"),
		os.Getenv("POSTCODE"),
	})

	switch collectionType {
	case RUBBISH:
		return alexa.NewSimpleResponse("Your next rubbish collection is on", dates["rubbish"])
	case RECYCLING:
		return alexa.NewSimpleResponse("Your next recyling collection is on", dates["recycling"])
	case FOOD:
		return alexa.NewSimpleResponse("Your next food waste collection is on", dates["food"])
	default:
		return alexa.NewSimpleResponse("Your next collections are on ",
			dates["rubbish"]+" for rubbish, "+dates["recycling"]+" for recycling "+"and "+dates["food"]+" for food")
	}

}

func main() {
	lambda.Start(Handler)
}
