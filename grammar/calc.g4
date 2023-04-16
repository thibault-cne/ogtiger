grammar calc;

/*
 * Lexer rules
 */

// Tokens
DIGIT : [0-9];
PLUS : '+';
MINUS : '-';
MULTIPLY : '*';
DIVIDE : '/';
LPAREN : '(';
RPAREN : ')';

// Ignored characters
WS : [ \t\n\r]+ -> skip;

/*
 * Parser rules
 */

// Entry rule
expr : term ((MULTIPLY | DIVIDE) term)*;

term : factor ((PLUS | MINUS) factor)*;

factor : LPAREN expr RPAREN
       | DIGIT+;