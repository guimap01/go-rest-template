package handlers

import (
	"encoding/json"
	"go-api-rest/src/models"
	"go-api-rest/src/util"
	"io/ioutil"
	"net/http"
)

func (suite *util.TestSuiteEnv) Test_GetBooks_EmptyResult() {
	req, w := setGetBooksRouter(suite.db)
	a := suite.Assert()
	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Book{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := models.Book{}
	a.Equal(expected, actual)
}
