(window.webpackJsonp=window.webpackJsonp||[]).push([[54],{261:function(s,e,a){"use strict";a.r(e);var t=a(0),n=Object(t.a)({},(function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("ContentSlotsDistributor",{attrs:{"slot-key":s.$parent.slotKey}},[a("h1",{attrs:{id:"version-0-3-1-update"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#version-0-3-1-update"}},[s._v("#")]),s._v(" Version "),a("code",[s._v("0.3.1")]),s._v(" update")]),s._v(" "),a("p",[s._v("The following guide explains how you can update your validator node from "),a("code",[s._v("v0.3.0")]),s._v(" to "),a("code",[s._v("v0.3.1")]),s._v(" seamlessly without loosing a single block validated in the process.")]),s._v(" "),a("h2",{attrs:{id:"phase-1-updating-desmos"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#phase-1-updating-desmos"}},[s._v("#")]),s._v(" Phase 1 - Updating Desmos")]),s._v(" "),a("ol",[a("li",[a("p",[s._v("Go into your Desmos code folder. If you didn't change the suggested folder path, it should be "),a("code",[s._v("~/desmos")]),s._v(":")]),s._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" ~/desmos\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br")])])]),s._v(" "),a("li",[a("p",[s._v("Fetch the changes and checkout tag "),a("code",[s._v("v0.3.1")]),s._v(":")]),s._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("git")]),s._v(" fetch -a\n"),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("git")]),s._v(" checkout tags/v0.3.1\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br"),a("span",{staticClass:"line-number"},[s._v("2")]),a("br")])])]),s._v(" "),a("li",[a("p",[s._v("Install the latest update:")]),s._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[s._v("make")]),s._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v("\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br")])])]),s._v(" "),a("li",[a("p",[s._v("Verify the update has been installed properly:")]),s._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[s._v("desmoscli version\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# Should ouput: 0.3.1")]),s._v("\n\ndesmosd version\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# Should ouput: 0.3.1")]),s._v("\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br"),a("span",{staticClass:"line-number"},[s._v("2")]),a("br"),a("span",{staticClass:"line-number"},[s._v("3")]),a("br"),a("span",{staticClass:"line-number"},[s._v("4")]),a("br"),a("span",{staticClass:"line-number"},[s._v("5")]),a("br")])])])]),s._v(" "),a("h2",{attrs:{id:"phase-2-update-your-node"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#phase-2-update-your-node"}},[s._v("#")]),s._v(" Phase 2 - Update your node")]),s._v(" "),a("p",[s._v("In order to update your node all you have to do is simply stopping it and re-starting it.")]),s._v(" "),a("p",[s._v("If you are running it as a server, you can run the following commands:")]),s._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[s._v("systemctl restart desmosd\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br")])]),a("p",[s._v("In order to make sure everything worked properly, you can run the following:")]),s._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[s._v("systemctl status desmosd\n")])]),s._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[s._v("1")]),a("br")])])])}),[],!1,null,null,null);e.default=n.exports}}]);