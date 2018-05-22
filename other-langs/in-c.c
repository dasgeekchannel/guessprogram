#include <stdlib.h>
#include <stdio.h>
#include <time.h>

int randomNumber(void) {
	int i, n;
    time_t t;
    srand((unsigned) time(&t));
    return rand() % 100;
}

void intro(void) {
    printf("------------------------------------\n");
	printf("       Guess The Number Game        \n");
	printf("------------------------------------\n");
	printf("Guess a number between 0 and 100:\n");
}

int userInput(void) {
    int data;
    int result = scanf ("%d",&data);
    return data;
}

int main(int argc, char *argv[]) {
	intro();
	int theNumber = randomNumber();
	for ( int x = 0 ; x < 1000 ; x++ ) { // if u need a thousand guesses .....
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
