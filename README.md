This is a simple cmd level quiz app which takes a csv file of questions, answers (using a default csv file with few mathematical problems), 
and loop through it and ask questions and shows your score at the end.
You can provide your preferred time limit in seconds and challenge yourself just as you want!!

Run command "./quiz -h" to get the format of csv file and available commands

Run the application by using the following commands
1. "go build ."
2. "./quiz -csv="example.csv" -limit=2

This simple app is made of go concurrency and doesn't use any third party dependencies. 
