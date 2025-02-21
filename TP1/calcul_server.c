#include "calcul.h"
#include <math.h>
#include <stdio.h>

/*
 * Dummy function, does nothing.
 */
void *
calnull_1_svc(void *argp, struct svc_req *rqstp)
{
    static char *result = "NULL";
    return (void *)&result;
}

/*
 * Calculate a^b.
 */
long *
puiss_1_svc(Param *argp, struct svc_req *rqstp)
{
    static long result;
    result = pow(argp->a, argp->b); // Calculate a^b
    return &result;
}

/*
 * Convert a decimal number to binary and return as a string.
 */

char **dec2bin_1_svc(int *argp, struct svc_req *rqstp)
{
    static char *result = NULL;  // Static pointer to hold the result
    static char buffer[64];      // Static buffer to store the binary string

    if (argp == NULL) {
        printf("Received NULL argument in dec2bin_1_svc\n");
        return NULL;  // Return NULL if the argument is invalid
    }

    int decimal = *argp;
    int index = 0;

    // Convert decimal to binary string
    if (decimal == 0) {
        snprintf(buffer, sizeof(buffer), "0");
    } else {
        while (decimal > 0) {
            buffer[index++] = (decimal % 2) + '0';
            decimal /= 2;
        }
        buffer[index] = '\0';

        // Reverse the string to get correct binary representation
        int start = 0, end = index - 1;
        while (start < end) {
            char temp = buffer[start];
            buffer[start] = buffer[end];
            buffer[end] = temp;
            start++;
            end--;
        }
    }

    result = buffer;  // Point the result to the buffer
    return &result;   // Return the address of the static pointer
}




/*
 * Convert a decimal number to hexadecimal.
 */
char **dec2hex_1_svc(int *argp, struct svc_req *rqstp) {
    static char *result = NULL;  // Static pointer to hold the result
    static char buffer[64];      // Static buffer to store the hexadecimal string

    if (argp == NULL) {
        printf("Received NULL argument in dec2hex_1_svc\n");
        return NULL;  // Return NULL if the argument is invalid
    }

    snprintf(buffer, sizeof(buffer), "%X", *argp);  // Format as hexadecimal (uppercase)

    result = buffer;  // Point the result to the buffer
    return &result;   // Return the address of the static pointer
}
