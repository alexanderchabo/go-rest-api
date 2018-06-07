package main

const (
	dbuser = "postgres"
	dbpass = ""
	dbname = "APP_DB"
)

func main() {
	a := App{}
	a.Initialize(dbuser, dbpass, dbname)
	a.Run(":8080")
}
