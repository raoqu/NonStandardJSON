grammar NJson;

json
 : object
 | array
 ;

object
 : '{' pair (',' pair)* '}'
 | '{' '}'
 ;

pair
 : STRING ':' value
 ;

array
 : '[' value (',' value)* ']'
 | '[' ']'
 ;

value
 : STRING
 | NUMBER
 | 'true'
 | 'false'
 | 'null'
 | object
 | array
 ;

STRING
 : '"' (ESC | SAFECODEPOINT)* '"'
 ;

ESC
 : '\\' (["\\/bfnrt] | UNICODE)
 ;

UNICODE
 : 'u' HEX HEX HEX HEX
 ;

HEX
 : [0-9a-fA-F]
 ;

NUMBER
 : '-'? INT '.' [0-9]+ EXP? 
 | '-'? INT EXP
 | '-'? INT
 ;

INT
 : '0' 
 | [1-9] [0-9]*
 ;

EXP
 : [eE] [+\-]? [0-9]+
 ;

SAFECODEPOINT
 : ~["\\]
 ;

WS: [ \t\n\r] + -> skip ; // skip spaces, tabs, newlines
