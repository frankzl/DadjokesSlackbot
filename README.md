# DadJokes Slack bot

### Requirements
1. go
2. gorilla/mux go package ``go get -u github.com/gorilla/mux``
3. ngrok

### Structure
1. API to tell the dad jokes (we are redirecting https://icanhazdadjoke.com/ to our custom API)
2. Slack app to serve as a client that connects to our API

### Run
Running the API
```
go run .	# Start API locally on localhost:8080
ngrok http 8080 # Forwards requests on public server to our local machine
```

Creating the App

1. Optional: Create your own workspace to test stuff
2. Create a Slack app from scratch through api.slack.com/apps
3. Install your app to your workspace using the website
4. Add features and functionality -> slash commands
5. Add slash command: /dadjokes
6. Fill in the URL where your API is running (URL where ngrok serves)
7. Try /dadjokes in any chat of your workspace
8. ****Insert laugh\*\***
