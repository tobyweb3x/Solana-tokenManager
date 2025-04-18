// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func forwardAndBackwardBtn(currDialogID, prevDialogID, nextDialogID string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"w-full mt-10 flex justify-end items-center gap-4\"><button id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%sBackBtn", currDialogID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/dialogs.templ`, Line: 7, Col: 53}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" tabindex=\"10\" _=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("on click call #%s.close() then call #%s.showModal() then halt", currDialogID, prevDialogID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/dialogs.templ`, Line: 7, Col: 178}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" class=\"py-[1.197rem] px-[1.197rem] rounded-3xl bg-[#FFFFFF1A]\"><img src=\"/public/assets/svg/continue.svg\" alt=\"\" class=\"rotate-180 w-[0.545rem] m-[0.276rem] aspect-square\"></button> <button disabled type=\"submit\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%sContinueBtn", currDialogID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/dialogs.templ`, Line: 10, Col: 80}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\" _=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("on click  call handleSubmit('%s', '%sForm') then call #%s.close() then call #%s.showModal() then halt", currDialogID, currDialogID, currDialogID, nextDialogID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/dialogs.templ`, Line: 10, Col: 259}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" class=\"opacity-40 flex gap-2 items-center py-[1.219rem] px-[2.196rem] rounded-3xl bg-CreateToken22\"><p class=\"text-sm font-bold tracking-[0.001em] text-[#FDFDFD]\">Continue </p><img src=\"/public/assets/svg/continue.svg\" alt=\"\" class=\"w-[0.545rem] m-[0.276rem] aspect-square\"></button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func dialogPageNumberAndCloseBtn(pageNumber, dialogID string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<section class=\"flex justify-between items-center gap-[28.25rem] text-[#A3A3A3]\"><p class=\"font-medium text-2xl leading-[1.816rem] text-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(pageNumber)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/dialogs.templ`, Line: 19, Col: 77}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</p><button tabindex=\"20\" class=\"outline-none\" _=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("on click call %s.close() then call %sForm.reset() then remove .gradientRadio from <%s .wrapperDivOnInput/> then add .opacity-40 to %sContinueBtn then add @disabled to %sContinueBtn", dialogID, dialogID, dialogID, dialogID, dialogID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/dialogs.templ`, Line: 20, Col: 294}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\"><img src=\"/public/assets/svg/close.svg\" alt=\"\" class=\"w-[1.766rem] aspect-square m-[0.366rem]\"></button></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func dialogTitleText(titleText string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<p class=\"font-normal text-2xl text-center my-10 text-white\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(titleText)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/dialogs.templ`, Line: 27, Col: 73}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func dialogPageOne() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<dialog id=\"dialogPageOne\" class=\"bg-[#0A041F] px-10 pt-[4.531rem] pb-[10.531rem] rounded-3xl border-solid border-[0.0625rem] border-[#FFFFFF0D] backdrop:backdrop-blur-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogPageNumberAndCloseBtn("1/4", "#dialogPageOne").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogTitleText("What Token Standard?").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<form action=\"\" id=\"dialogPageOneForm\"><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex justify-center items-center gap-6\" onclick=\"handleRadioClickDialogOne(this)\"><input type=\"radio\" name=\"pageOne\" id=\"token22\" value=\"token22\" tabindex=\"1\" class=\"w-[1.275rem] aspect-square opacity-70 rounded-[0.0231rem] cursor-pointer \"> <label for=\"token22\" class=\"cursor-pointer\"><p class=\"text-[#EBEBEB] font-normal text-[1.375rem] leading-7 -tracking-[0.0025em] text-left\">Token22</p><p class=\"text-[#A3A3A3] font-normal text-sm tracking-[0.004em] text-left\">Lorem ipsum dolor sit amet consectetur adipiscing elit Ut et massa mi. Aliquam in hendrerit urna. Pellentesque sit amet sapien fringilla, mattis</p></label></div><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex justify-center items-center gap-6\" onclick=\"handleRadioClickDialogOne(this)\"><input type=\"radio\" name=\"pageOne\" id=\"zkTokens\" value=\"zkTokens\" tabindex=\"1\" class=\"w-[1.275rem] aspect-square opacity-70 rounded-[0.0231rem] cursor-pointer\"> <label for=\"zkTokens\" class=\"cursor-pointer\"><p class=\"text-[#EBEBEB] font-normal text-[1.375rem] leading-7 -tracking-[0.0025em] text-left\">zk-Compressed Tokens</p><p class=\"text-[#A3A3A3] font-normal text-sm tracking-[0.004em] text-left\">Lorem ipsum dolor sit amet consectetur adipiscing elit Ut et massa mi. Aliquam in hendrerit urna. Pellentesque sit amet sapien fringilla, mattis</p></label></div><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex justify-center items-center gap-6\" onclick=\"handleRadioClickDialogOne(this)\"><input type=\"radio\" name=\"pageOne\" id=\"nanoTokens\" value=\"nanoTokens\" tabindex=\"1\" class=\"w-[1.275rem] aspect-square opacity-70 rounded-[0.0231rem] cursor-pointer\"> <label for=\"nanoTokens\" class=\"cursor-pointer\"><p class=\"text-[#EBEBEB] font-normal text-[1.375rem] leading-7 -tracking-[0.0025em] text-left\">NanoTokens</p><p class=\"text-[#A3A3A3] font-normal text-sm tracking-[0.004em] text-left\">Lorem ipsum dolor sit amet consectetur adipiscing elit Ut et massa mi. Aliquam in hendrerit urna. Pellentesque sit amet sapien fringilla, mattis</p></label></div><button disabled type=\"submit\" id=\"dialogPageOneContinueBtn\" _=\"on click call handleSubmit(&#39;dialogPageOne&#39;, &#39;dialogPageOneForm&#39;) then call #dialogPageOne.close() then call #dialogPageTwo.showModal() then halt\" class=\"opacity-40 ml-auto mt-10 flex gap-2 items-center py-[1.219rem] px-[2.196rem] cursor-pointer rounded-3xl bg-CreateToken22\"><p class=\"text-sm font-bold tracking-[0.001em] text-[#FDFDFD]\">Continue</p><img src=\"/public/assets/svg/continue.svg\" alt=\"\" class=\"w-[0.545rem] m-[0.276rem] aspect-square\"></button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = scriptAndStylesForDialogs().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func dialogPageTwo() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<dialog id=\"dialogPageTwo\" class=\"bg-[#0A041F] px-10 pt-[4.531rem] pb-[10.531rem] rounded-3xl border-solid border-[0.0625rem] border-[#FFFFFF0D] backdrop:backdrop-blur-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogPageNumberAndCloseBtn("2/4", "#dialogPageTwo").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogTitleText("What type of token is this?").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "<form action=\"\" id=\"dialogPageTwoForm\"><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex justify-center items-center gap-6\" onclick=\"handleRadioClickDialogTwo(this)\"><input type=\"radio\" name=\"pageTwo\" id=\"tokens\" value=\"tokens\" tabindex=\"1\" class=\"w-[1.275rem] aspect-square opacity-70 rounded-[0.0231rem] cursor-pointer \"> <label for=\"tokens\" class=\"cursor-pointer\"><p class=\"text-[#EBEBEB] font-normal text-[1.375rem] leading-7 -tracking-[0.0025em] text-left\">Fungible Tokens</p><p class=\"text-[#A3A3A3] font-normal text-sm tracking-[0.004em] text-left\">Lorem ipsum dolor sit amet consectetur adipiscing elit Ut et massa mi. Aliquam in hendrerit urna. Pellentesque sit amet sapien fringilla, mattis</p></label></div><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex justify-center items-center gap-6\" onclick=\"handleRadioClickDialogTwo(this)\"><input type=\"radio\" name=\"pageTwo\" id=\"sft\" value=\"sft\" tabindex=\"1\" class=\"w-[1.275rem] aspect-square opacity-70 rounded-[0.0231rem] cursor-pointer\"> <label for=\"sft\" class=\"cursor-pointer\"><p class=\"text-[#EBEBEB] font-normal text-[1.375rem] leading-7 -tracking-[0.0025em] text-left\">SemiFungible Tokens (SFTs)</p><p class=\"text-[#A3A3A3] font-normal text-sm tracking-[0.004em] text-left\">Lorem ipsum dolor sit amet consectetur adipiscing elit Ut et massa mi. Aliquam in hendrerit urna. Pellentesque sit amet sapien fringilla, mattis</p></label></div><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex justify-center items-center gap-6\" onclick=\"handleRadioClickDialogTwo(this)\"><input type=\"radio\" name=\"pageTwo\" id=\"nft\" value=\"nft\" tabindex=\"1\" class=\"w-[1.275rem] aspect-square opacity-70 rounded-[0.0231rem] cursor-pointer\"> <label for=\"nft\" class=\"cursor-pointer\"><p class=\"text-[#EBEBEB] font-normal text-[1.375rem] leading-7 -tracking-[0.0025em] text-left\">NonFungible Tokens (NFTs)</p><p class=\"text-[#A3A3A3] font-normal text-sm tracking-[0.004em] text-left\">Lorem ipsum dolor sit amet consectetur adipiscing elit Ut et massa mi. Aliquam in hendrerit urna. Pellentesque sit amet sapien fringilla, mattis</p></label></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = forwardAndBackwardBtn("dialogPageTwo", "dialogPageOne", "dialogPageThree").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = scriptAndStylesForDialogs().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func dialogPageThree() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var13 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var13 == nil {
			templ_7745c5c3_Var13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<dialog id=\"dialogPageThree\" class=\"bg-[#0A041F] px-10 pt-[4.531rem] pb-20 rounded-3xl border-solid border-[0.0625rem] border-[#FFFFFF0D] backdrop:backdrop-blur-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogPageNumberAndCloseBtn("3/4", "#dialogPageThree").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogTitleText("Lets know more about your token").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "<form id=\"dialogPageThreeForm\" class=\"text-[#EBEBEB]\"><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex flex-col gap-4\"><label for=\"tokenName\" class=\"text-lg font-normal tracking-[0.005em]\">Token names</label> <input required type=\"text\" name=\"pageThree\" id=\"tokenName\" value=\"\" tabindex=\"1\" placeholder=\"Enter Token name\" class=\"w-[33.563rem] aspect-[537/56] pl-2 rounded-3xl bg-inherit border-solid border-[#D1D1D1] border-[0.0313rem] outline-none\n\t\t\t\t\tplaceholder:text-[#616161] placeholder:text-xs placeholder:font-bold placeholder:tracking-[0.005em]\"></div><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex flex-col gap-4\"><label for=\"tokenSymbol\" class=\"text-lg font-normal tracking-[0.005em]\">Token symbol</label> <input required type=\"text\" name=\"pageThree\" id=\"tokenSymbol\" value=\"\" tabindex=\"1\" placeholder=\"Enter Token symbol\" class=\"w-[33.563rem] aspect-[537/56] pl-2 rounded-3xl bg-inherit border-solid border-[#D1D1D1] border-[0.0313rem] outline-none\n\t\t\t\t\tplaceholder:text-[#616161] placeholder:text-xs placeholder:font-bold placeholder:tracking-[0.005em]\"></div><div class=\"wrapperDivOnInput py-6 pl-6 pr-[3.063rem] flex flex-col gap-4\"><label for=\"tokenAvatar\" class=\"text-lg font-normal tracking-[0.005em]\">Token Image</label> <input required type=\"url\" name=\"pageThree\" id=\"tokenAvatar\" value=\"\" tabindex=\"1\" placeholder=\"Enter Token image URL\" class=\"w-[33.563rem] aspect-[537/56] pl-2 rounded-3xl bg-inherit border-solid border-[#D1D1D1] border-[0.0313rem] outline-none\n\t\t\t\t\tplaceholder:text-[#616161] placeholder:text-xs placeholder:font-bold placeholder:tracking-[0.005em]\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = forwardAndBackwardBtn("dialogPageThree", "dialogPageTwo", "dialogPageFour").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = scriptAndStylesForDialogs().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<style>\n\t\t// \t input:not(:placeholder-shown):valid {\n\t\t// \t\toutline: none;\n\t\t// \t\tborder-style: solid;\n\t\t// \t\tborder-width: 0.0313rem;\n\t\t// \t\tborder-image-source: linear-gradient(90deg, #AD6AFF 0%, #12DB88 100%);\n\t\t// \t\tpadding: 5px;\n\t\t// \t}\n\t\t\tinput:not(:placeholder-shown):invalid {\n\t\t\t\toutline: none;\n\t\t\t\tborder-style: solid;\n\t\t\t\tborder-width: 0.0313rem;\n\t\t\t\tborder-color: #FF4B4B;\n\t\t\t}\n\t\t// \tinput:focus:invalid {\n\t\t// \t\toutline: none;\n\t\t// \t\tborder-style: solid;\n\t\t// \t\tborder-width: 0.0313rem;\n\t\t// \t\toutline-color: yellow;\n\t\t// \t}\n\t\t</style></dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func dialogPageFour() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var14 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var14 == nil {
			templ_7745c5c3_Var14 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "<dialog id=\"dialogPageFour\" class=\"bg-[#0A041F] px-10 pt-[4.531rem] pb-[10.531rem] rounded-3xl border-solid border-[0.0313rem] border-[#FFFFFF0D] backdrop:backdrop-blur-sm overflow-y-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogPageNumberAndCloseBtn("4/4", "#dialogPageFour").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = dialogTitleText("Generate mint address").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<section class=\"grid grid-cols-2 gap-4 justify-center items-center bg-[#FFFFFF1A] p-2 rounded-3xl\"><button autofocus class=\"forRandom py-[1.188rem] px-[3.984rem] rounded-3xl border-[0.103rem] border-transparent border-solid focus:outline-none\"><p class=\"text-[#FFFFFF] text-nowrap text-sm text-center tracking-[0.001em] font-bold\">Generate Randomss</p></button> <button class=\"forCustomise py-[1.188rem] px-[3.984rem] focus:rounded-3xl focus:border-[0.103rem] focus:border-transparent focus:border-solid focus:outline-none\"><p class=\"text-[#FFFFFF] text-nowrap text-sm text-center tracking-[0.001em] font-bold\">Customize</p></button></section><section class=\"random mt-[7.5rem] flex justify-center items-center \"><div><p class=\"text-[1.375rem] leading-7 font-normal -tracking-[0.0025em] text-[#EBEBEB] text-center\">A random mint address would be <br>generated for your token</p><section class=\"w-full mt-10 flex justify-center items-center gap-4\"><button _=\"on click call #dialogPageFour.close() then call #dialogPageThree.showModal() then halt\" class=\"py-[1.197rem] px-[1.197rem] rounded-3xl bg-[#FFFFFF1A]\"><img src=\"/public/assets/svg/continue.svg\" alt=\"\" class=\"rotate-180 w-[0.545rem] m-[0.276rem] aspect-square\"></button> <button hx-post=\"/showMintExtensions\" hx-target=\"#mainContent\" hx-vals=\"js:{tokenInfo: generateRandomKeyPair()}\" hx-swap=\"innerHTML\" hx-push-url=\"true\" _=\"on click call #dialogPageFour.close()\" class=\"flex gap-2 items-center py-[1.219rem] px-[2.196rem] rounded-3xl bg-CreateToken22\"><p class=\"text-sm font-bold tracking-[0.001em] text-[#FDFDFD]\">Continue </p><img src=\"/public/assets/svg/continue.svg\" alt=\"\" class=\"w-[0.545rem] m-[0.276rem] aspect-square\"></button></section></div></section><section class=\"customise mt-10 hidden\"><form action=\"\" id=\"dialogPageFourForm\" class=\"grid grid-rows-4 gap-6 justify-center text-[#A3A3A3]\"><div class=\"wrapperDivOnInput grid grid-cols-dialogInput gap-8 justify-stretch items-center\"><label for=\"startWith\" class=\"text-xl font-normal tracking-[0.005em]\">Start with</label> <input type=\"text\" name=\"pageFour\" id=\"startWith\" placeholder=\"YOU CAN LEAVE EMPTY\" class=\"w-[33.563rem] aspect-[537/56] ml-[1.21rem] placeholder:text-[#616161] placeholder:text-sm  placeholder:font-extrabold placeholder:tracking-[0.005em] pl-4 bg-inherit rounded-3xl border-solid border-[#D1D1D1] border-[0.0313rem] outline-white\"></div><div class=\"wrapperDivOnInput grid grid-cols-dialogInput gap-8 justify-stretch items-center\"><label for=\"endWith\" class=\"text-xl font-normal tracking-[0.005em]\">End with</label> <input type=\"text\" name=\"pageFour\" id=\"endWith\" placeholder=\"YOU CAN LEAVE EMPTY\" class=\"w-[33.563rem] aspect-[537/56] ml-[1.21rem] placeholder:text-[#616161] placeholder:text-sm  placeholder:font-extrabold placeholder:tracking-[0.005em] pl-4 bg-inherit rounded-3xl border-solid border-[#D1D1D1] border-[0.0313rem] outline-white\"></div><div class=\"wrapperDivOnInput flex gap-4 justify-between items-center\"><div class=\"flex gap-2 items-center\"><input type=\"checkbox\" name=\"pageFour\" id=\"ignoreCase\" class=\"w-4 aspect-square accent-emerald-500/25 \"> <label for=\"ignoreCase\" class=\"text-xl font-normal tracking-[0.005em]\">Ignore case</label></div><div class=\"flex gap-2 items-center\"><input type=\"checkbox\" name=\"pageFour\" id=\"count\" class=\"w-4 aspect-square  accent-emerald-500/25\"> <label for=\"count\" class=\"text-xl font-normal tracking-[0.005em]\">Count</label></div><div class=\"flex gap-2 items-center\"><input type=\"number\" name=\"pageFour\" id=\"numThreads\" min=\"2\" max=\"10\" step=\"1\" value=\"8\" class=\"w-7 aspect-square text-black font-medium text-center outline-none\"> <label for=\"numThreads\" class=\"text-xl font-normal tracking-[0.005em]\">Number Of Threads</label></div><div class=\"flex gap-2 items-center\"><input checked type=\"checkbox\" name=\"pageFour\" id=\"outPutFile\" class=\"w-4 aspect-square accent-emerald-500/25\"> <label for=\"outPutFile\" class=\"text-xl font-normal tracking-[0.005em]\">Output File</label></div></div><div class=\"wrapperDivOnInput grid grid-cols-dialogInput gap-8 justify-stretch items-center\"><label for=\"passPhrase\" class=\"text-xl font-normal tracking-[0.005em]\">PassPhrase</label> <input type=\"text\" name=\"pageFour\" id=\"passPhrase\" placeholder=\"YOU CAN LEAVE EMPTY\" class=\"w-[33.563rem] aspect-[537/56] ml-[1.21rem] placeholder:text-[#616161] placeholder:text-sm  placeholder:font-extrabold placeholder:tracking-[0.005em] pl-4 bg-inherit rounded-3xl border-solid border-[#D1D1D1] border-[0.0313rem] outline-white\"></div><div class=\"wrapperDivOnInput grid grid-cols-dialogInput gap-8 justify-stretch items-center\"><label for=\"derivationPath\" class=\"text-xl font-normal tracking-[0.005em]\">Derivation path</label> <input type=\"text\" name=\"pageFour\" id=\"derivationPath\" placeholder=\"EXAMPLE: m/44&#39;/501&#39;/0&#39;/0&#39;\" class=\"w-[33.563rem] aspect-[537/56] ml-[1.21rem] placeholder:text-[#616161] placeholder:text-sm  placeholder:font-extrabold placeholder:tracking-[0.005em] pl-4 bg-inherit rounded-3xl border-solid border-[#D1D1D1] border-[0.0313rem] outline-white\"></div><div class=\"grid grid-cols-2 justify-center items-center gap-8\"><div class=\"wrappDivOnInput wrappDivOnSelect\"><label for=\"language\" class=\"text-lg font-normal tracking-[0.005em]\">SeedPhrase language</label> <select name=\"pageFour\" id=\"language\" class=\"ml-auto text-xs font-normal tracking-[0.005em] bg-inherit outline-none\"><option value=\"english\" class=\"text-xs font-semibold bg-[#0A041F]\">in English</option> <option value=\"chinese-simplified\" class=\"text-xs font-semibold bg-[#0A041F]\">in Chinese-simplified</option> <option value=\"chinese-traditional\" class=\"text-xs font-semibold bg-[#0A041F]\">in Chinese-traditional</option> <option value=\"japanese\" class=\"text-xs font-semibold bg-[#0A041F]\">in Japanese</option> <option value=\"spanish\" class=\"text-xs font-semibold bg-[#0A041F]\">in Spanish</option> <option value=\"korean\" class=\"text-xs font-semibold bg-[#0A041F]\">in Korean</option> <option value=\"french\" class=\"text-xs font-semibold bg-[#0A041F]\">in French</option> <option value=\"italian\" class=\"text-xs font-semibold bg-[#0A041F]\">in Italian</option></select></div><div class=\"wrappDivOnInput wrappDivOnSelect\"><label for=\"wordCount\" class=\"text-lg font-normal tracking-[0.005em]\">SeedPhrase WordCount</label> <select name=\"pageFour\" id=\"wordCount\" class=\"ml-auto text-sm font-normal tracking-[0.005em] bg-inherit outline-none\"><option value=\"12\" class=\"text-sm font-bold bg-[#0A041F]\">12-word seedphrase</option> <option value=\"15\" class=\"text-sm font-bold bg-[#0A041F]\">15-word seedphrase</option> <option value=\"18\" class=\"text-sm font-bold bg-[#0A041F]\">18-word seedphrase</option> <option value=\"21\" class=\"text-sm font-bold bg-[#0A041F]\">21-word seedphrase</option> <option value=\"24\" class=\"text-sm font-bold bg-[#0A041F]\">24-word seedphrase</option></select></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = forwardAndBackwardBtn("dialogPageFour", "dialogPageThree", "dialogPageFour").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</form></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = scriptAndStylesForDialogs().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "</dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

