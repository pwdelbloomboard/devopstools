# systemctl

https://man7.org/linux/man-pages/man1/systemctl.1.html

* systemctl - Control the systemd system and service manager

systemctl may be used to introspect and control the state of the
       "systemd" system and service manager. Please refer to systemd(1)
       for an introduction into the basic concepts and functionality
       this tool manages.

https://man7.org/linux/man-pages/man1/systemd.1.html

systemd is a system and service manager for Linux operating
       systems. When run as first process on boot (as PID 1), it acts as
       init system that brings up and maintains userspace services.
       Separate instances are started for logged-in users to start their
       services.

       systemd is usually not invoked directly by the user, but is
       installed as the /sbin/init symlink and started during early
       boot. The user manager instances are started automatically
       through the user@.service(5) service.

       For compatibility with SysV, if the binary is called as init and
       is not the first process on the machine (PID is not 1), it will
       execute telinit and pass all command line arguments unmodified.
       That means init and telinit are mostly equivalent when invoked
       from normal login sessions. See telinit(8) for more information.

       When run as a system instance, systemd interprets the
       configuration file system.conf and the files in system.conf.d
       directories; when run as a user instance, systemd interprets the
       configuration file user.conf and the files in user.conf.d
       directories. See systemd-system.conf(5) for more information.
