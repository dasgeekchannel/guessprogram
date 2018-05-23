#include <stdlib.h>
#include <stdio.h>
#include <time.h>

const int highNumber = 1000;

int randomNumber(void) {
    time_t t;
    srand((unsigned) time(&t));
    return rand() % highNumber;
}

void intro(void) {
    printf("------------------------------------\n");
    printf("       Guess The Number Game        \n");
    printf("------------------------------------\n");
    printf("Guess a number between 0 and %d:\n", highNumber);
}

int userInput(void) {
    int data;
    scanf ("%d",&data);
    return data;
}

int main(void) {
    intro();
    int theNumber = randomNumber();
    for ( int x = 0 ; x < highNumber ; x++ ) { // if u need a thousand guesses .....
        int theGuess = userInput();
        if ( theGuess < theNumber ) {
            printf("%d is too low\n", theGuess);
        }
        if ( theGuess > theNumber ) {
            printf("%d is too high\n", theGuess);
        }
        if ( theGuess == theNumber ) {
            printf("you win\nyour score was %d\n", x+1);
            return 0;
        }
    }
}
