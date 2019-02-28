package main

type Router struct {
	RouterMap map[string]InterfacePage
}

func (r *Router) Init() {
	r.RouterMap = make(map[string]InterfacePage)

	r.RouterMap["index"] = IndexPage{
		Common: PageCommon{
			Title:   "Bean Test - 主页",
			Path:    "/index",
			TplFile: "pages/index.html",
		},
	}

	r.RouterMap["error"] = ErrorPage{
		Common: PageCommon{
			Title:   "Bean Test - 错误",
			Path:    "/error",
			TplFile: "pages/error.html",
		},
	}

	r.RouterMap["welcome"] = WelcomePage{
		Common: PageCommon{
			Title:   "Bean Test - 欢迎",
			Path:    "/welcome",
			TplFile: "pages/welcome.html",
		},
	}

	r.RouterMap["arithmetic"] = ArithmeticPage{
		Common: PageCommon{
			Title:   "Bean Test - 算术",
			Path:    "/arithmetic",
			TplFile: "pages/arithmetic.html",
		},
	}

	r.RouterMap["geometry"] = GeometryPage{
		Common: PageCommon{
			Title:   "Bean Test - 几何",
			Path:    "/geometry",
			TplFile: "pages/geometry.html",
		},
	}

	r.RouterMap["wikipedia"] = WikipediaPage{
		Common: PageCommon{
			Title:   "Bean Test - 知识库",
			Path:    "/wikipedia",
			TplFile: "pages/wikipedia.html",
		},
	}
}

func (r *Router) GetPage(path string) InterfacePage {
	if inter, ok := r.RouterMap[path]; ok {
		return inter
	}

	return nil
}
