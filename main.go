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
	r_complex, err := p.Parse(txn)
	if err != nil {
		panic(err)
	}
	PrintStatements(r_complex)

	// Simple statement
	sel := "SELECT * FROM accounts"

	r_sel, err := p.Parse(sel)
	if err != nil {
		panic(err)
	}
	PrintStatements(r_sel)
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
