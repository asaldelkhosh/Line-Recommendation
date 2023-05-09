package http

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/sweetie-pie/line-recommendation/internal/model"
	"github.com/sweetie-pie/line-recommendation/internal/port/mysql"

	"github.com/gofiber/fiber/v2"
)

const (
	XUp   = 1000
	XDown = -1000
	YUp   = 1000
	YDown = -1000
)

type Handler struct {
	Repository *mysql.MySQL
}

func (h *Handler) CreateRoute(c *fiber.Ctx) error {
	tmp := c.Query("count", "4")
	count, _ := strconv.Atoi(tmp)

	nodes, err := h.Repository.GetNodes()
	if err != nil {
		log.Println(err)

		return fiber.ErrInternalServerError
	}

	for i := 0; i < count; i++ {
		// index of src and dest nodes
		srcIndex := rand.Intn(len(nodes))
		destIndex := srcIndex

		// make sure they are not the same
		for {
			destIndex = rand.Intn(len(nodes))
			if destIndex != srcIndex {
				break
			}
		}

		route := model.Route{
			StartID: nodes[srcIndex],
			StopID:  nodes[destIndex],
		}

		if err := h.Repository.InsertRoute(&route); err != nil {
			log.Println(err)

			return fiber.ErrInternalServerError
		}
	}

	return c.Status(fiber.StatusOK).SendString("Created!")
}

func (h *Handler) CreateNode(c *fiber.Ctx) error {
	tmp := c.Query("count", "4")
	count, _ := strconv.Atoi(tmp)

	for i := 0; i < count; i++ {
		node := model.Node{
			X: int64(rand.Intn(XUp-XDown) + XDown),
			Y: int64(rand.Intn(YUp-YDown) + YDown),
		}

		if err := h.Repository.InsertNode(&node); err != nil {
			log.Println(err)

			return fiber.ErrInternalServerError
		}
	}

	return c.Status(fiber.StatusOK).SendString("Created!")
}

func (h *Handler) Search(c *fiber.Ctx) error {
	tmp := c.Query("count", "4")
	count, _ := strconv.Atoi(tmp)

	for i := 0; i < count; i++ {
		search := model.Search{
			StartX: int64(rand.Intn(XUp-XDown) + XDown),
			StartY: int64(rand.Intn(YUp-YDown) + YDown),
			StopX:  int64(rand.Intn(XUp-XDown) + XDown),
			StopY:  int64(rand.Intn(YUp-YDown) + YDown),
		}

		if err := h.Repository.InsertSearch(&search); err != nil {
			log.Println(err)

			return fiber.ErrInternalServerError
		}
	}

	return c.Status(fiber.StatusOK).SendString("Created!")
}

func (h *Handler) Data(c *fiber.Ctx) error {
	routes, err := h.Repository.GetRoutes()
	if err != nil {
		log.Panicln(err)

		return fiber.ErrInternalServerError
	}

	routesResponse := make([]RouteResponse, 0)

	for _, route := range routes {
		src, _ := h.Repository.GetNode(route.StartID)
		dest, _ := h.Repository.GetNode(route.StopID)

		routesResponse = append(routesResponse, RouteResponse{
			ID:    route.ID,
			Start: *src,
			Stop:  *dest,
		})
	}

	searches, err := h.Repository.GetSearches()
	if err != nil {
		log.Panicln(err)

		return fiber.ErrInternalServerError
	}

	searchResponse := make([]SearchResponse, 0)

	for _, node := range searches {
		searchResponse = append(searchResponse, SearchResponse{
			X1: node.StartX,
			X2: node.StopX,
			Y1: node.StartY,
			Y2: node.StopY,
		})
	}

	return c.JSON(&Response{
		Routes:   routesResponse,
		Searches: searchResponse,
	})
}
