<script setup>
import {reactive} from 'vue'
import {MinWindow,SetH,KillSelf,Set,GetMsgs,GetLen,Del,Add,RedomSay} from '../../wailsjs/go/main/App'
import {SwitchButton,Minus,Edit,Delete,Check,Close,Refresh} from '@element-plus/icons-vue'

const todos = reactive({
  newTodo: "",
  len: 0,
  msgs: [],
  edit: -1,
  edTodo:"",
  height: 600,
  say:"",
})
getMsgs()
getLen()
redomSay()
window.onload = initH

function initH(){
  setTimeout(reSetH,100)
  // console.log(todos.height)
}

function redomSay(){
  RedomSay().then( result => {
    todos.say = result
  })
}

function add() {
  Add(todos.newTodo).then(() => {
    getLen()
    getMsgs()
    todos.newTodo = ""
  })
  // setTimeout 在reSetH之前等待dom更新
  // 不知道有没有更好的办法
  setTimeout(reSetH,100)
}

function del(i) {
  Del(i).then(() =>{
    getLen()
    getMsgs()
  })
  setTimeout(reSetH,100)
}

function set(i,todo){
  Set(i,todo).then(() =>{
    getLen()
    getMsgs()
    todos.edit = -1
  })
}

function getMsgs(){
  GetMsgs().then(result =>{
    todos.msgs = result
  }
  )
}

function getLen(){
  GetLen().then( result => {
    todos.len = result
  })
}

function reSetH(){
  // setTimeout(1000)
  todos.height = document.documentElement.getElementsByTagName("main")[0].offsetHeight+200
  // alert(todos.height)
  SetH(todos.height).then(() =>{})
}
</script>

<template>
  <!-- <main> -->
    <!-- <a @click="KillSelf()"><Close/></a> -->
    <el-header id="todoHeader" style="height: 100px;">
      <el-icon class="header-icon" @click="MinWindow()" :size="20"><Minus/></el-icon>
      <el-icon class="header-icon" @click="KillSelf()" :size="20"><Close/></el-icon>
      <!-- <a @click="MinWindow()" class="Document"><Minus/></a> -->
      <!-- <button @click="reSetH()">setH</button> -->
      <!-- <button @click="setH()">test</button> -->
    <!-- <p>"h:"{{ todos.height }}</p> -->
    <!-- <h2 >todos</h2> -->
      <div>
        <input v-model="todos.newTodo" @keyup.enter = "add()" placeholder = "做点什么..."/>
        <button @click="add()" :size="10">add</button>
      </div>
      <div>
      待做: {{ todos.len }}
      </div>
    </el-header>
   

    <el-main style="background-color:rgba(128,110,133,0.5);text-align:center;width: 400px;">
    <div id="lists">
      <!-- <ul> -->
      <el-row :gutter="30" class = "todo-li" type="flex" justify="end" v-for="(i,item) of todos.msgs" v-bind:key="item">
      <!--<nobr>-->
        <el-col :span="19">
            <p v-if="todos.edit != item">
              <span @dblclick="todos.edit = item;todos.edTodo=todos.msgs[item]" style="color: aliceblue;">{{ i }}</span>
            </p>
            <p v-else>
              <el-input label-width="10px" type="text" @keyup.enter = "set(item,todos.edTodo)" v-model="todos.edTodo"/>
            </p>
          </el-col>
          <el-col :span="5">
            <p v-if="todos.edit != item">
              <el-icon class="edit-button" @click = "todos.edit = item;todos.edTodo=todos.msgs[item] " :size="20"><Edit/></el-icon>
              <el-icon id="finished-button" @click = "del(item)" :size="20"><Delete/></el-icon>
              <!-- <el-icon id="delete-button" @click = "del(item)" :size="20"><Delete/></el-icon> -->
            </p>
            <p v-else>
              <el-icon slot="reference" class="edit-button" @click = "set(item,todos.edTodo)" :size="20"><Check/></el-icon>
              <el-icon id="edit-quxiao-button" @click="todos.edit = -1" :size="20"><Close/></el-icon>
            </p>
          </el-col>
      <!--</nobr>-->
        </el-row>
      <!-- </ul> -->
    </div>
  </el-main>
    
  <el-footer id="footer">
    <!-- <img id="logo" alt="Wails logo" src="../assets/images/logo-universal.png"/> -->
    <!-- <p>如我所愿,一切皆好</p> -->
    <p id="say">{{ todos.say }}</p>
    <el-icon id="Refresh" @click="redomSay()"><Refresh/></el-icon>
  </el-footer>
  <!-- </main> -->
