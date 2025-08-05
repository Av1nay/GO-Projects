	
 List of files and directories in the current directory: ls:
hello.sh
listing.sh
theDirectory

 Appends a / to directory names: 	
 Command:ls -p
 Output:
hello.sh
listing.sh
theDirectory/


List of files and directories using only commas as separators.
Command: ls -p | sed 's/,$//'
hello.sh,listing.sh,theDirectory/

List of files and directories using only commasand spaces as separators.
Command: ls -p | sed 's/,$//; s/,/, /g
hello.sh
listing.sh
theDirectory/


List of files and directories using only commasand spaces as separators.
Command: ls -p| tr | sed 's/,$//; s/,/, /g
