package handlers

import (
	"encoding/xml"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nurbekchymbaev/sanctions-go-app/database"
	"github.com/nurbekchymbaev/sanctions-go-app/models"
	"github.com/nurbekchymbaev/sanctions-go-app/util"
	"log"
)

const individualType = "Individual"

type SdnXml struct {
	SdnList []SdnEntry `xml:"sdnEntry"`
}

type SdnEntry struct {
	ID        string    `xml:"uid"`
	Firstname string    `xml:"firstName"`
	Lastname  string    `xml:"lastName"`
	Type      string    `xml:"sdnType"`
	Title     string    `xml:"title"`
	Remarks   string    `xml:"remarks"`
	AkaList   []AkaList `xml:"akaList"`
}

type AkaList struct {
	Aka []Aka `xml:"aka"`
}

type Aka struct {
	ID        string `xml:"uid"`
	Category  string `xml:"category"`
	Firstname string `xml:"firstName"`
	Lastname  string `xml:"lastName"`
}

func Update(c *fiber.Ctx) error {

	var sdnXml SdnXml
	bytes, _ := util.GetRemoteXML("https://www.treasury.gov/ofac/downloads/sdn.xml")
	err := xml.Unmarshal(bytes, &sdnXml)
	if err != nil {
		log.Println(err)
		return c.Status(503).JSON(&fiber.Map{
			"result": false,
			"info":   "service unavailable",
			"code":   503,
		})
	}

	var entries []models.Entry
	var names []models.Names
	for _, sdnEntry := range sdnXml.SdnList {
		if sdnEntry.Type == individualType {
			entry := new(models.Entry)
			entry.ID = util.ConvertToUint(sdnEntry.ID)
			entry.Firstname = sdnEntry.Firstname
			entry.Lastname = sdnEntry.Lastname
			entry.Remarks = sdnEntry.Remarks
			entry.Title = sdnEntry.Title
			entries = append(entries, *entry)
			for _, akaList := range sdnEntry.AkaList {
				for _, aka := range akaList.Aka {
					name := new(models.Names)
					name.ID = util.ConvertToUint(aka.ID)
					name.Category = aka.Category
					name.Firstname = aka.Firstname
					name.Lastname = aka.Lastname
					name.EntryID = entry.ID
					names = append(names, *name)
				}
			}
		}
	}

	database.DB.Db.Exec(fmt.Sprintf("select pg_try_advisory_lock(%d)", database.LockKey))
	defer database.DB.Db.Exec(fmt.Sprintf("select pg_advisory_unlock(%d)", database.LockKey))

	trx := database.DB.Db.Begin()
	trx.Exec("truncate table entries")
	trx.Exec("truncate table names")
	trx.CreateInBatches(entries, len(entries))
	trx.CreateInBatches(names, len(names))
	trx.Commit()

	return c.Status(200).JSON(&fiber.Map{
		"result": true,
		"info":   "",
		"code":   200,
	})
}
