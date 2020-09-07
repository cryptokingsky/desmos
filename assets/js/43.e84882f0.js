(window.webpackJsonp=window.webpackJsonp||[]).push([[43],{396:function(e,t,a){"use strict";a.r(t);var s=a(9),i=Object(s.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"attachment"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#attachment"}},[e._v("#")]),e._v(" Attachment")]),e._v(" "),a("p",[e._v("The "),a("code",[e._v("Attachment")]),e._v(" type contains the details of a single attachment file object that can be associated within a "),a("RouterLink",{attrs:{to:"/types/posts/post.html"}},[a("code",[e._v("Post")])]),e._v(" object. With attachment, you can add images and multimedia files (vocals, documents, videos, etc.) to posts.")],1),e._v(" "),a("p",[e._v("Following you will find a description of all the fields it is composed of.")]),e._v(" "),a("h2",{attrs:{id:"uri"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#uri"}},[e._v("#")]),e._v(" "),a("code",[e._v("URI")])]),e._v(" "),a("p",[e._v("The first field of an "),a("code",[e._v("Attachment")]),e._v(" is the "),a("code",[e._v("URI")]),e._v(" field. This field should contain the URI of the attachment file that is represented.")]),e._v(" "),a("p",[e._v("When creating a "),a("code",[e._v("Post")]),e._v(" on the chain, a regular expression checks this "),a("code",[e._v("URI")]),e._v(". If the check does not pass, the post will not be stored and an error will be thrown instead:")]),e._v(" "),a("div",{staticClass:"language-phpregexp line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("^(?:http(s)?://)[\\w.-]+(?:\\.[\\w.-]+)+[\\w\\-._~:/?#[\\]@!$&'()*+,;=.]+$\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("h2",{attrs:{id:"mimetype"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#mimetype"}},[e._v("#")]),e._v(" "),a("code",[e._v("MimeType")])]),e._v(" "),a("p",[e._v("The second field of an "),a("code",[e._v("Attachment")]),e._v(" is the "),a("code",[e._v("MimeType")]),e._v(" field. This one allows you to specify the "),a("a",{attrs:{href:"https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types",target:"_blank",rel:"noopener noreferrer"}},[e._v("MIME type"),a("OutboundLink")],1),e._v(" of the included media file.")]),e._v(" "),a("p",[e._v("No check is ever performed on this field's values, and any string is accepted as valid as long as it is not empty. That being said, please make sure to use a valid MIME type each time you specify it as it will make it easier for other apps to read your data.")]),e._v(" "),a("h2",{attrs:{id:"tags"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#tags"}},[e._v("#")]),e._v(" "),a("code",[e._v("Tags")])]),e._v(" "),a("p",[a("code",[e._v("Tags")]),e._v(" is the third field of a "),a("code",[e._v("Attachment")]),e._v(". This field allows you to tag any user on a particular attachment.\nThis field can be omitted and the system will check that every tag inside the array "),a("code",[e._v("Tags")]),e._v(" is a valid "),a("code",[e._v("bech32")]),e._v("\nencoded "),a("code",[e._v("address")]),e._v(" of desmos.\nE.g.\n"),a("code",[e._v("desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr")])])])}),[],!1,null,null,null);t.default=i.exports}}]);