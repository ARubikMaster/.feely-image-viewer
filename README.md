# .feely image file format viewer



## Linux


* 1: **Download bin/feely_viewer**


* 2: **Make the feely executable runnable**
```chmod +x feely_viewer```


* 3: **Add executable to path (optional, but reccomended)**

How to do for bash (If you use a different shell look up how to add a folder to path):

```echo 'export PATH="$PATH:/path/to/executable/folder"' >> ~/.bashrc ```


* 4: **To view a .feely image**

If added to path:

```feely_viewer /path/to/image.feely```

If not, cd to the executable location and run:

```./feely_viewer /path/to/image.feely```


You can also set your file explorer to open .feely files with feely viewer by default

## Windows


* 1: **Download bin/feely.exe**


* 2: **Add executable folder to path (optional but recommended)**

You can see an example here: https://windowsloop.com/how-to-add-to-windows-path/


* 3: **To view a .feely image**

If added to path:

```feely_viewer C:\path\to\image.feely```

If not, cd to the executable location and run:

```.\feely_viewer C:\path\to\image.feely```

Please note you might get an antivirus warning, as the executable is not signed.

You can also set your file explorer to open .feely files with feely viewer by default