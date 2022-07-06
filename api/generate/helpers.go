package main

import "github.com/valeralabs/jenny/jen"

func addCheckFunc(f *jen.File) {
	// func check(err error, save *error) {
	f.Func().ID("check").Parameters(
		jen.ID("err").Error(),
		jen.ID("save").Op("*").Error(),
	).Block(
		// 	if err != nil {
		jen.If(jen.ID("err").Op("!=").Nil()).Block(
			// 		*save = err
			jen.ID("*save").Op("=").ID("err"),
		),
	)
}

func addServers(f *jen.File) {
	f.Type().ID("Server").String()
	f.Const().Definitions(
		jen.ID("ValeraMN").ID("Server").Op("=").Literal("mn.stx.valera.sh"),
		jen.ID("ValeraTN").ID("Server").Op("=").Literal("tn.stx.valera.sh"),
		jen.ID("HiroMN").ID("Server").Op("=").Literal("stacks-node-api.mainnet.stacks.co"),
		jen.ID("HiroTN").ID("Server").Op("=").Literal("stacks-node-api.testnet.stacks.co"),
	)
}

func addMakeGetReqFunc(f *jen.File) {
	// func makeGetReq(url string, returnedErr *error) ([]byte, int) {
	f.Func().ID("makeGetReq").Parameters(
		jen.ID("url").String(),
		jen.ID("returnedErr").Op("*").Error(),
	).Parameters(
		jen.Index().Byte(),
		jen.Int(),
	).Block(
		// 	*returnedErr = nil
		jen.ID("*returnedErr").Op("=").Nil(),
		// 	resp, err := http.Get(url)
		jen.ID("resp").Op(",").ID("err").Op(":=").Qualified("net/http", "Get").Call(jen.ID("url")),
		// 	check(err, returnedErr)
		jen.ID("check").Call(jen.ID("err"), jen.ID("returnedErr")),
		jen.Line(),

		// 	defer resp.Body.Close()
		jen.Defer().ID("resp").Dot("Body").Dot("Close").Call(),
		jen.Line(),

		// 	body, err := ioutil.ReadAll(resp.Body)
		jen.ID("body").Op(",").ID("err").Op(":=").Qualified("io/ioutil", "ReadAll").Call(jen.ID("resp").Dot("Body")),
		// 	check(err, returnedErr)
		jen.ID("check").Call(jen.ID("err"), jen.ID("returnedErr")),
		jen.Line(),

		// 	return body, resp.StatusCode
		jen.Return(jen.ID("body"), jen.ID("resp").Dot("StatusCode")),
	)
}

func addMakePostReqFunc(f *jen.File) {
	// func makePostReq(url string, sendBody string, returnedErr *error) ([]byte, int) {
	f.Func().ID("makePostReq").Parameters(
		jen.ID("url").String(),
		jen.ID("sendBody").String(),
		jen.ID("returnedErr").Op("*").Error(),
	).Parameters(
		jen.Index().Byte(),
		jen.Int(),
	).Block(
		// 	*returnedErr = nil
		jen.ID("*returnedErr").Op("=").Nil(),
		// 	resp, err := http.Post(url, "application/json", strings.NewReader(sendBody))
		jen.ID("resp").Op(",").ID("err").Op(":=").Qualified("net/http", "Post").Call(
			jen.ID("url"), 
			jen.Literal("application/json"),
			jen.Qualified("strings", "NewReader").Call(jen.ID("sendBody")),
		),
		// 	check(err, returnedErr)
		jen.ID("check").Call(jen.ID("err"), jen.ID("returnedErr")),
		jen.Line(),

		// 	defer resp.Body.Close()
		jen.Defer().ID("resp").Dot("Body").Dot("Close").Call(),
		jen.Line(),

		// 	body, err := ioutil.ReadAll(resp.Body)
		jen.ID("body").Op(",").ID("err").Op(":=").Qualified("io/ioutil", "ReadAll").Call(jen.ID("resp").Dot("Body")),
		// 	check(err, returnedErr)
		jen.ID("check").Call(jen.ID("err"), jen.ID("returnedErr")),
		jen.Line(),

		// 	return body, resp.StatusCode
		jen.Return(jen.ID("body"), jen.ID("resp").Dot("StatusCode")),
	)
}