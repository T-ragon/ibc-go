"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[2613],{62502:(e,n,s)=>{s.r(n),s.d(n,{assets:()=>a,contentTitle:()=>t,default:()=>h,frontMatter:()=>o,metadata:()=>c,toc:()=>l});var i=s(85893),r=s(11151);const o={title:"IBC-Go v7 to v8",sidebar_label:"IBC-Go v7 to v8",sidebar_position:11,slug:"/migrations/v7-to-v8"},t="Migrating from v7 to v8",c={id:"migrations/v7-to-v8",title:"IBC-Go v7 to v8",description:"This guide provides instructions for migrating to version v8.0.0 of ibc-go.",source:"@site/versioned_docs/version-v8.2.x/05-migrations/11-v7-to-v8.md",sourceDirName:"05-migrations",slug:"/migrations/v7-to-v8",permalink:"/v8/migrations/v7-to-v8",draft:!1,unlisted:!1,tags:[],version:"v8.2.x",sidebarPosition:11,frontMatter:{title:"IBC-Go v7 to v8",sidebar_label:"IBC-Go v7 to v8",sidebar_position:11,slug:"/migrations/v7-to-v8"},sidebar:"defaultSidebar",previous:{title:"IBC-Go v7.2 to v7.3",permalink:"/v8/migrations/v7_2-to-v7_3"},next:{title:"IBC-Go v8 to v8.1",permalink:"/v8/migrations/v8-to-v8_1"}},a={},l=[{value:"Chains",id:"chains",level:2},{value:"Cosmos SDK v0.50 upgrade",id:"cosmos-sdk-v050-upgrade",level:3},{value:"Authority",id:"authority",level:3},{value:"Testing package",id:"testing-package",level:3},{value:"Params migration",id:"params-migration",level:3},{value:"Governance V1 migration",id:"governance-v1-migration",level:3},{value:"Transfer migration",id:"transfer-migration",level:3},{value:"IBC Apps",id:"ibc-apps",level:2},{value:"ICS20 - Transfer",id:"ics20---transfer",level:3},{value:"ICS27 - Interchain Accounts",id:"ics27---interchain-accounts",level:3},{value:"Relayers",id:"relayers",level:2},{value:"IBC Light Clients",id:"ibc-light-clients",level:2}];function d(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,r.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"migrating-from-v7-to-v8",children:"Migrating from v7 to v8"}),"\n",(0,i.jsxs)(n.p,{children:["This guide provides instructions for migrating to version ",(0,i.jsx)(n.code,{children:"v8.0.0"})," of ibc-go."]}),"\n",(0,i.jsx)(n.p,{children:"There are four sections based on the four potential user groups of this document:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:[(0,i.jsx)(n.a,{href:"#migrating-from-v7-to-v8",children:"Migrating from v7 to v8"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:[(0,i.jsx)(n.a,{href:"#chains",children:"Chains"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#cosmos-sdk-v050-upgrade",children:"Cosmos SDK v0.50 upgrade"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#authority",children:"Authority"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#testing-package",children:"Testing package"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#params-migration",children:"Params migration"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#governance-v1-migration",children:"Governance V1 migration"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#transfer-migration",children:"Transfer migration"})}),"\n"]}),"\n"]}),"\n",(0,i.jsxs)(n.li,{children:[(0,i.jsx)(n.a,{href:"#ibc-apps",children:"IBC Apps"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#ics20---transfer",children:"ICS20 - Transfer"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#ics27---interchain-accounts",children:"ICS27 - Interchain Accounts"})}),"\n"]}),"\n"]}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#relayers",children:"Relayers"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"#ibc-light-clients",children:"IBC Light Clients"})}),"\n"]}),"\n"]}),"\n"]}),"\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.strong,{children:"Note:"})," ibc-go supports golang semantic versioning and therefore all imports must be updated on major version releases."]}),"\n",(0,i.jsx)(n.h2,{id:"chains",children:"Chains"}),"\n",(0,i.jsxs)(n.p,{children:["The type of the ",(0,i.jsx)(n.code,{children:"PortKeeper"})," field of the IBC keeper have been changed to ",(0,i.jsx)(n.code,{children:"*portkeeper.Keeper"}),":"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-diff",children:"// Keeper defines each ICS keeper for IBC\ntype Keeper struct {\n  // implements gRPC QueryServer interface\n  types.QueryServer\n\n  cdc codec.BinaryCodec\n\n  ClientKeeper     clientkeeper.Keeper\n  ConnectionKeeper connectionkeeper.Keeper\n  ChannelKeeper    channelkeeper.Keeper\n- PortKeeper       portkeeper.Keeper\n+ PortKeeper       *portkeeper.Keeper\n  Router           *porttypes.Router\n\n  authority string\n}\n"})}),"\n",(0,i.jsxs)(n.p,{children:["See ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/4703/files#diff-d18972debee5e64f16e40807b2ae112ddbe609504a93ea5e1c80a5d489c3a08a",children:"this PR"})," for the changes required in ",(0,i.jsx)(n.code,{children:"app.go"}),"."]}),"\n",(0,i.jsxs)(n.p,{children:["An extra parameter ",(0,i.jsx)(n.code,{children:"totalEscrowed"})," of type ",(0,i.jsx)(n.code,{children:"sdk.Coins"})," has been added to transfer module's ",(0,i.jsxs)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/transfer/types/genesis.go#L10",children:[(0,i.jsx)(n.code,{children:"NewGenesisState"})," function"]}),". This parameter specifies the total amount of tokens that are in the module's escrow accounts."]}),"\n",(0,i.jsx)(n.h3,{id:"cosmos-sdk-v050-upgrade",children:"Cosmos SDK v0.50 upgrade"}),"\n",(0,i.jsxs)(n.p,{children:["Version ",(0,i.jsx)(n.code,{children:"v8.0.0"})," of ibc-go upgrades to Cosmos SDK v0.50. Please follow the ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/cosmos-sdk/blob/v0.50.1/UPGRADING.md",children:"Cosmos SDK v0.50 upgrading guide"})," to account for its API breaking changes."]}),"\n",(0,i.jsx)(n.h3,{id:"authority",children:"Authority"}),"\n",(0,i.jsxs)(n.p,{children:["An authority identifier (e.g. an address) needs to be passed in the ",(0,i.jsx)(n.code,{children:"NewKeeper"})," functions of the following keepers:"]}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["You must pass the ",(0,i.jsx)(n.code,{children:"authority"})," to the ica/host keeper (implemented in ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3520",children:"#3520"}),"). See ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3520/files#diff-d18972debee5e64f16e40807b2ae112ddbe609504a93ea5e1c80a5d489c3a08a",children:"diff"}),":"]}),"\n"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-diff",children:"// app.go\n\n// ICA Host keeper\napp.ICAHostKeeper = icahostkeeper.NewKeeper(\n  appCodec, keys[icahosttypes.StoreKey], app.GetSubspace(icahosttypes.SubModuleName),\n  app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack\n  app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,\n  app.AccountKeeper, scopedICAHostKeeper, app.MsgServiceRouter(),\n+ authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n)\n"})}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["You must pass the ",(0,i.jsx)(n.code,{children:"authority"})," to the ica/controller keeper (implemented in ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3590",children:"#3590"}),"). See ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3590/files#diff-d18972debee5e64f16e40807b2ae112ddbe609504a93ea5e1c80a5d489c3a08a",children:"diff"}),":"]}),"\n"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-diff",children:"// app.go\n\n// ICA Controller keeper\napp.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(\n  appCodec, keys[icacontrollertypes.StoreKey], app.GetSubspace(icacontrollertypes.SubModuleName),\n  app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack\n  app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,\n  scopedICAControllerKeeper, app.MsgServiceRouter(),\n+ authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n)\n"})}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["You must pass the ",(0,i.jsx)(n.code,{children:"authority"})," to the ibctransfer keeper (implemented in ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3553",children:"#3553"}),"). See ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3553/files#diff-d18972debee5e64f16e40807b2ae112ddbe609504a93ea5e1c80a5d489c3a08a",children:"diff"}),":"]}),"\n"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-diff",children:"// app.go\n\n// Create Transfer Keeper and pass IBCFeeKeeper as expected Channel and PortKeeper\n// since fee middleware will wrap the IBCKeeper for underlying application.\napp.TransferKeeper = ibctransferkeeper.NewKeeper(\n  appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),\n  app.IBCFeeKeeper, // ISC4 Wrapper: fee IBC middleware\n  app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,\n  app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,\n+ authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n)\n"})}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["You should pass the ",(0,i.jsx)(n.code,{children:"authority"})," to the IBC keeper (implemented in ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3640",children:"#3640"})," and ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3650",children:"#3650"}),"). See ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3640/files#diff-d18972debee5e64f16e40807b2ae112ddbe609504a93ea5e1c80a5d489c3a08a",children:"diff"}),":"]}),"\n"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-diff",children:"// app.go\n\n// IBC Keepers\napp.IBCKeeper = ibckeeper.NewKeeper(\n  appCodec, \n  keys[ibcexported.StoreKey],\n  app.GetSubspace(ibcexported.ModuleName),\n  app.StakingKeeper,\n  app.UpgradeKeeper,\n  scopedIBCKeeper,\n+ authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n)\n"})}),"\n",(0,i.jsxs)(n.p,{children:["The authority determines the transaction signer allowed to execute certain messages (e.g. ",(0,i.jsx)(n.code,{children:"MsgUpdateParams"}),")."]}),"\n",(0,i.jsx)(n.h3,{id:"testing-package",children:"Testing package"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["The function ",(0,i.jsx)(n.code,{children:"SetupWithGenesisAccounts"})," has been removed."]}),"\n",(0,i.jsxs)(n.li,{children:["The function ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/testing/path.go#L66",children:(0,i.jsx)(n.code,{children:"RelayPacketWithResults"})})," has been added. This function returns the result of the packet receive transaction, the acknowledgement written on the receiving chain, an error if a relay step fails or the packet commitment does not exist on either chain."]}),"\n"]}),"\n",(0,i.jsx)(n.h3,{id:"params-migration",children:"Params migration"}),"\n",(0,i.jsx)(n.p,{children:"Params are now self managed in the following submodules:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["ica/controller ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3590",children:"#3590"})]}),"\n",(0,i.jsxs)(n.li,{children:["ica/host ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3520",children:"#3520"})]}),"\n",(0,i.jsxs)(n.li,{children:["ibc/connection ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3650",children:"#3650"})]}),"\n",(0,i.jsxs)(n.li,{children:["ibc/client ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3640",children:"#3640"})]}),"\n",(0,i.jsxs)(n.li,{children:["ibc/transfer ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/3553",children:"#3553"})]}),"\n"]}),"\n",(0,i.jsxs)(n.p,{children:["Each module has a corresponding ",(0,i.jsx)(n.code,{children:"MsgUpdateParams"})," message with a ",(0,i.jsx)(n.code,{children:"Params"})," which can be specified in full to update the modules' ",(0,i.jsx)(n.code,{children:"Params"}),"."]}),"\n",(0,i.jsxs)(n.p,{children:["Legacy params subspaces must still be initialised in app.go in order to successfully migrate from `x/params`` to the new self-contained approach. See reference ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/testing/simapp/app.go#L1007-L1012",children:"this"})," for reference."]}),"\n",(0,i.jsxs)(n.p,{children:["For new chains which do not rely on migration of parameters from ",(0,i.jsx)(n.code,{children:"x/params"}),", an expected interface has been added for each module. This allows chain developers to provide ",(0,i.jsx)(n.code,{children:"nil"})," as the ",(0,i.jsx)(n.code,{children:"legacySubspace"})," argument to ",(0,i.jsx)(n.code,{children:"NewKeeper"})," functions."]}),"\n",(0,i.jsx)(n.h3,{id:"governance-v1-migration",children:"Governance V1 migration"}),"\n",(0,i.jsxs)(n.p,{children:["Proposals have been migrated to ",(0,i.jsx)(n.a,{href:"https://docs.cosmos.network/v0.50/modules/gov#messages",children:"gov v1 messages"})," (see ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/4620",children:"#4620"}),"). The proposal ",(0,i.jsx)(n.code,{children:"ClientUpdateProposal"})," has been deprecated and ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/proto/ibc/core/client/v1/tx.proto#L121-L134",children:(0,i.jsx)(n.code,{children:"MsgRecoverClient"})})," should be used instead. Likewise, the proposal ",(0,i.jsx)(n.code,{children:"UpgradeProposal"})," has been deprecated and ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/proto/ibc/core/client/v1/tx.proto#L139-L154",children:(0,i.jsx)(n.code,{children:"MsgIBCSoftwareUpgrade"})})," should be used instead. Both proposals will be removed in the next major release."]}),"\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.code,{children:"MsgRecoverClient"})," and ",(0,i.jsx)(n.code,{children:"MsgIBCSoftwareUpgrade"})," will only be allowed to be executed if the signer is the authority designated at the time of instantiating the IBC keeper. So please make sure that the correct authority is provided to the IBC keeper."]}),"\n",(0,i.jsxs)(n.p,{children:["Remove the ",(0,i.jsx)(n.code,{children:"UpgradeProposalHandler"})," and ",(0,i.jsx)(n.code,{children:"UpdateClientProposalHandler"})," from the ",(0,i.jsx)(n.code,{children:"BasicModuleManager"}),":"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-diff",children:"app.BasicModuleManager = module.NewBasicManagerFromManager(\n  app.ModuleManager,\n  map[string]module.AppModuleBasic{\n    genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),\n    govtypes.ModuleName: gov.NewAppModuleBasic(\n      []govclient.ProposalHandler{\n      paramsclient.ProposalHandler,\n-     ibcclientclient.UpdateClientProposalHandler,\n-     ibcclientclient.UpgradeProposalHandler,\n    },\n  ),\n})\n"})}),"\n",(0,i.jsxs)(n.p,{children:["Support for in-flight legacy recover client proposals (i.e. ",(0,i.jsx)(n.code,{children:"ClientUpdateProposal"}),") will be made for v8, but chains should use ",(0,i.jsx)(n.code,{children:"MsgRecoverClient"})," only afterwards to avoid in-flight client recovery failing when upgrading to v9. See ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/issues/4721",children:"this issue"})," for more information."]}),"\n",(0,i.jsx)(n.p,{children:"Please note that ibc-go offers facilities to test an ibc-go upgrade:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["All e2e tests of the repository can be ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/c5bac5e03a0eae449b9efe0d312258115c1a1e85/e2e/README.md#running-tests-with-custom-images",children:"run with custom Docker chain images"}),"."]}),"\n",(0,i.jsxs)(n.li,{children:["An ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/c5bac5e03a0eae449b9efe0d312258115c1a1e85/e2e/README.md#importable-workflow",children:"importable workflow"})," that can be used from any other repository to test chain upgrades."]}),"\n"]}),"\n",(0,i.jsx)(n.h3,{id:"transfer-migration",children:"Transfer migration"}),"\n",(0,i.jsxs)(n.p,{children:["An ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/transfer/module.go#L136",children:"automatic migration handler"})," is configured in the transfer module to set the ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/cosmos-sdk/blob/v0.50.1/proto/cosmos/bank/v1beta1/bank.proto#L96-L125",children:"denomination metadata"})," for the IBC denominations of all vouchers minted by the transfer module."]}),"\n",(0,i.jsx)(n.h2,{id:"ibc-apps",children:"IBC Apps"}),"\n",(0,i.jsx)(n.h3,{id:"ics20---transfer",children:"ICS20 - Transfer"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["The function ",(0,i.jsx)(n.code,{children:"IsBound"})," has been renamed to ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/transfer/keeper/keeper.go#L98",children:(0,i.jsx)(n.code,{children:"hasCapability"})})," and made unexported."]}),"\n"]}),"\n",(0,i.jsx)(n.h3,{id:"ics27---interchain-accounts",children:"ICS27 - Interchain Accounts"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["Functions ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/27-interchain-accounts/types/codec.go#L32",children:(0,i.jsx)(n.code,{children:"SerializeCosmosTx"})})," and ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/27-interchain-accounts/types/codec.go#L76",children:(0,i.jsx)(n.code,{children:"DeserializeCosmosTx"})})," now accept an extra parameter ",(0,i.jsx)(n.code,{children:"encoding"})," of type ",(0,i.jsx)(n.code,{children:"string"})," that specifies the format in which the transaction messages are marshaled. Both ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/27-interchain-accounts/types/metadata.go#L14-L17",children:"protobuf and proto3 JSON formats"})," are supported."]}),"\n",(0,i.jsxs)(n.li,{children:["The function ",(0,i.jsx)(n.code,{children:"IsBound"})," of controller submodule has been renamed to ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/27-interchain-accounts/controller/keeper/keeper.go#L111",children:(0,i.jsx)(n.code,{children:"hasCapability"})})," and made unexported."]}),"\n",(0,i.jsxs)(n.li,{children:["The function ",(0,i.jsx)(n.code,{children:"IsBound"})," of host submodule has been renamed to ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/blob/v8.0.0/modules/apps/27-interchain-accounts/host/keeper/keeper.go#L94",children:(0,i.jsx)(n.code,{children:"hasCapability"})})," and made unexported."]}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"relayers",children:"Relayers"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["Getter functions in ",(0,i.jsx)(n.code,{children:"MsgChannelOpenInitResponse"}),", ",(0,i.jsx)(n.code,{children:"MsgChannelOpenTryResponse"}),", ",(0,i.jsx)(n.code,{children:"MsgTransferResponse"}),", ",(0,i.jsx)(n.code,{children:"MsgRegisterInterchainAccountResponse"})," and ",(0,i.jsx)(n.code,{children:"MsgSendTxResponse"})," have been removed. The fields can be accessed directly."]}),"\n",(0,i.jsxs)(n.li,{children:[(0,i.jsx)(n.code,{children:"channeltypes.EventTypeTimeoutPacketOnClose"})," (where ",(0,i.jsx)(n.code,{children:"channeltypes"})," is an import alias for ",(0,i.jsx)(n.code,{children:'"github.com/cosmos/ibc-go/v8/modules/core/04-channel"'}),") has been removed, since core IBC does not emit any event with this key."]}),"\n",(0,i.jsxs)(n.li,{children:["Attribute with key ",(0,i.jsx)(n.code,{children:"counterparty_connection_id"})," has been removed from event with key ",(0,i.jsx)(n.code,{children:"connectiontypes.EventTypeConnectionOpenInit"})," (where ",(0,i.jsx)(n.code,{children:"connectiontypes"})," is an import alias for ",(0,i.jsx)(n.code,{children:'"github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"'}),") and attribute with key ",(0,i.jsx)(n.code,{children:"counterparty_channel_id"})," has been removed from event with key ",(0,i.jsx)(n.code,{children:"channeltypes.EventTypeChannelOpenInit"})," (where ",(0,i.jsx)(n.code,{children:"channeltypes"})," is an import alias for ",(0,i.jsx)(n.code,{children:'"github.com/cosmos/ibc-go/v8/modules/core/04-channel"'}),") since both (counterparty connection ID and counterparty channel ID) are empty on ",(0,i.jsx)(n.code,{children:"ConnectionOpenInit"})," and ",(0,i.jsx)(n.code,{children:"ChannelOpenInit"})," respectively."]}),"\n",(0,i.jsxs)(n.li,{children:["As part of the migration to ",(0,i.jsx)(n.a,{href:"#governance-v1-migration",children:"governance V1 messages"})," the following changes in events have been made:"]}),"\n"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-diff",children:'// IBC client events vars\nvar (\n  EventTypeCreateClient          = "create_client"\n  EventTypeUpdateClient          = "update_client"\n  EventTypeUpgradeClient         = "upgrade_client"\n  EventTypeSubmitMisbehaviour    = "client_misbehaviour"\n- EventTypeUpdateClientProposal  = "update_client_proposal"\n- EventTypeUpgradeClientProposal = "upgrade_client_proposal"\n+ EventTypeRecoverClient              = "recover_client"\n+ EventTypeScheduleIBCSoftwareUpgrade = "schedule_ibc_software_upgrade"\n  EventTypeUpgradeChain               = "upgrade_chain"\n)\n'})}),"\n",(0,i.jsx)(n.h2,{id:"ibc-light-clients",children:"IBC Light Clients"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["Functions ",(0,i.jsx)(n.code,{children:"Pretty"})," and ",(0,i.jsx)(n.code,{children:"String"})," of type ",(0,i.jsx)(n.code,{children:"MerklePath"})," have been ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/pull/4459/files#diff-dd94ec1dde9b047c0cdfba204e30dad74a81de202e3b09ac5b42f493153811af",children:"removed"}),"."]}),"\n"]})]})}function h(e={}){const{wrapper:n}={...(0,r.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},11151:(e,n,s)=>{s.d(n,{Z:()=>c,a:()=>t});var i=s(67294);const r={},o=i.createContext(r);function t(e){const n=i.useContext(o);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function c(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:t(e.components),i.createElement(o.Provider,{value:n},e.children)}}}]);