package service

import (
	"fmt"
	"log"

	"inflix-main/config"
	"inflix-main/model"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllAlbums - fetch all Albums
func GetAllAlbums(album *[]model.Album) (err error) {
	if err = config.DB.Find(album).Error; err != nil {
		return err
	}
	return nil
}

//CreateAlbum - creates an album
func CreateAlbum(album *model.Album) (err error) {
	log.Println("Album Input:", album)
	if err = config.DB.Create(album).Error; err != nil {
		return err
	}
	return nil
}

//GetAlbum fetch one alb um
func GetAlbum(album *model.Album, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(album).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAlbum - modifies an album
func UpdateAlbum(album *model.Album, id string) (err error) {
	fmt.Println(album)
	config.DB.Save(album)
	return nil
}

//DeleteAlbum deletes a given album name
func DeleteAlbum(album *model.Album, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(album)
	return nil
}
