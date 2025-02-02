package service

import (
	"log"
	"strconv"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetVenues(c *gin.Context) {
	r, err := s.VenueService.GetVenues()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(r, err))
}

func (s *Service) GetVenueDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	r, err := s.VenueService.GetVenueById(int32(id))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(r, err))
}

func (s *Service) GetVenueQrCode(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	r, err := s.VenueService.GetVenueById(int32(id))
	if nil != err {
		c.JSON(500, err.Error())
		return
	}
	log.Println(r)
}
