(window.webpackJsonp=window.webpackJsonp||[]).push([[21],{223:function(e,t,a){"use strict";a.r(t);var o=a(0),s=Object(o.a)({},(function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"validator-faq"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#validator-faq"}},[e._v("#")]),e._v(" Validator FAQ")]),e._v(" "),a("div",{staticClass:"custom-block warning"},[a("p",{staticClass:"custom-block-title"},[e._v("Disclaimer")]),e._v(" "),a("p",[e._v("This is work in progress. Mechanisms and values are susceptible to change.")])]),e._v(" "),a("h2",{attrs:{id:"general-concepts"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#general-concepts"}},[e._v("#")]),e._v(" General Concepts")]),e._v(" "),a("h3",{attrs:{id:"what-is-a-validator"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-is-a-validator"}},[e._v("#")]),e._v(" What is a validator?")]),e._v(" "),a("p",[a("a",{attrs:{href:"../.."}},[e._v("Desmos")]),e._v(" is based on "),a("a",{attrs:{href:"https://tendermint.com/docs/introduction/what-is-tendermint.html",target:"_blank",rel:"noopener noreferrer"}},[e._v("Tendermint"),a("OutboundLink")],1),e._v(", which relies on a set of validators to secure the network. The role of validators is to run a full-node and participate in consensus by broadcasting votes which contain cryptographic signatures signed by their private key. Validators commit new blocks in the blockchain and receive revenue in exchange for their work. They must also participate in governance by voting on proposals. Validators are weighted according to their total stake.")]),e._v(" "),a("h3",{attrs:{id:"what-is-staking"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-is-staking"}},[e._v("#")]),e._v(" What is 'staking'?")]),e._v(" "),a("p",[e._v("Desmos is a public Proof-Of-Stake (PoS) blockchain, meaning that the weight of validators is determined by the amount of staking tokens (Desmos tokens) bonded as collateral. These tokens can be self-delegated directly by the validator or delegated to them by other Desmos tokens holders.")]),e._v(" "),a("p",[e._v("Any user in the system can declare their intention to become a validator by sending a "),a("code",[e._v("create-validator")]),e._v(" transaction. From there, they become validator candidates.")]),e._v(" "),a("p",[e._v("The weight (i.e. voting power) of a validator determines whether or not they are an active validator. Initially, only the top 100 validators with the most voting power will be active validators.")]),e._v(" "),a("h3",{attrs:{id:"what-is-a-full-node"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-is-a-full-node"}},[e._v("#")]),e._v(" What is a full-node?")]),e._v(" "),a("p",[e._v("A full-node is a program that fully validates transactions and blocks of a blockchain. It is distinct from a light-node that only processes block headers and a small subset of transactions. Running a full-node requires more resources than a light-node but is necessary in order to be a validator. In practice, running a full-node only implies running a non-compromised and up-to-date version of the software with low network latency and without downtime.")]),e._v(" "),a("p",[e._v("Of course, it is possible and encouraged for users to run full-nodes even if they do not plan to be validators.")]),e._v(" "),a("h3",{attrs:{id:"what-is-a-delegator"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-is-a-delegator"}},[e._v("#")]),e._v(" What is a delegator?")]),e._v(" "),a("p",[e._v("Delegators are Desmos tokens holders who cannot, or do not want to run a validator themselves. Token holders can delegate Desmos tokens to a validator and obtain a part of their revenue in exchange (for more detail on how revenue is distributed, see "),a("a",{attrs:{href:"#what-is-the-incentive-to-stake?"}},[a("strong",[e._v("What is the incentive to stake?")])]),e._v(" and "),a("a",{attrs:{href:"#what-are-validators-commission?"}},[a("strong",[e._v("What are validators commission?")])]),e._v(" sections below).")]),e._v(" "),a("p",[e._v("Because they share revenue with their validators, delegators also share risks. Should a validator misbehave, each of their delegators will be partially slashed in proportion to their delegated stake. This is why delegators should perform due diligence on validators before delegating, as well as spreading their stake over multiple validators.")]),e._v(" "),a("p",[e._v("Delegators play a critical role in the system, as they are responsible for choosing validators. Being a delegator is not a passive role: delegators should actively monitor the actions of their validators and participate in governance. For more, read the "),a("a",{attrs:{href:"https://cosmos.network/resources/delegators",target:"_blank",rel:"noopener noreferrer"}},[e._v("delegator's faq"),a("OutboundLink")],1),e._v(".")]),e._v(" "),a("h2",{attrs:{id:"becoming-a-validator"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#becoming-a-validator"}},[e._v("#")]),e._v(" Becoming a Validator")]),e._v(" "),a("h3",{attrs:{id:"how-to-become-a-validator"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-to-become-a-validator"}},[e._v("#")]),e._v(" How to become a validator?")]),e._v(" "),a("p",[e._v("Any participant in the network can signal that they want to become a validator by sending a "),a("code",[e._v("create-validator")]),e._v(" transaction, where they must fill out the following parameters:")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Validator's "),a("code",[e._v("PubKey")]),e._v(":")]),e._v(" The private key associated with this Tendermint "),a("code",[e._v("PubKey")]),e._v(" is used to sign "),a("em",[e._v("prevotes")]),e._v(" and "),a("em",[e._v("precommits")]),e._v(".")]),e._v(" "),a("li",[a("strong",[e._v("Validator's Address:")]),e._v(" Application level address. This is the address used to identify your validator publicly. The private key associated with this address is used to delegate, unbond, claim rewards, and participate in governance.")]),e._v(" "),a("li",[a("strong",[e._v("Validator's name (moniker)")])]),e._v(" "),a("li",[a("strong",[e._v("Validator's website (Optional)")])]),e._v(" "),a("li",[a("strong",[e._v("Validator's description (Optional)")])]),e._v(" "),a("li",[a("strong",[e._v("Initial commission rate")]),e._v(": The commission rate on block rewards and fees charged to delegators.")]),e._v(" "),a("li",[a("strong",[e._v("Maximum commission:")]),e._v(" The maximum commission rate which this validator can charge. This parameter cannot be changed after "),a("code",[e._v("create-validator")]),e._v(" is processed.")]),e._v(" "),a("li",[a("strong",[e._v("Commission max change rate:")]),e._v(" The maximum daily increase of the validator commission. This parameter cannot be changed after "),a("code",[e._v("create-validator")]),e._v(" is processed.")]),e._v(" "),a("li",[a("strong",[e._v("Minimum self-delegation:")]),e._v(" Minimum amount of Desmos tokens the validator needs to have bonded at all time. If the validator's self-delegated stake falls below this limit, their entire staking pool will unbond.")])]),e._v(" "),a("p",[e._v("Once a validator is created, Desmos tokens holders can delegate tokens to them, effectively adding stake to their pool. The total stake of an address is the combination of Desmos tokens bonded by delegators and Desmos tokens self-bonded by the entity which designated themselves.")]),e._v(" "),a("p",[e._v("Out of all validator candidates that signaled themselves, the 100 with the most total stake are the ones who are designated as validators. They become "),a("strong",[e._v("validators")]),e._v(" If a validator's total stake falls below the top 100 then that validator loses their validator privileges: they don't participate in consensus and generate rewards any more. Over time, the maximum number of validators will increase, according to the following schedule ("),a("em",[e._v("Note: this schedule can be changed by governance")]),e._v("):")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Year 0:")]),e._v(" 100")]),e._v(" "),a("li",[a("strong",[e._v("Year 1:")]),e._v(" 113")]),e._v(" "),a("li",[a("strong",[e._v("Year 2:")]),e._v(" 127")]),e._v(" "),a("li",[a("strong",[e._v("Year 3:")]),e._v(" 144")]),e._v(" "),a("li",[a("strong",[e._v("Year 4:")]),e._v(" 163")]),e._v(" "),a("li",[a("strong",[e._v("Year 5:")]),e._v(" 184")]),e._v(" "),a("li",[a("strong",[e._v("Year 6:")]),e._v(" 208")]),e._v(" "),a("li",[a("strong",[e._v("Year 7:")]),e._v(" 235")]),e._v(" "),a("li",[a("strong",[e._v("Year 8:")]),e._v(" 265")]),e._v(" "),a("li",[a("strong",[e._v("Year 9:")]),e._v(" 300")]),e._v(" "),a("li",[a("strong",[e._v("Year 10:")]),e._v(" 300")])]),e._v(" "),a("h2",{attrs:{id:"testnet"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#testnet"}},[e._v("#")]),e._v(" Testnet")]),e._v(" "),a("h3",{attrs:{id:"how-can-i-join-the-testnet"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-can-i-join-the-testnet"}},[e._v("#")]),e._v(" How can I join the testnet?")]),e._v(" "),a("p",[e._v("The Testnet is a great environment to test your validator setup before launch.")]),e._v(" "),a("p",[e._v("We view testnet participation as a great way to signal to the community that you are ready and able to operate a validator. You can find all relevant information about the testnet "),a("router-link",{attrs:{to:"/join-testnet.html"}},[e._v("here")]),e._v(" and "),a("a",{attrs:{href:"https://github.com/desmos-labs/testnets",target:"_blank",rel:"noopener noreferrer"}},[e._v("here"),a("OutboundLink")],1),e._v(".")],1),e._v(" "),a("h3",{attrs:{id:"what-are-the-different-types-of-keys"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-the-different-types-of-keys"}},[e._v("#")]),e._v(" What are the different types of keys?")]),e._v(" "),a("p",[e._v("In short, there are two types of keys:")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Tendermint Key")]),e._v(": This is a unique key used to sign consensus votes.\n"),a("ul",[a("li",[e._v("It is associated with a public key "),a("code",[e._v("desmosvalconspub")]),e._v(" (Get this value with "),a("code",[e._v("desmosd tendermint show-validator")]),e._v(")")]),e._v(" "),a("li",[e._v("It is generated when the node is created with "),a("code",[e._v("desmosd init")]),e._v(".")])])]),e._v(" "),a("li",[a("strong",[e._v("Application key")]),e._v(": This key is created from "),a("code",[e._v("desmoscli")]),e._v(" and used to sign transactions. Application keys are associated with a public key prefixed by "),a("code",[e._v("desmospub")]),e._v(" and an address prefixed by "),a("code",[e._v("desmos")]),e._v(". Both are derived from account keys generated by "),a("code",[e._v("desmoscli keys add")]),e._v(".")])]),e._v(" "),a("p",[e._v("Note: A validator's operator key is directly tied to an application key, but\nuses reserved prefixes solely for this purpose: "),a("code",[e._v("desmosvaloper")]),e._v(" and "),a("code",[e._v("desmosvaloperpub")])]),e._v(" "),a("h3",{attrs:{id:"what-are-the-different-states-a-validator-can-be-in"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-the-different-states-a-validator-can-be-in"}},[e._v("#")]),e._v(" What are the different states a validator can be in?")]),e._v(" "),a("p",[e._v("After a validator is created with a "),a("code",[e._v("create-validator")]),e._v(" transaction, they can be in three states:")]),e._v(" "),a("ul",[a("li",[a("code",[e._v("in validator set")]),e._v(": Validator is in the active set and participates in consensus. Validator is earning rewards and can be slashed for misbehaviour.")]),e._v(" "),a("li",[a("code",[e._v("jailed")]),e._v(": Validator misbehaved and is in jail, i.e. outside of the validator set. If the jailing is due to being offline for too long, the validator can send an "),a("code",[e._v("unjail")]),e._v(" transaction in order to re-enter the validator set. If the jailing is due to double signing, the validator cannot unjail.")]),e._v(" "),a("li",[a("code",[e._v("unbonded")]),e._v(": Validator is not in the active set, and therefore not signing blocs. Validator cannot be slashed, and does not earn any reward. It is still possible to delegate Desmos tokens to this validator. Un-delegating from an "),a("code",[e._v("unbonded")]),e._v(" validator is immediate.")])]),e._v(" "),a("h3",{attrs:{id:"what-is-self-delegation-how-can-i-increase-my-self-delegation"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-is-self-delegation-how-can-i-increase-my-self-delegation"}},[e._v("#")]),e._v(" What is 'self-delegation'? How can I increase my 'self-delegation'?")]),e._v(" "),a("p",[e._v("Self-delegation is delegation from a validator to themselves. This amount can be increases by sending a "),a("code",[e._v("delegate")]),e._v(" transaction from your validator's "),a("code",[e._v("application")]),e._v(" application key.")]),e._v(" "),a("h3",{attrs:{id:"is-there-a-minimum-amount-of-desmos-tokens-that-must-be-delegated-to-be-an-active-bonded-validator"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#is-there-a-minimum-amount-of-desmos-tokens-that-must-be-delegated-to-be-an-active-bonded-validator"}},[e._v("#")]),e._v(" Is there a minimum amount of Desmos tokens that must be delegated to be an active (=bonded) validator?")]),e._v(" "),a("p",[e._v("The minimum is "),a("code",[e._v("1 desmos")]),e._v(".")]),e._v(" "),a("h3",{attrs:{id:"how-will-delegators-choose-their-validators"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-will-delegators-choose-their-validators"}},[e._v("#")]),e._v(" How will delegators choose their validators?")]),e._v(" "),a("p",[e._v("Delegators are free to choose validators according to their own subjective criteria. This said, criteria anticipated to be important include:")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Amount of self-delegated Desmos tokens:")]),e._v(" Number of Desmos tokens a validator self-delegated to themselves. A validator with a higher amount of self-delegated Desmos tokens has more skin in the game, making them more liable for their actions.")]),e._v(" "),a("li",[a("strong",[e._v("Amount of delegated Desmos tokens:")]),e._v(" Total number of Desmos tokens delegated to a validator. A high voting power shows that the community trusts this validator, but it also means that this validator is a bigger target for hackers. Bigger validators also decrease the decentralisation of the network.")]),e._v(" "),a("li",[a("strong",[e._v("Commission rate:")]),e._v(" Commission applied on revenue by validators before it is distributed to their delegators.")]),e._v(" "),a("li",[a("strong",[e._v("Track record:")]),e._v(" Delegators will likely look at the track record of the validators they plan to delegate to. This includes seniority, past votes on proposals, historical average uptime and how often the node was compromised.")])]),e._v(" "),a("p",[e._v("Apart from these criteria, there will be a possibility for validators to signal a website address to complete their resume. Validators will need to build reputation one way or another to attract delegators. For example, it would be a good practice for validators to have their setup audited by third parties. Note though, that the Tendermint team will not approve or conduct any audit themselves. For more on due diligence, see "),a("a",{attrs:{href:"https://medium.com/@interchain_io/3d0faf10ce6f",target:"_blank",rel:"noopener noreferrer"}},[e._v("this blog post"),a("OutboundLink")],1)]),e._v(" "),a("h2",{attrs:{id:"responsibilities"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#responsibilities"}},[e._v("#")]),e._v(" Responsibilities")]),e._v(" "),a("h3",{attrs:{id:"do-validators-need-to-be-publicly-identified"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#do-validators-need-to-be-publicly-identified"}},[e._v("#")]),e._v(" Do validators need to be publicly identified?")]),e._v(" "),a("p",[e._v("No, they do not. Each delegator will value validators based on their own criteria. Validators will be able to register a website address when they nominate themselves so that they can advertise their operation as they see fit. Some delegators may prefer a website that clearly displays the team operating the validator and their resume, while others might prefer anonymous validators with positive track records.")]),e._v(" "),a("h3",{attrs:{id:"what-are-the-responsibilities-of-a-validator"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-the-responsibilities-of-a-validator"}},[e._v("#")]),e._v(" What are the responsibilities of a validator?")]),e._v(" "),a("p",[e._v("Validators have two main responsibilities:")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Be able to constantly run a correct version of the software:")]),e._v(" Validators need to make sure that their servers are always online and their private keys are not compromised.")]),e._v(" "),a("li",[a("strong",[e._v("Actively participate in governance:")]),e._v(" Validators are required to vote on every proposal.")])]),e._v(" "),a("p",[e._v("Additionally, validators are expected to be active members of the community. They should always be up-to-date with the current state of the ecosystem so that they can easily adapt to any change.")]),e._v(" "),a("h3",{attrs:{id:"what-does-participate-in-governance-entail"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-does-participate-in-governance-entail"}},[e._v("#")]),e._v(" What does 'participate in governance' entail?")]),e._v(" "),a("p",[e._v("Validators and delegators on the Cosmos Hub can vote on proposals to change operational parameters (such as the block gas limit), coordinate upgrades, or make a decision on any given matter.")]),e._v(" "),a("p",[e._v("Validators play a special role in the governance system. Being the pillars of the system, they are required to vote on every proposal. It is especially important since delegators who do not vote will inherit the vote of their validator.")]),e._v(" "),a("h3",{attrs:{id:"what-does-staking-imply"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-does-staking-imply"}},[e._v("#")]),e._v(" What does staking imply?")]),e._v(" "),a("p",[e._v("Staking Desmos tokens can be thought of as a safety deposit on validation activities. When a validator or a delegator wants to retrieve part or all of their deposit, they send an "),a("code",[e._v("unbonding")]),e._v(" transaction. Then, Desmos tokens undergo a "),a("strong",[e._v("3 weeks unbonding period")]),e._v(" during which they are liable to being slashed for potential misbehaviour committed by the validator before the unbonding process started.")]),e._v(" "),a("p",[e._v("Validators, and by association delegators, receive block rewards, fees, and have the right to participate in governance. If a validator misbehaves, a certain portion of their total stake is slashed. This means that every delegator that bonded Desmos tokens to this validator gets penalized in proportion to their bonded stake. Delegators are therefore incentivized to delegate to validators that they anticipate will function safely.")]),e._v(" "),a("h3",{attrs:{id:"can-a-validator-run-away-with-their-delegators-desmos-tokens"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#can-a-validator-run-away-with-their-delegators-desmos-tokens"}},[e._v("#")]),e._v(" Can a validator run away with their delegators' Desmos tokens?")]),e._v(" "),a("p",[e._v("By delegating to a validator, a user delegates voting power. The more voting power a validator have, the more weight they have in the consensus and governance processes. This does not mean that the validator has custody of their delegators' Desmos tokens. "),a("strong",[e._v("By no means can a validator run away with its delegator's funds")]),e._v(".")]),e._v(" "),a("p",[e._v("Even though delegated funds cannot be stolen by their validators, delegators are still liable if their validators misbehave.")]),e._v(" "),a("h3",{attrs:{id:"how-often-will-a-validator-be-chosen-to-propose-the-next-block-does-it-go-up-with-the-quantity-of-bonded-desmos-tokens"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-often-will-a-validator-be-chosen-to-propose-the-next-block-does-it-go-up-with-the-quantity-of-bonded-desmos-tokens"}},[e._v("#")]),e._v(" How often will a validator be chosen to propose the next block? Does it go up with the quantity of bonded Desmos tokens?")]),e._v(" "),a("p",[e._v("The validator that is selected to propose the next block is called proposer. Each proposer is selected deterministically, and the frequency of being chosen is proportional to the voting power (i.e. amount of bonded Desmos tokens) of the validator. For example, if the total bonded stake across all validators is 100 Desmos tokens and a validator's total stake is 10 Desmos tokens, then this validator will proposer ~10% of the blocks.")]),e._v(" "),a("h3",{attrs:{id:"will-validators-of-the-cosmos-hub-ever-be-required-to-validate-other-zones-in-the-cosmos-ecosystem"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#will-validators-of-the-cosmos-hub-ever-be-required-to-validate-other-zones-in-the-cosmos-ecosystem"}},[e._v("#")]),e._v(" Will validators of the Cosmos Hub ever be required to validate other zones in the Cosmos ecosystem?")]),e._v(" "),a("p",[e._v("Yes, they will. If governance decides so, validators of the Cosmos hub may be required to validate additional zones in the Cosmos ecosystem.")]),e._v(" "),a("h2",{attrs:{id:"incentives"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#incentives"}},[e._v("#")]),e._v(" Incentives")]),e._v(" "),a("h3",{attrs:{id:"what-is-the-incentive-to-stake"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-is-the-incentive-to-stake"}},[e._v("#")]),e._v(" What is the incentive to stake?")]),e._v(" "),a("p",[e._v("Each member of a validator's staking pool earns different types of revenue:")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Block rewards:")]),e._v(" Native tokens of applications run by validators (e.g. Desmos tokens on the Cosmos Hub) are inflated to produce block provisions. These provisions exist to incentivize Desmos tokens holders to bond their stake, as non-bonded Desmos tokens will be diluted over time.")]),e._v(" "),a("li",[a("strong",[e._v("Transaction fees:")]),e._v(" The Cosmos Hub maintains a whitelist of token that are accepted as fee payment. The initial fee token is the "),a("code",[e._v("Desmos tokens")]),e._v(".")])]),e._v(" "),a("p",[e._v("This total revenue is divided among validators' staking pools according to each validator's weight. Then, within each validator's staking pool the revenue is divided among delegators in proportion to each delegator's stake. A commission on delegators' revenue is applied by the validator before it is distributed.")]),e._v(" "),a("h3",{attrs:{id:"what-is-the-incentive-to-run-a-validator"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-is-the-incentive-to-run-a-validator"}},[e._v("#")]),e._v(" What is the incentive to run a validator ?")]),e._v(" "),a("p",[e._v("Validators earn proportionally more revenue than their delegators because of commissions.")]),e._v(" "),a("p",[e._v("Validators also play a major role in governance. If a delegator does not vote, they inherit the vote from their validator. This gives validators a major responsibility in the ecosystem.")]),e._v(" "),a("h3",{attrs:{id:"what-are-validators-commission"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-validators-commission"}},[e._v("#")]),e._v(" What are validators commission?")]),e._v(" "),a("p",[e._v("Revenue received by a validator's pool is split between the validator and their delegators. The validator can apply a commission on the part of the revenue that goes to their delegators. This commission is set as a percentage. Each validator is free to set their initial commission, maximum daily commission change rate and maximum commission. The Cosmos Hub enforces the parameter that each validator sets. Only the commission rate can change after the validator is created.")]),e._v(" "),a("h3",{attrs:{id:"how-are-block-rewards-distributed"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-are-block-rewards-distributed"}},[e._v("#")]),e._v(" How are block rewards distributed?")]),e._v(" "),a("p",[e._v("Block rewards are distributed proportionally to all validators relative to their voting power. This means that even though each validator gains Desmos tokens with each reward, all validators will maintain equal weight over time.")]),e._v(" "),a("p",[e._v("Let us take an example where we have 10 validators with equal voting power and a commission rate of 1%. Let us also assume that the reward for a block is 1000 Desmos tokens and that each validator has 20% of self-bonded Desmos tokens. These tokens do not go directly to the proposer. Instead, they are evenly spread among validators. So now each validator's pool has 100 Desmos tokens. These 100 Desmos tokens will be distributed according to each participant's stake:")]),e._v(" "),a("ul",[a("li",[e._v("Commission: "),a("code",[e._v("100*80%*1% = 0.8 Desmos tokens")])]),e._v(" "),a("li",[e._v("Validator gets: "),a("code",[e._v("100\\*20% + Commission = 20.8 Desmos tokens")])]),e._v(" "),a("li",[e._v("All delegators get: "),a("code",[e._v("100\\*80% - Commission = 79.2 Desmos tokens")])])]),e._v(" "),a("p",[e._v("Then, each delegator can claim their part of the 79.2 Desmos tokens in proportion to their stake in the validator's staking pool.")]),e._v(" "),a("h3",{attrs:{id:"how-are-fees-distributed"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-are-fees-distributed"}},[e._v("#")]),e._v(" How are fees distributed?")]),e._v(" "),a("p",[e._v("Fees are similarly distributed with the exception that the block proposer can get a bonus on the fees of the block they propose if they include more than the strict minimum of required precommits.")]),e._v(" "),a("p",[e._v("When a validator is selected to propose the next block, they must include at least 2/3 precommits of the previous block. However, there is an incentive to include more than 2/3 precommits in the form of a bonus. The bonus is linear: it ranges from 1% if the proposer includes 2/3rd precommits (minimum for the block to be valid) to 5% if the proposer includes 100% precommits. Of course the proposer should not wait too long or other validators may timeout and move on to the next proposer. As such, validators have to find a balance between wait-time to get the most signatures and risk of losing out on proposing the next block. This mechanism aims to incentivize non-empty block proposals, better networking between validators as well as to mitigate censorship.")]),e._v(" "),a("p",[e._v("Let's take a concrete example to illustrate the aforementioned concept. In this example, there are 10 validators with equal stake. Each of them applies a 1% commission rate and has 20% of self-delegated Desmos tokens. Now comes a successful block that collects a total of 1025.51020408 Desmos tokens in fees.")]),e._v(" "),a("p",[e._v("First, a 2% tax is applied. The corresponding Desmos tokens go to the reserve pool. Reserve pool's funds can be allocated through governance to fund bounties and upgrades.")]),e._v(" "),a("ul",[a("li",[a("code",[e._v("2% * 1025.51020408 = 20.51020408")]),e._v(" Desmos tokens go to the reserve pool.")])]),e._v(" "),a("p",[e._v("1005 Desmos tokens now remain. Let's assume that the proposer included 100% of the signatures in its block. It thus obtains the full bonus of 5%.")]),e._v(" "),a("p",[e._v("We have to solve this simple equation to find the reward R for each validator:")]),e._v(" "),a("p",[a("code",[e._v("9*R + R + R*5% = 1005 ⇔ R = 1005/10.05 = 100")])]),e._v(" "),a("ul",[a("li",[e._v("For the proposer validator:\n"),a("ul",[a("li",[e._v("The pool obtains "),a("code",[e._v("R + R * 5%")]),e._v(": 105 Desmos tokens")]),e._v(" "),a("li",[e._v("Commission: "),a("code",[e._v("105 * 80% * 1%")]),e._v(" = 0.84 Desmos tokens")]),e._v(" "),a("li",[e._v("Validator's reward: "),a("code",[e._v("105 * 20% + Commission")]),e._v(" = 21.84 Desmos tokens")]),e._v(" "),a("li",[e._v("Delegators' rewards: "),a("code",[e._v("105 * 80% - Commission")]),e._v(" = 83.16 Desmos tokens (each delegator will be able to claim its portion of these rewards in proportion to their stake)")])])]),e._v(" "),a("li",[e._v("For each non-proposer validator:\n"),a("ul",[a("li",[e._v("The pool obtains R: 100 Desmos tokens")]),e._v(" "),a("li",[e._v("Commission: "),a("code",[e._v("100 * 80% * 1%")]),e._v(" = 0.8 Desmos tokens")]),e._v(" "),a("li",[e._v("Validator's reward: "),a("code",[e._v("100 * 20% + Commission")]),e._v(" = 20.8 Desmos tokens")]),e._v(" "),a("li",[e._v("Delegators' rewards: "),a("code",[e._v("100 * 80% - Commission")]),e._v(" = 79.2 Desmos tokens (each delegator will be able to claim their portion of these rewards in proportion to their stake)")])])])]),e._v(" "),a("h3",{attrs:{id:"what-are-the-slashing-conditions"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-the-slashing-conditions"}},[e._v("#")]),e._v(" What are the slashing conditions?")]),e._v(" "),a("p",[e._v("If a validator misbehaves, their delegated stake will be partially slashed. There are currently two faults that can result in slashing of funds for a validator and their delegators:")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Double signing:")]),e._v(" If someone reports on chain A that a validator signed two blocks at the same height on chain A and chain B, and if chain A and chain B share a common ancestor, then this validator will get slashed by 5% on chain A.")]),e._v(" "),a("li",[a("strong",[e._v("Downtime:")]),e._v(" If a validator misses more than 95% of the last 10.000 blocks, they will get slashed by 0.01%.")])]),e._v(" "),a("h3",{attrs:{id:"do-validators-need-to-self-delegate-desmos-tokens"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#do-validators-need-to-self-delegate-desmos-tokens"}},[e._v("#")]),e._v(" Do validators need to self-delegate Desmos tokens?")]),e._v(" "),a("p",[e._v("Yes, they do need to self-delegate at least "),a("code",[e._v("1 Desmos tokens")]),e._v(". Even though there is no obligation for validators to self-delegate more than "),a("code",[e._v("1 Desmos tokens")]),e._v(", delegators should want their validator to have more self-delegated Desmos tokens in their staking pool. In other words, validators should have skin in the game.")]),e._v(" "),a("p",[e._v("In order for delegators to have some guarantee about how much skin-in-the-game their validator has, the latter can signal a minimum amount of self-delegated Desmos tokens. If a validator's self-delegation goes below the limit that it predefined, this validator and all of its delegators will unbond.")]),e._v(" "),a("h3",{attrs:{id:"how-to-prevent-concentration-of-stake-in-the-hands-of-a-few-top-validators"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-to-prevent-concentration-of-stake-in-the-hands-of-a-few-top-validators"}},[e._v("#")]),e._v(" How to prevent concentration of stake in the hands of a few top validators?")]),e._v(" "),a("p",[e._v("For now the community is expected to behave in a smart and self-preserving way. When a mining pool in Bitcoin gets too much mining power the community usually stops contributing to that pool. The Cosmos Hub will rely on the same effect initially. Other mechanisms are in place to smoothen this process as much as possible:")]),e._v(" "),a("ul",[a("li",[a("strong",[e._v("Penalty-free re-delegation:")]),e._v(" This is to allow delegators to easily switch from one validator to another, in order to reduce validator stickiness.")]),e._v(" "),a("li",[a("strong",[e._v("UI warning:")]),e._v(" Wallets can implement warnings that will be displayed to users if they want to delegate to a validator that already has a significant amount of staking power.")])]),e._v(" "),a("h2",{attrs:{id:"technical-requirements"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#technical-requirements"}},[e._v("#")]),e._v(" Technical Requirements")]),e._v(" "),a("h3",{attrs:{id:"what-are-hardware-requirements"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-hardware-requirements"}},[e._v("#")]),e._v(" What are hardware requirements?")]),e._v(" "),a("p",[e._v("Validators should expect to provision one or more data center locations with redundant power, networking, firewalls, HSMs and servers.")]),e._v(" "),a("p",[e._v("We expect that a modest level of hardware specifications will be needed initially and that they might rise as network use increases. Participating in the testnet is the best way to learn more.")]),e._v(" "),a("h3",{attrs:{id:"what-are-software-requirements"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-software-requirements"}},[e._v("#")]),e._v(" What are software requirements?")]),e._v(" "),a("p",[e._v("In addition to running a Cosmos Hub node, validators should develop monitoring, alerting and management solutions.")]),e._v(" "),a("h3",{attrs:{id:"what-are-bandwidth-requirements"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-bandwidth-requirements"}},[e._v("#")]),e._v(" What are bandwidth requirements?")]),e._v(" "),a("p",[e._v("The Cosmos network has the capacity for very high throughput relative to chains like Ethereum or Bitcoin.")]),e._v(" "),a("p",[e._v("We recommend that the data center nodes only connect to trusted full-nodes in the cloud or other validators that know each other socially. This relieves the data center node from the burden of mitigating denial-of-service attacks.")]),e._v(" "),a("p",[e._v("Ultimately, as the network becomes more heavily used, multigigabyte per day bandwidth is very realistic.")]),e._v(" "),a("h3",{attrs:{id:"what-does-running-a-validator-imply-in-terms-of-logistics"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-does-running-a-validator-imply-in-terms-of-logistics"}},[e._v("#")]),e._v(" What does running a validator imply in terms of logistics?")]),e._v(" "),a("p",[e._v("A successful validator operation will require the efforts of multiple highly skilled individuals and continuous operational attention. This will be considerably more involved than running a bitcoin miner for instance.")]),e._v(" "),a("h3",{attrs:{id:"how-to-handle-key-management"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-to-handle-key-management"}},[e._v("#")]),e._v(" How to handle key management?")]),e._v(" "),a("p",[e._v("Validators should expect to run an HSM that supports ed25519 keys. Here are potential options:")]),e._v(" "),a("ul",[a("li",[e._v("YubiHSM 2")]),e._v(" "),a("li",[e._v("Ledger Nano S")]),e._v(" "),a("li",[e._v("Ledger BOLOS SGX enclave")]),e._v(" "),a("li",[e._v("Thales nShield support")])]),e._v(" "),a("p",[e._v("The Tendermint team does not recommend one solution above the other. The community is encouraged to bolster the effort to improve HSMs and the security of key management.")]),e._v(" "),a("h3",{attrs:{id:"what-can-validators-expect-in-terms-of-operations"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-can-validators-expect-in-terms-of-operations"}},[e._v("#")]),e._v(" What can validators expect in terms of operations?")]),e._v(" "),a("p",[e._v("Running effective operation is the key to avoiding unexpectedly unbonding or being slashed. This includes being able to respond to attacks, outages, as well as to maintain security and isolation in your data center.")]),e._v(" "),a("h3",{attrs:{id:"what-are-the-maintenance-requirements"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#what-are-the-maintenance-requirements"}},[e._v("#")]),e._v(" What are the maintenance requirements?")]),e._v(" "),a("p",[e._v("Validators should expect to perform regular software updates to accommodate upgrades and bug fixes. There will inevitably be issues with the network early in its bootstrapping phase that will require substantial vigilance.")]),e._v(" "),a("h3",{attrs:{id:"how-can-validators-protect-themselves-from-denial-of-service-attacks"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#how-can-validators-protect-themselves-from-denial-of-service-attacks"}},[e._v("#")]),e._v(" How can validators protect themselves from denial-of-service attacks?")]),e._v(" "),a("p",[e._v("Denial-of-service attacks occur when an attacker sends a flood of internet traffic to an IP address to prevent the server at the IP address from connecting to the internet.")]),e._v(" "),a("p",[e._v("An attacker scans the network, tries to learn the IP address of various validator nodes and disconnect them from communication by flooding them with traffic.")]),e._v(" "),a("p",[e._v("One recommended way to mitigate these risks is for validators to carefully structure their network topology in a so-called sentry node architecture.")]),e._v(" "),a("p",[e._v("Validator nodes should only connect to full-nodes they trust because they operate them themselves or are run by other validators they know socially. A validator node will typically run in a data center. Most data centers provide direct links the networks of major cloud providers. The validator can use those links to connect to sentry nodes in the cloud. This shifts the burden of denial-of-service from the validator's node directly to its sentry nodes, and may require new sentry nodes be spun up or activated to mitigate attacks on existing ones.")]),e._v(" "),a("p",[e._v("Sentry nodes can be quickly spun up or change their IP addresses. Because the links to the sentry nodes are in private IP space, an internet based attacked cannot disturb them directly. This will ensure validator block proposals and votes always make it to the rest of the network.")]),e._v(" "),a("p",[e._v("It is expected that good operating procedures on that part of validators will completely mitigate these threats.")]),e._v(" "),a("p",[e._v("For more on sentry node architecture, see "),a("a",{attrs:{href:"https://forum.cosmos.network/t/sentry-node-architecture-overview/454",target:"_blank",rel:"noopener noreferrer"}},[e._v("this"),a("OutboundLink")],1),e._v(".")])])}),[],!1,null,null,null);t.default=s.exports}}]);