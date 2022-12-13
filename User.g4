grammar User;

term: user EOF;

user:
	'userset' '(' ID ',' ID ',' ID ')'
	| 'object' '(' ID ',' ID ')'
	| ID;

ID: [a-zA-Z][a-zA-Z0-9_-]*;
WS: [ \t\n\r\f]+ -> skip;
