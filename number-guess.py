#!/usr/bin/python3
import random

highNumber = 100


def intro():
    print('------------------------------------')
    print('         Guess The Number Game')
    print('------------------------------------')
    print('Guess a number between 0 and '+str(highNumber)+':')


the_number = random.randint(0, highNumber)
intro()
score = 0
while True:
    score += 1
    guess_text = input()
    guess = int(guess_text)
    if guess < the_number:
        print(guess_text, "is too low")
    elif guess > the_number:
        print(guess_text, "is too High")
    else:
        print('you win!\nyour score was', score)
        exit(0)
