grammar tiger;

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
expr : term ((PLUS | MINUS) term)*;

term : factor ((MULTIPLY | DIVIDE) factor)*;

factor : LPAREN expr RPAREN
       | DIGIT+;