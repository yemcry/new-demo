<template>
  <div class="msg_box">
    <div class="search">
      <el-input prefix-icon="el-icon-search" 
      placeholder="请输入内容" size="small" suffix-icon="el-icon-plus">
    </el-input>
    </div>
    <div v-if="page==='msg'" class="msg_list">
      <div class="default">
        <img src="@/assets/tiktok.png"/>
        <div class="default_r">
          <span>聊天交友群</span>
          <span>测试测试测试</span>
        </div>
      </div>
      <ChatItem v-for="v in chats" :key="v.friend_id"
      :avatar="v.avatar_url ? require(`@/assets/tx/${v.avatar_url}`) :
      require('@/assets/tx/11.jpg')"
      :name="v.friend_name"
      :lastMessage="v.friend_id"
      :lastTime="lastTime"
      @chat-selected="handleChatSelected"
      @click="friendWindow(v)"
      />
    </div>

    <div v-if="page==='homi'" class="homi_box">
      <div class="homi_friend" @click="changeContent('friend')">
        <img :src="require('@/assets/svg/新的朋友.svg')">
        <span>新的朋友</span>
      </div>
      <div class="homi_group">
        <img :src="require('@/assets/svg/群聊.svg')">
        <span>群聊</span>
      </div>
      <div class="homi_list">
        <div @click="handleClick(friend)" v-for="friend in friends" :key="friend.id" >
          <HomiItem :name="friend.friend_name" 
          :avatar="friend.avatar_url ? require(`@/assets/tx/${friend.avatar_url}`) :
            require('@/assets/tx/11.jpg')"  />
        </div>
      </div>
    </div>
    <div v-if="page==='pyq'">
      朋友圈
    </div>
  </div>
</template>

<script src="./index.js"></script>

<style lang="scss" scoped src="./index.scss"></style>