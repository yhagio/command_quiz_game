# Quiz Game

### How to run

```bash
git clone 
cd quiz_game

go build .
./quiz_game
```

Customize quiz
```bash
# HELP menu
./quiz_game -h

# Use your own csv file
go build . && ./quiz_game -csv "your_custom_questions.csv" 

# Change default timer to 5 seconds
go build . && ./quiz_game -limit=5
```