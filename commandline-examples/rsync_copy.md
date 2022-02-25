# rsync

a fast, versatile, remote (and local) file-copying tool

# About

rsync provides a number of options that control how the command behaves. The most widely used options are:

-a, --archive, archive mode, equivalent to -rlptgoD. This option tells rsync to syncs directories recursively, transfer special and block devices, preserve symbolic links, modification times, groups, ownership, and permissions.
-z, --compress. This option forces rsync to compresses the data as it is sent to the destination machine. Use this option only if the connection to the remote machine is slow.
-P, equivalent to --partial --progress. When this option is used, rsync shows a progress bar during the transfer and keeps the partially transferred files. It is useful when transferring large files over slow or unstable network connections.
--delete. When this option is used, rsync deletes extraneous files from the destination location. It is useful for mirroring.
-q, --quiet. Use this option if you want to suppress non-error messages.
-e. This option allows you to choose a different remote shell. By default, rsync is configured to use ssh.

# Basic Usage

```
 rsync -a rsync.md rsync_copy.md
```

That command above directly copies rsync.md, to a new file called rsync_copy.md.

Running this command again simply overwrites the rsync_copy.md file with a newly updated file.