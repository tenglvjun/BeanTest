package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// ArithmeticPage struct for the page data
type ArithmeticPage struct {
	Common PageCommon
}

// Render render the page
func (a ArithmeticPage) Render(w http.ResponseWriter, req *http.Request) {
	RenderPage(a, w, req, true)
}

// GetPath return the URL path
func (a ArithmeticPage) GetPath() string {
	return a.Common.Path
}

// GetTplFile return the tempalte file path
func (a ArithmeticPage) GetTplFile() string {
	return a.Common.TplFile
}

// GetArithmetics generate questions and return.
func (a ArithmeticPage) GetArithmetics(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	ret := Result{}

	if method != "POST" {
		ret.Code = RESULT_FAILED
		ret.Msg = ERROR_MSG_INVALID_METHOD
		ret.Data = nil
	} else {
		err := req.ParseForm()
		if err != nil {
			ret.Code = RESULT_FAILED
			ret.Msg = ERROR_MSG_PARSE_PARAMS_FAILED
			ret.Data = nil
		} else {
			r, err := strconv.Atoi(req.FormValue("range"))
			if err != nil {
				ret.Code = RESULT_FAILED
				ret.Msg = ERROR_MSG_INVALID_PARAMS
				ret.Data = nil
			} else {
				count, err := strconv.Atoi(req.FormValue("count"))
				if err != nil {
					ret.Code = RESULT_FAILED
					ret.Msg = ERROR_MSG_INVALID_PARAMS
					ret.Data = nil
				} else {
					rand.Seed(time.Now().UnixNano())

					questions := make([]string, count)

					i := 0
					for {
						if i >= len(questions) {
							break
						}

						left := rand.Intn(r)
						if left == 0 {
							left++
						}

						right := rand.Intn(r - left + 1)
						if right == 0 {
							right++
						}

						op := ""

						sign := rand.Intn(2)

						if sign == ADD_OP {
							op = "+"
						} else if sign == SUBTRAC_OP {
							op = "-"
						} else {
							continue
						}

						if (op == "-") && (left < right) {
							tmp := left
							left = right
							right = tmp
						}

						newQuestion := fmt.Sprintf("%d %s %d = ", left, op, right)

						needAdd := true

						for index := range questions {
							if questions[index] == newQuestion {
								needAdd = false
								break
							}
						}

						if needAdd {
							questions[i] = newQuestion
							i++
						}
					}

					ret.Code = RESULT_OK
					ret.Msg = SUCCESS_MSG
					ret.Data = questions
				}
			}
		}
	}
	OutputJson(ret, w)
}
