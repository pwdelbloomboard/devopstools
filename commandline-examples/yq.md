# yq

https://mikefarah.gitbook.io/yq/

yq is a portable command-line YAML processor (https://github.com/mikefarah/yq/) 
See https://mikefarah.gitbook.io/yq/ for detailed documentation and examples.

Usage:
  yq [flags]
  yq [command]

Available Commands:
  eval             Apply the expression to each document in each yaml file in sequence
  
  eval-all         Loads _all_ yaml documents of _all_ yaml files and runs expression once
  
  help             Help about any command
  
  shell-completion Generate completion script

Flags:
  -C, --colors                 force print with colors
  
  -e, --exit-status            set exit status if there are no matches or null or false is returned
  
  -f, --front-matter string    (extract|process) first input as 
  
  yaml front-matter. Extract will pull out the yaml content, process will run the expression against the yaml content, leaving the remaining data intact
      --header-preprocess      Slurp any header comments and separators before processing expression. This is a workaround for go-yaml to persist header content properly. (default true)
  
  -h, --help                   help for yq
  
  -I, --indent int             sets indent level for output (default 2)
  
  -i, --inplace                update the yaml file inplace of first yaml file given.
  
  -M, --no-colors              force print with no colors
  
  -N, --no-doc                 Don't print document separators (---)
  
  -n, --null-input             Don't read input, simply evaluate the expression given. Useful for creating yaml docs from scratch.
  
  -o, --output-format string   [yaml|y|json|j|props|p] output format type. (default "yaml")
  
  -P, --prettyPrint            pretty print, shorthand for '... style = ""'
      --unwrapScalar           unwrap scalar, print the value with no quotes, colors or comments (default true)
  
  -v, --verbose                verbose mode
  
  -V, --version                Print version information and quit

Use "yq [command] --help" for more information about a command.