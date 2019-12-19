## Usage

### call the interface 
./main

### specify configuration file 
./main cofigFile.txt

user could type some sample inputs to get server information, type "quit()" to exit


## Explanations

### loadConfig():  
load the configuration file, build up a Trie tree.

### findRoute(input string) -> server string:  
serach the input with the built Trie tree, return the server accroding to the "most specific match wins" rule

### Notice
Here, it is assumed that every address is formatted as  "<customer_id>.<country>.<state>.<city>", and any part might be specified. In our implementation, we use a variable matchPriority to measure how "specific" the rule is.

And, if any rule is specified from the first element(customer_id) to itself or any other element, we don't need to use that variable.