</template>

<style scoped>
/*
body{
  background-image:url('gradient2.png');
}
*/
#footer{
  color: #CB8986;
}

#footer:hover #Refresh{
  opacity: 0.8;
}

#footer:hover #say{
  opacity: 1;
}

#Refresh{
  opacity: 0.3;
}

#say{
  opacity: 0.7;
}

#todo{
  color: #ccc;
}

.el-row {
  margin-bottom: 0px;
}
.el-col {
  border-radius: 10px;
}

#todoHeader:hover .header-icon{
  opacity: 1;
}

.header-icon{
  opacity:0.3;
}
.header-icon:hover{
  color: rgb(214, 44, 44);
  vertical-align: right;
}

#lists{
  /* background-color:rgba(255,255,255,0.6); */
  /* border: 2px solid rgba(0,0,0,0.6); */
  border-radius: 5px;
  /* margin: auto; */
  text-align: left;
  max-height: 500px;
  min-height: 300px;
  overflow-y: hidden;
  overflow-y: scroll;
  /* opacity: 0.3; */
  /* width: 10; */
}

#lists::-webkit-scrollbar {
    height: 0;
    width: 0;
}

#edit-input{
  /* width: 10; */
  border: transparent;
}

span {
  position: relative;
  /* left: 20px; */
  word-wrap: break-word;
  display: inline-block;
  color:black;
  border-bottom:2px solid grey;
  width: 240px;
}

#edit-quxiao-button{
  position: absolute;
  bottom: 15px;
  right: 24px;
  background-color: rgba(255, 255, 255, 0.603);
  border: none;
  /* color: white; */
  color: black;
  border: 2px solid #484848;
  padding: 5px 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 3px;
  border-radius: 20%;
}

#edit-quxiao-button:hover {
    box-shadow: 0px 0px 0px 2.5px rgba(255,255,255,0.2);
}

.edit-button{
  position: absolute;
  bottom: 15px;
  right: 60px;
  background-color: rgba(255, 255, 255, 0.603);
  border: none;
  /* color: white; */
  color: black;
  border: 2px solid #484848;
  padding: 5px 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 3px;
  border-radius: 20%;
  /* text-align:; */
}

.edit-button:hover {
    box-shadow: 0px 0px 0px 2.5px rgba(255,255,255,0.2);
}

/*
#add-button{
  background-color: rgba(255, 255, 255);
  border: none;
  color: black;
  padding: 5px 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 3px;
  border-radius: 20%;
}
#add-button:hover {
    box-shadow: 0px 0px 0px 2.5px rgba(255,255,255,0.2);
}
*/


#delete-button{
  position: absolute;
  bottom: 11.5px;
  right: 8px;
  /* background-color: #fff; */
  border: none;
  color: rgb(128,128,128);
  /* color: black; */
  /* border: px solid #484848; */
  padding: 5px 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 3px;
  border-radius: 20%;
}

#delete-button:hover {
    /* box-shadow: 0px 0px 0px 2.5px rgba(255,255,255,0.2); */
    color: rgb(252, 99, 99);
}

#finished-button{
  position: absolute;
  bottom: 15px;
  right: 24px;
  background-color: rgba(255, 255, 255, 0.603);
  border: none;
  /* color: white; */
  color: black;
  border: 2px solid #484848;
  padding: 5px 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 3px;
  border-radius: 20%;
}

#finished-button:hover {
    box-shadow: 0px 0px 0px 2.5px rgba(255,255,255,0.2);
}

#todoHeader {
  user-select:none;
  --wails-draggable:drag;
  display: block;
  width: 100%;
  height: 5 px;
  color: grey;
}
</style>
