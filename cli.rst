cli
###

**micropesce** has a minimal cli interface.

synopsis
========

    micropesce <options...> <db_file>

    <db_file> is your instance's sqlite3 database file.

    -wizard
       instead of the usual micropesce instance, start a cli
       configuration wizard. if you pass <db_file>, the wizard will
       modify the configuration in that file. else a new db file will
       be created.

    -version
       display micropesce version and quit.
