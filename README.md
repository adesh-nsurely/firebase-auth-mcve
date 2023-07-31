# firebase-auth-mcve

This is a Minimal, Reproducible Example that is created to reproduce an issue we are facing with firebase auth, where the custom claims set on a user are not reflecting.
The code creates x number of users in firebase simultaneously and sets custom claims on those users.
The code then waits for a few seconds and then fetches the user again to see if the custom claims have been set or not.

Upon testing multiple times, following things have been observed:

- When multiple users are created simultaneously and custom claims are set on these users, the claims are not reflected on some users
- When a single user is created at a time, the issue does not seem to occur
- This issue does not seem to be occurring with all firebase accounts. We have 2 firebase accounts, onc created 2 years back and one created recently. The issue only occurs on the newly created firebase account

## Setup

#### Initiate the go project 

`go mod init github.com/username/repo-name`

#### Install firebase dependency

`go get firebase.google.com/go/v4`

#### Download the firebase credentials file and rename it to credentials.json

You can get the credentials from here this link https://console.firebase.google.com/project/_/settings/serviceaccounts/adminsdk

#### Run the code

`go run main.go`


## Output from testing

#### Firebase project 1
The custom claims are set correctly when running the code on this firebase project
```
➜  firebase-auth-mcve git:(main) ✗ go run main.go
2023-07-31 13:37:23: Successfully created user azkxgykzz3uib2u@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user gzgpkdgbuwrggqs@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user 1wwvfwrshmt6wbq@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user uy0l2vg8fcvrr3i@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user u7ap9zv9yeoi600@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user zkblm2c4lvlre8i@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user 5cri3nrtfgd4zwc@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user dezyaabbbf5qnsg@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user dhdxnqstpmat3fg@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user kcqlgzgfc6xw2t9@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user dyxfoq9zp1xopa2@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user j5tyaop4p4eicii@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user qbgl914dow2d9wq@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user bzy46bkb7ehcdox@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user azk33g5xq9okz2n@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user tffbjaso0rntiv1@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user bagh2nbqwohnv1m@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user svt6axzlzg7bkr6@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user j0gs2er0j8f10r4@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user oan9wllsy5tayt4@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user guqrotilqmxkf9q@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user tho79lcekkgis45@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user tfbrokapo4vnxkg@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user hol1tfuutspeuek@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user w2xuqsk9zuuyitb@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user awscqzkdcgibqjp@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user gxyimu767zvvhe6@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user wtgqbpgb7au6eod@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user dowvjleqjgrzlyg@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user v51ezfaxmdwo03h@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user qbwhvulwk6lvv6h@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user osauxpo5ygaeirz@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user q1xdyr5dxodsu04@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user lf1msyfvifhs3ue@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user hpa6nem5uoitca1@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user fl69oe97xpcqahk@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user s2ifpmzz7be4n5h@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user 5vtyluddwurtcsp@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user bkqyvqohaysbxlk@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user fxj28e11imdxyid@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user axdjmfcdexjwjk6@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user iwbindsxcqmmcxt@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user glkg5wbphsm9wcr@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user mehyyk2t4nxbeh6@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user mofqnk9mr4bmlog@nsurely.com and default claims null
2023-07-31 13:37:23: Successfully created user kauhvwnjfdhvxrw@nsurely.com and default claims null
2023-07-31 13:37:24: Successfully created user dv1ctwtek1juhdo@nsurely.com and default claims null
2023-07-31 13:37:24: Successfully created user xwjs7btoec69j0q@nsurely.com and default claims null
2023-07-31 13:37:24: Successfully created user byh373wxfidfleq@nsurely.com and default claims null
2023-07-31 13:37:24: Successfully created user tkaayd5xqhkbofg@nsurely.com and default claims null
2023-07-31 13:37:29: Successfully fetched user 1wwvfwrshmt6wbq@nsurely.com custom claims: {"orgGroup":"qziekv0CLbx916FXP2sh","orgs":["f6nKg3o4ugSvhMM1F57b","0qDfHpJSs9yhCAxnxO2N"]}
2023-07-31 13:37:29: Successfully fetched user azkxgykzz3uib2u@nsurely.com custom claims: {"orgGroup":"UajQCRHxLFYdylfu3Cu6","orgs":["OxPBuFbEINo1ThSQzN8I","TueFvtuaWYUvYGaO039j"]}
2023-07-31 13:37:29: Successfully fetched user svt6axzlzg7bkr6@nsurely.com custom claims: {"orgGroup":"jXNeMcDwjf3nF4ZifrLi","orgs":["nKbXig4grVLjFglt1VDq","8PiPXVvJ7wvq7azX9JDd"]}
2023-07-31 13:37:29: Successfully fetched user w2xuqsk9zuuyitb@nsurely.com custom claims: {"orgGroup":"0w9RkVrzCV1OAimBSZ4c","orgs":["klyg72E5Lfd8eZSxDbAy","iy2YTSvSdZKXqIMHlR45"]}
2023-07-31 13:37:29: Successfully fetched user uy0l2vg8fcvrr3i@nsurely.com custom claims: {"orgGroup":"YYQmiwKVOQDuoCfVnqQH","orgs":["9fpD1B3kvVNqyzjABPwa","LrloTIUORqspd5qnL01T"]}
2023-07-31 13:37:29: Successfully fetched user gzgpkdgbuwrggqs@nsurely.com custom claims: {"orgGroup":"IhRxYIEyKhJ5s2t0SQb8","orgs":["MXMkcLcgqVuRMdSWC5m3","kifzgIF0MInAhlawC5q9"]}
2023-07-31 13:37:29: Successfully fetched user azk33g5xq9okz2n@nsurely.com custom claims: {"orgGroup":"NuelAAEjJF33IoPDOBVz","orgs":["VgUV7pDA07u9zPIccQwR","sSyZ0fSqfXwGKaRELcuI"]}
2023-07-31 13:37:29: Successfully fetched user dezyaabbbf5qnsg@nsurely.com custom claims: {"orgGroup":"vHNy7fGxuWdYNg78kKj6","orgs":["DUQn5FWOgAAGOEHHe8jv","EhdjMWMZPJQ47UHZgv4v"]}
2023-07-31 13:37:29: Successfully fetched user j0gs2er0j8f10r4@nsurely.com custom claims: {"orgGroup":"gR5Pzki5vIBcQMOlw0nA","orgs":["9AzqAalSDnko5nppjEhK","jrYZqrEeXZfDTVepkk2w"]}
2023-07-31 13:37:29: Successfully fetched user kcqlgzgfc6xw2t9@nsurely.com custom claims: {"orgGroup":"5f4FBax2V78yJM332DiB","orgs":["TM2RYaTRd2KKKXxGKZZm","XCaXt9K044eNUSsfaarq"]}
2023-07-31 13:37:29: Successfully fetched user u7ap9zv9yeoi600@nsurely.com custom claims: {"orgGroup":"Xe4ffvBpC2kX07oZSwGU","orgs":["jV6vUkzClGlLTxG0iDJe","nhTzg3MZhiIyKxQjGveA"]}
2023-07-31 13:37:29: Successfully fetched user oan9wllsy5tayt4@nsurely.com custom claims: {"orgGroup":"apejhOGBi4lWSNKiUXCq","orgs":["28Qn7TzyeBUhHEWiQRpa","Of2n5aD7F7BGjhjGxraY"]}
2023-07-31 13:37:29: Successfully fetched user qbgl914dow2d9wq@nsurely.com custom claims: {"orgGroup":"XhW2z4EY7nbuQvKMk8UE","orgs":["TluwWn1KWr5bp3aMEGsA","yCZspWEjOcPJ97OJfxqc"]}
2023-07-31 13:37:29: Successfully fetched user tfbrokapo4vnxkg@nsurely.com custom claims: {"orgGroup":"d0ewoA89zON1jBDf0Yda","orgs":["3F9EoPCg9t3qr6rP6tlv","Wwi3H6dcNjjAkDJcrgJi"]}
2023-07-31 13:37:29: Successfully fetched user wtgqbpgb7au6eod@nsurely.com custom claims: {"orgGroup":"mK08Y5YACIvVyqQs9AkG","orgs":["7CqmSUJVNq1PpImcbltU","R7NZo9bUMLE40xBJT8M4"]}
2023-07-31 13:37:29: Successfully fetched user dyxfoq9zp1xopa2@nsurely.com custom claims: {"orgGroup":"sf0HHVa5lKRxIOc7e9BR","orgs":["z4xCG6aAx7tXMbh85ick","PQI8JlqBBTRbD6mUh16J"]}
2023-07-31 13:37:29: Successfully fetched user zkblm2c4lvlre8i@nsurely.com custom claims: {"orgGroup":"VUBFecNEGage5h35Wn7h","orgs":["CrwPVMu3I0d8fLmkBb40","iQwLlxuo1cLfNfV9t3Ch"]}
2023-07-31 13:37:29: Successfully fetched user hol1tfuutspeuek@nsurely.com custom claims: {"orgGroup":"QgWpMkqLWw7CHXV6C0gg","orgs":["d2NzPw6CZsJV6hOqpQTY","mykzZ7u19glFiCBYoMKy"]}
2023-07-31 13:37:29: Successfully fetched user guqrotilqmxkf9q@nsurely.com custom claims: {"orgGroup":"wrIFQA9D22VcyOCxAX8n","orgs":["1wBxuAuer2kMgPtiaaxf","LMIFXxaml5puzbmWwhkK"]}
2023-07-31 13:37:29: Successfully fetched user bkqyvqohaysbxlk@nsurely.com custom claims: {"orgGroup":"MDfi6lFHsJp6F99RHDZy","orgs":["uMtnG6N5JTpFpTt2Stko","nWL87Ca6tlbW6711Y0pq"]}
2023-07-31 13:37:29: Successfully fetched user fxj28e11imdxyid@nsurely.com custom claims: {"orgGroup":"fZRwiGI0SvZb6PfH39PM","orgs":["WEQ5U8kSMXUbIWKcrg0y","qMiKO0PJWQ2q57ORkUFF"]}
2023-07-31 13:37:29: Successfully fetched user qbwhvulwk6lvv6h@nsurely.com custom claims: {"orgGroup":"ZrC7M89RSH40VoetYBkY","orgs":["YQ6c0PJEJdsIf1VrAPJ7","U7QzKLn7dkJSQe1od0bV"]}
2023-07-31 13:37:29: Successfully fetched user dhdxnqstpmat3fg@nsurely.com custom claims: {"orgGroup":"OiuxX6uX17ZtZEuAJLbT","orgs":["A5rfaiK7l9fSfcRv6ySp","yvlOUu97GyoN5KBi0hp3"]}
2023-07-31 13:37:29: Successfully fetched user gxyimu767zvvhe6@nsurely.com custom claims: {"orgGroup":"JRaHE91ePJAJojxDOg4t","orgs":["gQehKQxD2xTVIkcHcNDC","ICSgNxEREpkz5wztSnIk"]}
2023-07-31 13:37:29: Successfully fetched user j5tyaop4p4eicii@nsurely.com custom claims: {"orgGroup":"CtomoIJeueAfQ0gPiUnN","orgs":["vFIEicHNjVWZkwM2ksWg","AbGEUY8hfn8oF8JsUh62"]}
2023-07-31 13:37:29: Successfully fetched user lf1msyfvifhs3ue@nsurely.com custom claims: {"orgGroup":"2XPhd5iQGQIKw2rUug9o","orgs":["iBCf9yTrdFz9JMgiDQOg","YqWyjSBbv1qr9OD5BUm5"]}
2023-07-31 13:37:29: Successfully fetched user awscqzkdcgibqjp@nsurely.com custom claims: {"orgGroup":"c8uVOyndkfYEVaDo1zCn","orgs":["Dv7m8NnTh8Kx8mvAayju","oLD8PhAkqKOfRyMmM6oR"]}
2023-07-31 13:37:29: Successfully fetched user bzy46bkb7ehcdox@nsurely.com custom claims: {"orgGroup":"vImDv7NEiNK1M9f0vlPn","orgs":["AKjB2JhylgLesXqpqdrT","s2vyc3KEibQ1aBrH7fnX"]}
2023-07-31 13:37:29: Successfully fetched user s2ifpmzz7be4n5h@nsurely.com custom claims: {"orgGroup":"FuT4gtIdFJwmM749RHZf","orgs":["EAyfRRpWc1JOksWnoEvE","x3KKkLhDBl4CGOUjI1rq"]}
2023-07-31 13:37:29: Successfully fetched user 5cri3nrtfgd4zwc@nsurely.com custom claims: {"orgGroup":"9OCu4nAI74K7pRRiQoPk","orgs":["Oo7ujNX5vLwuKBGpwdqj","vRedkrwCmTWEcLdwiNZL"]}
2023-07-31 13:37:29: Successfully fetched user q1xdyr5dxodsu04@nsurely.com custom claims: {"orgGroup":"6tv2WK8p3dIIkO7aCZ0l","orgs":["ldOsQ2xteQDZK3tf2kPA","PkYoAiReYfhF5DQ4NAWb"]}
2023-07-31 13:37:29: Successfully fetched user dowvjleqjgrzlyg@nsurely.com custom claims: {"orgGroup":"AmgIsAlS5ith8vrE6o98","orgs":["CNrFtw0ApM09ybcC3t9l","LP8ajGQwSFOkrOpl3bwo"]}
2023-07-31 13:37:29: Successfully fetched user osauxpo5ygaeirz@nsurely.com custom claims: {"orgGroup":"Eeke3RTjpDkw4R5gWIXe","orgs":["Eb8Im4CjSJume3G0I0UO","g0ujhq1ue7DNW6SxYCdE"]}
2023-07-31 13:37:29: Successfully fetched user 5vtyluddwurtcsp@nsurely.com custom claims: {"orgGroup":"PE8jahdDOQgDvBJoXU2d","orgs":["sEvh5d9Lx3NDVAJsy4Lx","oTQHWOK6S63lljkcPBqH"]}
2023-07-31 13:37:29: Successfully fetched user iwbindsxcqmmcxt@nsurely.com custom claims: {"orgGroup":"14gKClbpqyGnmebYcv89","orgs":["lwrP9tQmiEzBlgLQQstg","vkGPsdfQeh49tjZzcxPT"]}
2023-07-31 13:37:29: Successfully fetched user tho79lcekkgis45@nsurely.com custom claims: {"orgGroup":"hSBG5BiFxQ7zJXGTAW8i","orgs":["TA2iN4Q59McY1RzVXx2N","QV8dFqbphMWt6lSfbgvz"]}
2023-07-31 13:37:29: Successfully fetched user hpa6nem5uoitca1@nsurely.com custom claims: {"orgGroup":"35OfCl0nTkfOEYvOobNz","orgs":["qEktlh8DMGEiltkecO5b","Ey4v0lMlQRo1Giwa2zC0"]}
2023-07-31 13:37:29: Successfully fetched user v51ezfaxmdwo03h@nsurely.com custom claims: {"orgGroup":"QPulg5kSMIlFbNY8T2Eb","orgs":["vM5vKTXCDZ22MCxzlzET","I7nFl8ahrTMyo7sChAku"]}
2023-07-31 13:37:29: Successfully fetched user fl69oe97xpcqahk@nsurely.com custom claims: {"orgGroup":"ajIBnaVYoSzow6a1Fzg6","orgs":["8HzAhhsvI8aayQWpaDeR","uI4jUbckRoFAj2jmjY7k"]}
2023-07-31 13:37:29: Successfully fetched user axdjmfcdexjwjk6@nsurely.com custom claims: {"orgGroup":"r0BiyJbLFByPNtoLQM0V","orgs":["yUGu2tQ60f1VXfWr5z0O","5r0U1m5N7x5ARPvCqpx3"]}
2023-07-31 13:37:29: Successfully fetched user mofqnk9mr4bmlog@nsurely.com custom claims: {"orgGroup":"weQkVURE53jtJfY2TRDh","orgs":["xftgrGu5jtdQMpsbO3Ql","5PFztJtS7L8P794m9fjk"]}
2023-07-31 13:37:29: Successfully fetched user mehyyk2t4nxbeh6@nsurely.com custom claims: {"orgGroup":"wpL06aTkQvBnkwAhMEcw","orgs":["WI5t13j7to4pMXfT2HGf","BWBtSLT3VOwYumG9Xaqh"]}
2023-07-31 13:37:29: Successfully fetched user glkg5wbphsm9wcr@nsurely.com custom claims: {"orgGroup":"HdpA5X4EJUXx6l6KKsmU","orgs":["LN6t9attgtU7z7XkZA8d","yKMfekzM0tSoEaPiqOCG"]}
2023-07-31 13:37:29: Successfully fetched user kauhvwnjfdhvxrw@nsurely.com custom claims: {"orgGroup":"fmgRlo3VyEYeMQddDrBp","orgs":["H5nyZYKt9OQDaipOJDjA","6dMGOoWvw7CQ9mF71ES0"]}
2023-07-31 13:37:29: Successfully fetched user bagh2nbqwohnv1m@nsurely.com custom claims: {"orgGroup":"R3VdTMDMFKKQu0JEwsVn","orgs":["XXvNClhlFSzdFLdgEWnd","nBAxv5C3un92TfLGeJQ1"]}
2023-07-31 13:37:29: Successfully fetched user dv1ctwtek1juhdo@nsurely.com custom claims: {"orgGroup":"csRHqgqg8Amyz3laoMqR","orgs":["O7fLEhprXzDxyV0rrFAf","Uw6bZqQsCj0xvQhvE3CZ"]}
2023-07-31 13:37:29: Successfully fetched user xwjs7btoec69j0q@nsurely.com custom claims: {"orgGroup":"7YBsWPoAMXH9kCerWlLS","orgs":["eAID9wd7MLTvJwZ9SHEC","KnMYiVRPZtbYavFMnScv"]}
2023-07-31 13:37:29: Successfully fetched user byh373wxfidfleq@nsurely.com custom claims: {"orgGroup":"fDqJQ3sSn7QCr6uU15Gp","orgs":["0Tu3skL1MhmWe77q28o5","lnV1MSu9MuQAJxaKrrtp"]}
2023-07-31 13:37:29: Successfully fetched user tkaayd5xqhkbofg@nsurely.com custom claims: {"orgGroup":"09kZF2Dv1uqYRBElViJo","orgs":["wdg0MKJQROa3jo5QD4pz","H9EfiCdsPMZdfnF7fPfF"]}
2023-07-31 13:37:29: Successfully fetched user tffbjaso0rntiv1@nsurely.com custom claims: {"orgGroup":"VegYhPz3RxOhntld7Ppl","orgs":["HuMCozeP9DERlVnWbpHO","QGm2gKOntFtwaNRxHRhl"]}
```

