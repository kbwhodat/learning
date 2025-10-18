#include <stdio.h>
#include <stdlib.h>
#include <libpq-fe.h>
#include "settings.h"

int main (int argc, char *argv[])
{
    // Connect to the database.
    PGconn *conn = PQconnectdb("host=" HOSTNAME " dbname=" DATABASE " user=" USERNAME " password=" PASSWORD);

    if (PQstatus(conn) == CONNECTION_BAD) {
        fprintf(stderr, "Connection to database failed: %s\n", PQerrorMessage(conn));
        PQfinish(conn);
        exit(0);
    }

    PGresult *query = PQexec(conn, "select *  from cd.facilities;");
    //int num_of_rows = PQntuples(query);

    // printf("%d ", num_of_rows);

    // PQclear(query);
    // printf("clearing query sotrage");
    PQprintOpt options = {
        .header = 1,
        .align = 1,
        .standard = 0,
        .html3 = 0,
        .expanded = 0,
        .pager = 0,
        .fieldSep = "|",
        .tableOpt = "",
        .caption = "User List",
        .fieldName = NULL,
    };

    FILE *f = fopen("results", "w");
    PQprint(f, query, &options);
    fclose(f);

    FILE *file = fopen("results", "r");
    int c;
    if (file) {
        while ((c = getc(file)) != EOF)
            putchar(c);
        fclose(file);
    }

    // Close connection.
    PQfinish(conn);
    printf("Disconnected!\n");

    return 0;
}
