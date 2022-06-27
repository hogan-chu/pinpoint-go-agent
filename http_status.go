package pinpoint

import "strconv"

type httpStatusCode interface {
	isError(code int) bool
}

type httpStatusInformational struct{}

func newHttpStatusInformational() *httpStatusInformational {
	return &httpStatusInformational{}
}

func (h *httpStatusInformational) isError(code int) bool {
	return 100 <= code && code <= 199
}

type httpStatusSuccess struct{}

func newHttpStatusSuccess() *httpStatusSuccess {
	return &httpStatusSuccess{}
}

func (h *httpStatusSuccess) isError(code int) bool {
	return 200 <= code && code <= 299
}

type httpStatusRedirection struct{}

func newHttpStatusRedirection() *httpStatusRedirection {
	return &httpStatusRedirection{}
}

func (h *httpStatusRedirection) isError(code int) bool {
	return 300 <= code && code <= 399
}

type httpStatusClientError struct{}

func newHttpStatusClientError() *httpStatusClientError {
	return &httpStatusClientError{}
}

func (h *httpStatusClientError) isError(code int) bool {
	return 400 <= code && code <= 499
}

type httpStatusServerError struct{}

func newHttpStatusServerError() *httpStatusServerError {
	return &httpStatusServerError{}
}

func (h *httpStatusServerError) isError(code int) bool {
	return 500 <= code && code <= 599
}

type httpStatusDefault struct {
	statusCode int
}

func newHttpStatusDefault(code int) *httpStatusDefault {
	return &httpStatusDefault{
		statusCode: code,
	}
}

func (h *httpStatusDefault) isError(code int) bool {
	return h.statusCode == code
}

type httpStatusError struct {
	errors []httpStatusCode
}

func newHttpStatusError(config Config) *httpStatusError {
	return &httpStatusError{
		errors: setupHttpStatusErrors(config),
	}
}

func setupHttpStatusErrors(config Config) []httpStatusCode {
	var errors []httpStatusCode

	for _, s := range config.Http.StatusCodeErrors {
		if s == "5xx" {
			errors = append(errors, newHttpStatusServerError())
		} else if s == "4xx" {
			errors = append(errors, newHttpStatusClientError())
		} else if s == "3xx" {
			errors = append(errors, newHttpStatusRedirection())
		} else if s == "2xx" {
			errors = append(errors, newHttpStatusSuccess())
		} else if s == "1xx" {
			errors = append(errors, newHttpStatusInformational())
		} else {
			c, e := strconv.Atoi(s)
			if e != nil {
				c = -1
			}
			errors = append(errors, newHttpStatusDefault(c))
		}
	}

	return errors
}

func (h *httpStatusError) isError(code int) bool {
	for _, h := range h.errors {
		if h.isError(code) {
			return true
		}
	}
	return false
}