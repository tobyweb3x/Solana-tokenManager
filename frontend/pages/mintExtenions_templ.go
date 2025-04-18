// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func MintExtensionsPartial() templ.Component {
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
		templ_7745c5c3_Err = backButtonAndTextDirection(
			[]DirectionParams{
				{Direction: "Dashboard", Url: "/"},
				{Direction: "Mint Extensions", Url: "/showMintExtensions"},
				{Direction: "Select Mint Extenisons"},
			},
		).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"grid grid-cols-[55%_45%] gap-6 pr-10 h-full w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = extensionsView().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = asideElement().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<style>\n\t\t\t.swiper-slide {\n\t\t\t\twidth: fit-content;\n\t\t\t}\n\t\t</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = swipperScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func MintExtensionsPage() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = IndexPage(MintExtensionsPartial()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func swipperScript() templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<script>\n\t\tconst swiper = new Swiper( '#swiper', {\n\t\t\tdirection: 'horizontal',\n\t\t\tslidesPerView: 'auto',\n\t\t\tspaceBetween: 8,\n\t\t\tmousewheel: true,\n\t\t\tnavigation: {\n\t\t\t\tnextEl: '.swiper-nextEl',\n\t\t\t\tprevEl: '.swiper-prevEl',\n\t\t\t},\n\t\t\t// centerInsufficientSlides: true,\n\t\t\t// watchOverflow: true,\n\t\t\t// autoplay: {\n\t\t\t// \tdelay: 1000,\n\t\t\t// \tdisableOnInteraction: false,\n\t\t\t// },\n\t\t} );\n\t\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func extensionsView() templ.Component {
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<section class=\"h-full\"><section class=\"flex flex-col gap-[1.489rem] font-normal\"><p class=\"text-white font-ArchivoSemiCondensed_SemiBold text-[2.25rem] leading-[2.75rem] text-left\">Select Mint extension</p><p class=\"text-[#A3A3A3] text-base tracking-[0.005em] text-left text-wrap\">Select Mint extensions from this list below, and fill the necessary dialogue, so we can create your token22</p></section><section id=\"swiper\" class=\"swiper mt-[2.481rem]\"><div class=\"swiper-wrapper flex items-center gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := range mintExtensions {
			if lastItem := len(mintExtensions) - 1; i != lastItem {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<div class=\"swiper-slide text-[#EBEBEB] font-bold text-[0.744rem] leading-[0.901rem] text-left\"><button class=\"gradientRadio rounded-3xl py-[1.215rem] px-[2.063rem] text-nowrap\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(mintExtensions[i].ExtensionName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/mintExtenions.templ`, Line: 73, Col: 41}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</button></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<div class=\"swiper-slide text-[#EBEBEB] font-bold text-[0.744rem] leading-[0.901rem] text-left mr-[4.50rem]\"><button class=\"gradientRadio rounded-3xl py-[1.215rem] px-[2.063rem] text-nowrap\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(mintExtensions[i].ExtensionName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/mintExtenions.templ`, Line: 79, Col: 41}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</button></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</div></section><section class=\"swipers-Btn mt-[1.03rem] mr-[0.263rem] ml-auto flex justify-center items-center gap-[0.6801rem] w-fit\"><button class=\"p-[0.726rem] rounded-full bg-[#FFFFFF4D] swiper-prevEl\"><img src=\"/public/assets/svg/ArrowLeft.svg\" alt=\"\" class=\"w-[1.09rem] aspect-square\"></button> <button class=\"p-[0.726rem] rounded-full bg-CreateToken22 swiper-nextEl\"><img src=\"/public/assets/svg/ArrowLeft.svg\" alt=\"\" class=\"w-[1.09rem] aspect-square rotate-180\"></button></section><section class=\"mt-[1.503rem] rounded-3xl border-[0.019rem] border-solid border-[#A3A3A3] p-4 flex justify-between gap-2 h-3/4 overflow-y-auto\"><div class=\"\"><p class=\"font-medium text-xl text-nowrap\">Metadata Extension</p><form action=\"\" class=\"mt-4\"><fieldset class=\"fieldset bg-base-100 border border-base-300 rounded-box max-w-max px-2\"><label class=\"fieldset-label\"><input type=\"checkbox\" class=\"toggle toggle-xl gradientRadio toggle-primary peer\"> <span class=\"transition-colors peer-checked:text-green-200 peer-checked:text-transparent peer-checked:text-clip\">Enable metadata extension</span></label></fieldset><div class=\"mt-2 flex flex-col gap-3\"><label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Update Authority</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Mint</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Name</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Symbol</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">URI</span></label><div class=\"flex gap-3\"><label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Key</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Value</span></label></div><button type=\"button\" class=\"btn btn-accent\">Add more metadata</button></div></form></div><div class=\"\"><p class=\"font-medium text-xl text-nowrap\">Metadata Pointer Extension</p><form action=\"\" class=\"mt-4\"><fieldset class=\"fieldset bg-base-100 border border-base-300 rounded-box max-w-max px-2\"><label class=\"fieldset-label\"><input type=\"checkbox\" class=\"toggle toggle-xl gradientRadio toggle-primary peer\"> <span class=\"transition-colors peer-checked:text-green-200 peer-checked:text-transparent peer-checked:text-clip\">Enable metadata extension</span></label></fieldset><div class=\"mt-2 flex flex-col gap-3\"><label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Update Authority</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Update Authority</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Update Authority</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Update Authority</span></label> <label class=\"input\"><input type=\"text\" placeholder=\"\" class=\"input input-xs\"> <span class=\"label\">Update Authority</span></label></div></form></div></section></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func asideElement() templ.Component {
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
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<aside class=\"rounded-2xl border-[0.019rem] border-solid border-[#A3A3A3]\"><div></div></aside>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
