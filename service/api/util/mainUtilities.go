package util

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// func GetLoggedInUser(w http.ResponseWriter, ctx reqcontext.RequestContext, ADB database.AppDatabase) User {

// 	if !ctx.Authenticated {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return User{}
// 	}

// 	// FoundUser, err := AllUsers[ctx.UserID]
// 	FoundUser, err := ADB.GetUser(ctx.UserID)
// 	if err != nil {
// 		ctx.Logger.Warningln("User with ID: ", ctx.UserID, " requested but not found in database!")
// 		w.WriteHeader(http.StatusNotFound)
// 		return User{}
// 	}

// 	return FoundUser
// }

// Function to check if a user is found in a list of usernames
func IsUserFoundInList(usernames []int, userIDToCheck int) bool {
	for _, user := range usernames {
		if user == userIDToCheck {
			return true
		}
	}
	return false
}

func UserBelongsToConversation(conv Conversation, user User) bool {
	// check if the person already belongs to the group or not
	return IsUserFoundInList(conv.Participants, user.Id)

}

func GetBasicUserPicture() ([]byte, error) {
	return GetPicture("/basic_user_picture.png")
}

func GetBasicGroupPicture() ([]byte, error) {
	return GetPicture("/basic_group_picture.png")
}
func GetPicture(imagePath string) ([]byte, error) {
	cwd, _ := os.Getwd()
	// filePath := filepath.Join(filepath.Dir(filepath.Dir(cwd)), "/public/images", imagePath)
	filePath := filepath.Join(filepath.Dir(cwd), "/public/images/default", imagePath)

	fmt.Println("current WD: ", filepath.Dir(cwd))
	fmt.Println("Attempting to open file:", filePath)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []byte{}, errors.New("image file does not exist: " + filePath)
	}

	// Open the file
	imageFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", filePath, "Error:", err)
		return []byte{}, errors.New("failed to open image file: " + err.Error())
	}
	defer imageFile.Close()

	// Read the file contents into a byte slice
	binaryData, err := io.ReadAll(imageFile)
	if err != nil {
		return []byte{}, errors.New("failed to read image file: " + err.Error())
	}

	return binaryData, nil
}
