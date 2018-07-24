# Quiz Game

### How to run

```bash
git clone git@github.com:yhagio/command_quiz_game.git quiz
cd quiz

go build .
./quiz
```

Customize quiz
```bash
# HELP menu
./quiz -h

# Use your own csv file
go build . && ./quiz -csv "your_custom_questions.csv" 

# Change default timer to 5 seconds
go build . && ./quiz -limit=5
```