import{_ as t,u as e,P as s,S as a,Q as n,R as o}from"./index.1709626208081.js";import{d as u,a1 as l,l as r,k as i,aa as d,a as h,Y as m,S as c,ah as g,o as p}from"./vue.1709626208081.js";var f=t(u({name:"limitsFrontEndPage",setup(){const t=e(),u=l({val:"",userAuth:""}),h=r((()=>t.state.userInfos.userInfos.roles)),m=r((()=>t.state.userInfos.userInfos.authBtnList));return i((()=>{u.userAuth=t.state.userInfos.userInfos.roles[0]})),{getRoles:h,getAuthBtnList:m,onRadioChange:async()=>{s();let e=[],l=[],r=["admin"],i=["btn.add","btn.del","btn.edit","btn.link"],d=["common"],h=["btn.add","btn.link"];"admin"===u.userAuth?(e=r,l=i):(e=d,l=h);const m={userName:u.userAuth,photo:"admin"===u.userAuth?"https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=1813762643,1914315241&fm=26&gp=0.jpg":"https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=317673774,2961727727&fm=26&gp=0.jpg",time:(new Date).getTime(),roles:e,authBtnList:l};a.set("userInfo",m),t.dispatch("userInfos/setUserInfos",m),await n(),o()},...d(u)}}}),[["render",function(t,e,s,a,n,o){const u=g("el-alert"),l=g("el-radio-button"),r=g("el-radio-group"),i=g("el-card");return p(),h("div",null,[m(u,{title:"温馨提示：此权限页面代码及效果只作为演示使用，若出现不可逆转的bug，请尝试 `F5` 刷新页面。若实际项目中非要实现此用户权限切换功能，\n      请在切换方法 `onRadioChange` 最后面添加刷新代码 `window.location.reload()`。 请注意：按钮权限页面中的演示2（指令模式）、演示3（函数模式）\n      切换用户时无法动态演示，想要动态演示，请按 `F5` 或者添加 `window.location.reload()`。",type:"warning",closable:!1}),m(u,{title:`当前用户页面权限：[${t.getRoles}]，当前用户按钮权限：[${t.getAuthBtnList}]`,type:"success",closable:!1,"class":"mt15"},null,8,["title"]),m(i,{shadow:"nover",header:"切换用户演示，前端控制不同用户显示不同页面、按钮权限","class":"mt15"},{"default":c((()=>[m(r,{modelValue:t.userAuth,"onUpdate:modelValue":e[0]||(e[0]=e=>t.userAuth=e),onChange:t.onRadioChange},{"default":c((()=>[m(l,{label:"admin"}),m(l,{label:"common"})])),_:1},8,["modelValue","onChange"])])),_:1})])}]]);export{f as default};