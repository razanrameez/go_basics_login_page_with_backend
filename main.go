package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Serve the login page
    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", nil)
    })

    // Handle login form submission
    r.POST("/login", func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")

        // Predefined username and password for validation
        if username == "admin" && password == "password" {
            // Set a session or cookie here
            c.SetCookie("session_token", "your-session-token", 3600, "/", "localhost", false, true)
            c.Redirect(http.StatusFound, "/home")
        } else {
            c.HTML(http.StatusOK, "login.html", gin.H{
                "error": "Incorrect username or password",
            })
        }
    })

    // Serve the home page
    r.GET("/home", func(c *gin.Context) {
        sessionToken, err := c.Cookie("session_token")
        if err != nil || sessionToken != "your-session-token" {
            c.Redirect(http.StatusFound, "/login")
            return
        }

        c.HTML(http.StatusOK, "home.html", nil)
    })

    // Handle sign out
    r.POST("/logout", func(c *gin.Context) {
        // Clear the session or cookie
        c.SetCookie("session_token", "", -1, "/", "localhost", false, true)
        c.Redirect(http.StatusFound, "/login")
    })

    r.LoadHTMLGlob("templates/*")
    r.Run(":8080")
}
