IMDb-Go
=======



1 DESCRIPTION
-------------

**IMDb-Go** is a program written in _Go_ to look up IMDb information of a movie directly from the command line.

2 BUILD
---------

###2.1 Installation

To install the program, simply use:

		sudo make install


###2.2 Uninstallation

To uninstall the program, simply use:

		sudo make uninstall


3 HOW TO RUN
-------------

Lets say you want to find the information for **Harry Potter** which released on **2010**:

		imdb -m="Harry Potter" -y=2010


__*Note:*__ Make sure the moviename is in inverted commas. The year is not necessary (It will give the latest result), but helps in case of movies with similar name like in above example.


4 SCREENSHOTS
--------------

![Screenshot](https://github.com/gauravmittal1995/IMDb/blob/master/Screenshot/Screenshot-IMDb.png)


5 HELP
-------

For help on the usage or flags, simply use one of the following:

        imdb --help
        imdb -h


6 STATUS
--------

Currently only Movies are supported. Tv Shows, Animes, etc will be added soon.

7 LICENSE
---------

This project uses the MIT Open-Source License. See the [__LICENSE__] [1] file for more information.

[1]: https://github.com/gauravmittal1995/IMDb/blob/master/LICENSE "LICENSE"  
