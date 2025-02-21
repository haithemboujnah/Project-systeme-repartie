#include "calcul.h"
#include <stdio.h>

void
cal_bin_prog_1(char *host, int a, int b, int decimal, int decimaly, int choice)
{
    CLIENT *clnt;
    void *result_1;
    char *calnull_1_arg;
    long *result_2;
    Param puiss_1_arg;
    char **result_3;
    int dec2bin_1_arg;
    char **result_4;
    int dec2hex_1_arg;

    // Create client handle
    clnt = clnt_create(host, CAL_BIN_PROG, CAL_VERS_ONE, "udp");
    if (clnt == NULL) {
        clnt_pcreateerror(host);
        exit(1);
    }

    switch(choice) {
        case 1:
            // Call PUISS (calculate a^b)
            puiss_1_arg.a = a;
            puiss_1_arg.b = b;
            result_2 = puiss_1(&puiss_1_arg, clnt);
            if (result_2 == (long *)NULL) {
                clnt_perror(clnt, "call failed");
            } else {
                printf("Power: %ld\n", *result_2);
            }
            break;

        case 2:
            // Call DEC2BIN (convert decimal to binary)
            dec2bin_1_arg = decimal;
            result_3 = dec2bin_1(&dec2bin_1_arg, clnt);
            if (result_3 == NULL) {
                clnt_perror(clnt, "call failed");
            } else {
                printf("Binary: %s\n", *result_3);
            }
            break;

        case 3:
            // Call DEC2HEX (convert decimal to hexadecimal)
            dec2hex_1_arg = decimaly;
			result_4 = dec2hex_1(&dec2hex_1_arg, clnt);
			if (result_4 == NULL) {
				clnt_perror(clnt, "call failed");
			} else {
				printf("Hexadecimal: %s\n", *result_4);
			}
            break;

        default:
            printf("Invalid choice! Please choose between 1, 2, or 3.\n");
            break;
    }

    // Destroy client handle
    clnt_destroy(clnt);
}

int
main(int argc, char *argv[])
{
    char *host;
    int a, b, decimal, decimaly;
    char *endptr;
    int choice;

    host = argv[1];
    printf("Host: %s\n", host);

    while (1) {
        // Ask for the user's choice
        printf("\nChoose an operation:\n");
        printf("1: Power calculation (a^b)\n");
        printf("2: Decimal to Binary conversion\n");
        printf("3: Decimal to Hexadecimal conversion\n");
        printf("0: Exit\n");
        printf("Enter your choice (1/2/3/0): ");
        scanf("%d", &choice);

        if (choice == 0) {
            break;  // Exit the loop if the user chooses 0
        }

        // Validate choice
        if (choice < 1 || choice > 3) {
            fprintf(stderr, "Invalid choice! Please choose 1, 2, or 3.\n");
            continue;
        }

        // Ask for the relevant arguments based on the choice
        if (choice == 1) {
            printf("Enter value for a: ");
            scanf("%d", &a);
            printf("Enter value for b: ");
            scanf("%d", &b);
        } 

        if (choice == 2) {
            printf("Enter value for decimal: ");
            scanf("%d", &decimal);
        }

        // For DEC2HEX, we use a different variable `decimaly` for hex conversion
        if (choice == 3) {
            printf("Enter value for decimaly: ");
            scanf("%d", &decimaly);
        }

        // Call the chosen function
        cal_bin_prog_1(host, a, b, decimal, decimaly, choice);
    }

    printf("Exiting program...\n");
    exit(0);
}
