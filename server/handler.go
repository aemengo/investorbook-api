package server

import (
	"fmt"
	"github.com/aemengo/investorbook-api/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (s *server) connections(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	item := ps.ByName("investorId")
	investorId, err := strconv.Atoi(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errorMessage())
		return
	}

	investors, err := s.con.Connections(investorId)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errorMessage())
		return
	}

	fmt.Fprint(w, present(investors))
}

func (s *server) mutualConnections(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	item := ps.ByName("investorId")
	investorId, err := strconv.Atoi(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errorMessage())
		return
	}

	item = ps.ByName("otherInvestorId")
	mutualInvestorId, err := strconv.Atoi(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errorMessage())
		return
	}

	investors, err := s.con.MutualConnections(investorId, mutualInvestorId)
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errorMessage())
		return
	}

	fmt.Fprint(w, present(investors))
}

func (s *server) search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	item := ps.ByName("investorId")
	investorId, err := strconv.Atoi(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errorMessage())
		return
	}

	investors, err := s.con.Search(investorId, ps.ByName("q"))
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, errorMessage())
		return
	}

	fmt.Fprint(w, presentWithConnection(investors))
}

func present(investors []models.Investor) string {
	result := []string{}
	for _, investor := range investors {
		result = append(result, investor.Name)
	}

	return strings.Join(result, "\n")
}

func presentWithConnection(investors []models.Investor) string {
	result := []string{}
	for _, investor := range investors {
		result = append(result, fmt.Sprintf("%s %s", investor.Name, investor.ConnectionDegree))
	}

	return strings.Join(result, "\n")
}

func errorMessage() string {
	return "Error: There was an error."
}
