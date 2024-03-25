// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"

func mouseOverToolTip(toolTip string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_mouseOverToolTip_0c4e`,
		Function: `function __templ_mouseOverToolTip_0c4e(toolTip){const
toolTipDiv=document.querySelector(".social-tool-tip");toolTip.length>0?(toolTipDiv.classList.add("tooltip-removed"),toolTipDiv.classList.remove("tooltip-added"),setTimeout(()=>{toolTipDiv.textContent=toolTip,toolTipDiv.classList.add("tooltip-added"),toolTipDiv.classList.remove("tooltip-removed")},300)):(toolTipDiv.classList.add("tooltip-removed"),toolTipDiv.classList.remove("tooltip-added"),setTimeout(()=>{toolTipDiv.textContent=""},300));
}`,
		Call:       templ.SafeScript(`__templ_mouseOverToolTip_0c4e`, toolTip),
		CallInline: templ.SafeScriptInline(`__templ_mouseOverToolTip_0c4e`, toolTip),
	}
}