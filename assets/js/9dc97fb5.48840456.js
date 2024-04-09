"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[6715],{90938:(e,i,n)=>{n.r(i),n.d(i,{assets:()=>c,contentTitle:()=>s,default:()=>p,frontMatter:()=>l,metadata:()=>r,toc:()=>o});var a=n(85893),t=n(11151);const l={title:"Capability Module",sidebar_label:"Capability Module",sidebar_position:12,slug:"/ibc/capability-module"},s="Capability Module",r={id:"ibc/capability-module",title:"Capability Module",description:"Overview",source:"@site/docs/01-ibc/12-capability-module.md",sourceDirName:"01-ibc",slug:"/ibc/capability-module",permalink:"/main/ibc/capability-module",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:12,frontMatter:{title:"Capability Module",sidebar_label:"Capability Module",sidebar_position:12,slug:"/ibc/capability-module"},sidebar:"defaultSidebar",previous:{title:"Troubleshooting",permalink:"/main/ibc/troubleshooting"},next:{title:"Overview",permalink:"/main/apps/transfer/overview"}},c={},o=[{value:"Overview",id:"overview",level:2},{value:"Initialization",id:"initialization",level:2},{value:"Contents",id:"contents",level:2},{value:"Concepts",id:"concepts",level:2},{value:"Capabilities",id:"capabilities",level:3},{value:"Stores",id:"stores",level:3},{value:"State",id:"state",level:2},{value:"Persisted KV store",id:"persisted-kv-store",level:3},{value:"In-memory KV store",id:"in-memory-kv-store",level:3}];function d(e){const i={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",ol:"ol",p:"p",pre:"pre",ul:"ul",...(0,t.a)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(i.h1,{id:"capability-module",children:"Capability Module"}),"\n",(0,a.jsx)(i.h2,{id:"overview",children:"Overview"}),"\n",(0,a.jsxs)(i.p,{children:[(0,a.jsx)(i.code,{children:"modules/capability"})," is an implementation of a Cosmos SDK module, per ",(0,a.jsx)(i.a,{href:"https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-003-dynamic-capability-store.md",children:"ADR 003"}),", that allows for provisioning, tracking, and authenticating multi-owner capabilities at runtime."]}),"\n",(0,a.jsx)(i.p,{children:"The keeper maintains two states: persistent and ephemeral in-memory. The persistent\nstore maintains a globally unique auto-incrementing index and a mapping from\ncapability index to a set of capability owners that are defined as a module and\ncapability name tuple. The in-memory ephemeral state keeps track of the actual\ncapabilities, represented as addresses in local memory, with both forward and reverse indexes.\nThe forward index maps module name and capability tuples to the capability name. The\nreverse index maps between the module and capability name and the capability itself."}),"\n",(0,a.jsx)(i.p,{children:'The keeper allows the creation of "scoped" sub-keepers which are tied to a particular\nmodule by name. Scoped keepers must be created at application initialization and\npassed to modules, which can then use them to claim capabilities they receive and\nretrieve capabilities which they own by name, in addition to creating new capabilities\n& authenticating capabilities passed by other modules. A scoped keeper cannot escape its scope,\nso a module cannot interfere with or inspect capabilities owned by other modules.'}),"\n",(0,a.jsx)(i.p,{children:"The keeper provides no other core functionality that can be found in other modules\nlike queriers, REST and CLI handlers, and genesis state."}),"\n",(0,a.jsx)(i.h2,{id:"initialization",children:"Initialization"}),"\n",(0,a.jsx)(i.p,{children:"During application initialization, the keeper must be instantiated with a persistent\nstore key and an in-memory store key."}),"\n",(0,a.jsx)(i.pre,{children:(0,a.jsx)(i.code,{className:"language-go",children:"type App struct {\n  // ...\n\n  capabilityKeeper *capability.Keeper\n}\n\nfunc NewApp(...) *App {\n  // ...\n\n  app.capabilityKeeper = capabilitykeeper.NewKeeper(codec, persistentStoreKey, memStoreKey)\n}\n"})}),"\n",(0,a.jsx)(i.p,{children:"After the keeper is created, it can be used to create scoped sub-keepers which\nare passed to other modules that can create, authenticate, and claim capabilities.\nAfter all the necessary scoped keepers are created and the state is loaded, the\nmain capability keeper must be sealed to prevent further scoped keepers from\nbeing created."}),"\n",(0,a.jsx)(i.pre,{children:(0,a.jsx)(i.code,{className:"language-go",children:"func NewApp(...) *App {\n  // ...\n\n  // Creating a scoped keeper\n  scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)\n\n  // Seal the capability keeper to prevent any further modules from creating scoped\n  // sub-keepers.\n  app.capabilityKeeper.Seal()\n\n  return app\n}\n"})}),"\n",(0,a.jsx)(i.h2,{id:"contents",children:"Contents"}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsxs)(i.li,{children:[(0,a.jsx)(i.a,{href:"#capability-module",children:(0,a.jsx)(i.code,{children:"modules/capability"})}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsx)(i.li,{children:(0,a.jsx)(i.a,{href:"#overview",children:"Overview"})}),"\n",(0,a.jsx)(i.li,{children:(0,a.jsx)(i.a,{href:"#initialization",children:"Initialization"})}),"\n",(0,a.jsx)(i.li,{children:(0,a.jsx)(i.a,{href:"#contents",children:"Contents"})}),"\n",(0,a.jsxs)(i.li,{children:[(0,a.jsx)(i.a,{href:"#concepts",children:"Concepts"}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsx)(i.li,{children:(0,a.jsx)(i.a,{href:"#capabilities",children:"Capabilities"})}),"\n",(0,a.jsx)(i.li,{children:(0,a.jsx)(i.a,{href:"#stores",children:"Stores"})}),"\n"]}),"\n"]}),"\n",(0,a.jsxs)(i.li,{children:[(0,a.jsx)(i.a,{href:"#state",children:"State"}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsx)(i.li,{children:(0,a.jsx)(i.a,{href:"#persisted-kv-store",children:"Persisted KV store"})}),"\n",(0,a.jsx)(i.li,{children:(0,a.jsx)(i.a,{href:"#in-memory-kv-store",children:"In-memory KV store"})}),"\n"]}),"\n"]}),"\n"]}),"\n"]}),"\n"]}),"\n",(0,a.jsx)(i.h2,{id:"concepts",children:"Concepts"}),"\n",(0,a.jsx)(i.h3,{id:"capabilities",children:"Capabilities"}),"\n",(0,a.jsxs)(i.p,{children:["Capabilities are multi-owner. A scoped keeper can create a capability via ",(0,a.jsx)(i.code,{children:"NewCapability"}),"\nwhich creates a new unique, unforgeable object-capability reference. The newly\ncreated capability is automatically persisted; the calling module need not call\n",(0,a.jsx)(i.code,{children:"ClaimCapability"}),". Calling ",(0,a.jsx)(i.code,{children:"NewCapability"})," will create the capability with the\ncalling module and name as a tuple to be treated the capabilities first owner."]}),"\n",(0,a.jsxs)(i.p,{children:["Capabilities can be claimed by other modules which add them as owners. ",(0,a.jsx)(i.code,{children:"ClaimCapability"}),"\nallows a module to claim a capability key which it has received from another\nmodule so that future ",(0,a.jsx)(i.code,{children:"GetCapability"})," calls will succeed. ",(0,a.jsx)(i.code,{children:"ClaimCapability"})," MUST\nbe called if a module which receives a capability wishes to access it by name in\nthe future. Again, capabilities are multi-owner, so if multiple modules have a\nsingle Capability reference, they will all own it. If a module receives a capability\nfrom another module but does not call ",(0,a.jsx)(i.code,{children:"ClaimCapability"}),", it may use it in the executing\ntransaction but will not be able to access it afterwards."]}),"\n",(0,a.jsxs)(i.p,{children:[(0,a.jsx)(i.code,{children:"AuthenticateCapability"})," can be called by any module to check that a capability\ndoes in fact correspond to a particular name (the name can be un-trusted user input)\nwith which the calling module previously associated it."]}),"\n",(0,a.jsxs)(i.p,{children:[(0,a.jsx)(i.code,{children:"GetCapability"})," allows a module to fetch a capability which it has previously\nclaimed by name. The module is not allowed to retrieve capabilities which it does\nnot own."]}),"\n",(0,a.jsx)(i.h3,{id:"stores",children:"Stores"}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsx)(i.li,{children:"MemStore"}),"\n",(0,a.jsx)(i.li,{children:"KeyStore"}),"\n"]}),"\n",(0,a.jsx)(i.h2,{id:"state",children:"State"}),"\n",(0,a.jsx)(i.h3,{id:"persisted-kv-store",children:"Persisted KV store"}),"\n",(0,a.jsxs)(i.ol,{children:["\n",(0,a.jsx)(i.li,{children:"Global unique capability index"}),"\n",(0,a.jsx)(i.li,{children:"Capability owners"}),"\n"]}),"\n",(0,a.jsx)(i.p,{children:"Indexes:"}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsxs)(i.li,{children:["Unique index: ",(0,a.jsx)(i.code,{children:'[]byte("index") -> []byte(currentGlobalIndex)'})]}),"\n",(0,a.jsxs)(i.li,{children:["Capability Index: ",(0,a.jsx)(i.code,{children:'[]byte("capability_index") | []byte(index) -> ProtocolBuffer(CapabilityOwners)'})]}),"\n"]}),"\n",(0,a.jsx)(i.h3,{id:"in-memory-kv-store",children:"In-memory KV store"}),"\n",(0,a.jsxs)(i.ol,{children:["\n",(0,a.jsx)(i.li,{children:"Initialized flag"}),"\n",(0,a.jsx)(i.li,{children:"Mapping between the module and capability tuple and the capability name"}),"\n",(0,a.jsx)(i.li,{children:"Mapping between the module and capability name and its index"}),"\n"]}),"\n",(0,a.jsx)(i.p,{children:"Indexes:"}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsxs)(i.li,{children:["Initialized flag: ",(0,a.jsx)(i.code,{children:'[]byte("mem_initialized")'})]}),"\n",(0,a.jsxs)(i.li,{children:["RevCapabilityKey: ",(0,a.jsx)(i.code,{children:'[]byte(moduleName + "/rev/" + capabilityName) -> []byte(index)'})]}),"\n",(0,a.jsxs)(i.li,{children:["FwdCapabilityKey: ",(0,a.jsx)(i.code,{children:'[]byte(moduleName + "/fwd/" + capabilityPointerAddress) -> []byte(capabilityName)'})]}),"\n"]})]})}function p(e={}){const{wrapper:i}={...(0,t.a)(),...e.components};return i?(0,a.jsx)(i,{...e,children:(0,a.jsx)(d,{...e})}):d(e)}},11151:(e,i,n)=>{n.d(i,{Z:()=>r,a:()=>s});var a=n(67294);const t={},l=a.createContext(t);function s(e){const i=a.useContext(l);return a.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function r(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:s(e.components),a.createElement(l.Provider,{value:i},e.children)}}}]);