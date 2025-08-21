# GO-QUIZ-TIMER
A simple timer based command-line quiz application using Go that reads questions from a CSV file.

## 📌 Features  
- Reads quiz questions and answers from a CSV file  
- Timer support (default: 5 seconds per question)  
- Counts and displays your score  
- Supports custom CSV files and timer duration  

## 📂 Project Files  
- `main.go` → Go source code for the quiz app  
- `quiz.csv` → Sample quiz questions in CSV format  

## ⚙️ Installation  
1. Clone the repository:  
   ```bash
   git clone https://github.com/<your-username>/go-quiz-timer.git
   cd go-quiz-timer

2. Build the application:
   ```bash
   go build -o quiz

3. Run the quiz:
   ```bash
   ./quiz -f quiz.csv

## 📂 CSV File Format
Each line of the CSV file contains a question and its correct answer, separated by a comma.
Example:
```text
5+5,10
7*3,21
12-4,8
```
➡️ Format: `question,answer`
➡️ Example: `5+5,10`

## 🎮 Usage
Run the program with default sample CSV file:
```bash
  ./quiz -f quiz.csv
```
Runs the quiz using quiz.csv with 5 seconds per question.

Specify a custom CSV file:
```bash
  ./quiz -f myquiz.csv
```
Runs the quiz with a different question file.

## 📝 Example Output
```text
You have 30 seconds to answer 3 questions
Problem 1: 5+5= 10
Problem 2: 7*3= 21
Problem 3: 12-4= 8

Your result is 3 out of 3
Press enter to exit
```
