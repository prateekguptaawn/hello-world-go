# Enhanced Hello World In Go

### Requirements:
#### Create a program, hello, which will do the following:
1. When run as: `./hello` will output:  `Hello, World!`
2. When run as `./hello Rob` will output: `Hello, Rob!` (reads command line arguments)
3. When run as `HELLO_NAME=Sally ./hello` will output: `Hello, Sally!` (reads from environment variable HELLO_NAME)
4. When run as `HELLO_MSG=Bonjour ./hello`  will output: `Bonjour, World!` (reads from environment variable HELLO_MSG to change the greeting).
5. When run as `HELLO_MSG='Greetings & Salutations' HELLO_NAME=Prateek ./hello` will output: `Greetings & Salutations Prateek!`
6. Uses zap.SugaredLogger to output formatted debug logs:
   1. If no environment or command line args are used; OR 
   2. The value of the command line argument if one is present; OR 
   3. The value of the relevant environment variables (`HELLO_MSG, HELLO_NAME` ) if supplied.
      

Hint: zap logging is found here: https://github.com/uber-go/zap