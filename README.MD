# ASCII ART OUTPUT
## Description
This project involves creating a command-line program that generates text-based output based on user inputs and writes the result to a file specified by a flag.The project is based on our Ascii-Art and it generates art for the printable string passed.
## Authors
- [Maina Ann](https://github.com/nyagooh)
- [Ray ogwel](https://github.com/anxielray)
- [Allan Kamau](https://github.com/Githaiga22)
## Usage

### Recquirements
 - Go (version 1.15 or higher recommended)
- Access to a terminal or command prompt


## Functionality
### Command-Line Interface

The program accepts command-line arguments to determine its operation.
The basic format for running the program is:
```bash
go run . [OPTION] [STRING] [BANNER]
```
### Output File

The --output=<fileName.txt> flag specifies the name of the file where the program's output will be written.
<fileName.txt> should be replaced with the desired file name.
The flag must follow the exact format --output=<fileName.txt>. Any deviation from this format will trigger a usage message.
### Usage Message
If the flag is not in the correct format or incase of any error, the program will display the following usage message:
```bash
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```
### Arguments and Options

The program can handle different types of [OPTION] and [BANNER] based on the ASCII-art projects implemented.
It must also handle the case where only a single [STRING] argument is provided, without any [OPTION] or [BANNER]

### Implementation Details

The program should be able to generate ASCII-art or text-based representations based on the input string and optional textual files.The textual file include:
 - standard.txt
 - thinkertoy.txt
 - shadow.txt

The output will be formatted according to the specified options and then written to the file designated by the --output flag.

### Running the Application

1. Clone the repository:
```bash
git clone https://learn.zone01kisumu.ke/git/rogwel/ascii-art-output

cd ascii-art-output
```
2. Run the program
```bash
go run  . --output=filename.txt hello standard
```
The program will create a file called filename.txt and write the string hello in its Ascii Art format depending on the file specified.To see whats written in the file:
```bash
cat -e filename.txt
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
```
```bash
go run . --output=filename.txt hello
```
In this case our program will take the default banner file that is standard and write the string hello in its ascii art form in the file filename.txt
``` bash
cat - e filename.txt
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
```
since the program is built on other Ascii art programs,the program will run in the following instances:

```bash
go run . hello

 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
 ```
The default banner file is standard.txt

```bash
go run . hello thinkertoy


o        o o     
|        | |     
O--o o-o | | o-o 
|  | |-' | | | | 
o  o o-o o o o-o 
                 
                 
```
The program prints the string in  specified banner file format
## Contributing

Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.
