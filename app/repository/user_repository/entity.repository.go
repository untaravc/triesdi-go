package user_repository

type User struct {
	ID         string 
	Preference string 
	WeightUnit string 
	HeightUnit string 
	Weight     *int    
	Height     *int   
	Email	   string 
	Name       string 
	ImageUri   string 
}