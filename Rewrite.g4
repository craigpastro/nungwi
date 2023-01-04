grammar Rewrite;

term: rewrite EOF;

rewrite:
	'self'
	| 'computedUserset' '(' ID ')'
	| 'tupleToUserset' '(' ID ',' ID ')'
	| 'union' '(' rewrite ',' rewrite ')'
	| 'intersection' '(' rewrite ',' rewrite ')'
	| 'exclusion' '(' rewrite ',' rewrite ')';

ID: [a-zA-Z0-9_-]+;
WS: [ \t\n\r\f]+ -> skip;
