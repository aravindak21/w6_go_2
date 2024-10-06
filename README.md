# w6_go_2
🎬 Movie Collection CRUD API in Go 🎥<br/>
This is a simple CRUD (Create, Read, Update, Delete) application built in Go to manage a collection of movies. The application uses the built-in net/http package to create and manage movies with fields like Title, Director, Year, and Genre. 🚀<br/>
✨ Features<br/>
* Create a new movie in the collection.<br/>
* Read all movies or a single movie by its ID.<br/>
* Update the details of an existing movie.<br/>
* Delete a movie from the collection by its ID.<br/>
No external libraries are used in this application, making it lightweight and easy to understand. 😎<br/>

---

🛠️ Setup<br/>
1. clone the repo
   ```
      git clone https://github.com/aravindak21/w6_go_2.git
      cd movie-crud-api
   ```
2. Run the application:
   ```
       go run main.go
    ```
3. The server will start on http://localhost:8080. You can test it using curl or tools like Postman. I have tested here with curl command<br/>
   
---

📋 Endpoints<br/>
1️⃣ Create a Movie<br/>
Method: POST<br/>
Endpoint: /movies<br/>
Description: Create a new movie by providing the title, director, year, and genre.<br/>
Example Request:
```
    curl -X POST -H "Content-Type: application/json" \
    -d '{"title": "Inception", "director": "Christopher Nolan", "year": 2010, "genre": "Sci-Fi"}' \
    http://localhost:8080/movies

```
Code Snippet:

```
    func createMovie(w http.ResponseWriter, r *http.Request) {
    var movie Movie
    json.NewDecoder(r.Body).Decode(&movie)
    movie.ID = nextMovieID
    nextMovieID++
    movies = append(movies, movie)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(movie)
}

 ```
What It Does: It decodes the incoming JSON request, assigns an auto-incremented ID to the new movie, and appends it to the movie collection.<br/>

---

2️⃣ Get All Movies<br/>
Method: GET<br/>
Endpoint: /movies<br/>
Description: Retrieves the entire list of movies in the collection.<br/>
Example Request:<br/>

```
    curl http://localhost:8080/movies
```
Code Snippet:
```
    func getMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movies)
}
```
What It Does: It encodes and returns the entire movies slice as a JSON array.<br/>

---

3️⃣ Get a Movie by ID<br/>
Method: GET<br/>
Endpoint: /movies/{id}<br/>
Description: Retrieves a single movie by its ID from the collection.<br/>
Example Request:
```
    curl http://localhost:8080/movies/1
```
Code Snippet:
```
func getMovie(w http.ResponseWriter, r *http.Request) {
    id, err := extractID(r.URL.Path)
    if err != nil {
        http.Error(w, "Invalid Movie ID", http.StatusBadRequest)
        return
    }

    for _, movie := range movies {
        if movie.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(movie)
            return
        }
    }
    http.Error(w, "Movie not found", http.StatusNotFound)
}
```
What It Does: It extracts the movie ID from the URL, searches the *movies* slice for a matching movie, and returns it if found.

---

4️⃣ Update a Movie<br/>
Method: PUT<br/>
Endpoint: /movies/{id}<br/>
Description: Updates an existing movie's details (Title, Director, Year, Genre) based on its ID.<br/>
Example Request:
```
    curl -X PUT -H "Content-Type: application/json" \
    -d '{"title": "Inception", "director": "Christopher Nolan", "year": 2010, "genre": "Thriller"}' \
    http://localhost:8080/movies/1
```
Code Snippet:
```
    func updateMovie(w http.ResponseWriter, r *http.Request) {
    id, err := extractID(r.URL.Path)
    var updatedMovie Movie
    json.NewDecoder(r.Body).Decode(&updatedMovie)

    for i, movie := range movies {
        if movie.ID == id {
            movies[i].Title = updatedMovie.Title
            movies[i].Director = updatedMovie.Director
            movies[i].Year = updatedMovie.Year
            movies[i].Genre = updatedMovie.Genre
            json.NewEncoder(w).Encode(movies[i])
            return
        }
    }
    http.Error(w, "Movie not found", http.StatusNotFound)
}
```
What It Does: It decodes the updated movie details from the request body and updates the corresponding movie in the collection.<br/>

---

5️⃣ Delete a Movie<br/>
Method: DELETE<br/>
Endpoint: /movies/{id}<br/>
Description: Deletes a movie from the collection based on its ID.<br/>
Example Request:
```
    curl -X DELETE http://localhost:8080/movies/1
```
Code Snippet:
```
    func deleteMovie(w http.ResponseWriter, r *http.Request) {
    id, err := extractID(r.URL.Path)

    for i, movie := range movies {
        if movie.ID == id {
            movies = append(movies[:i], movies[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "Movie not found", http.StatusNotFound)
}
```
What It Does: It removes the movie with the matching ID from the movies slice and sends a *204 No Content* status if successful.

---

⚙️ Helper Functions<br/>
extractID Function<br/>
This helper function extracts the movie ID from the URL path. It splits the URL and converts the ID part from a string to an integer.<br/>
Code Snippet:
```
    func extractID(path string) (int, error) {
    parts := strings.Split(path, "/")
    if len(parts) < 3 {
        return 0, fmt.Errorf("invalid path")
    }
    return strconv.Atoi(parts[2])
}
```

---

🚀 Running the Application<br/>
1. Make sure you have Go installed.<br/>
2. Clone the repository, and run the application:
```
    go run main.go
```
3. The application will run on *http://localhost:8080*<br/>

---

# OUTPUT<br/>
Run man.go<br/>
<img width="1260" alt="Screenshot 2024-10-03 at 11 41 43 PM" src="https://github.com/user-attachments/assets/50643492-0d05-4334-bdbe-164240cad6d6">

---

## Curl command and ouput<br/>

<img width="1260" alt="Screenshot 2024-10-03 at 11 41 31 PM" src="https://github.com/user-attachments/assets/594c992b-e8d4-4003-a510-0491b23023ba">




