# lsof

* List of open files

## Usage

All sorts of usages in terms of showing what files are open:

```
lsof -p processID
```

For example:

```
$ lsof -p 78692
COMMAND   PID                          USER   FD   TYPE DEVICE  SIZE/OFF                NODE NAME
zsh     78692 patrick.delaneybloomboard.com  cwd    DIR    1,5      1920             1468593 /Users/patrick.delaneybloomboard.com/Projects/devopstools/commandline-examples
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5   1347856 1152921500312764850 /bin/zsh
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    137696 1152921500312767269 /usr/lib/zsh/5.8/zsh/langinfo.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    138640 1152921500312767301 /usr/lib/zsh/5.8/zsh/terminfo.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    579888 1152921500312767305 /usr/lib/zsh/5.8/zsh/zle.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    341920 1152921500312767253 /usr/lib/zsh/5.8/zsh/complete.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    174352 1152921500312767315 /usr/lib/zsh/5.8/zsh/zutil.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    139088 1152921500312767261 /usr/lib/zsh/5.8/zsh/datetime.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    138880 1152921500312767295 /usr/lib/zsh/5.8/zsh/stat.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    174752 1152921500312767287 /usr/lib/zsh/5.8/zsh/parameter.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    211600 1152921500312767255 /usr/lib/zsh/5.8/zsh/complist.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    138448 1152921500312767289 /usr/lib/zsh/5.8/zsh/regex.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    139600 1152921500312767273 /usr/lib/zsh/5.8/zsh/mathfunc.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    138048 1152921500312767307 /usr/lib/zsh/5.8/zsh/zleparameter.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5    208784 1152921500312767257 /usr/lib/zsh/5.8/zsh/computil.so
zsh     78692 patrick.delaneybloomboard.com  txt    REG    1,5   2547856 1152921500312767024 /usr/lib/dyld
zsh     78692 patrick.delaneybloomboard.com    0u   CHR   16,0 0t1637831                 891 /dev/ttys000
zsh     78692 patrick.delaneybloomboard.com    1u   CHR   16,0 0t1637831                 891 /dev/ttys000
zsh     78692 patrick.delaneybloomboard.com    2u   CHR   16,0 0t1637831                 891 /dev/ttys000
zsh     78692 patrick.delaneybloomboard.com   10u   CHR   16,0  0t107836                 891 /dev/ttys000
```

Basically this shows all of the files open for the zshell process, the number of which was identified with the command, "ps" - essentially, this shows all of the files open on zshell, the size, node, and location for each.