# Output Redirection

* “>” and “>>” both are output (STDOUT) direction operators, however, they differ in the following ways.

* “>” overwrites an already existing file or a new file is created providing the mentioned file name isn’t there in the directory. This means that while making changes in a file you need to overwrite certain any existing data, use the “>” operator.

```
echo “Welcome to Linux” > my_file_1.txt
```

After executing the above command, you’ll find that a text file “my_file_1.txt” is created in the directory. It’ll contain the text “Welcome to Linux” in it.  If there was already a file called, "my_file_1.txt" it would have been overwritten.

* “>>” operator appends an already present file or creates a new file if that file name doesn’t exist in the directory. 


# Resources

* [output redirection > and >>](https://www.shells.com/l/en-US/tutorial/Difference-between-%E2%80%9C%3E%E2%80%9D-and-%E2%80%9C%3E%3E%E2%80%9D-in-Linux#)