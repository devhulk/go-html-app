# Go Front-End App (Task List)

## Intro

I wanted to experiment with golang's html templating. I wanted to do it as native Go as possible ie trying not to use frameworks just pure golang htmltemplating. 

I am using tailwinds for styling and go/html for data rendering. 


## How to run

[Download Tailwind CSS CLI](https://tailwindcss.com/blog/standalone-cli)

I have three terminal panes running. One for dev, one for tailwinds, and one for the go server.

Pane One (Compiles css and watches for changes)
```
tailwindcss -i src/main.css -o src/output.css --watch
```

Pane Two (Run Go Server)
```
go run main.go
```


Next Steps:
- Add Postgres Integration
- Test deployment with Waypoint


