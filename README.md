## Small back story
I've started with Rust, but I thought that it would be a good idea to check how it's main rival is doing - ofc I mean Golang (just want to be clear cause you could thought about C++, or even C). This time I used a little different approach, the course from [gophercises](https://gophercises.com/) has an higher entry level - you should have a basic programming knowledge and not to be afraid of the cold documentation.

## About
This app is ready to read the questions from the given csv file, which should contain the questions with answers written like this:`1+1,2`. The questions from the file will be presented to the user, his task will be to answer them without exceeding the specified time limit 

The two possible flags to use:

| flag name | description |
| --- | --- |
| `csv` <br><br>Default: `"problems.csv"`| With this flag you can enter the name of the file with the questions |
| `limit` <br><br>Default: `30`| With this flag you can set how much time you will have to answer to the question (the number corresponds to seconds) |

So full command should look like this: `./main -csv="problems.csv" -limit=5`

`main` - is an binary file - So you don't need to download the Go to check how the code works :wink: