package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"googlemaps.github.io/maps"
)

type Handler struct {
	mapclient *maps.Client
}

type request struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
type reponse struct {
	Distance    string `json:"distance"`
	Duration    string `json:"duration"`
	Destination string `json:"destination"`
	Origin      string `json:"origin"`
}
type urlresponse struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []struct {
		Elements []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}

func NewHandler(client maps.Client) *Handler {
	h := &Handler{
		mapclient: &client,
	}

	return h
}

type TransitMode string

func (h *Handler) GetWithSDK(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, token")
	(w).Header().Set("Content-Type", "application/json")

	req := request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to unmarhsal json ", http.StatusInternalServerError)
		return
	}
	fmt.Println("request is", req)
	distmtr := &maps.DistanceMatrixRequest{
		Origins:      []string{req.Origin},
		Destinations: []string{req.Destination},
	}

	resp, err := h.mapclient.DistanceMatrix(context.Background(), distmtr)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Unable to to use googl distance matrix api ", http.StatusInternalServerError)
		return
	}

	res := reponse{
		Distance:    resp.Rows[0].Elements[0].Distance.HumanReadable,
		Duration:    resp.Rows[0].Elements[0].Duration.String(),
		Destination: resp.DestinationAddresses[0],
		Origin:      resp.OriginAddresses[0],
	}
	buf, err := json.Marshal(res)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Unable to marshal ", http.StatusInternalServerError)
		return
	}
	w.Write(buf)
	return

}

func (h *Handler) GetWithUrl(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, token")
	(w).Header().Set("Content-Type", "application/json")

	req := request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to unmarhsal json ", http.StatusInternalServerError)
		return
	}

	url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=Washington,%20DC&destinations=New%20York%20City,%20NY&units=metric&key=AIzaSyCNYxu7FcuVSYRaRZxumdf1JnyvvTa9F38"
	method := "GET"

	client := &http.Client{}
	greq, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	urlres, err := client.Do(greq)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer urlres.Body.Close()

	urlresponse := urlresponse{}
	if err := json.NewDecoder(urlres.Body).Decode(&urlresponse); err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to unmarhsal json ", http.StatusInternalServerError)
		return
	}
	res := reponse{
		Distance:    urlresponse.Rows[0].Elements[0].Distance.Text,
		Duration:    urlresponse.Rows[0].Elements[0].Duration.Text,
		Destination: urlresponse.DestinationAddresses[0],
		Origin:      urlresponse.OriginAddresses[0],
	}
	fmt.Println(res)
	buf, err := json.Marshal(res)
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Unable to marshal ", http.StatusInternalServerError)
		return
	}
	w.Write(buf)
	return

}
