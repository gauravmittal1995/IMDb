/*
 *  Copyright (c) 2015 Gaurav Mittal
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 *
 */

package main

import (
        "encoding/json"
        "fmt"
        "net/http"
        "os"
        "flag"
        "regexp"
       )

var Movie struct {
    Response    string
    Type        string
    Title       string
    Year        string
    Rated       string
    Released    string
    Runtime     string
    Genre       string
    Director    string
    Writer      string
    Actors      string
    Plot        string
    Language    string
    Country     string
    ImdbRating  string
    Metascore   string
    ImdbVotes   string
    ImdbID      string
    Error       string
}

func incorrectUsage() {
    fmt.Printf("Incorrect Usage. Use `imdb -h` or `imdb --help` for correct usage\n")
    os.Exit(0)
}

func main() {

    var movieName string
    flag.StringVar(&movieName, "m", "moviename", "Name of Movie")

    var movieYear string
    flag.StringVar(&movieYear, "y", "0000", "Releasing Year")

    flag.Parse()

    var req string

    regex, _ := regexp.Compile(" ")
    regex = regexp.MustCompile(" ")
    movieName = regex.ReplaceAllString(movieName, "+")

    if (movieName == "moviename") {
        incorrectUsage() 
    } else if movieYear == "0000" {
        req = "http://www.omdbapi.com/?t=" + movieName + "&type=movie&plot=full&r=json"
    } else {
        req = "http://www.omdbapi.com/?t=" + movieName + "&y=" + movieYear + "&type=movie&plot=full&r=json"
    }

    r, _ := http.Get(req)
    defer r.Body.Close()

    dec := json.NewDecoder(r.Body)
    dec.Decode(&Movie)

    fmt.Printf("\n")

    if Movie.Response == "True" {
        fmt.Printf("%s (%s)\n", Movie.Title, Movie.Year)
        fmt.Printf("%s | %s | %s | %s (%s)\n", Movie.Rated, Movie.Runtime, Movie.Genre, Movie.Released, Movie.Country)
        fmt.Printf("Ratings: %s/10 from %s votes. \t Metascore: %s/100\n", Movie.ImdbRating, Movie.ImdbVotes, Movie.Metascore)
        fmt.Printf("\n%s\n", Movie.Plot)
        fmt.Printf("\nDirector: %s\n", Movie.Director)
        fmt.Printf("Writer: %s\n", Movie.Writer)
        fmt.Printf("Stars: %s\n", Movie.Actors)
        fmt.Printf("\nImdb Page: http://www.imdb.com/title/%s\n", Movie.ImdbID)
    } else {
        fmt.Printf("%s\n", Movie.Error)
    }

    fmt.Printf("\n")
}
