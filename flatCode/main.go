package main

func main() {
	runProject()
	/*urlExample := "postgres://MyFlat:12345@localhost:5436/test_db"
	_, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stderr, "ABLE to connect to database: %v\n", err)
		fmt.Println("ready")
	}*/
	/*defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	os.Exit(1)
	}

	fmt.Println(name, weight)*/
}
