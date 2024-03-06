import{d as e,a1 as l,h as a,k as t,aa as o,R as r,S as u,j as d,aC as s,aE as m,ah as n,o as i,Y as p,b as f,a as _,a7 as b,F as c,T as y,Z as V,W as v,V as h}from"./vue.1709626208081.js";import{_ as g}from"./index.170962620808127.js";import{a as w}from"./index.170962620808130.js";import{a as x}from"./index.170962620808122.js";import{_ as k,E as U}from"./index.1709626208081.js";const N=e({name:"tunnelCreate",components:{codeEditor:g},props:{type:{type:String,"default":""}},setup(e){const{proxy:r}=d(),u=s(),n=m(),{network_tunnel_type:i,tunnel_serial_baudrate:p,tunnel_serial_databits:f,tunnel_serial_stopbits:_,tunnel_serial_parity:b,network_protocols:c}=r.useDict("network_tunnel_type","tunnel_serial_baudrate","tunnel_serial_databits","tunnel_serial_stopbits","tunnel_serial_parity","network_protocols"),y=l({resourceModalPro:{mode:"",content:""},detail:{},activeViewName:["1","2","3","4","5"],messageData:[],form:{id:"",name:"新建通道",types:"serial",status:!1,addr:"",serial:{baud_rate:"9600",data_bits:"6",stop_bits:"1",parity:"0"},retry:{enable:!0,timeout:30,maximum:0},protoccol:{name:"SagooMqtt",options:{}},heartbeat:{enable:!1,hex:"",regex:"^\\w+$",text:"",timeout:30}}});x.product.getTypesAll({types:"protocol"}).then((e=>{y.messageData=e||[]}));const V=a("first"),v=a("mirrorRef"),h=()=>{const e=u.params&&u.params.id;w.tunnel.getDetail({id:e}).then((e=>{const{id:l,name:a,types:t,status:o,addr:r,serial:u,retry:d,protoccol:s,heartbeat:m}=e;if(y.form.name=a,y.form.types=t,y.form.addr=r,y.form.status=o,y.form.serial=JSON.parse(u||"{}"),y.form.retry=JSON.parse(d||"{}"),y.form.heartbeat=JSON.parse(m||"{}"),y.form.protoccol=s?JSON.parse(s):{name:"Modbus RTU",options:{}},y.form.id=l,s){let e=JSON.stringify(JSON.parse(s).options);y.resourceModalPro.content=JSON.stringify(JSON.parse(e),null,4),v.value.setValue(y.resourceModalPro.content)}}))};return t((()=>{h()})),{mirrorRef:v,activeName:V,getDetail:h,network_tunnel_type:i,tunnel_serial_baudrate:p,tunnel_serial_databits:f,tunnel_serial_stopbits:_,tunnel_serial_parity:b,network_protocols:c,submit:()=>{0==y.form.serial.parity&&(y.form.serial.rs485=!1,delete y.form.serial.port),1!=y.form.serial.parity&&2!=y.form.serial.parity||(y.form.serial.port=null,delete y.form.serial.rs485),w.tunnel.editItem({...y.form}).then((e=>{U.success("修改成功"),n.go(-1)}))},...o(e),...o(y)}}}),S={"class":"collapse-wrap"},C={"class":"collapse-wrap"},M={style:{position:"absolute",right:"20px",top:"20px"}};var J=k(N,[["render",function(e,l,a,t,o,d){const s=n("el-input"),m=n("el-form-item"),g=n("el-option"),w=n("el-select"),x=n("el-switch"),k=n("el-form"),U=n("el-collapse-item"),N=n("el-button"),J=n("el-input-number"),O=n("codeEditor"),j=n("el-collapse"),D=n("el-tab-pane"),P=n("el-tabs"),R=n("el-card");return i(),r(R,{"class":"system-dic-container",style:{position:"relative"}},{"default":u((()=>[p(P,{modelValue:e.activeName,"onUpdate:modelValue":l[14]||(l[14]=l=>e.activeName=l)},{"default":u((()=>[p(D,{label:"编辑通道",name:"first"},{"default":u((()=>[p(j,{modelValue:e.activeViewName,"onUpdate:modelValue":l[13]||(l[13]=l=>e.activeViewName=l)},{"default":u((()=>[p(U,{title:"基本信息",name:"1"},{"default":u((()=>[f("div",S,[p(k,{style:{width:"600px",margin:"0 auto"},model:e.form,"label-width":"68px"},{"default":u((()=>[p(m,{label:"名称"},{"default":u((()=>[p(s,{modelValue:e.form.name,"onUpdate:modelValue":l[0]||(l[0]=l=>e.form.name=l),"show-word-limit":"",maxlength:"20",placeholder:"请填写名称"},null,8,["modelValue"])])),_:1}),p(m,{label:"类型"},{"default":u((()=>[p(w,{modelValue:e.form.types,"onUpdate:modelValue":l[1]||(l[1]=l=>e.form.types=l),placeholder:"请选择类型"},{"default":u((()=>[(i(!0),_(c,null,b(e.network_tunnel_type,(e=>(i(),r(g,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),y(p(m,{label:"地址"},{"default":u((()=>[p(s,{modelValue:e.form.addr,"onUpdate:modelValue":l[2]||(l[2]=l=>e.form.addr=l),placeholder:"端口号，IP:端口"},null,8,["modelValue"])])),_:1},512),[[V,"serial"!=e.form.types]]),p(m,{label:"启用"},{"default":u((()=>[p(x,{modelValue:e.form.status,"onUpdate:modelValue":l[3]||(l[3]=l=>e.form.status=l),"active-value":1,"inactive-value":0},null,8,["modelValue"])])),_:1})])),_:1},8,["model"])])])),_:1}),y(p(U,{title:"串口参数",name:"2"},{"default":u((()=>[f("div",C,[p(k,{style:{width:"600px",margin:"0 auto"},model:e.form,"label-width":"68px"},{"default":u((()=>[p(m,{label:"端口"},{"default":u((()=>[p(N,null,{"default":u((()=>[v("/dev/ttyS0")])),_:1})])),_:1}),p(m,{label:"波特率"},{"default":u((()=>[p(w,{modelValue:e.form.serial.baud_rate,"onUpdate:modelValue":l[4]||(l[4]=l=>e.form.serial.baud_rate=l),placeholder:"请选择波特率"},{"default":u((()=>[(i(!0),_(c,null,b(e.tunnel_serial_baudrate,(e=>(i(),r(g,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),p(m,{label:"数据位"},{"default":u((()=>[p(w,{modelValue:e.form.serial.data_bits,"onUpdate:modelValue":l[5]||(l[5]=l=>e.form.serial.data_bits=l),placeholder:"请选择数据位"},{"default":u((()=>[(i(!0),_(c,null,b(e.tunnel_serial_databits,(e=>(i(),r(g,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),p(m,{label:"停止位"},{"default":u((()=>[p(w,{modelValue:e.form.serial.stop_bits,"onUpdate:modelValue":l[6]||(l[6]=l=>e.form.serial.stop_bits=l),placeholder:"请选择停止位"},{"default":u((()=>[(i(!0),_(c,null,b(e.tunnel_serial_stopbits,(e=>(i(),r(g,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1}),p(m,{label:"检验位"},{"default":u((()=>[p(w,{modelValue:e.form.serial.parity,"onUpdate:modelValue":l[7]||(l[7]=l=>e.form.serial.parity=l),placeholder:"请选择检验位"},{"default":u((()=>[(i(!0),_(c,null,b(e.tunnel_serial_parity,(e=>(i(),r(g,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["model"])])])),_:1},512),[[V,"serial"==e.form.types]]),y(p(U,{title:"心跳包",name:"4"},{"default":u((()=>[p(k,{style:{width:"600px",margin:"0 auto"},model:e.form,"label-width":"68px"},{"default":u((()=>[p(m,{label:"启用"},{"default":u((()=>[p(x,{modelValue:e.form.heartbeat.enable,"onUpdate:modelValue":l[8]||(l[8]=l=>e.form.heartbeat.enable=l)},null,8,["modelValue"])])),_:1})])),_:1},8,["model"])])),_:1},512),[[V,"serial"!=e.form.type]]),"serial"==e.form.types||"tcp-client"==e.form.types||"udp-client"==e.form.types?(i(),r(U,{key:0,title:"断线重连",name:"4"},{"default":u((()=>[p(k,{style:{width:"600px",margin:"0 auto"},model:e.form,"label-width":"68px"},{"default":u((()=>[p(m,{label:"启用"},{"default":u((()=>[p(x,{modelValue:e.form.retry.enable,"onUpdate:modelValue":l[9]||(l[9]=l=>e.form.retry.enable=l)},null,8,["modelValue"])])),_:1}),p(m,{label:"间隔"},{"default":u((()=>[p(J,{modelValue:e.form.retry.timeout,"onUpdate:modelValue":l[10]||(l[10]=l=>e.form.retry.timeout=l),min:0,onChange:e.handleChange},null,8,["modelValue","onChange"])])),_:1}),p(m,{label:"最大次数"},{"default":u((()=>[p(J,{modelValue:e.form.retry.maximum,"onUpdate:modelValue":l[11]||(l[11]=l=>e.form.retry.maximum=l),min:0,onChange:e.handleChange},null,8,["modelValue","onChange"])])),_:1})])),_:1},8,["model"])])),_:1})):h("",!0),p(U,{title:"协议适配",name:"5"},{"default":u((()=>[p(k,{style:{width:"600px",margin:"0 auto"},model:e.form,"label-width":"68px"},{"default":u((()=>[p(m,{label:"协议"},{"default":u((()=>[p(w,{modelValue:e.form.protoccol.name,"onUpdate:modelValue":l[12]||(l[12]=l=>e.form.protoccol.name=l),placeholder:"请选择协议适配"},{"default":u((()=>[(i(!0),_(c,null,b(e.messageData,(e=>(i(),r(g,{key:e.types,label:e.title,value:e.name},null,8,["label","value"])))),128)),p(g,{label:"Sagoo Mqtt",value:"SagooMqtt"})])),_:1},8,["modelValue"])])),_:1}),p(m,{label:"协议参数"},{"default":u((()=>[p(O,{"class":"params",ref:"mirrorRef",mode:e.resourceModalPro.mode,content:e.resourceModalPro.content},null,8,["mode","content"])])),_:1})])),_:1},8,["model"])])),_:1})])),_:1},8,["modelValue"])])),_:1})])),_:1},8,["modelValue"]),f("div",M,[p(N,{size:"medium",onClick:l[15]||(l[15]=l=>e.$router.replace("/iotmanager/network/tunnel"))},{"default":u((()=>[v("取消")])),_:1}),p(N,{onClick:e.submit,size:"medium",type:"primary"},{"default":u((()=>[v("提交")])),_:1},8,["onClick"])])])),_:1})}],["__scopeId","data-v-a23cc652"]]);export{J as default};