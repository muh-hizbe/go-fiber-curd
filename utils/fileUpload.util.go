package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type File struct {
	Name string `default:"here.png"`
	Path string `default:"./public/images/"`
	Age  int    `default:"23"`
}

const DefaultPathAssetImage = "./public/images/"

func HandleRemoveFile(opt File) (*File, error) {
	result := &File{}

	// SET NAME VALUE WHEN CONDITION NAME NOT EMPTY STRING
	if opt.Name != "" {
		result.Name = opt.Name
	} else {
		defaultName, _ := GetDefault("Name", opt)
		result.Name = defaultName
	}

	// SET PATH VALUE WHEN CONDITION BELOW
	if opt.Path != "" {
		result.Path = opt.Path
	} else {
		//result.Path = DefaultPathAssetImage

		defaultPath, err := GetDefault("Paths", opt)
		if err != nil {
			log.Println(err)
		}

		result.Path = defaultPath
		log.Println("default path ::: ", result.Path)
		log.Println("default name ::: ", result.Name)
	}

	// COMBINE PATH AND FILENAME TO BE ONE STRING
	fullPathName := result.Path + result.Name

	// REMOVE FILE FROM SPECIFIC ASSET DIRECTORY
	err := os.Remove(fullPathName)
	if err != nil {
		log.Println(err)
		return result, err
	}

	return result, nil
}

func HandleFileUpload(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	log.Println(file)

	if err != nil {
		log.Println("Error image upload ==> ", err)
		log.Println("Will skip this upload proccess ")
		//return c.Status(500).JSON(fiber.Map{
		//	"status": "error",
		//	"message": "failed upload image",
		//})
	}

	var filename string
	if file != nil {
		filename = file.Filename

		// STORE IMAGE TO THE ASSET DIRECTORY
		err := c.SaveFile(file, fmt.Sprintf("./public/images/%s", file.Filename))
		if err != nil {
			log.Println("No image stored in asset directory.")
		}
		log.Println("Successfully store image to the assets directory")
	} else {
		log.Println("No image stored in asset directory.")
	}

	log.Println("filename ::: ", filename)

	// PASSING VARIABLE filename TO NEXT HANDLER USING context Locals
	c.Locals("filename", filename)
	return c.Next()
}
