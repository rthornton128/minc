# minc
Minimal C
---------

Minimal C is a toy language intended to help facilitate learning how to write
a compiler. In fact, while based primarily on C it is not a subset of C and
will differ in some respects.

Sources written for Minimal C will likely not compile with a C compiler and 
will not be API/ABI compatible.

MinC is also intended to be implemented incrementally with each stage
adding additional functionality to the language.

Language Overview: Phase 1
--------------------------

MinC contains a single type: void.

Identifiers must start with a letter and may contain one or more letters or
numbers. They may also contain an underscore.

Braces {} are used to encapsulate statement blocks.
Parenthases are used to encapsulate function arguments.

The only valid program is: void main() {}

