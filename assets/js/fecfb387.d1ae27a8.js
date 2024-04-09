"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[1500],{50232:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>a,default:()=>d,frontMatter:()=>i,metadata:()=>o,toc:()=>l});var r=t(85893),s=t(11151);const i={},a="ADR 025: IBC Passive Channels",o={id:"adr-025-ibc-passive-channels",title:"ADR 025: IBC Passive Channels",description:"Changelog",source:"@site/architecture/adr-025-ibc-passive-channels.md",sourceDirName:".",slug:"/adr-025-ibc-passive-channels",permalink:"/architecture/adr-025-ibc-passive-channels",draft:!1,unlisted:!1,tags:[],version:"current",frontMatter:{},sidebar:"defaultSidebar",previous:{title:"ADR 015: IBC Packet Receiver",permalink:"/architecture/adr-015-ibc-packet-receiver"},next:{title:"ADR 026: IBC Client Recovery Mechanisms",permalink:"/architecture/adr-026-ibc-client-recovery-mechanisms"}},c={},l=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Context",id:"context",level:2},{value:"Handling Channel Open Attempts",id:"handling-channel-open-attempts",level:3},{value:"Decision",id:"decision",level:2},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}];function h(e){const n={a:"a",code:"code",em:"em",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",ul:"ul",...(0,s.a)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(n.h1,{id:"adr-025-ibc-passive-channels",children:"ADR 025: IBC Passive Channels"}),"\n",(0,r.jsx)(n.h2,{id:"changelog",children:"Changelog"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsx)(n.li,{children:'2021-04-23: Change status to "deprecated"'}),"\n",(0,r.jsx)(n.li,{children:"2020-05-23: Provide sample Go code and more details"}),"\n",(0,r.jsx)(n.li,{children:"2020-05-18: Initial Draft"}),"\n"]}),"\n",(0,r.jsx)(n.h2,{id:"status",children:"Status"}),"\n",(0,r.jsx)(n.p,{children:(0,r.jsx)(n.em,{children:"deprecated"})}),"\n",(0,r.jsx)(n.h2,{id:"context",children:"Context"}),"\n",(0,r.jsxs)(n.p,{children:['The current "naive" IBC Relayer strategy currently establishes a single predetermined IBC channel atop a single connection between two clients (each potentially of a different chain).  This strategy then detects packets to be relayed by watching for ',(0,r.jsx)(n.code,{children:"send_packet"})," and ",(0,r.jsx)(n.code,{children:"recv_packet"})," events matching that channel, and sends the necessary transactions to relay those packets."]}),"\n",(0,r.jsx)(n.p,{children:'We wish to expand this "naive" strategy to a "passive" one which detects and relays both channel handshake messages and packets on a given connection, without the need to know each channel in advance of relaying it.'}),"\n",(0,r.jsxs)(n.p,{children:["In order to accomplish this, we propose adding more comprehensive events to expose channel metadata for each transaction sent from the ",(0,r.jsx)(n.code,{children:"x/ibc/core/04-channel/keeper/handshake.go"})," and ",(0,r.jsx)(n.code,{children:"x/ibc/core/04-channel/keeper/packet.go"})," modules."]}),"\n",(0,r.jsxs)(n.p,{children:["Here is an example of what would be in ",(0,r.jsx)(n.code,{children:"ChanOpenInit"}),":"]}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",children:'const (\n  EventTypeChannelMeta = "channel_meta"\n  AttributeKeyAction = "action"\n  AttributeKeyHops = "hops"\n  AttributeKeyOrder = "order"\n  AttributeKeySrcPort = "src_port"\n  AttributeKeySrcChannel = "src_channel"\n  AttributeKeySrcVersion = "src_version"\n  AttributeKeyDstPort = "dst_port"\n  AttributeKeyDstChannel = "dst_channel"\n  AttributeKeyDstVersion = "dst_version"\n)\n// ...\n// Emit Event with Channel metadata for the relayer to pick up and\n// relay to the other chain\n// This appears immediately before the successful return statement.\nctx.EventManager().EmitEvents(sdk.Events{\n  sdk.NewEvent(\n    types.EventTypeChannelMeta,\n    sdk.NewAttribute(types.AttributeKeyAction, "open_init"),\n    sdk.NewAttribute(types.AttributeKeySrcConnection, connectionHops[0]),\n    sdk.NewAttribute(types.AttributeKeyHops, strings.Join(connectionHops, ",")),\n    sdk.NewAttribute(types.AttributeKeyOrder, order.String()),\n    sdk.NewAttribute(types.AttributeKeySrcPort, portID),\n    sdk.NewAttribute(types.AttributeKeySrcChannel, channelID),\n    sdk.NewAttribute(types.AttributeKeySrcVersion, version),\n    sdk.NewAttribute(types.AttributeKeyDstPort, counterparty.GetPortID()),\n    sdk.NewAttribute(types.AttributeKeyDstChannel, counterparty.GetChannelID()),\n    // The destination version is not yet known, but a value is necessary to pad\n    // the event attribute offsets\n    sdk.NewAttribute(types.AttributeKeyDstVersion, ""),\n  ),\n})\n'})}),"\n",(0,r.jsxs)(n.p,{children:['These metadata events capture all the "header" information needed to route IBC channel handshake transactions without requiring the client to query any data except that of the connection ID that it is willing to relay.  It is intended that ',(0,r.jsx)(n.code,{children:"channel_meta.src_connection"})," is the only event key that needs to be indexed for a passive relayer to function."]}),"\n",(0,r.jsx)(n.h3,{id:"handling-channel-open-attempts",children:"Handling Channel Open Attempts"}),"\n",(0,r.jsxs)(n.p,{children:["In the case of the passive relayer, when one chain sends a ",(0,r.jsx)(n.code,{children:"ChanOpenInit"}),', the relayer should inform the other chain of this open attempt and allow that chain to decide how (and if) it continues the handshake.  Once both chains have actively approved the channel opening, then the rest of the handshake can happen as it does with the current "naive" relayer.']}),"\n",(0,r.jsxs)(n.p,{children:["To implement this behavior, we propose replacing the ",(0,r.jsx)(n.code,{children:"cbs.OnChanOpenTry"})," callback with a new ",(0,r.jsx)(n.code,{children:"cbs.OnAttemptChanOpenTry"})," callback which explicitly handles the ",(0,r.jsx)(n.code,{children:"MsgChannelOpenTry"}),", usually by resulting in a call to ",(0,r.jsx)(n.code,{children:"keeper.ChanOpenTry"}),".  The typical implementation, in ",(0,r.jsx)(n.code,{children:"x/ibc-transfer/module.go"}),' would be compatible with the current "naive" relayer, as follows:']}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",children:'func (am AppModule) OnAttemptChanOpenTry(\n  ctx sdk.Context,\n  chanKeeper channel.Keeper,\n  portCap *capability.Capability,\n  msg channel.MsgChannelOpenTry,\n) (*sdk.Result, error) {\n  // Require portID is the portID transfer module is bound to\n  boundPort := am.keeper.GetPort(ctx)\n  if boundPort != msg.PortID {\n    return nil, sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", msg.PortID, boundPort)\n  }\n\n  // BEGIN NEW CODE\n  // Assert our protocol version, overriding the relayer\'s suggestion.\n  msg.Version = types.Version\n  // Continue the ChanOpenTry.\n  res, chanCap, err := channel.HandleMsgChannelOpenTry(ctx, chanKeeper, portCap, msg)\n  if err != nil {\n    return nil, err\n  }\n  // END OF NEW CODE\n\n  // ... the rest of the callback is similar to the existing OnChanOpenTry\n  // but uses msg.* directly.\n'})}),"\n",(0,r.jsxs)(n.p,{children:["Here is how this callback would be used, in the implementation of ",(0,r.jsx)(n.code,{children:"x/ibc/handler.go"}),":"]}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",children:'// ...\ncase channel.MsgChannelOpenTry:\n  // Lookup module by port capability\n  module, portCap, err := k.PortKeeper.LookupModuleByPort(ctx, msg.PortID)\n  if err != nil {\n    return nil, sdkerrors.Wrap(err, "could not retrieve module from port-id")\n  }\n  // Retrieve callbacks from router\n  cbs, ok := k.Router.GetRoute(module)\n  if !ok {\n    return nil, sdkerrors.Wrapf(port.ErrInvalidRoute, "route not found to module: %s", module)\n  }\n  // Delegate to the module\'s OnAttemptChanOpenTry.\n  return cbs.OnAttemptChanOpenTry(ctx, k.ChannelKeeper, portCap, msg)\n'})}),"\n",(0,r.jsxs)(n.p,{children:["The reason we do not have a more structured interaction between ",(0,r.jsx)(n.code,{children:"x/ibc/handler.go"})," and the port's module (to explicitly negotiate versions, etc) is that we do not wish to constrain the app module to have to finish handling the ",(0,r.jsx)(n.code,{children:"MsgChannelOpenTry"})," during this transaction or even this block."]}),"\n",(0,r.jsx)(n.h2,{id:"decision",children:"Decision"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsx)(n.li,{children:'Expose events to allow "passive" connection relayers.'}),"\n",(0,r.jsx)(n.li,{children:"Enable application-initiated channels via such passive relayers."}),"\n",(0,r.jsx)(n.li,{children:"Allow port modules to control how to handle open-try messages."}),"\n"]}),"\n",(0,r.jsx)(n.h2,{id:"consequences",children:"Consequences"}),"\n",(0,r.jsx)(n.h3,{id:"positive",children:"Positive"}),"\n",(0,r.jsx)(n.p,{children:"Makes channels into a complete\xa0application-level abstraction."}),"\n",(0,r.jsx)(n.p,{children:"Applications have full control over initiating and accepting channels, rather than expecting a relayer to tell them when to do so."}),"\n",(0,r.jsx)(n.p,{children:"A passive relayer does not have to know what kind of channel (version string, ordering constraints, firewalling logic) the application supports.  These are negotiated directly between applications."}),"\n",(0,r.jsx)(n.h3,{id:"negative",children:"Negative"}),"\n",(0,r.jsx)(n.p,{children:"Increased event size for IBC messages."}),"\n",(0,r.jsx)(n.h3,{id:"neutral",children:"Neutral"}),"\n",(0,r.jsx)(n.p,{children:"More IBC events are exposed."}),"\n",(0,r.jsx)(n.h2,{id:"references",children:"References"}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:["The Agoric VM's IBC handler currently ",(0,r.jsxs)(n.a,{href:"https://github.com/Agoric/agoric-sdk/blob/904b3a0423222a1b32893453e44bbde598473960/packages/cosmic-swingset/lib/ag-solo/vats/ibc.js#L546",children:["accommodates ",(0,r.jsx)(n.code,{children:"attemptChanOpenTry"})]})]}),"\n"]})]})}function d(e={}){const{wrapper:n}={...(0,s.a)(),...e.components};return n?(0,r.jsx)(n,{...e,children:(0,r.jsx)(h,{...e})}):h(e)}},11151:(e,n,t)=>{t.d(n,{Z:()=>o,a:()=>a});var r=t(67294);const s={},i=r.createContext(s);function a(e){const n=r.useContext(i);return r.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function o(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:a(e.components),r.createElement(i.Provider,{value:n},e.children)}}}]);