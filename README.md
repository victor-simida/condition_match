# condition_match
Condition_match is a structure to describe some conditions for input key-values, condition_match use bit-field instead of "&&" and "||" symbols, for example, byte '011' means a=xxx||b=xxx&&c=xxx&&d=xxx.

The "&&" operator has higher precedence than the "||".
