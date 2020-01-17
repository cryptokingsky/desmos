(window.webpackJsonp=window.webpackJsonp||[]).push([[31],{240:function(a,s,t){"use strict";t.r(s);var e=t(0),n=Object(e.a)({},(function(){var a=this,s=a.$createElement,t=a._self._c||s;return t("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[t("h1",{attrs:{id:"run-a-validator-on-desmos"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#run-a-validator-on-desmos"}},[a._v("#")]),a._v(" Run a Validator on Desmos")]),a._v(" "),t("h2",{attrs:{id:"what-is-a-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#what-is-a-validator"}},[a._v("#")]),a._v(" What is a Validator?")]),a._v(" "),t("p",[t("router-link",{attrs:{to:"/validators/overview.html"}},[a._v("Validators")]),a._v(" are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable or sign blocks at the same height. Please read about "),t("router-link",{attrs:{to:"/validators/validator-faq.html#how-can-validators-protect-themselves-from-denial-of-service-attacks"}},[a._v("Sentry Node Architecture")]),a._v(" to protect your node from DDOS attacks and to ensure high-availability.")],1),a._v(" "),t("div",{staticClass:"custom-block danger"},[t("p",{staticClass:"custom-block-title"},[a._v("Warning")]),a._v(" "),t("p",[a._v("If you want to become a validator for the "),t("code",[a._v("mainnet")]),a._v(", you should "),t("router-link",{attrs:{to:"/validators/security.html"}},[a._v("research security")]),a._v(".")],1)]),a._v(" "),t("h2",{attrs:{id:"running-a-fullnode"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#running-a-fullnode"}},[a._v("#")]),a._v(" Running a Fullnode")]),a._v(" "),t("p",[a._v("To become a validator, you must first have "),t("code",[a._v("desmosd")]),a._v(" and "),t("code",[a._v("desmoscli")]),a._v(" installed and be able to run a fullnode. You can first "),t("router-link",{attrs:{to:"/fullnode/overview.html"}},[a._v("setup your fullnode")]),a._v(" if you haven't yet.")],1),a._v(" "),t("h2",{attrs:{id:"create-your-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#create-your-validator"}},[a._v("#")]),a._v(" Create Your Validator")]),a._v(" "),t("p",[a._v("Your "),t("code",[a._v("desmosvalconspub")]),a._v(" (Desmos Validator Consensus Pubkey) can be used to create a new validator by staking tokens. You can find your validator pubkey by running:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmosd tendermint show-validator\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("To create your validator, just use the following command:")]),a._v(" "),t("div",{staticClass:"custom-block warning"},[t("p",{staticClass:"custom-block-title"},[a._v("WARNING")]),a._v(" "),t("p",[a._v("Don't use more "),t("code",[a._v("desmos")]),a._v(" than you have!")])]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmoscli tx staking create-validator "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --amount"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("1000000udaric "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --pubkey"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token variable"}},[t("span",{pre:!0,attrs:{class:"token variable"}},[a._v("$(")]),a._v("udaricd tendermint show-validator"),t("span",{pre:!0,attrs:{class:"token variable"}},[a._v(")")])]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --moniker"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"choose a moniker"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("chain_id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --commission-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.10"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --commission-max-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.20"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --commission-max-change-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.01"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --min-self-delegation"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"1"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --gas"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"auto"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --gas-prices"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.025udaric"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key_name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br"),t("span",{staticClass:"line-number"},[a._v("4")]),t("br"),t("span",{staticClass:"line-number"},[a._v("5")]),t("br"),t("span",{staticClass:"line-number"},[a._v("6")]),t("br"),t("span",{staticClass:"line-number"},[a._v("7")]),t("br"),t("span",{staticClass:"line-number"},[a._v("8")]),t("br"),t("span",{staticClass:"line-number"},[a._v("9")]),t("br"),t("span",{staticClass:"line-number"},[a._v("10")]),t("br"),t("span",{staticClass:"line-number"},[a._v("11")]),t("br"),t("span",{staticClass:"line-number"},[a._v("12")]),t("br")])]),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[a._v("TIP")]),a._v(" "),t("p",[a._v("When specifying commission parameters, the "),t("code",[a._v("commission-max-change-rate")]),a._v(" is used to measure % "),t("em",[a._v("point")]),a._v(" change over the "),t("code",[a._v("commission-rate")]),a._v(". E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.")])]),a._v(" "),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[a._v("TIP")]),a._v(" "),t("p",[t("code",[a._v("Min-self-delegation")]),a._v(" is a stritly positive integer that represents the minimum amount of self-delegated voting power your validator must always have. A "),t("code",[a._v("min-self-delegation")]),a._v(" of 1 means your validator will never have a self-delegation lower than "),t("code",[a._v("1daric")]),a._v(", or "),t("code",[a._v("1000000udaric")])])]),a._v(" "),t("p",[a._v("You can confirm that you are in the validator set by using a third party explorer.")]),a._v(" "),t("h2",{attrs:{id:"participate-in-genesis-as-a-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#participate-in-genesis-as-a-validator"}},[a._v("#")]),a._v(" Participate in Genesis as a Validator")]),a._v(" "),t("p",[a._v("If you want to participate in genesis as a validator, you need to justify that\nyou have some stake at genesis, create one (or multiple) transactions to bond this stake to your validator address, and include this transaction in the genesis file.")]),a._v(" "),t("p",[a._v("Your "),t("code",[a._v("desmosvalconspub")]),a._v(" can be used to create a new validator by staking tokens. You can find your validator pubkey by running:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmosd tendermint show-validator\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("Next, craft your "),t("code",[a._v("desmosd gentx")]),a._v(" command.")]),a._v(" "),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[a._v("TIP")]),a._v(" "),t("p",[a._v("A "),t("code",[a._v("gentx")]),a._v(" is a JSON file carrying a self-delegation. All genesis transactions are collected by a "),t("code",[a._v("genesis coordinator")]),a._v(" and validated against an initial "),t("code",[a._v("genesis.json")]),a._v(".")])]),a._v(" "),t("div",{staticClass:"custom-block warning"},[t("p",{staticClass:"custom-block-title"},[a._v("Note")]),a._v(" "),t("p",[a._v("Don't use more "),t("code",[a._v("desmos")]),a._v(" than you have!")])]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmosd gentx "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --amount "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("amount_of_delegation_desmos"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --commission-rate "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("commission_rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --commission-max-rate "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("commission_max_rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --commission-max-change-rate "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("commission_max_change_rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --pubkey "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("consensus_pubkey"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --name "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key_name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br"),t("span",{staticClass:"line-number"},[a._v("4")]),t("br"),t("span",{staticClass:"line-number"},[a._v("5")]),t("br"),t("span",{staticClass:"line-number"},[a._v("6")]),t("br"),t("span",{staticClass:"line-number"},[a._v("7")]),t("br")])]),t("div",{staticClass:"custom-block tip"},[t("p",{staticClass:"custom-block-title"},[a._v("TIP")]),a._v(" "),t("p",[a._v("When specifying commission parameters, the "),t("code",[a._v("commission-max-change-rate")]),a._v(" is used to measure % "),t("em",[a._v("point")]),a._v(" change over the "),t("code",[a._v("commission-rate")]),a._v(". E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.")])]),a._v(" "),t("p",[a._v("You can then submit your "),t("code",[a._v("gentx")]),a._v(" on the "),t("a",{attrs:{href:"https://github.com/desmos/launch",target:"_blank",rel:"noopener noreferrer"}},[a._v("launch repository"),t("OutboundLink")],1),a._v(". These "),t("code",[a._v("gentx")]),a._v(" will be used to form the final genesis file.")]),a._v(" "),t("h2",{attrs:{id:"edit-validator-description"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#edit-validator-description"}},[a._v("#")]),a._v(" Edit Validator Description")]),a._v(" "),t("p",[a._v("You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below. If a flag is not included in the command the field will default to empty ("),t("code",[a._v("--moniker")]),a._v(" defaults to the machine name) if the field has never been set or remain the same if it has been set in the past.")]),a._v(" "),t("p",[a._v("The <key_name> specifies which validator you are editing. If you choose to not include certain flags, remember that the --from flag must be included to identify the validator to update.")]),a._v(" "),t("p",[a._v("The "),t("code",[a._v("--identity")]),a._v(" can be used as to verify identity with systems like Keybase or UPort. When using with Keybase "),t("code",[a._v("--identity")]),a._v(" should be populated with a 16-digit string that is generated with a "),t("a",{attrs:{href:"https://keybase.io",target:"_blank",rel:"noopener noreferrer"}},[a._v("keybase.io"),t("OutboundLink")],1),a._v(" account. It's a cryptographically secure method of verifying your identity across multiple online networks. The Keybase API allows us to retrieve your Keybase avatar. This is how you can add a logo to your validator profile.")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmoscli tx staking edit-validator\n  --moniker"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"choose a moniker"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --website"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"https://desmos.network"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --identity"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),a._v("6A0D65E29A4CBC8E "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --details"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"To infinity and beyond!"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("chain_id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --gas"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"auto"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --gas-prices"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.025desmos"')]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key_name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n  --commission-rate"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"0.10"')]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br"),t("span",{staticClass:"line-number"},[a._v("4")]),t("br"),t("span",{staticClass:"line-number"},[a._v("5")]),t("br"),t("span",{staticClass:"line-number"},[a._v("6")]),t("br"),t("span",{staticClass:"line-number"},[a._v("7")]),t("br"),t("span",{staticClass:"line-number"},[a._v("8")]),t("br"),t("span",{staticClass:"line-number"},[a._v("9")]),t("br"),t("span",{staticClass:"line-number"},[a._v("10")]),t("br")])]),t("p",[t("strong",[a._v("Note")]),a._v(": The "),t("code",[a._v("commission-rate")]),a._v(" value must adhere to the following invariants:")]),a._v(" "),t("ul",[t("li",[a._v("Must be between 0 and the validator's "),t("code",[a._v("commission-max-rate")])]),a._v(" "),t("li",[a._v("Must not exceed the validator's "),t("code",[a._v("commission-max-change-rate")]),a._v(" which is maximum\n% point change rate "),t("strong",[a._v("per day")]),a._v(". In other words, a validator can only change\nits commission once per day and within "),t("code",[a._v("commission-max-change-rate")]),a._v(" bounds.")])]),a._v(" "),t("h2",{attrs:{id:"view-validator-description"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#view-validator-description"}},[a._v("#")]),a._v(" View Validator Description")]),a._v(" "),t("p",[a._v("View the validator's information with this command:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmoscli query staking validator "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("account_desmos"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmoscli tx slashing unjail "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n\t--from"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("key_name"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("\\")]),a._v("\n\t--chain-id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("<")]),a._v("chain_id"),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v(">")]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br")])]),t("h2",{attrs:{id:"confirm-your-validator-is-running"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#confirm-your-validator-is-running"}},[a._v("#")]),a._v(" Confirm Your Validator is Running")]),a._v(" "),t("p",[a._v("Your validator is active if the following command returns anything:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmoscli query tendermint-validator-set "),t("span",{pre:!0,attrs:{class:"token operator"}},[a._v("|")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token function"}},[a._v("grep")]),a._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[a._v('"'),t("span",{pre:!0,attrs:{class:"token variable"}},[t("span",{pre:!0,attrs:{class:"token variable"}},[a._v("$(")]),a._v("desmosd tendermint show-validator"),t("span",{pre:!0,attrs:{class:"token variable"}},[a._v(")")])]),a._v('"')]),a._v("\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("You should now see your validator in one of the Desmos explorers. You are looking for the "),t("code",[a._v("bech32")]),a._v(" encoded "),t("code",[a._v("address")]),a._v(" in the "),t("code",[a._v("~/.desmosd/config/priv_validator.json")]),a._v(" file.")]),a._v(" "),t("div",{staticClass:"custom-block warning"},[t("p",{staticClass:"custom-block-title"},[a._v("Note")]),a._v(" "),t("p",[a._v("To be in the validator set, you need to have more total voting power than the 100th validator.")])]),a._v(" "),t("h2",{attrs:{id:"halting-your-validator"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#halting-your-validator"}},[a._v("#")]),a._v(" Halting Your Validator")]),a._v(" "),t("p",[a._v("When attempting to perform routine maintenance or planning for an upcoming coordinated\nupgrade, it can be useful to have your validator systematically and gracefully halt.\nYou can achieve this by either setting the "),t("code",[a._v("halt-height")]),a._v(" to the height at which\nyou want your node to shutdown or by passing the "),t("code",[a._v("--halt-height")]),a._v(" flag to "),t("code",[a._v("desmosd")]),a._v(".\nThe node will shutdown with a zero exit code at that given height after committing\nthe block.")]),a._v(" "),t("h2",{attrs:{id:"common-problems"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#common-problems"}},[a._v("#")]),a._v(" Common Problems")]),a._v(" "),t("h3",{attrs:{id:"problem-1-my-validator-has-voting-power-0"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#problem-1-my-validator-has-voting-power-0"}},[a._v("#")]),a._v(" Problem #1: My validator has "),t("code",[a._v("voting_power: 0")])]),a._v(" "),t("p",[a._v("Your validator has become jailed. Validators get jailed, i.e. get removed from the active validator set, if they do not vote on "),t("code",[a._v("500")]),a._v(" of the last "),t("code",[a._v("10000")]),a._v(" blocks, or if they double sign.")]),a._v(" "),t("p",[a._v("If you got jailed for downtime, you can get your voting power back to your validator. First, if "),t("code",[a._v("desmosd")]),a._v(" is not running, start it up again:")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmosd start\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("Wait for your full node to catch up to the latest block. Then, you can "),t("a",{attrs:{href:"#unjail-validator"}},[a._v("unjail your validator")])]),a._v(" "),t("p",[a._v("Lastly, check your validator again to see if your voting power is back.")]),a._v(" "),t("div",{staticClass:"language-bash line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[a._v("desmoscli status\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br")])]),t("p",[a._v("You may notice that your voting power is less than it used to be. That's because you got slashed for downtime!")]),a._v(" "),t("h3",{attrs:{id:"problem-2-my-desmosd-crashes-because-of-too-many-open-files"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#problem-2-my-desmosd-crashes-because-of-too-many-open-files"}},[a._v("#")]),a._v(" Problem #2: My "),t("code",[a._v("desmosd")]),a._v(" crashes because of "),t("code",[a._v("too many open files")])]),a._v(" "),t("p",[a._v("The default number of files Linux can open (per-process) is "),t("code",[a._v("1024")]),a._v(". "),t("code",[a._v("desmosd")]),a._v(" is known to open more than "),t("code",[a._v("1024")]),a._v(" files. This causes the process to crash. A quick fix is to run "),t("code",[a._v("ulimit -n 4096")]),a._v(" (increase the number of open files allowed) and then restart the process with "),t("code",[a._v("desmosd start")]),a._v(". If you are using "),t("code",[a._v("systemd")]),a._v(" or another process manager to launch "),t("code",[a._v("desmosd")]),a._v(" this may require some configuration at that level. A sample "),t("code",[a._v("systemd")]),a._v(" file to fix this issue is below:")]),a._v(" "),t("div",{staticClass:"language-toml line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-toml"}},[t("code",[t("span",{pre:!0,attrs:{class:"token comment"}},[a._v("# /etc/systemd/system/desmosd.service")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),t("span",{pre:!0,attrs:{class:"token table class-name"}},[a._v("Unit")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("Description")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("Desmos Full Node\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("After")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("network"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(".")]),a._v("target\n\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),t("span",{pre:!0,attrs:{class:"token table class-name"}},[a._v("Service")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("Type")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("simple\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("User")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("ubuntu\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("WorkingDirectory")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("/home/ubuntu\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("ExecStart")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("/home/ubuntu/go/bin/desmosd start\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("Restart")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("on-failure\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("RestartSec")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("3")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("LimitNOFILE")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),t("span",{pre:!0,attrs:{class:"token number"}},[a._v("4096")]),a._v("\n\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("[")]),t("span",{pre:!0,attrs:{class:"token table class-name"}},[a._v("Install")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("]")]),a._v("\n"),t("span",{pre:!0,attrs:{class:"token key property"}},[a._v("WantedBy")]),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v("=")]),a._v("multi-user"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[a._v(".")]),a._v("target\n")])]),a._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[a._v("1")]),t("br"),t("span",{staticClass:"line-number"},[a._v("2")]),t("br"),t("span",{staticClass:"line-number"},[a._v("3")]),t("br"),t("span",{staticClass:"line-number"},[a._v("4")]),t("br"),t("span",{staticClass:"line-number"},[a._v("5")]),t("br"),t("span",{staticClass:"line-number"},[a._v("6")]),t("br"),t("span",{staticClass:"line-number"},[a._v("7")]),t("br"),t("span",{staticClass:"line-number"},[a._v("8")]),t("br"),t("span",{staticClass:"line-number"},[a._v("9")]),t("br"),t("span",{staticClass:"line-number"},[a._v("10")]),t("br"),t("span",{staticClass:"line-number"},[a._v("11")]),t("br"),t("span",{staticClass:"line-number"},[a._v("12")]),t("br"),t("span",{staticClass:"line-number"},[a._v("13")]),t("br"),t("span",{staticClass:"line-number"},[a._v("14")]),t("br"),t("span",{staticClass:"line-number"},[a._v("15")]),t("br"),t("span",{staticClass:"line-number"},[a._v("16")]),t("br")])])])}),[],!1,null,null,null);s.default=n.exports}}]);