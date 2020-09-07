(window.webpackJsonp=window.webpackJsonp||[]).push([[39],{392:function(s,a,e){"use strict";e.r(a);var t=e(9),n=Object(t.a)({},(function(){var s=this,a=s.$createElement,e=s._self._c||a;return e("ContentSlotsDistributor",{attrs:{"slot-key":s.$parent.slotKey}},[e("h1",{attrs:{id:"installing-rocksdb"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#installing-rocksdb"}},[s._v("#")]),s._v(" Installing RocksDB")]),s._v(" "),e("p",[s._v("By default, Desmos uses "),e("a",{attrs:{href:"https://github.com/google/leveldb",target:"_blank",rel:"noopener noreferrer"}},[s._v("LevelDB"),e("OutboundLink")],1),s._v(" as its database backend engine. However, since version "),e("code",[s._v("v0.6.0")]),s._v(" we've also added the possibility of optionally using "),e("a",{attrs:{href:"https://github.com/facebook/rocksdb",target:"_blank",rel:"noopener noreferrer"}},[s._v("Facebook's RocksDB"),e("OutboundLink")],1),s._v(", which, although still being experimental, is know to be faster and could lead to lower syncing times. If you want to try out RocksDB (which we suggest you to do) you can take a look at our "),e("RouterLink",{attrs:{to:"/fullnode/rocksdb-installation.html"}},[s._v("RocksDB installation guide")]),s._v(" before proceeding further.")],1),s._v(" "),e("p",[s._v("The following guide allows you to install "),e("a",{attrs:{href:"https://github.com/facebook/rocksdb",target:"_blank",rel:"noopener noreferrer"}},[s._v("Facebook's RocksDB"),e("OutboundLink")],1),s._v(" inside your local machine, so that you will be able to use as Desmos' backend engine for better performances.")]),s._v(" "),e("h2",{attrs:{id:"_1-install-the-dependencies"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#_1-install-the-dependencies"}},[s._v("#")]),s._v(" 1. Install the dependencies")]),s._v(" "),e("p",[s._v("The first thing to do is to install all the dependencies for RocksDB. Following you will find the installation guide for both Ubuntu and MacOS. If you have a different operative system you can refer to the "),e("a",{attrs:{href:"https://github.com/facebook/rocksdb/blob/master/INSTALL.md",target:"_blank",rel:"noopener noreferrer"}},[s._v("official installation guide"),e("OutboundLink")],1)]),s._v(" "),e("Tabs",{attrs:{type:"border-card"}},[e("Tab",{attrs:{label:"Linux"}},[e("div",{staticClass:"language-bash line-numbers-mode"},[e("pre",{pre:!0,attrs:{class:"language-bash"}},[e("code",[e("span",{pre:!0,attrs:{class:"token function"}},[s._v("sudo")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("apt-get")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v(" -y "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n  libgflags-dev "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n  libsnappy-dev "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n  zlib1g-dev "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n  libbz2-dev "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n  libzstd-dev "),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v("\\")]),s._v("\n  liblz4-dev\n")])]),s._v(" "),e("div",{staticClass:"line-numbers-wrapper"},[e("span",{staticClass:"line-number"},[s._v("1")]),e("br"),e("span",{staticClass:"line-number"},[s._v("2")]),e("br"),e("span",{staticClass:"line-number"},[s._v("3")]),e("br"),e("span",{staticClass:"line-number"},[s._v("4")]),e("br"),e("span",{staticClass:"line-number"},[s._v("5")]),e("br"),e("span",{staticClass:"line-number"},[s._v("6")]),e("br"),e("span",{staticClass:"line-number"},[s._v("7")]),e("br")])])]),s._v(" "),e("Tab",{attrs:{label:"MacOS"}},[e("div",{staticClass:"language-bash line-numbers-mode"},[e("pre",{pre:!0,attrs:{class:"language-bash"}},[e("code",[e("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# Only needed if you are first time developing of your machine")]),s._v("\nxcode-select --install\n\n"),e("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# Install the dependencies")]),s._v("\nbrew tap homebrew/versions"),e("span",{pre:!0,attrs:{class:"token punctuation"}},[s._v(";")]),s._v(" brew "),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v(" gcc48 --use-llvm\n")])]),s._v(" "),e("div",{staticClass:"line-numbers-wrapper"},[e("span",{staticClass:"line-number"},[s._v("1")]),e("br"),e("span",{staticClass:"line-number"},[s._v("2")]),e("br"),e("span",{staticClass:"line-number"},[s._v("3")]),e("br"),e("span",{staticClass:"line-number"},[s._v("4")]),e("br"),e("span",{staticClass:"line-number"},[s._v("5")]),e("br")])])])],1),s._v(" "),e("h2",{attrs:{id:"_2-install-rocksdb"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#_2-install-rocksdb"}},[s._v("#")]),s._v(" 2. Install RocksDB")]),s._v(" "),e("p",[s._v("After having installed the dependencies, you need to install RocksDB. Again, following you will find the Linux and MacOS instructions. If you are running another OS, please refer to the "),e("a",{attrs:{href:"https://github.com/facebook/rocksdb/blob/master/INSTALL.md",target:"_blank",rel:"noopener noreferrer"}},[s._v("official documentation"),e("OutboundLink")],1),s._v(".")]),s._v(" "),e("Tabs",{attrs:{type:"border-card"}},[e("Tab",{attrs:{label:"Linux"}},[e("div",{staticClass:"language-bash line-numbers-mode"},[e("pre",{pre:!0,attrs:{class:"language-bash"}},[e("code",[e("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# Clone RocksDB")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("git")]),s._v(" clone https://github.com/facebook/rocksdb.git "),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v("&&")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" rocksdb\n\n"),e("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# Build RocksDB")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token assign-left variable"}},[s._v("DEBUG_LEVEL")]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),e("span",{pre:!0,attrs:{class:"token number"}},[s._v("0")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("make")]),s._v(" shared_lib\n\n"),e("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# Install RocksDB so that Desmos can access it")]),s._v("\n"),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("sudo")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("make")]),s._v(" install-shared\n")])]),s._v(" "),e("div",{staticClass:"line-numbers-wrapper"},[e("span",{staticClass:"line-number"},[s._v("1")]),e("br"),e("span",{staticClass:"line-number"},[s._v("2")]),e("br"),e("span",{staticClass:"line-number"},[s._v("3")]),e("br"),e("span",{staticClass:"line-number"},[s._v("4")]),e("br"),e("span",{staticClass:"line-number"},[s._v("5")]),e("br"),e("span",{staticClass:"line-number"},[s._v("6")]),e("br"),e("span",{staticClass:"line-number"},[s._v("7")]),e("br"),e("span",{staticClass:"line-number"},[s._v("8")]),e("br")])])]),s._v(" "),e("Tab",{attrs:{label:"MacOS"}},[e("div",{staticClass:"language-bash line-numbers-mode"},[e("pre",{pre:!0,attrs:{class:"language-bash"}},[e("code",[s._v("brew "),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v(" rocksdb\n")])]),s._v(" "),e("div",{staticClass:"line-numbers-wrapper"},[e("span",{staticClass:"line-number"},[s._v("1")]),e("br")])])])],1),s._v(" "),e("p",[s._v("Once the installation has finished, you will be able to compile Desmos using the following command:")]),s._v(" "),e("div",{staticClass:"language-bash line-numbers-mode"},[e("pre",{pre:!0,attrs:{class:"language-bash"}},[e("code",[e("span",{pre:!0,attrs:{class:"token function"}},[s._v("make")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v(" "),e("span",{pre:!0,attrs:{class:"token assign-left variable"}},[s._v("DB_BACKEND")]),e("span",{pre:!0,attrs:{class:"token operator"}},[s._v("=")]),s._v("rocksdb\n")])]),s._v(" "),e("div",{staticClass:"line-numbers-wrapper"},[e("span",{staticClass:"line-number"},[s._v("1")]),e("br")])])],1)}),[],!1,null,null,null);a.default=n.exports}}]);