#### Firebase project 2
The custom claims are not always set correctly when running on the second project. This firebase project was setup few months back

```
➜  firebase-auth-mcve git:(main) ✗ go run main.go
2023-07-31 13:40:37: Successfully created user ap5g2rvhxdpxg8i@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user xqqniomb0nlcnf0@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user scrcrnexzivcxbl@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user xsipnbmxvq7mp0q@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user q1cpeuomw7sxfq9@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user ouagb6h7hurseoh@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user gglaoupcblaiu03@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user b5mosg42u2f0m5j@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user d0i6sxug0xamdjn@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user exrudumaq92oni4@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user zsibxunt4nrbbxu@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user lr4jnd3lhp3xxqo@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user nzlphrs75jhfp3c@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user ahgyist5zlx73dy@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user 1ze0wjjxybrkra1@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user qid6ofpzg3xpok4@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user bd0tsiextmc8vvp@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user ikmiexnbu4jpno6@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user 0fi2vys59sjb3jq@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user jgakm8bajnfvbun@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user c8h00uget6xxjq6@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user npr414ungljumnb@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user wogfkbneb1vpqyv@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user d6vovqg88fuu86r@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user cyhehcqacg2w17n@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user rw803qqijhyh6tm@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user 80diejntngthsuw@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user i2pld6xdbebgxxw@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user zfwof7gnwmp0mgc@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user mapap7jfnb2f96c@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user hy4cici9qk2terv@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user rlzcb0wfnzatydu@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user saouacsr2go9pxt@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user jippyz51qknlllb@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user qkzxxvu4tbmexhj@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user vyi4qmjdhk7u4zz@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user wafypfhrjhcstou@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user dc2xuoahlywpkst@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user ykshqm4vu29bzj4@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user xd18trmp4riojom@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user mxvq37e1tzsbbxu@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user 05jrzo2srbybgrf@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user cxskjblh9iuvj5t@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user wrylcgq73rhskgu@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user 690psqoyqzq2iuy@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user wqmpdnksr8gxrl6@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user 3kpx7dnl7igv7ng@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user gtow9m2fayxvlfm@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user n6shhehvru4v4nl@nsurely.com and default claims null
2023-07-31 13:40:37: Successfully created user lnj8pca6jsnfqrj@nsurely.com and default claims null
2023-07-31 13:40:42: Successfully fetched user xqqniomb0nlcnf0@nsurely.com custom claims: {"isAdmin":true,"tenantId":"M7KJQZfuWgdkgEalRxND"}
2023-07-31 13:40:43: Successfully fetched user ap5g2rvhxdpxg8i@nsurely.com custom claims: {"orgGroup":"ART7ofBoZHV8dgvvuxoi","orgs":["TE6sHURtjky4AeKlJTYu","deM5rkbNSp0YMujYFg0i"]}
2023-07-31 13:40:43: Successfully fetched user xsipnbmxvq7mp0q@nsurely.com custom claims: {"orgGroup":"obXg9Ye2FhwYcohYR8uO","orgs":["CTT2x2Sm8SUKvj4JRXKB","HcGCx9S9B5A8bN8BJ7Fh"]}
2023-07-31 13:40:43: Successfully fetched user scrcrnexzivcxbl@nsurely.com custom claims: {"isAdmin":true,"tenantId":"dqJabwNTd5OAETrKOdoy"}
2023-07-31 13:40:43: Successfully fetched user zsibxunt4nrbbxu@nsurely.com custom claims: {"orgGroup":"LehLjpCHXREFvu1uZHn9","orgs":["UyBLbBUyIBNY9W9a5J6K","xCFtv3jWapF8XDEFyTGm"]}
2023-07-31 13:40:43: Successfully fetched user b5mosg42u2f0m5j@nsurely.com custom claims: {"isAdmin":true,"tenantId":"l6q8qB5nubaxVqw7poz2"}
2023-07-31 13:40:43: Successfully fetched user ouagb6h7hurseoh@nsurely.com custom claims: {"orgGroup":"3ZlMciTGSv89rL2F7yMH","orgs":["M6YQFfL2lhktSHjg9Hqq","5NHuXq7QdgsQnMQeFusq"]}
2023-07-31 13:40:43: Successfully fetched user 1ze0wjjxybrkra1@nsurely.com custom claims: {"orgGroup":"6ypg7OHeJwusVbJ2RifT","orgs":["xo94pcCD6FJlwbErOoTk","1nf5qIbqvwhbeMFuCkfp"]}
2023-07-31 13:40:43: Successfully fetched user bd0tsiextmc8vvp@nsurely.com custom claims: {"orgGroup":"iPzoKHJlj1i0B9zYdqtK","orgs":["gKmjWuJZISHp4x3UMmra","cW1RT1wMzji0kKS3umEM"]}
2023-07-31 13:40:43: Successfully fetched user q1cpeuomw7sxfq9@nsurely.com custom claims: {"orgGroup":"ueKHV7iPfKFQOgcUtuSQ","orgs":["oPK3nAzAkehDUvBxCXoT","zcIKXf4g23iNHbOnszEw"]}
2023-07-31 13:40:43: Successfully fetched user lr4jnd3lhp3xxqo@nsurely.com custom claims: {"orgGroup":"0w1hYLVV6SxPac3IwVgE","orgs":["EbuEY4bFawSsrvhAbxPz","0h6UEPFu2qGoN8wT4ywm"]}
2023-07-31 13:40:43: Successfully fetched user qid6ofpzg3xpok4@nsurely.com custom claims: {"orgGroup":"xlN3lDh3Xgj8Qf7lNu0y","orgs":["v1UAwxbfAnAcrCncn95C","tdWvmpqWm1g2TAPxKfIo"]}
2023-07-31 13:40:43: Successfully fetched user d0i6sxug0xamdjn@nsurely.com custom claims: {"orgGroup":"fVCwICVR4vytTj6nH5cn","orgs":["VjOuR3pixIVSAJT8XKo8","Xy9eyJOqmn5Ytha5JzbV"]}
2023-07-31 13:40:43: Successfully fetched user ikmiexnbu4jpno6@nsurely.com custom claims: {"orgGroup":"hkvo49ROcvLoSpT1ultL","orgs":["3iZ09KmBy8I5ERMskkTq","zZOBIx3ooUb6So24EcL5"]}
2023-07-31 13:40:43: Successfully fetched user d6vovqg88fuu86r@nsurely.com custom claims: {"orgGroup":"1mzJHSzPwjnU10pqRGhS","orgs":["lEiwkMQ6FdUMytzcw1Uw","DiJ8A2A2gDIxH4EoavKq"]}
2023-07-31 13:40:43: Successfully fetched user ahgyist5zlx73dy@nsurely.com custom claims: {"orgGroup":"kVHGw5o77LPmW0qDjSvO","orgs":["Dt9LWOWBNqsyO0d7Voj5","z8IEYdc1CS3DtJnLA9u1"]}
2023-07-31 13:40:43: Successfully fetched user nzlphrs75jhfp3c@nsurely.com custom claims: {"orgGroup":"YwHJMKoAJZLu4tMU3LRV","orgs":["so6LKItB5VvFjxr38YLe","iTwyKwVo6XdEB5DvuF9q"]}
2023-07-31 13:40:43: Successfully fetched user gglaoupcblaiu03@nsurely.com custom claims: {"orgGroup":"eS0Unekt4vWOfYS5OJQN","orgs":["E1Zl2SLIocMSXNnbgz9d","aQsHjBtD9VN5uQrL5y1v"]}
2023-07-31 13:40:43: Successfully fetched user wogfkbneb1vpqyv@nsurely.com custom claims: {"orgGroup":"ATQCsZ3yPPUtZLqmrjQW","orgs":["F0C2HzaliyMlC60iM0BS","C8gHE8kWikI7gUM0myWn"]}
2023-07-31 13:40:43: Successfully fetched user jgakm8bajnfvbun@nsurely.com custom claims: {"orgGroup":"J7zgLvUNgNpnUEEj6fID","orgs":["1PXtCj8xGMOOLnEdSUDs","vHJ6Vq20hzyFZCk6Q7tv"]}
2023-07-31 13:40:43: Successfully fetched user c8h00uget6xxjq6@nsurely.com custom claims: {"orgGroup":"1NHuSqvYfuJcSUJIdPUM","orgs":["fbk3z6xAUaOIQ0d6NowJ","6jZdAyHXpJIMBnwlzpqx"]}
2023-07-31 13:40:43: Successfully fetched user i2pld6xdbebgxxw@nsurely.com custom claims: {"orgGroup":"W1r6y4G6s7ewPiKi0dm2","orgs":["bZ90gAwF6bpVVeOVGfO4","x9Bv3IeOgu6QefRwgjJw"]}
2023-07-31 13:40:43: Successfully fetched user 0fi2vys59sjb3jq@nsurely.com custom claims: {"orgGroup":"OBp3G6JhRmIQyiPAXmwW","orgs":["8z6zwtX6UercVjHLOTx6","Ntti178wqOmDsbEmflZW"]}
2023-07-31 13:40:43: Successfully fetched user exrudumaq92oni4@nsurely.com custom claims: {"orgGroup":"k08cGJMrgbZcCNuBmDfv","orgs":["3njEJy2uBYJclBTEYYJX","55YgOAgtnrN5wVpt4Ra5"]}
2023-07-31 13:40:43: Successfully fetched user zfwof7gnwmp0mgc@nsurely.com custom claims: {"isAdmin":true,"tenantId":"KCVTsHlCvvnOL8RcpC6k"}
2023-07-31 13:40:43: Successfully fetched user rw803qqijhyh6tm@nsurely.com custom claims: {"isAdmin":true,"tenantId":"fCeOwxli3dxfgPPCmBsp"}
2023-07-31 13:40:43: Successfully fetched user qkzxxvu4tbmexhj@nsurely.com custom claims: {"orgGroup":"aOqWSioFJ2v8MRcSSg3p","orgs":["kRYXk9xYFYzvstPo8zHb","rtgBwFl0fzXUwyZAE9sN"]}
2023-07-31 13:40:43: Successfully fetched user mapap7jfnb2f96c@nsurely.com custom claims: {"orgGroup":"HN06QmeOzlxbbIlP0ELM","orgs":["4ooFePKhjbWyJ9K1TsyS","MKsYrEtviREfCTb0JvWz"]}
2023-07-31 13:40:43: Successfully fetched user rlzcb0wfnzatydu@nsurely.com custom claims: {"orgGroup":"UOydXyw0qtS8ivqEWZll","orgs":["fE1nDmjhiSjZA1bkJck9","d8rL75fIqMMYoOy4f9lL"]}
2023-07-31 13:40:43: Successfully fetched user cyhehcqacg2w17n@nsurely.com custom claims: {"orgGroup":"4TlcM6fBC61B50JjtgGN","orgs":["ZQpeL2lIkQ2GwtfDYrCR","YIKmpz3AVgLbigqFpF5B"]}
2023-07-31 13:40:43: Successfully fetched user jippyz51qknlllb@nsurely.com custom claims: {"orgGroup":"mpU8uu1HgECDG9N2MvbL","orgs":["9N345TEk9ZiDfmJtaAWA","7XmS3T3PXxDbEseu8Kah"]}
2023-07-31 13:40:43: Successfully fetched user saouacsr2go9pxt@nsurely.com custom claims: {"orgGroup":"5wTsqF4cEmG4RpxgbN21","orgs":["pvHnYV1nlJchWkGfMmkQ","7sRvfUUQmpDbVlG8EeYV"]}
2023-07-31 13:40:43: Successfully fetched user hy4cici9qk2terv@nsurely.com custom claims: {"orgGroup":"kjeuIcP9oMNffFeH1LF6","orgs":["GzRzytGh3FIGT88Tr7Hc","Pa2BeYepcqVMCehoGIlB"]}
2023-07-31 13:40:43: Successfully fetched user dc2xuoahlywpkst@nsurely.com custom claims: {"orgGroup":"bPzcWats41jGKEparkey","orgs":["YVvud6cFYC34jHVR8JgM","KUq8Hj48ZeR4idr7p1yH"]}
2023-07-31 13:40:43: Successfully fetched user xd18trmp4riojom@nsurely.com custom claims: {"orgGroup":"KEtghLUm1Zvnmt8ELSHl","orgs":["xreLn8aSMSEutZphZYb6","7HAVGjAbI66KHLwbIM0R"]}
2023-07-31 13:40:43: Successfully fetched user vyi4qmjdhk7u4zz@nsurely.com custom claims: {"orgGroup":"oj8XbPy98KenWTz9Qkuf","orgs":["rc33kVr8zrSgRxwzlBbd","jw3YsFlyUovmOvqRw04T"]}
2023-07-31 13:40:43: Successfully fetched user mxvq37e1tzsbbxu@nsurely.com custom claims: {"orgGroup":"RygrzBVEGpWHgYvhB6Kq","orgs":["2XWnXksVedqqaGsB7bdv","JRrJzIAypzjVuiXGts8j"]}
2023-07-31 13:40:43: Successfully fetched user ykshqm4vu29bzj4@nsurely.com custom claims: {"orgGroup":"nQjnLOITP6OxzjZOrM9m","orgs":["QbLM8bdfeY4IbOf3ODeW","lqWB0LDLOoHbF51ozxx6"]}
2023-07-31 13:40:43: Successfully fetched user wqmpdnksr8gxrl6@nsurely.com custom claims: {"orgGroup":"j2CUqqkFL6Zfb7PokDTn","orgs":["hOxY4NleGmkg2nuSUCza","87uksAyb90o9NOoN9IuQ"]}
2023-07-31 13:40:43: Successfully fetched user wafypfhrjhcstou@nsurely.com custom claims: {"orgGroup":"ckKCDHOKSItDazVoHUTH","orgs":["dJGcHyT5oe8lj418I7ec","bbzNQT9hLB8w8ZOFrMkE"]}
2023-07-31 13:40:43: Successfully fetched user gtow9m2fayxvlfm@nsurely.com custom claims: {"orgGroup":"s6qtWDI38lPGrnE4gSbw","orgs":["pk7om9WEIlWgTrgM2DRJ","Z1ew5lNwLjpKhCo4payP"]}
2023-07-31 13:40:43: Successfully fetched user cxskjblh9iuvj5t@nsurely.com custom claims: {"orgGroup":"kmx84tiOcP1xiEWTjdbK","orgs":["mljSG84FsVcktf1vVOUs","3lboKHYX486e4yPY1CdX"]}
2023-07-31 13:40:43: Successfully fetched user 05jrzo2srbybgrf@nsurely.com custom claims: {"orgGroup":"8wsFUsuMN7a2Av1N6zSh","orgs":["DZ4UmjbT7jZMOkUIU5XU","i4cRGTDs7JTar2sM5PVp"]}
2023-07-31 13:40:43: Successfully fetched user wrylcgq73rhskgu@nsurely.com custom claims: {"orgGroup":"oKXzKYAshzzyhdW35sc8","orgs":["KrekvFH01qrqOdPs6qdX","zAitL9T3sXyy4r406QaF"]}
2023-07-31 13:40:43: Successfully fetched user 3kpx7dnl7igv7ng@nsurely.com custom claims: {"orgGroup":"do0mCJ2WRtyrpVFVo0lC","orgs":["QSWrTazwjzSsN1fMa1Fq","H4OxskGXUGOStltUiFv1"]}
2023-07-31 13:40:43: Successfully fetched user n6shhehvru4v4nl@nsurely.com custom claims: {"orgGroup":"2goFgaL26eEPA5jVpoxi","orgs":["bgzO94DK1Iq6HaUfqY06","4Z7nkj8npzt9Mkkr1mAu"]}
2023-07-31 13:40:43: Successfully fetched user 690psqoyqzq2iuy@nsurely.com custom claims: {"orgGroup":"9RaB3t7pVfITfldaO2id","orgs":["QKNMOAYHRxz9QcnLgZXd","WlBV7LYMBGw7WeO2KJ5M"]}
2023-07-31 13:40:43: Successfully fetched user lnj8pca6jsnfqrj@nsurely.com custom claims: {"orgGroup":"CGJXwc9Wgx45NHytIqXb","orgs":["5xb5NspcjPigKjIfASJB","u0eN1a3MWJnpwIN6UYnS"]}
2023-07-31 13:40:43: Successfully fetched user 80diejntngthsuw@nsurely.com custom claims: {"isAdmin":true,"tenantId":"a0bnh03gKHam42NXCqR4"}
2023-07-31 13:40:43: Successfully fetched user npr414ungljumnb@nsurely.com custom claims: {"isAdmin":true,"tenantId":"xERhf9qiCHIRmDwiYW4e"}
```