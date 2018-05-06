#import random
import random

print('------------------------------------')
print('         Guess The Number Game')
print('------------------------------------')
print()
# alt + enter will try to autofix with red bulb in pycharm
the_number = random.randint(0,100)
guess = -1


while guess != the_number:
    guess_text = input('Guess a number between 0 and 100: ')
    # conver to an integer
    guess = int(guess_text)

    if guess < the_number:
        print("Too low")
    elif guess > the_number:
        print("Too High")
    else:
        print('you win!')
