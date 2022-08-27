This is a simple cmd level quiz app which takes a csv file of questions, answers (using a default csv file with few mathematical problems), 
and loop through it and ask questions and shows your score at the end.
You can provide your preferred time limit in seconds and challenge yourself just as you want!!
This simple app is made of go concurrency and doesn't use any third party dependencies.

Default folder name is "quiz". If you wanna use any other folder name, replace the word "quiz" with your folder name in the below commands.
Run command "./quiz -h" to get the format of csv file and available commands.

Run the application by using the following commands.
1. "go build ."
2. "./quiz -limit=2"
3. If you wanna provide your own csv file (make sure it's in the correct format), copy it to the root folder and run the following command 
"./quiz csv="yourCsvFileName.csv" -limit=2" 
