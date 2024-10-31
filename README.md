# Norsk Klokka GO

Norsk Klokka GO is a simple command-line application that helps users practice telling time in Norwegian. The program generates random times in a 24-hour format and checks user input against correct and accepted Norwegian phrases for those times.

## Features

- **Random Time Generation**: The program randomly generates a time in the 24-hour format (00:00 to 23:59) for users to practice.
  
- **Norwegian Time Formats**: Users can respond with various Norwegian phrases to indicate the time, accommodating several accepted formats, such as:
  - Exact hour (e.g., "klokka ni")
  - Minutes past the hour (e.g., "ti over ni")
  - Half past (e.g., "halv ti")
  - Minutes to the next hour (e.g., "fem på ti")
  - Quarter past/to (e.g., "kvart over ni" or "kvart på ti")

- **Correctness Feedback**: After a user inputs their response, the program checks if the answer is correct, providing immediate feedback. It indicates if the response is right or wrong.

- **Alternative Accepted Answers**: If the user's answer is correct, the program also lists alternative accepted formats for the same time. If incorrect, users can press 's' to see the correct answer.

- **User Interaction**: The program prompts users to try again after a wrong answer or offers an option to exit the program.
