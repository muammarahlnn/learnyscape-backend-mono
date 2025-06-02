package pageutil

import (
	"fmt"
	"learnyscape-backend-mono/pkg/dto"
	"net/http"
	"net/url"
	"strconv"
)

const linkFormat = "%s%s?%s"

func NewLinks(req *http.Request, page, size, totalItem, totalPage int) *dto.Links {
	queries := req.URL.Query()
	host := req.Host
	path := req.URL.Path
	if req.TLS != nil {
		host = fmt.Sprintf("https://%s", host)
	} else {
		host = fmt.Sprintf("http://%s", host)
	}

	setPageQuery(queries, page)
	selfLink := fmt.Sprintf(linkFormat, host, path, queries.Encode())

	setPageQuery(queries, 1)
	firstLink := fmt.Sprintf(linkFormat, host, path, queries.Encode())

	if totalPage > 0 {
		setPageQuery(queries, totalPage)
	} else {
		setPageQuery(queries, 1)
	}
	lastLink := fmt.Sprintf(linkFormat, host, path, queries.Encode())

	return &dto.Links{
		Self:  selfLink,
		First: firstLink,
		Last:  lastLink,
		Prev:  createPrevLink(queries, host, path, page),
		Next:  createNextLink(queries, host, path, page, totalPage),
	}
}

func createNextLink(queries url.Values, host, path string, page, totalPage int) string {
	if page+1 < totalPage {
		setPageQuery(queries, page+1)
	} else {
		if totalPage > 0 {
			setPageQuery(queries, totalPage)
		} else {
			setPageQuery(queries, 1)
		}
	}

	return fmt.Sprintf(linkFormat, host, path, queries.Encode())
}

func createPrevLink(queries url.Values, host, path string, page int) string {
	if page > 1 {
		setPageQuery(queries, page-1)
	} else {
		setPageQuery(queries, 1)
	}

	return fmt.Sprintf(linkFormat, host, path, queries.Encode())
}

func setPageQuery(queries url.Values, page int) {
	queries.Set("page", strconv.Itoa(page))
}
