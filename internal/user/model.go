package user
// TODO | COMPLETELY CHANGE MODEL FOR WORKING WITH METHEOROLOGY DATA
type User struct {
	ID 				string 	`json:"id" bson:"_id,omitempty"`
	Username 		string	`json:"username" bson:"username"`
	PasswordHash 	string	`json:"-" bson:"password"`
	Email 			string	`json:"email" bson:"email"`
}

type CreateUserDTO struct {
	Email 			string	`json:"email"`
	Username 		string	`json:"username"`
	PasswordHash 	string	`json:"password"`
}
