/* calcul.x */
program CAL_BIN_PROG {
    version CAL_VERS_ONE {
        void CALNULL(void) = 0;
        long PUISS(Param) = 1;
        string DEC2BIN(int) = 2;
        string DEC2HEX(int) = 3;
    } = 1;
} = 0x20000001; 

struct Param {
    int a;
    int b;
};