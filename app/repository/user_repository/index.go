package user_repository

// type Repository interface {
// 	CreateUser(email, password string) (Auth, error)
// 	GetUserByEmail(email string) (Auth, error)
// }

// func getUsers(c *gin.Context) {
// 	rows, err := db_config.DB.Query("SELECT id, name, email FROM users")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	users := []map[string]interface{}{}
// 	for rows.Next() {
// 		var id int
// 		var name, email string
// 		if err := rows.Scan(&id, &name, &email); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		users = append(users, gin.H{"id": id, "name": name, "email": email})
// 	}

// 	c.JSON(http.StatusOK, users)
// }
