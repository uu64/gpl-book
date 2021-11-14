# ch07/ex08

```
$ go run ./main.go
default
Title               Artist                 Album                     Year  Length  
-----               ------                 -----                     ----  ------  
Go                  Delilah                From the Roots Up         2012  3m38s   
Go                  Moby                   Moby                      1992  3m37s   
Go Ahead            Alicia Keys            As I Am                   2007  4m36s   
Time Machine        Alicia Keys            ALICIA                    2020  4m26s   
If I Ain't Got You  Alicia Keys            The Diary of Alicia Keys  2003  3m48s   
Ready 2 Go          Martin Solveig         Smash                     2011  4m24s   
Ready 2 Go          Martin Solveig         Smash                     2011  4m24s   
Under the Bridge    Red Hot Chili Peppers  Blood Sugar Sex Magik     1992  4m24s   

multi-tier sort (Year > Artist)
Title               Artist                 Album                     Year  Length  
-----               ------                 -----                     ----  ------  
Go                  Moby                   Moby                      1992  3m37s   
Under the Bridge    Red Hot Chili Peppers  Blood Sugar Sex Magik     1992  4m24s   
If I Ain't Got You  Alicia Keys            The Diary of Alicia Keys  2003  3m48s   
Go Ahead            Alicia Keys            As I Am                   2007  4m36s   
Ready 2 Go          Martin Solveig         Smash                     2011  4m24s   
Ready 2 Go          Martin Solveig         Smash                     2011  4m24s   
Go                  Delilah                From the Roots Up         2012  3m38s   
Time Machine        Alicia Keys            ALICIA                    2020  4m26s   

sort.Stable (Year > Artist)
Title               Artist                 Album                     Year  Length  
-----               ------                 -----                     ----  ------  
Go                  Moby                   Moby                      1992  3m37s   
Under the Bridge    Red Hot Chili Peppers  Blood Sugar Sex Magik     1992  4m24s   
If I Ain't Got You  Alicia Keys            The Diary of Alicia Keys  2003  3m48s   
Go Ahead            Alicia Keys            As I Am                   2007  4m36s   
Ready 2 Go          Martin Solveig         Smash                     2011  4m24s   
Ready 2 Go          Martin Solveig         Smash                     2011  4m24s   
Go                  Delilah                From the Roots Up         2012  3m38s   
Time Machine        Alicia Keys            ALICIA                    2020  4m26s   

```