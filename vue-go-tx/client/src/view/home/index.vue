<template>
  <div class="body">
    <Banner 
    :fullAvatarUrl="fullAvatarUrl"
    />
    <Msg 
    :friends="friends"
    @getFriendInfo="getFriendInfo"
    :chats="chats"
    @friendWindow="friendWindow"
    />
    <div class="chat_box" v-if="content==='chat' && currentChatWindow">
      <div class="chat_t">
        <div class="chat_t_l">
          <div class="el-icon-arrow-left"></div>
          <h4>{{ currentChatWindow.friend_name }}</h4>
        </div>
        <div class="chat_t_r">
          <div class="el-icon-bell"></div>
          <div class="el-icon-more"></div>
        </div>
      </div>
      <div class="chat_c">
        <h4>当前用户: {{ username }}</h4>
          <div class="messages" ref="messages">
            <div v-for="(msg, index) in filteredMessages" :key="index" 
                 :data-sender-id="msg.username === username ? 'user' : 'nouser'" 
                 class="messages_item">
              <p v-if="shouldShowTimestamp(index)">{{ formatTime(msg.timestamp) }}</p>
              <div class="messages_item_b">
                <img :src="msg.avatar_url ? require(`@/assets/tx/${msg.avatar_url}`) :
                require(`@/assets/tx/11.jpg`)"/>
                <div class="messages_item_b_r">
                  <span>{{ msg.username === username ? username : msg.username }}</span>
                  <el-tag effect="plain" type="warning">{{ msg.message }}</el-tag>
                </div>
              </div>
            </div>
          </div>
      </div>
      <div class="chat_b">
        <input v-model="msg" placeholder="请输入内容" @keyup.enter="sendMessage"  class="input-field"/>
        <el-button type="primary"  @click="sendMessage">发送</el-button>
      </div>
    </div>
    <div class="friend_box" v-if="content==='friend'">
      <div class="friend_title">
        <div class="el-icon-arrow-left"></div>
        <h4>新的朋友</h4>
      </div>
      <el-input size="small" prefix-icon="el-icon-search" 
       placeholder="请输入好友名字"  v-model="search" class="friend_input"
       suffix-icon="el-icon-plus" @keyup.enter.native="addFriends">
      </el-input>
      <div class="friend_list">
        <div v-for="apply in apply_friends" :key="apply.id" >
          <FriendItem :name="apply.username"
           :remark="apply.remark"
           :avatar="require('@/assets/tiktok.png')"
           @agreeFriend="agreeFriend"
           />
        </div>
      </div>
    </div>
    <div class="friend_info_box" v-if="content==='friend_info'">
      <div class="friend_info_title">
        <img :src="friend.avatar_url ? require(`@/assets/tx/${friend.avatar_url}`) :
            require('@/assets/tx/11.jpg')">
        <div class="friend_info_title_right">
          <p>{{ friend.friend_name }}</p>
          <p>昵称:&nbsp;{{ friend.friend_name }}</p>
          <p>微信号:&nbsp;{{ friend.friend_id }}</p>
          <p>地区:&nbsp;四川成都</p>
        </div>
      </div>
      <div class="friend_info_button">
        <el-button type="success" @click="friendWindow(friend)">发消息</el-button>
        <el-button type="danger">删除</el-button>
      </div>
    </div>
  </div>
</template>

<script src="./index.js"></script>

<style lang="scss" scoped src="./index.scss"></style>