// var dialogScriptsAndStyles = templ.NewOnceHandle()
func scriptAndStylesForDialogs() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var15 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var15 == nil {
			templ_7745c5c3_Var15 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "<style>\n\t\t.gradientRadio {\n\t\t\tbackground: \n\t\t\t\tlinear-gradient(#1b012a 0 0) padding-box, /* this is your background */\n\t\t\t\tlinear-gradient(90deg, #AD6AFF 0%, #12DB88 100%) border-box;\n\t\t\tborder-radius: 16px;\n\t\t\tborder: 1px solid transparent;\n\t\t\topacity: 0px;\n\t\t}\n\t\t</style><script defer>\n\t\tlet TokenInfo = {}\n\n\t// dialog 1/4\n\t\tfunction handleRadioClickDialogOne(param) {\n\t\t\tconst dialogPageOneContinueBtn = document.getElementById('dialogPageOneContinueBtn');\n\t\t\tdialogPageOneContinueBtn.disabled = false;\n\t\t\tdialogPageOneContinueBtn.classList.remove('opacity-40');\n\n\n\t\t\tconst parentDivs = document.querySelectorAll('#dialogPageOne .wrapperDivOnInput');\n\t\t\tparentDivs.forEach(div => {\n\t\t\t\tdiv.classList.remove('gradientRadio');\n\t\t\t});\n\n\t\t\tparam.classList.add('gradientRadio');\n\t\t}\n\n\t// dialog 2/4\n\t\tfunction handleRadioClickDialogTwo(param) {\n\t\t\tconst dialogPageTwoContinueBtn = document.getElementById('dialogPageTwoContinueBtn');\n\t\t\tdialogPageTwoContinueBtn.disabled = false;\n\t\t\tdialogPageTwoContinueBtn.classList.remove('opacity-40');\n\n\n\t\t\tconst parentDivs = document.querySelectorAll('#dialogPageTwo .wrapperDivOnInput');\n\t\t\tparentDivs.forEach(div => {\n\t\t\t\tdiv.classList.remove('gradientRadio');\n\t\t\t});\n\n\t\t\tparam.classList.add('gradientRadio');\n\t\t}\n\n\t// dialog 3/4\n\t\tdocument.addEventListener('DOMContentLoaded', () => {\n\t\t\tconst dialogPageThreeForm = document.getElementById('dialogPageThreeForm');\n\t\t\tconst continueButton = document.getElementById('dialogPageThreeContinueBtn');\n\t\t\t\t\n\t\t\tfunction checkFormValidity() {\n\t\t\t\tif (dialogPageThreeForm.checkValidity()) {\n\t\t\t\t\tcontinueButton.disabled = false;\n\t\t\t\t\tcontinueButton.classList.remove('opacity-40');\n\t\t\t\t} else {\n\t\t\t\t\tcontinueButton.disabled = true;\n\t\t\t\t\tcontinueButton.classList.add('opacity-40');\n\t\t\t\t}\n\t\t\t}\n\n\t\t\tdialogPageThreeForm.addEventListener('input', checkFormValidity);\n\n\t// dialog 4/4\n\t\t\tconst forRandomBtn = document.querySelector('.forRandom');\n\t\t\tconst forCustomiseBtn = document.querySelector('.forCustomise');\n\t\t\tconst randomSection = document.querySelector('.random');\n\t\t\tconst customiseSection = document.querySelector('.customise');\n\n\t\t\tforRandomBtn.addEventListener('focus', () => {\n\t\t\t\trandomSection.style.display = 'flex';\n\t\t\t\tforRandomBtn.classList.add('gradientRadio')\n\t\t\t\tcustomiseSection.style.display = 'none';\n\t\t\t\tforCustomiseBtn.classList.remove('gradientRadio')\n\t\t\t});\n\n\t\t\tforCustomiseBtn.addEventListener('focus', () => {\n\t\t\t\tcustomiseSection.style.display = 'flex';\n\t\t\t\tforCustomiseBtn.classList.add('gradientRadio')\n\t\t\t\trandomSection.style.display = 'none';\n\t\t\t\tforRandomBtn.classList.remove('gradientRadio')\n\t\t\t\tconst x = document.getElementById('dialogPageFour')\n\t\t\t\tx.classList.remove('pb-[10.531rem]')\n\t\t\t\tx.classList.add('pb-[5.266rem]')\n\t\t\t});\n\t\t});\n\n\t// input collection\n\t\tfunction handleSubmit(dialogID, dialogFormID) {\n\t\t\tevent.preventDefault();\n\t\t\t\n\t\t\tconst form = document.getElementById(dialogFormID)\n\t\t\tconst formData = new FormData(form);\n\t\t\t// console.log(formData)\n\n\t\t\tlet selectedValue;\n\t\t\tswitch (dialogID) {\n\t\t\t\tcase \"dialogPageOne\":\n\t\t\t\t\tselectedValue = formData.get('pageOne');\n\t\t\t\t\tTokenInfo.tokenStandard = selectedValue\n\t\t\t\t\tbreak\n\n\t\t\t\tcase \"dialogPageTwo\":\n\t\t\t\t\tselectedValue = formData.get('pageTwo');\n\t\t\t\t\tTokenInfo.tokenType = selectedValue\n\t\t\t\t\tbreak\n\n\t\t\t\tcase \"dialogPageThree\":\n\t\t\t\t\tselectedValue = formData.getAll('pageThree');\n\t\t\t\t\tconst propertyNames = ['tokenName', 'tokenSymbol', 'tokenUrl'];\n\t\t\t\t\tselectedValue.forEach((value, index) => {\n\t\t\t\t\t\tif (index < propertyNames.length) {\n\t\t\t\t\t\t\tTokenInfo[propertyNames[index]] = value;\n\t\t\t\t\t\t}\n\t\t\t\t\t});\n\t\t\t\t\tbreak\n\t\t\t\t\t\n\t\t\t\tcase \"dialogPageFour\":\n\t\t\t\t\tbreak\n\t\t\t}\n\n\t\t\t// console.log('Selected Value:', selectedValue);\n\t\t\tconsole.log('TokenInfo:', TokenInfo);\n\t\t}\n\n\t\tfunction generateRandomKeyPair() {\n\t\t\treturn createRandomKeyPair(TokenInfo)\t\n\t\t}\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
