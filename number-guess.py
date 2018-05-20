#!/usr/bin/python3
import random

def intro ():
    print('------------------------------------')
    print('         Guess The Number Game')
    print('------------------------------------')
    print('Guess a number between 0 and 100:')


the_number = random.randint(0, 100)
intro()

while True:
    guess_text = input()
    guess = int(guess_text)
    if guess < the_number:
        print(guess_text, "is too low")
    elif guess > the_number:
        print(guess_text, "is too High")
    else:
        print('you win!')
        exit(0)
