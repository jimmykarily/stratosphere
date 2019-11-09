# Stratosphere

A cross platform desktop application that allows you to read the data from your Amazfit Stratos smartwatch.

## Why

- Because there is not Desktop application that does that.
- Because if the official application is no longer supported, you will still have your watch.
- Because the official application is missing some data the watch shows you (e.g. avg pace per kilometer)
- Because free software allows you to do more with your data. Your data, your code, your choices.
- Because it's fun.

## Scope

This application is being tested only with Amazfit Stratos. The same company produces more watches that can be used with the same mobile application so it is possible that this application works with those watches too. Contributions are welcome (a database from another watch for example).

## Prerequisites

This application will not try to talk to your watch (for now). You need to export the database yourself and feed it to this application. See the section below on how to achieve that.

## Extracting the database from your watch

_Based on this original post: https://forum.xda-developers.com/smartwatch/amazfit/xiaomi-huami-amazfit-export-data-t3533292_

First you need to have adb installed on your system. Google around on how to do that. Connect your watch to the charging base and that to your machines USB.

You now need to find what the relevant package is called in your watch. In my case, the package to backup had a different name in my watch compared to the instructions above. To find the right package run `adb shell` and from the shell run `pm list packages | grep sport` to find the app. It was `com.huami.watch.newsport` in my case. Exit the adb shell.

The backup command will not fail if you specify the wrong application but it will generate and empty backup file. Make sure your backup file is not an empty file. No get the package's backup:


```
$ adb backup -f export_data.ab -noapk com.huami.watch.newsport
```

(Yes, you can use this process to backup and restore the data on your watch if you ever need to factory reset your watch.  Simply call `adb restore export_data.ab` to restore)

Now you need to convert the data to a tarball.

Download the abe tool from here: https://github.com/nelenkov/android-backup-extractor/releases
and then run the following command to convert the backup file to the tar format:

```
java -jar abe-all.jar unpack export_data.ab export_data.tar
```

Inside the tarball is the database file you need.  Extract the tar file and find the sqlite database file:

```
tar xvf export_data.tar
```

in my case it was under: `apps/com.huami.watch.newsport/db/sport_data.db`

TODO: Create helper script to automate all the above or use adb bindings to do that from within the app (https://godoc.org/github.com/zach-klippenstein/goadb).

## Compilation

TODO:

- Create docker image to help compilation without having to install all X dependencies
- Try compilation for Mac and Windows and add instructions
- Add instructions on how to compile this application

## Links

- Backup data instructions:
  - https://forum.xda-developers.com/smartwatch/amazfit/xiaomi-huami-amazfit-export-data-t3533292
  - https://forum.xda-developers.com/smartwatch/amazfit/xiaomi-huami-amazfit-export-data-t3533292/page17
- abe tool: https://github.com/nelenkov/android-backup-extractor/releases

## Notes

- Find logs of a running package (app): https://stackoverflow.com/a/9869609
  e.g. `adb logcat | grep -F "`adb shell ps | grep com.huami.watch.wifiuploaddata  | tr -s [:space:] ' ' | cut -d' ' -f2`"`
