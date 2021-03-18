package main

import (
	"fmt"
	"strings"

	p "github.com/pganalyze/pg_query_go/v2"
)

func main() {
	// Transaction statement
	txn := `
BEGIN;
UPDATE accounts SET balance = balance - 100.00
    WHERE name = 'Alice';
SAVEPOINT my_savepoint;
UPDATE accounts SET balance = balance + 100.00
    WHERE name = 'Bob';
-- oops ... forget that and use Wally's account
ROLLBACK TO my_savepoint;
UPDATE accounts SET balance = balance + 100.00
    WHERE name = 'Wally';
COMMIT;
`
	t, err := p.Parse(txn)
	if err != nil {
		panic(err)
	}
	PrintStatements(t)

	// Simple statement
	s := "SELECT * FROM accounts"
	r_sel, err := p.Parse(s)
	if err != nil {
		panic(err)
	}
	PrintStatements(r_sel)

	// Create statement
	v := `
CREATE VIEW myview AS
SELECT city, temp_lo, temp_hi, prcp, date, location
FROM weather, cities
WHERE city = name;
`
	r_create, err := p.Parse(v)
	if err != nil {
		panic(err)
	}
	PrintStatements(r_create)

}

func PrintStatements(r *p.ParseResult) {
	fmt.Println("---")
	if len(r.Stmts) > 0 {
		s := r.Stmts[0].Stmt.String()
		fmt.Println(s[:strings.IndexByte(s, ':')])
	}
	for k, v := range r.GetStmts() {
		fmt.Println(k+1, ".", v)
	}
	fmt.Println("")
}
