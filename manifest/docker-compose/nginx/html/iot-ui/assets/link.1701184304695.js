import{L as f,a2 as p,l as m,i as h,ab as k,a as d,b as g,Z as c,P as R,aD as v,o as M}from"./vue.1701184304695.js";import{_ as w,u as y}from"./index.1701184304695.js";const $=f({name:"layoutLinkView",setup(){const e=v(),a=y(),t=p({currentRouteMeta:{isLink:"",linkUrl:"",title:""}}),r=m(()=>{let{isTagsview:s}=a.state.themeConfig.themeConfig;return s?"114px":"80px"});return h(()=>e.path,()=>{t.currentRouteMeta=e.meta},{immediate:!0}),{setLinkHeight:r,...k(t)}}}),L=["href"];function _(e,a,t,r,s,b){var n,o,i,l,u;return M(),d("div",{class:"layout-view-bg-white flex layout-view-link",style:R({height:`calc(100vh - ${e.setLinkHeight}`})},[g("a",{href:(n=e.currentRouteMeta)==null?void 0:n.linkUrl,target:"_blank",rel:"opener",class:"flex-margin"},c(((o=e.currentRouteMeta)==null?void 0:o.title.indexOf("."))>0?e.$t((i=e.currentRouteMeta)==null?void 0:i.title):(l=e.currentRouteMeta)==null?void 0:l.title)+"\uFF1A"+c((u=e.currentRouteMeta)==null?void 0:u.linkUrl),9,L)],4)}var S=w($,[["render",_]]);export{S as default};